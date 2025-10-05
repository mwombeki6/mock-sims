package handlers

import (
	"github.com/mwombeki6/mock-sims/internal/config"
	"gorm.io/gorm"
)

// Handlers aggregates all handler groups
type Handlers struct {
	OAuth   *OAuthHandler
	Student *StudentHandler
	Faculty *FacultyHandler
	Course  *CourseHandler
	Admin   *AdminHandler
	Docs    *DocsHandler
}

// New creates a new Handlers instance
func New(db *gorm.DB, cfg *config.Config) *Handlers {
	return &Handlers{
		OAuth:   NewOAuthHandler(db, cfg),
		Student: NewStudentHandler(db, cfg),
		Faculty: NewFacultyHandler(db, cfg),
		Course:  NewCourseHandler(db, cfg),
		Admin:   NewAdminHandler(db, cfg),
		Docs:    NewDocsHandler(),
	}
}
