package database

import (
	"fmt"
	"log"

	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect establishes database connection
func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("✅ Database connected successfully")
	return db, nil
}

// Migrate runs auto-migration for all models
func Migrate(db *gorm.DB) error {
	log.Println("🔄 Running database migrations...")

	err := db.AutoMigrate(
		// Organizational structure
		&models.College{},
		&models.Department{},
		&models.Program{},
		&models.Venue{},
		&models.Semester{},

		// People
		&models.User{},
		&models.Student{},
		&models.Faculty{},
		&models.Admin{},

		// Academic
		&models.Course{},
		&models.Lecture{},
		&models.Enrollment{},
		&models.Grade{},
		&models.CourseAssignment{},

		// OAuth
		&models.OAuthClient{},
		&models.OAuthAuthorizationCode{},
		&models.OAuthAccessToken{},

		// Webhooks & Payments
		&models.WebhookLog{},
		&models.Payment{},
	)

	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("✅ Migrations completed successfully")
	return nil
}
