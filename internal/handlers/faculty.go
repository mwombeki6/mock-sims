package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/services"
	"gorm.io/gorm"
)

type FacultyHandler struct {
	db             *gorm.DB
	cfg            *config.Config
	facultyService *services.FacultyService
}

func NewFacultyHandler(db *gorm.DB, cfg *config.Config) *FacultyHandler {
	return &FacultyHandler{
		db:             db,
		cfg:            cfg,
		facultyService: services.NewFacultyService(db, cfg),
	}
}

// GetMe returns authenticated faculty member's profile
// GET /api/faculty/me
func (h *FacultyHandler) GetMe(c *fiber.Ctx) error {
	// Get user ID from auth middleware
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	// Get faculty profile
	faculty, err := h.facultyService.GetFacultyByUserID(userID.(uint))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response
	return c.JSON(fiber.Map{
		"faculty_id":  faculty.ID,
		"staff_id":    faculty.StaffID,
		"name":        faculty.FirstName + " " + faculty.MiddleName + " " + faculty.LastName,
		"first_name":  faculty.FirstName,
		"middle_name": faculty.MiddleName,
		"last_name":   faculty.LastName,
		"email":       faculty.User.Email,
		"rank":        faculty.Rank,
		"specialization": faculty.Specialization,
		"department": fiber.Map{
			"name":    faculty.Department.Name,
			"code":    faculty.Department.Code,
			"college": faculty.Department.College.Name,
		},
	})
}

// GetCourses returns faculty's teaching assignments
// GET /api/faculty/:id/courses
func (h *FacultyHandler) GetCourses(c *fiber.Ctx) error {
	facultyIDStr := c.Params("id")
	facultyID, err := strconv.ParseUint(facultyIDStr, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid faculty ID",
		})
	}

	// Get course assignments
	assignments, err := h.facultyService.GetFacultyCourses(uint(facultyID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response
	var courses []fiber.Map
	for _, assignment := range assignments {
		courses = append(courses, fiber.Map{
			"course_code": assignment.Course.Code,
			"course_name": assignment.Course.Name,
			"credits":     assignment.Course.Credits,
			"level":       assignment.Course.Level,
			"department":  assignment.Course.Department.Name,
			"role":        assignment.Role,
			"semester":    assignment.Semester.Name,
		})
	}

	return c.JSON(fiber.Map{
		"faculty_id": facultyID,
		"courses":    courses,
		"total":      len(courses),
	})
}

// SubmitCAMarks submits Continuous Assessment marks
// POST /api/faculty/courses/:id/ca-marks
func (h *FacultyHandler) SubmitCAMarks(c *fiber.Ctx) error {
	// Get user ID from auth middleware
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	// Get faculty profile
	faculty, err := h.facultyService.GetFacultyByUserID(userID.(uint))
	if err != nil {
		return c.Status(403).JSON(fiber.Map{
			"error": "not a faculty member",
		})
	}

	// Get course ID
	courseIDStr := c.Params("id")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid course ID",
		})
	}

	// Parse request body
	type CAMarkEntry struct {
		StudentID uint    `json:"student_id"`
		CAMarks   float64 `json:"ca_marks"`
	}

	var request struct {
		Marks []CAMarkEntry `json:"marks"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// Convert to service format
	var marks []struct {
		StudentID uint
		CAMarks   float64
	}
	for _, entry := range request.Marks {
		marks = append(marks, struct {
			StudentID uint
			CAMarks   float64
		}{
			StudentID: entry.StudentID,
			CAMarks:   entry.CAMarks,
		})
	}

	// Submit CA marks
	if err := h.facultyService.SubmitCAMarks(faculty.ID, uint(courseID), marks); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "CA marks submitted successfully",
		"course_id": courseID,
		"marks_count": len(marks),
	})
}
