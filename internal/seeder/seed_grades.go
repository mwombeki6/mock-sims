package seeder

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mwombeki6/mock-sims/internal/models"
	"github.com/mwombeki6/mock-sims/internal/utils"
)

// SeedGrades generates transcripts for completed semesters and payment history for students.
func (s *Seeder) SeedGrades() error {
	var previousSemester models.Semester
	if err := s.db.Where("is_current = ?", false).Order("start_date DESC").First(&previousSemester).Error; err != nil {
		// If there is no previous semester there is nothing to grade yet
		return nil
	}

	var enrollments []models.Enrollment
	if err := s.db.Where("semester_id = ?", previousSemester.ID).Find(&enrollments).Error; err != nil {
		return err
	}

	if len(enrollments) == 0 {
		return nil
	}

	submittedAt := previousSemester.EndDate
	if submittedAt.IsZero() {
		submittedAt = time.Now().AddDate(0, -2, 0)
	}

	for _, enrollment := range enrollments {
		grade := models.Grade{EnrollmentID: enrollment.ID}
		total := randomFloat(52, 89)
		if enrollment.ID%7 == 0 {
			total = randomFloat(36, 55) // add some borderline grades
		}

		grade.CAMarks = total * 0.4
		grade.FinalExam = total * 0.6
		grade.TotalMarks = grade.CAMarks + grade.FinalExam
		grade.LetterGrade = letterFromTotal(grade.TotalMarks)
		grade.GradePoint = gradePointFromLetter(grade.LetterGrade)
		remark := remarkFromLetter(grade.LetterGrade)
		grade.Remarks = remark
		grade.StudentID = enrollment.StudentID
		grade.CourseID = enrollment.CourseID
		grade.SubmittedAt = &submittedAt

		if err := s.db.FirstOrCreate(&grade, models.Grade{EnrollmentID: enrollment.ID}).Error; err != nil {
			return err
		}
	}

	return nil
}

// SeedPayments mirrors SIMS invoices across academic years.
func (s *Seeder) SeedPayments() error {
	var students []models.Student
	if err := s.db.Find(&students).Error; err != nil {
		return err
	}
	if len(students) == 0 {
		return nil
	}

	var semesters []models.Semester
	if err := s.db.Order("start_date ASC").Find(&semesters).Error; err != nil {
		return err
	}

	academicYears := uniqueAcademicYears(semesters)
	if len(academicYears) == 0 {
		academicYears = []string{"2024/2025"}
	}

	sequence := 1
	for _, student := range students {
		for _, year := range academicYears {
			for idx := 0; idx < 3; idx++ {
				description := paymentDescriptions[(sequence+idx)%len(paymentDescriptions)]
				amount := amountForDescription(description)
				method := paymentMethods[(sequence+idx)%len(paymentMethods)]
				status := paymentStatuses[(sequence+idx)%len(paymentStatuses)]
				semesterName := semesterLabel(idx)
				paymentDate := paymentDateForYear(year, idx)

				payment := models.Payment{
					StudentID:     student.ID,
					InvoiceNumber: utils.GenerateInvoiceNumber(sequence + idx),
					Amount:        amount,
					PaymentMethod: method,
					PaymentDate:   paymentDate,
					AcademicYear:  year,
					Semester:      semesterName,
					Status:        status,
					ReceiptNumber: fmt.Sprintf("RCP%06d", sequence+idx),
				}

				if err := s.db.FirstOrCreate(&payment, models.Payment{InvoiceNumber: payment.InvoiceNumber}).Error; err != nil {
					return err
				}
			}
			sequence += 3
		}
	}

	return nil
}

func letterFromTotal(total float64) string {
	switch {
	case total >= 70:
		return "A"
	case total >= 65:
		return "B+"
	case total >= 60:
		return "B"
	case total >= 50:
		return "C"
	case total >= 40:
		return "D"
	default:
		return "F"
	}
}

func gradePointFromLetter(letter string) float64 {
	switch letter {
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

func remarkFromLetter(letter string) string {
	switch letter {
	case "A":
		return "PASS"
	case "B+":
		return "PASS"
	case "B":
		return "PASS"
	case "C":
		return "PASS"
	case "D":
		return "SUPP"
	default:
		return "RETAKE"
	}
}

func uniqueAcademicYears(semesters []models.Semester) []string {
	seen := make(map[string]struct{})
	var years []string
	for _, semester := range semesters {
		if semester.AcademicYear == "" {
			continue
		}
		if _, ok := seen[semester.AcademicYear]; !ok {
			years = append(years, semester.AcademicYear)
			seen[semester.AcademicYear] = struct{}{}
		}
	}
	if len(years) == 0 {
		return years
	}
	return years
}

func amountForDescription(description string) float64 {
	switch strings.ToLower(description) {
	case "tuition fee":
		return 900000
	case "accommodation fee":
		return 107100
	case "administrative fee":
		return 200000
	case "meals fee":
		return 350000
	case "library fee":
		return 50000
	case "ict service fee":
		return 60000
	default:
		return 100000
	}
}

func semesterLabel(idx int) string {
	if idx%2 == 0 {
		return "Semester I"
	}
	return "Semester II"
}

func paymentDateForYear(academicYear string, idx int) time.Time {
	// Academic year format YYYY/YYYY
	parts := strings.Split(academicYear, "/")
	year := time.Now().Year()
	if len(parts) == 2 {
		if parsed, err := strconv.Atoi(parts[0]); err == nil {
			year = parsed
		}
	}
	month := time.Month((idx % 6) + 1) // spread payments across months
	day := (idx%20 + 5)
	return time.Date(year, month, day, 10, 30, 0, 0, time.UTC)
}
