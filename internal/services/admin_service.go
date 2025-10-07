package services

import (
	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/models"
	"gorm.io/gorm"
)

type AdminService struct {
	db             *gorm.DB
	cfg            *config.Config
	webhookService *WebhookService
}

func NewAdminService(db *gorm.DB, cfg *config.Config) *AdminService {
	return &AdminService{
		db:             db,
		cfg:            cfg,
		webhookService: NewWebhookService(db, cfg),
	}
}

// GetAllColleges retrieves all colleges
func (s *AdminService) GetAllColleges() ([]models.College, error) {
	var colleges []models.College
	err := s.db.
		Preload("Departments").
		Order("code").
		Find(&colleges).Error

	if err != nil {
		return nil, err
	}

	return colleges, nil
}

// GetAllDepartments retrieves all departments
func (s *AdminService) GetAllDepartments() ([]models.Department, error) {
	var departments []models.Department
	err := s.db.
		Preload("College").
		Preload("Programs").
		Order("code").
		Find(&departments).Error

	if err != nil {
		return nil, err
	}

	return departments, nil
}

// GetDepartmentsByCollege retrieves departments filtered by college
func (s *AdminService) GetDepartmentsByCollege(collegeID uint) ([]models.Department, error) {
	var departments []models.Department
	err := s.db.
		Where("college_id = ?", collegeID).
		Preload("College").
		Preload("Programs").
		Order("code").
		Find(&departments).Error

	if err != nil {
		return nil, err
	}

	return departments, nil
}

// GetAllPrograms retrieves all programs
func (s *AdminService) GetAllPrograms() ([]models.Program, error) {
	var programs []models.Program
	err := s.db.
		Preload("Department").
		Preload("Department.College").
		Order("degree_level, code").
		Find(&programs).Error

	if err != nil {
		return nil, err
	}

	return programs, nil
}

// GetProgramsByDepartment retrieves programs filtered by department
func (s *AdminService) GetProgramsByDepartment(departmentID uint) ([]models.Program, error) {
	var programs []models.Program
	err := s.db.
		Where("department_id = ?", departmentID).
		Preload("Department").
		Order("degree_level, code").
		Find(&programs).Error

	if err != nil {
		return nil, err
	}

	return programs, nil
}

// GetProgramsByLevel retrieves programs filtered by degree level
func (s *AdminService) GetProgramsByLevel(degreeLevel string) ([]models.Program, error) {
	var programs []models.Program
	err := s.db.
		Where("degree_level = ?", degreeLevel).
		Preload("Department").
		Preload("Department.College").
		Order("code").
		Find(&programs).Error

	if err != nil {
		return nil, err
	}

	return programs, nil
}

// CreateBulkEnrollments creates multiple enrollments at once
func (s *AdminService) CreateBulkEnrollments(enrollments []models.Enrollment) error {
	var createdEnrollments []models.Enrollment

	// Use transaction for bulk insert
	err := s.db.Transaction(func(tx *gorm.DB) error {
		for _, enrollment := range enrollments {
			// Check if enrollment already exists
			var existing models.Enrollment
			err := tx.Where("student_id = ? AND course_id = ? AND semester_id = ?",
				enrollment.StudentID, enrollment.CourseID, enrollment.SemesterID).
				First(&existing).Error

			if err == nil {
				// Enrollment exists, skip
				continue
			}

			// Create new enrollment
			if err := tx.Create(&enrollment).Error; err != nil {
				return err
			}

			// Track created enrollment for webhook notification
			createdEnrollments = append(createdEnrollments, enrollment)
		}
		return nil
	})

	if err != nil {
		return err
	}

	// Send webhook notifications for newly created enrollments (async, ignore errors)
	go func(enrollments []models.Enrollment) {
		for i := range enrollments {
			enrollment := enrollments[i]
			s.webhookService.SendEnrollmentCreated(&enrollment)
		}
	}(append([]models.Enrollment(nil), createdEnrollments...))

	return nil
}

// GetCurrentSemester retrieves the current active semester
func (s *AdminService) GetCurrentSemester() (*models.Semester, error) {
	var semester models.Semester
	err := s.db.Where("is_current = ?", true).First(&semester).Error

	if err != nil {
		return nil, err
	}

	return &semester, nil
}

// GetAllSemesters retrieves all semesters
func (s *AdminService) GetAllSemesters() ([]models.Semester, error) {
	var semesters []models.Semester
	err := s.db.Order("id DESC").Find(&semesters).Error

	if err != nil {
		return nil, err
	}

	return semesters, nil
}
