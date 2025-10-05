package seeder

import (
	"fmt"
	"time"

	"github.com/mwombeki6/mock-sims/internal/models"
)

// SeedGrades creates sample grades for students
func (s *Seeder) SeedGrades() error {
	// Get enrollments
	var enrollments []models.Enrollment
	s.db.Preload("Student").Preload("Course").Limit(10).Find(&enrollments)

	if len(enrollments) == 0 {
		return nil
	}

	// Sample grade data (CA marks, Final Exam marks)
	gradeData := []struct {
		CAMarks   float64
		FinalExam float64
	}{
		{CAMarks: 35, FinalExam: 52}, // Total: 87 (A)
		{CAMarks: 32, FinalExam: 48}, // Total: 80 (A)
		{CAMarks: 30, FinalExam: 45}, // Total: 75 (A)
		{CAMarks: 28, FinalExam: 40}, // Total: 68 (B+)
		{CAMarks: 25, FinalExam: 38}, // Total: 63 (B)
		{CAMarks: 22, FinalExam: 35}, // Total: 57 (C)
		{CAMarks: 20, FinalExam: 32}, // Total: 52 (C)
		{CAMarks: 18, FinalExam: 28}, // Total: 46 (D)
		{CAMarks: 15, FinalExam: 25}, // Total: 40 (D)
		{CAMarks: 12, FinalExam: 20}, // Total: 32 (F)
	}

	for i, enrollment := range enrollments {
		if i >= len(gradeData) {
			break
		}

		data := gradeData[i]
		totalMarks := data.CAMarks + data.FinalExam
		letterGrade := calculateLetterGrade(totalMarks)
		gradePoint := calculateGradePoint(letterGrade)

		now := time.Now()
		grade := models.Grade{
			EnrollmentID: enrollment.ID,
			StudentID:    enrollment.StudentID,
			CourseID:     enrollment.CourseID,
			CAMarks:      data.CAMarks,
			FinalExam:    data.FinalExam,
			TotalMarks:   totalMarks,
			LetterGrade:  letterGrade,
			GradePoint:   gradePoint,
			Remarks:      generateRemarks(letterGrade),
			SubmittedAt:  &now,
		}

		if err := s.db.FirstOrCreate(&grade, models.Grade{EnrollmentID: grade.EnrollmentID}).Error; err != nil {
			return err
		}
	}

	return nil
}

// SeedPayments creates sample payment records
func (s *Seeder) SeedPayments() error {
	// Get students
	var students []models.Student
	s.db.Limit(5).Find(&students)

	if len(students) == 0 {
		return nil
	}

	// Get current semester
	var semester models.Semester
	if err := s.db.Where("is_current = ?", true).First(&semester).Error; err != nil {
		return nil
	}

	// Create payment for each student
	for i, student := range students {
		payment := models.Payment{
			StudentID:     student.ID,
			InvoiceNumber: generateInvoiceNumber(i + 1),
			Amount:        1500000.00, // Full tuition
			PaymentMethod: "Bank",
			PaymentDate:   time.Now().Add(-15 * 24 * time.Hour), // 15 days ago
			AcademicYear:  semester.AcademicYear,
			Semester:      "Semester II",
			Status:        "completed",
			ReceiptNumber: generateReceiptNumber(i + 1),
		}

		if err := s.db.FirstOrCreate(&payment, models.Payment{InvoiceNumber: payment.InvoiceNumber}).Error; err != nil {
			return err
		}
	}

	return nil
}

// Helper functions

func calculateLetterGrade(totalMarks float64) string {
	if totalMarks >= 70 {
		return "A"
	} else if totalMarks >= 65 {
		return "B+"
	} else if totalMarks >= 60 {
		return "B"
	} else if totalMarks >= 50 {
		return "C"
	} else if totalMarks >= 40 {
		return "D"
	}
	return "F"
}

func calculateGradePoint(letterGrade string) float64 {
	switch letterGrade {
	case "A":
		return 5.0
	case "B+":
		return 4.0
	case "B":
		return 3.0
	case "C":
		return 2.0
	case "D":
		return 1.0
	default:
		return 0.0
	}
}

func generateRemarks(letterGrade string) string {
	switch letterGrade {
	case "A":
		return "Excellent performance"
	case "B+":
		return "Very good performance"
	case "B":
		return "Good performance"
	case "C":
		return "Satisfactory performance"
	case "D":
		return "Pass"
	default:
		return "Fail - supplementary required"
	}
}

func generateInvoiceNumber(sequence int) string {
	return fmt.Sprintf("MB%d", 1010231180+sequence)
}

func generateReceiptNumber(sequence int) string {
	return fmt.Sprintf("RCP%d", 2024001000+sequence)
}
