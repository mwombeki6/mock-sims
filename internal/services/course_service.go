package services

import (
	"errors"

	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/models"
	"gorm.io/gorm"
)

type CourseService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewCourseService(db *gorm.DB, cfg *config.Config) *CourseService {
	return &CourseService{
		db:  db,
		cfg: cfg,
	}
}

// ListCourses retrieves all courses with pagination
func (s *CourseService) ListCourses(page, limit int) ([]models.Course, int64, error) {
	var courses []models.Course
	var total int64

	// Count total
	s.db.Model(&models.Course{}).Count(&total)

	// Get paginated results
	offset := (page - 1) * limit
	err := s.db.
		Preload("Department").
		Preload("Department.College").
		Offset(offset).
		Limit(limit).
		Order("code").
		Find(&courses).Error

	if err != nil {
		return nil, 0, err
	}

	return courses, total, nil
}

// GetCourseByCode retrieves course by code
func (s *CourseService) GetCourseByCode(code string) (*models.Course, error) {
	var course models.Course
	err := s.db.
		Preload("Department").
		Preload("Department.College").
		Preload("Programs").
		Where("code = ?", code).
		First(&course).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("course not found")
		}
		return nil, err
	}

	return &course, nil
}

// GetCourseLectures retrieves lecture schedule for a course
func (s *CourseService) GetCourseLectures(courseCode string) ([]models.Lecture, error) {
	// First, find course ID
	var course models.Course
	if err := s.db.Where("code = ?", courseCode).First(&course).Error; err != nil {
		return nil, errors.New("course not found")
	}

	// Get lectures for current semester
	var lectures []models.Lecture
	err := s.db.
		Preload("Faculty").
		Preload("Venue").
		Preload("Semester").
		Joins("JOIN semesters ON semesters.id = lectures.semester_id").
		Where("lectures.course_id = ? AND semesters.is_current = ?", course.ID, true).
		Order("day_of_week, start_time").
		Find(&lectures).Error

	if err != nil {
		return nil, err
	}

	return lectures, nil
}

// GetCourseStudents retrieves students enrolled in a course
func (s *CourseService) GetCourseStudents(courseCode string) ([]models.Enrollment, error) {
	// Find course
	var course models.Course
	if err := s.db.Where("code = ?", courseCode).First(&course).Error; err != nil {
		return nil, errors.New("course not found")
	}

	// Get enrollments for current semester
	var enrollments []models.Enrollment
	err := s.db.
		Preload("Student").
		Preload("Student.Program").
		Preload("Semester").
		Joins("JOIN semesters ON semesters.id = enrollments.semester_id").
		Where("enrollments.course_id = ? AND semesters.is_current = ?", course.ID, true).
		Find(&enrollments).Error

	if err != nil {
		return nil, err
	}

	return enrollments, nil
}

// SearchCourses searches courses by name or code
func (s *CourseService) SearchCourses(query string) ([]models.Course, error) {
	var courses []models.Course
	err := s.db.
		Preload("Department").
		Where("code LIKE ? OR name LIKE ?", "%"+query+"%", "%"+query+"%").
		Limit(20).
		Find(&courses).Error

	if err != nil {
		return nil, err
	}

	return courses, nil
}
