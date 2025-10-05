package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/services"
	"gorm.io/gorm"
)

type StudentHandler struct {
	db             *gorm.DB
	cfg            *config.Config
	studentService *services.StudentService
}

func NewStudentHandler(db *gorm.DB, cfg *config.Config) *StudentHandler {
	return &StudentHandler{
		db:             db,
		cfg:            cfg,
		studentService: services.NewStudentService(db, cfg),
	}
}

// GetMe returns authenticated student's profile
// GET /api/students/me
func (h *StudentHandler) GetMe(c *fiber.Ctx) error {
	// Get user ID from auth middleware
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	// Get student profile
	student, err := h.studentService.GetStudentByUserID(userID.(uint))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Calculate current GPA
	gpa, _ := h.studentService.CalculateGPA(student.ID)
	student.GPA = gpa

	// Build response matching real SIMS structure
	return c.JSON(fiber.Map{
		"student_id":   student.ID,
		"reg_number":   student.RegNumber,
		"name":         student.FirstName + " " + student.MiddleName + " " + student.LastName,
		"first_name":   student.FirstName,
		"middle_name":  student.MiddleName,
		"last_name":    student.LastName,
		"email":        student.User.Email,
		"program": fiber.Map{
			"code":       student.Program.Code,
			"name":       student.Program.Name,
			"level":      student.Program.DegreeLevel,
			"department": student.Program.Department.Name,
			"college":    student.Program.Department.College.Name,
		},
		"year_of_study":      student.YearOfStudy,
		"gpa":                gpa,
		"enrollment_status":  student.EnrollmentStatus,
		"payment_status":     student.PaymentStatus,
		"admission_year":     student.AdmissionYear,
		"semester":           "2024/2025 - Semester II",
	})
}

// GetCourses returns student's enrolled courses
// GET /api/students/:id/courses
func (h *StudentHandler) GetCourses(c *fiber.Ctx) error {
	studentIDStr := c.Params("id")
	studentID, err := strconv.ParseUint(studentIDStr, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid student ID",
		})
	}

	// Get enrollments
	enrollments, err := h.studentService.GetStudentCourses(uint(studentID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response
	var courses []fiber.Map
	for _, enrollment := range enrollments {
		courses = append(courses, fiber.Map{
			"course_code": enrollment.Course.Code,
			"course_name": enrollment.Course.Name,
			"credits":     enrollment.Course.Credits,
			"level":       enrollment.Course.Level,
			"department":  enrollment.Course.Department.Name,
			"semester":    enrollment.Semester.Name,
			"status":      enrollment.Status,
			"enrolled_at": enrollment.EnrolledAt,
		})
	}

	return c.JSON(fiber.Map{
		"student_id": studentID,
		"courses":    courses,
		"total":      len(courses),
	})
}

// GetGrades returns student's grades
// GET /api/students/:id/grades
func (h *StudentHandler) GetGrades(c *fiber.Ctx) error {
	studentIDStr := c.Params("id")
	studentID, err := strconv.ParseUint(studentIDStr, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid student ID",
		})
	}

	// Get grades
	grades, err := h.studentService.GetStudentGrades(uint(studentID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response matching real SIMS
	var gradesList []fiber.Map
	for _, grade := range grades {
		gradesList = append(gradesList, fiber.Map{
			"course_code":   grade.Course.Code,
			"course_name":   grade.Course.Name,
			"credits":       grade.Course.Credits,
			"ca_marks":      grade.CAMarks,
			"final_exam":    grade.FinalExam,
			"total_marks":   grade.TotalMarks,
			"letter_grade":  grade.LetterGrade,
			"grade_point":   grade.GradePoint,
			"semester":      grade.Enrollment.Semester.Name,
			"remark":        grade.Remarks,
			"submitted_at":  grade.SubmittedAt,
		})
	}

	return c.JSON(fiber.Map{
		"student_id": studentID,
		"grades":     gradesList,
		"total":      len(gradesList),
	})
}

// GetTimetable returns student's class timetable
// GET /api/students/:id/timetable
func (h *StudentHandler) GetTimetable(c *fiber.Ctx) error {
	studentIDStr := c.Params("id")
	studentID, err := strconv.ParseUint(studentIDStr, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid student ID",
		})
	}

	// Get timetable
	lectures, err := h.studentService.GetStudentTimetable(uint(studentID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response
	var timetable []fiber.Map
	for _, lecture := range lectures {
		timetable = append(timetable, fiber.Map{
			"course_code": lecture.Course.Code,
			"course_name": lecture.Course.Name,
			"day":         lecture.DayOfWeek,
			"start_time":  lecture.StartTime,
			"end_time":    lecture.EndTime,
			"venue":       lecture.Venue.Building + " " + lecture.Venue.RoomNumber,
			"lecturer":    lecture.Faculty.FirstName + " " + lecture.Faculty.LastName,
			"semester":    lecture.Semester.Name,
		})
	}

	return c.JSON(fiber.Map{
		"student_id": studentID,
		"timetable":  timetable,
		"semester":   "2024/2025 - Semester II",
	})
}
