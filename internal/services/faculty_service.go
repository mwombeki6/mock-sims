package services

import (
	"errors"
	"time"

	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/models"
	"gorm.io/gorm"
)

type FacultyService struct {
	db             *gorm.DB
	cfg            *config.Config
	webhookService *WebhookService
}

func NewFacultyService(db *gorm.DB, cfg *config.Config) *FacultyService {
	return &FacultyService{
		db:             db,
		cfg:            cfg,
		webhookService: NewWebhookService(db, cfg),
	}
}

// GetFacultyByUserID retrieves faculty profile by user ID
func (s *FacultyService) GetFacultyByUserID(userID uint) (*models.Faculty, error) {
	var faculty models.Faculty
	err := s.db.
		Preload("User").
		Preload("Department").
		Preload("Department.College").
		Where("user_id = ?", userID).
		First(&faculty).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("faculty member not found")
		}
		return nil, err
	}

	return &faculty, nil
}

// GetFacultyCourses retrieves all course assignments for a faculty member
func (s *FacultyService) GetFacultyCourses(facultyID uint) ([]models.CourseAssignment, error) {
	var assignments []models.CourseAssignment
	err := s.db.
		Preload("Course").
		Preload("Course.Department").
		Preload("Semester").
		Where("faculty_id = ?", facultyID).
		Order("semester_id DESC").
		Find(&assignments).Error

	if err != nil {
		return nil, err
	}

	return assignments, nil
}

// GetCourseStudents retrieves all students enrolled in a course
func (s *FacultyService) GetCourseStudents(courseID uint, semesterID uint) ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	err := s.db.
		Preload("Student").
		Preload("Student.Program").
		Where("course_id = ? AND semester_id = ?", courseID, semesterID).
		Find(&enrollments).Error

	if err != nil {
		return nil, err
	}

	return enrollments, nil
}

// SubmitCAMarks submits Continuous Assessment marks for students
func (s *FacultyService) SubmitCAMarks(facultyID uint, courseID uint, marks []struct {
	StudentID uint
	CAMarks   float64
}) error {
	// Verify faculty is assigned to this course
	var assignment models.CourseAssignment
	err := s.db.
		Where("faculty_id = ? AND course_id = ?", facultyID, courseID).
		First(&assignment).Error

	if err != nil {
		return errors.New("you are not assigned to this course")
	}

	var updatedGrades []models.Grade

	// Update grades for each student
	for _, mark := range marks {
		var grade models.Grade
		var enrollment models.Enrollment

		// Find enrollment
		err := s.db.
			Where("student_id = ? AND course_id = ?", mark.StudentID, courseID).
			First(&enrollment).Error

		if err != nil {
			continue // Skip if enrollment not found
		}

		// Find or create grade record
		err = s.db.Where("enrollment_id = ?", enrollment.ID).First(&grade).Error
		now := time.Now()
		if err != nil {
			// Create new grade record
			grade = models.Grade{
				EnrollmentID: enrollment.ID,
				StudentID:    mark.StudentID,
				CourseID:     courseID,
				CAMarks:      mark.CAMarks,
				SubmittedAt:  &now,
			}
			s.db.Create(&grade)
		} else {
			// Update existing grade
			grade.CAMarks = mark.CAMarks
			grade.SubmittedAt = &now
			// Recalculate total if final exam marks exist
			if grade.FinalExam > 0 {
				grade.TotalMarks = grade.CAMarks + grade.FinalExam
				grade.LetterGrade = s.calculateLetterGrade(grade.TotalMarks)
				grade.GradePoint = s.calculateGradePoint(grade.LetterGrade)
			}
			s.db.Save(&grade)
		}

		// Track updated grade for webhook notification
		updatedGrades = append(updatedGrades, grade)
	}

	// Send webhook notifications for submitted grades (async, ignore errors)
	go func() {
		for _, grade := range updatedGrades {
			// Only send webhook if grade is complete (has both CA and final exam)
			if grade.TotalMarks > 0 && grade.LetterGrade != "" {
				s.webhookService.SendGradeSubmitted(&grade)
			}
		}
	}()

	return nil
}

// calculateLetterGrade converts total marks to letter grade
func (s *FacultyService) calculateLetterGrade(totalMarks float64) string {
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

// calculateGradePoint converts letter grade to grade point
func (s *FacultyService) calculateGradePoint(letterGrade string) float64 {
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
