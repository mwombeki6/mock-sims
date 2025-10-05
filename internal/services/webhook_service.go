package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/models"
	"github.com/mwombeki6/mock-sims/internal/utils"
	"gorm.io/gorm"
)

type WebhookService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewWebhookService(db *gorm.DB, cfg *config.Config) *WebhookService {
	return &WebhookService{
		db:  db,
		cfg: cfg,
	}
}

// WebhookPayload represents the structure of webhook notifications sent to LMS
type WebhookPayload struct {
	Event     string                 `json:"event"`
	Timestamp string                 `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
}

// SendEnrollmentCreated sends a webhook notification when a student is enrolled in a course
func (s *WebhookService) SendEnrollmentCreated(enrollment *models.Enrollment) error {
	// Build payload
	payload := WebhookPayload{
		Event:     "enrollment.created",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data: map[string]interface{}{
			"enrollment_id": enrollment.ID,
			"student_id":    enrollment.StudentID,
			"course_id":     enrollment.CourseID,
			"semester_id":   enrollment.SemesterID,
			"status":        enrollment.Status,
			"enrolled_at":   enrollment.EnrolledAt.Format(time.RFC3339),
		},
	}

	return s.sendWebhook(payload)
}

// SendEnrollmentUpdated sends a webhook notification when enrollment status changes
func (s *WebhookService) SendEnrollmentUpdated(enrollment *models.Enrollment) error {
	payload := WebhookPayload{
		Event:     "enrollment.updated",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data: map[string]interface{}{
			"enrollment_id": enrollment.ID,
			"student_id":    enrollment.StudentID,
			"course_id":     enrollment.CourseID,
			"semester_id":   enrollment.SemesterID,
			"status":        enrollment.Status,
			"updated_at":    time.Now().UTC().Format(time.RFC3339),
		},
	}

	return s.sendWebhook(payload)
}

// SendGradeSubmitted sends a webhook notification when grades are submitted
func (s *WebhookService) SendGradeSubmitted(grade *models.Grade) error {
	payload := WebhookPayload{
		Event:     "grade.submitted",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data: map[string]interface{}{
			"grade_id":      grade.ID,
			"enrollment_id": grade.EnrollmentID,
			"ca_marks":      grade.CAMarks,
			"final_exam":    grade.FinalExam,
			"total_marks":   grade.TotalMarks,
			"letter_grade":  grade.LetterGrade,
			"grade_point":   grade.GradePoint,
			"submitted_at":  grade.SubmittedAt.Format(time.RFC3339),
		},
	}

	return s.sendWebhook(payload)
}

// SendPaymentReceived sends a webhook notification when payment is received
func (s *WebhookService) SendPaymentReceived(payment *models.Payment) error {
	payload := WebhookPayload{
		Event:     "payment.received",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data: map[string]interface{}{
			"payment_id":      payment.ID,
			"student_id":      payment.StudentID,
			"invoice_number":  payment.InvoiceNumber,
			"amount":          payment.Amount,
			"payment_method":  payment.PaymentMethod,
			"payment_date":    payment.PaymentDate.Format(time.RFC3339),
			"academic_year":   payment.AcademicYear,
			"semester":        payment.Semester,
			"status":          payment.Status,
		},
	}

	return s.sendWebhook(payload)
}

// sendWebhook sends an HTTP POST request with HMAC signature to LMS webhook URL
func (s *WebhookService) sendWebhook(payload WebhookPayload) error {
	// Skip if webhook URL is not configured
	if s.cfg.LMSWebhookURL == "" {
		return nil
	}

	// Marshal payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal webhook payload: %w", err)
	}

	// Generate HMAC signature
	signature := utils.GenerateHMACSignature(payloadBytes, s.cfg.LMSWebhookSecret)

	// Create HTTP request
	req, err := http.NewRequest("POST", s.cfg.LMSWebhookURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create webhook request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-SIMS-Signature", signature)
	req.Header.Set("X-SIMS-Event", payload.Event)
	req.Header.Set("X-SIMS-Timestamp", payload.Timestamp)

	// Send request
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send webhook: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("webhook request failed with status: %d", resp.StatusCode)
	}

	// Log webhook delivery
	s.logWebhookDelivery(payload.Event, resp.StatusCode, nil)

	return nil
}

// logWebhookDelivery logs webhook delivery attempt to database
func (s *WebhookService) logWebhookDelivery(event string, statusCode int, err error) {
	log := models.WebhookLog{
		Event:      event,
		URL:        s.cfg.LMSWebhookURL,
		StatusCode: statusCode,
		SentAt:     time.Now(),
	}

	if err != nil {
		errMsg := err.Error()
		log.Error = &errMsg
	}

	// Save log (ignore errors to prevent webhook failures from affecting main operation)
	s.db.Create(&log)
}

// RetryFailedWebhooks attempts to retry webhooks that failed to deliver
func (s *WebhookService) RetryFailedWebhooks() error {
	var failedLogs []models.WebhookLog

	// Find webhooks that failed in the last 24 hours
	if err := s.db.Where("status_code >= 400 OR error IS NOT NULL").
		Where("sent_at > ?", time.Now().Add(-24*time.Hour)).
		Order("sent_at ASC").
		Limit(100).
		Find(&failedLogs).Error; err != nil {
		return err
	}

	// Retry each failed webhook
	for _, log := range failedLogs {
		// Reconstruct payload from log (simplified - in production, store full payload)
		payload := WebhookPayload{
			Event:     log.Event,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Data:      map[string]interface{}{},
		}

		// Attempt to resend
		if err := s.sendWebhook(payload); err == nil {
			// Mark as successfully retried
			s.db.Model(&log).Update("status_code", 200)
		}
	}

	return nil
}
