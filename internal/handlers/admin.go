package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/models"
	"github.com/mwombeki6/mock-sims/internal/services"
	"gorm.io/gorm"
)

type AdminHandler struct {
	db           *gorm.DB
	cfg          *config.Config
	adminService *services.AdminService
}

func NewAdminHandler(db *gorm.DB, cfg *config.Config) *AdminHandler {
	return &AdminHandler{
		db:           db,
		cfg:          cfg,
		adminService: services.NewAdminService(db, cfg),
	}
}

// GetColleges returns all colleges
// GET /api/colleges
func (h *AdminHandler) GetColleges(c *fiber.Ctx) error {
	colleges, err := h.adminService.GetAllColleges()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response
	var collegeList []fiber.Map
	for _, college := range colleges {
		collegeList = append(collegeList, fiber.Map{
			"id":         college.ID,
			"code":       college.Code,
			"name":       college.Name,
			"short_name": college.ShortName,
			"dean":       college.Dean,
			"departments_count": len(college.Departments),
		})
	}

	return c.JSON(fiber.Map{
		"colleges": collegeList,
		"total":    len(collegeList),
	})
}

// GetDepartments returns all departments
// GET /api/departments?college_id=1
func (h *AdminHandler) GetDepartments(c *fiber.Ctx) error {
	collegeID := c.QueryInt("college_id", 0)

	var departments []models.Department
	var err error

	if collegeID > 0 {
		departments, err = h.adminService.GetDepartmentsByCollege(uint(collegeID))
	} else {
		departments, err = h.adminService.GetAllDepartments()
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response
	var deptList []fiber.Map
	for _, dept := range departments {
		deptList = append(deptList, fiber.Map{
			"id":      dept.ID,
			"code":    dept.Code,
			"name":    dept.Name,
			"head":    dept.Head,
			"college": dept.College.Name,
			"programs_count": len(dept.Programs),
		})
	}

	return c.JSON(fiber.Map{
		"departments": deptList,
		"total":       len(deptList),
	})
}

// GetPrograms returns all programs
// GET /api/programs?department_id=1&level=Bachelor
func (h *AdminHandler) GetPrograms(c *fiber.Ctx) error {
	departmentID := c.QueryInt("department_id", 0)
	level := c.Query("level")

	var programs []models.Program
	var err error

	if departmentID > 0 {
		programs, err = h.adminService.GetProgramsByDepartment(uint(departmentID))
	} else if level != "" {
		programs, err = h.adminService.GetProgramsByLevel(level)
	} else {
		programs, err = h.adminService.GetAllPrograms()
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response
	var programList []fiber.Map
	for _, program := range programs {
		programList = append(programList, fiber.Map{
			"id":           program.ID,
			"code":         program.Code,
			"name":         program.Name,
			"degree_level": program.DegreeLevel,
			"nta_level":    program.NTALevel,
			"duration":     program.Duration,
			"tuition_fees": program.TuitionFees,
			"department":   program.Department.Name,
			"college":      program.Department.College.Name,
		})
	}

	return c.JSON(fiber.Map{
		"programs": programList,
		"total":    len(programList),
	})
}

// CreateEnrollments creates bulk enrollments
// POST /api/enrollments
func (h *AdminHandler) CreateEnrollments(c *fiber.Ctx) error {
	var request struct {
		Enrollments []models.Enrollment `json:"enrollments"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// Create enrollments
	if err := h.adminService.CreateBulkEnrollments(request.Enrollments); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "enrollments created successfully",
		"count":   len(request.Enrollments),
	})
}
