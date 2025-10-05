package seeder

import (
	"github.com/mwombeki6/mock-sims/internal/models"
	"github.com/mwombeki6/mock-sims/internal/utils"
)

// SeedUsers creates test users (students, faculty, admins)
func (s *Seeder) SeedUsers() error {
	// Get programs
	var programs []models.Program
	s.db.Limit(5).Find(&programs)

	// Get departments
	var departments []models.Department
	s.db.Limit(3).Find(&departments)

	// Hash password (password123 for all test users)
	hashedPassword, err := utils.HashPassword("password123")
	if err != nil {
		return err
	}

	// Create admin user
	adminUser := models.User{
		Email:    "admin@must.ac.tz",
		Password: hashedPassword,
		UserType: "admin",
		IsActive: true,
	}
	if err := s.db.FirstOrCreate(&adminUser, models.User{Email: adminUser.Email}).Error; err != nil {
		return err
	}

	admin := models.Admin{
		UserID:    adminUser.ID,
		FirstName: "System",
		LastName:  "Administrator",
		Role:      "Registrar",
	}
	if err := s.db.FirstOrCreate(&admin, models.Admin{UserID: admin.UserID}).Error; err != nil {
		return err
	}

	// Create faculty members
	facultyData := []struct {
		Email          string
		FirstName      string
		MiddleName     string
		LastName       string
		StaffID        string
		DepartmentID   uint
		Rank           string
		Specialization string
	}{
		{
			Email:          "joseph.mkunda@must.ac.tz",
			FirstName:      "Joseph",
			MiddleName:     "T",
			LastName:       "Mkunda",
			StaffID:        "MUST-F-001",
			DepartmentID:   departments[0].ID,
			Rank:           "Professor",
			Specialization: "Computer Networks",
		},
		{
			Email:          "devotha.nyambo@must.ac.tz",
			FirstName:      "Devotha",
			MiddleName:     "G",
			LastName:       "Nyambo",
			StaffID:        "MUST-F-002",
			DepartmentID:   departments[0].ID,
			Rank:           "Senior Lecturer",
			Specialization: "Software Engineering",
		},
		{
			Email:          "mussa.dida@must.ac.tz",
			FirstName:      "Mussa",
			MiddleName:     "Ally",
			LastName:       "Dida",
			StaffID:        "MUST-F-003",
			DepartmentID:   departments[0].ID,
			Rank:           "Lecturer",
			Specialization: "Artificial Intelligence",
		},
	}

	for _, fd := range facultyData {
		user := models.User{
			Email:    fd.Email,
			Password: hashedPassword,
			UserType: "faculty",
			IsActive: true,
		}
		if err := s.db.FirstOrCreate(&user, models.User{Email: user.Email}).Error; err != nil {
			return err
		}

		faculty := models.Faculty{
			UserID:         user.ID,
			StaffID:        fd.StaffID,
			FirstName:      fd.FirstName,
			MiddleName:     fd.MiddleName,
			LastName:       fd.LastName,
			DepartmentID:   fd.DepartmentID,
			Rank:           fd.Rank,
			Specialization: fd.Specialization,
		}
		if err := s.db.FirstOrCreate(&faculty, models.Faculty{UserID: faculty.UserID}).Error; err != nil {
			return err
		}
	}

	// Create students
	studentData := []struct {
		Email        string
		FirstName    string
		MiddleName   string
		LastName     string
		RegNumber    string
		ProgramID    uint
		YearOfStudy  int
		AdmissionYear int
	}{
		{
			Email:         "john.doe@must.ac.tz",
			FirstName:     "John",
			MiddleName:    "Matias",
			LastName:      "Doe",
			RegNumber:     "23100523050032",
			ProgramID:     programs[0].ID,
			YearOfStudy:   2,
			AdmissionYear: 2023,
		},
		{
			Email:         "jane.smith@must.ac.tz",
			FirstName:     "Jane",
			MiddleName:    "Neema",
			LastName:      "Smith",
			RegNumber:     "23100523050033",
			ProgramID:     programs[0].ID,
			YearOfStudy:   2,
			AdmissionYear: 2023,
		},
		{
			Email:         "peter.mwamba@must.ac.tz",
			FirstName:     "Peter",
			MiddleName:    "Joseph",
			LastName:      "Mwamba",
			RegNumber:     "23100523050034",
			ProgramID:     programs[0].ID,
			YearOfStudy:   2,
			AdmissionYear: 2023,
		},
		{
			Email:         "mary.kimaro@must.ac.tz",
			FirstName:     "Mary",
			MiddleName:    "Charles",
			LastName:      "Kimaro",
			RegNumber:     "24100523050001",
			ProgramID:     programs[0].ID,
			YearOfStudy:   1,
			AdmissionYear: 2024,
		},
		{
			Email:         "james.mollel@must.ac.tz",
			FirstName:     "James",
			MiddleName:    "Said",
			LastName:      "Mollel",
			RegNumber:     "24100523050002",
			ProgramID:     programs[0].ID,
			YearOfStudy:   1,
			AdmissionYear: 2024,
		},
	}

	for _, sd := range studentData {
		user := models.User{
			Email:    sd.Email,
			Password: hashedPassword,
			UserType: "student",
			IsActive: true,
		}
		if err := s.db.FirstOrCreate(&user, models.User{Email: user.Email}).Error; err != nil {
			return err
		}

		student := models.Student{
			UserID:           user.ID,
			RegNumber:        sd.RegNumber,
			FirstName:        sd.FirstName,
			MiddleName:       sd.MiddleName,
			LastName:         sd.LastName,
			ProgramID:        sd.ProgramID,
			YearOfStudy:      sd.YearOfStudy,
			GPA:              0.0,
			EnrollmentStatus: "active",
			PaymentStatus:    "paid",
			AdmissionYear:    sd.AdmissionYear,
		}
		if err := s.db.FirstOrCreate(&student, models.Student{UserID: student.UserID}).Error; err != nil {
			return err
		}
	}

	return nil
}
