package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/database"
	"github.com/mwombeki6/mock-sims/internal/handlers"
	"github.com/mwombeki6/mock-sims/internal/middleware"
)

// @title Mock SIMS API
// @version 1.0
// @description MUST Student Information Management System - OAuth 2.0 Server and REST API for LMS Integration
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/mwombeki6/mock-sims
// @contact.email support@must.ac.tz

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /
// @schemes http https

// @securityDefinitions.oauth2.authorizationCode OAuth2AuthCode
// @tokenUrl /oauth/token
// @authorizationUrl /oauth/authorize
// @scope.student.read Read student information
// @scope.courses.read Read course information

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Mock SIMS v1.0",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AllowedOrigins,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "mock-sims",
			"version": "1.0.0",
		})
	})

	// Initialize handlers
	h := handlers.New(db, cfg)

	// Documentation routes
	app.Get("/docs", h.Docs.ServeAPIDocs)
	app.Get("/swagger", h.Docs.ServeSwaggerUI)
	app.Get("/redoc", h.Docs.ServeReDoc)
	app.Get("/swagger.json", h.Docs.ServeSwaggerJSON)

	// OAuth routes (both GET and POST for authorize)
	app.Get("/oauth/authorize", h.OAuth.Authorize)
	app.Post("/oauth/authorize", h.OAuth.Authorize)
	app.Post("/oauth/token", h.OAuth.Token)

	// API routes (protected)
	api := app.Group("/api", middleware.AuthMiddleware(db, cfg))

	// Student endpoints
	api.Get("/students/me", h.Student.GetMe)
	api.Get("/students/:id/courses", h.Student.GetCourses)
	api.Get("/students/:id/grades", h.Student.GetGrades)
	api.Get("/students/:id/timetable", h.Student.GetTimetable)

	// Faculty endpoints
	api.Get("/faculty/me", h.Faculty.GetMe)
	api.Get("/faculty/:id/courses", h.Faculty.GetCourses)
	api.Post("/faculty/courses/:id/ca-marks", h.Faculty.SubmitCAMarks)

	// Course endpoints
	api.Get("/courses", h.Course.List)
	api.Get("/courses/:code", h.Course.Get)
	api.Get("/courses/:code/lectures", h.Course.GetLectures)
	api.Get("/courses/:code/students", h.Course.GetStudents)

	// Admin endpoints
	api.Get("/colleges", h.Admin.GetColleges)
	api.Get("/departments", h.Admin.GetDepartments)
	api.Get("/programs", h.Admin.GetPrograms)
	api.Post("/enrollments", h.Admin.CreateEnrollments)

	// Start server
	log.Printf("🚀 Mock SIMS starting on %s:%s", cfg.Host, cfg.Port)

	// Graceful shutdown
	go func() {
		if err := app.Listen(cfg.Host + ":" + cfg.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down Mock SIMS...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("✅ Mock SIMS stopped")
}
