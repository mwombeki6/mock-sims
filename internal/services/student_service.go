package services

import (
	"errors"

	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/models"
	"gorm.io/gorm"
)

type StudentService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewStudentService(db *gorm.DB, cfg *config.Config) *StudentService {
	return &StudentService{
		db:  db,
		cfg: cfg,
	}
}

// GetStudentByUserID retrieves student profile by user ID
func (s *StudentService) GetStudentByUserID(userID uint) (*models.Student, error) {
	var student models.Student
	err := s.db.
		Preload("User").
		Preload("Program").
		Preload("Program.Department").
		Preload("Program.Department.College").
		Where("user_id = ?", userID).
		First(&student).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("student not found")
		}
		return nil, err
	}

	return &student, nil
}

// GetStudentByRegNumber retrieves student by registration number
func (s *StudentService) GetStudentByRegNumber(regNumber string) (*models.Student, error) {
	var student models.Student
	err := s.db.
		Preload("User").
		Preload("Program").
		Preload("Program.Department").
		Preload("Program.Department.College").
		Where("reg_number = ?", regNumber).
		First(&student).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("student not found")
		}
		return nil, err
	}

	return &student, nil
}

// GetStudentCourses retrieves all courses for a student in current semester
func (s *StudentService) GetStudentCourses(studentID uint) ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	err := s.db.
		Preload("Course").
		Preload("Course.Department").
		Preload("Semester").
		Where("student_id = ?", studentID).
		Order("semester_id DESC, id DESC").
		Find(&enrollments).Error

	if err != nil {
		return nil, err
	}

	return enrollments, nil
}

// GetStudentGrades retrieves all grades for a student
func (s *StudentService) GetStudentGrades(studentID uint) ([]models.Grade, error) {
	var grades []models.Grade
	err := s.db.
		Preload("Course").
		Preload("Enrollment").
		Preload("Enrollment.Semester").
		Where("student_id = ?", studentID).
		Order("id DESC").
		Find(&grades).Error

	if err != nil {
		return nil, err
	}

	return grades, nil
}

// GetStudentTimetable retrieves student's class schedule
func (s *StudentService) GetStudentTimetable(studentID uint) ([]models.Lecture, error) {
	// Get student's current enrollments
	var enrollments []models.Enrollment
	err := s.db.
		Joins("JOIN semesters ON semesters.id = enrollments.semester_id").
		Where("enrollments.student_id = ? AND semesters.is_current = ?", studentID, true).
		Select("enrollments.course_id").
		Find(&enrollments).Error

	if err != nil {
		return nil, err
	}

	if len(enrollments) == 0 {
		return []models.Lecture{}, nil
	}

	// Extract course IDs
	courseIDs := make([]uint, len(enrollments))
	for i, enrollment := range enrollments {
		courseIDs[i] = enrollment.CourseID
	}

	// Get lectures for these courses
	var lectures []models.Lecture
	err = s.db.
		Preload("Course").
		Preload("Faculty").
		Preload("Venue").
		Preload("Semester").
		Where("course_id IN ? AND semester_id IN (SELECT id FROM semesters WHERE is_current = ?)", courseIDs, true).
		Order("day_of_week, start_time").
		Find(&lectures).Error

	if err != nil {
		return nil, err
	}

	return lectures, nil
}

// CalculateGPA calculates student's cumulative GPA
func (s *StudentService) CalculateGPA(studentID uint) (float64, error) {
	// Get grades with course information
	var results []struct {
		GradePoint float64
		Credits    int
	}

	err := s.db.Table("grades").
		Joins("JOIN enrollments ON enrollments.id = grades.enrollment_id").
		Joins("JOIN courses ON courses.id = enrollments.course_id").
		Where("grades.student_id = ? AND grades.letter_grade IS NOT NULL AND grades.letter_grade != ''", studentID).
		Select("grades.grade_point, courses.credits").
		Find(&results).Error

	if err != nil {
		return 0, err
	}

	if len(results) == 0 {
		return 0, nil
	}

	var totalPoints, totalCredits float64
	for _, result := range results {
		totalPoints += result.GradePoint * float64(result.Credits)
		totalCredits += float64(result.Credits)
	}

	if totalCredits == 0 {
		return 0, nil
	}

	return totalPoints / totalCredits, nil
}
