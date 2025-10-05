package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/services"
	"gorm.io/gorm"
)

type CourseHandler struct {
	db            *gorm.DB
	cfg           *config.Config
	courseService *services.CourseService
}

func NewCourseHandler(db *gorm.DB, cfg *config.Config) *CourseHandler {
	return &CourseHandler{
		db:            db,
		cfg:           cfg,
		courseService: services.NewCourseService(db, cfg),
	}
}

// List returns all courses
// GET /api/courses?page=1&limit=20
func (h *CourseHandler) List(c *fiber.Ctx) error {
	// Get pagination parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "20"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	// Get courses
	courses, total, err := h.courseService.ListCourses(page, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response
	var courseList []fiber.Map
	for _, course := range courses {
		courseList = append(courseList, fiber.Map{
			"code":        course.Code,
			"name":        course.Name,
			"credits":     course.Credits,
			"level":       course.Level,
			"description": course.Description,
			"department":  course.Department.Name,
			"college":     course.Department.College.Name,
		})
	}

	return c.JSON(fiber.Map{
		"courses": courseList,
		"total":   total,
		"page":    page,
		"limit":   limit,
	})
}

// Get returns course details by code
// GET /api/courses/:code
func (h *CourseHandler) Get(c *fiber.Ctx) error {
	courseCode := c.Params("code")

	// Get course
	course, err := h.courseService.GetCourseByCode(courseCode)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response
	return c.JSON(fiber.Map{
		"code":        course.Code,
		"name":        course.Name,
		"credits":     course.Credits,
		"level":       course.Level,
		"description": course.Description,
		"department": fiber.Map{
			"name": course.Department.Name,
			"code": course.Department.Code,
		},
		"college": fiber.Map{
			"name": course.Department.College.Name,
			"code": course.Department.College.Code,
		},
	})
}

// GetLectures returns course lectures/schedule
// GET /api/courses/:code/lectures
func (h *CourseHandler) GetLectures(c *fiber.Ctx) error {
	courseCode := c.Params("code")

	// Get lectures
	lectures, err := h.courseService.GetCourseLectures(courseCode)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response
	var lectureList []fiber.Map
	for _, lecture := range lectures {
		lectureList = append(lectureList, fiber.Map{
			"day":        lecture.DayOfWeek,
			"start_time": lecture.StartTime,
			"end_time":   lecture.EndTime,
			"venue":      lecture.Venue.Building + " " + lecture.Venue.RoomNumber,
			"capacity":   lecture.Venue.Capacity,
			"lecturer":   lecture.Faculty.FirstName + " " + lecture.Faculty.LastName,
			"semester":   lecture.Semester.Name,
		})
	}

	return c.JSON(fiber.Map{
		"course_code": courseCode,
		"lectures":    lectureList,
		"total":       len(lectureList),
	})
}

// GetStudents returns enrolled students for a course
// GET /api/courses/:code/students
func (h *CourseHandler) GetStudents(c *fiber.Ctx) error {
	courseCode := c.Params("code")

	// Get enrolled students
	enrollments, err := h.courseService.GetCourseStudents(courseCode)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response
	var students []fiber.Map
	for _, enrollment := range enrollments {
		students = append(students, fiber.Map{
			"student_id":  enrollment.Student.ID,
			"reg_number":  enrollment.Student.RegNumber,
			"name":        enrollment.Student.FirstName + " " + enrollment.Student.LastName,
			"program":     enrollment.Student.Program.Name,
			"year":        enrollment.Student.YearOfStudy,
			"status":      enrollment.Status,
			"enrolled_at": enrollment.EnrolledAt,
		})
	}

	return c.JSON(fiber.Map{
		"course_code": courseCode,
		"students":    students,
		"total":       len(students),
	})
}
