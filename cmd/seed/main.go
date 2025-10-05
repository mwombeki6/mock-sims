package main

import (
	"log"

	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/database"
	"github.com/mwombeki6/mock-sims/internal/seeder"
)

func main() {
	log.Println("Starting Mock SIMS seed...")

	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations first
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Create seeder
	s := seeder.NewSeeder(db, cfg)

	// Run seeder
	if err := s.SeedAll(); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	log.Println("✅ Mock SIMS seed completed successfully!")
}
