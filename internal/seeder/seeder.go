package seeder

import (
	"log"
	"time"

	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/models"
	"github.com/mwombeki6/mock-sims/internal/utils"
	"gorm.io/gorm"
)

type Seeder struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewSeeder(db *gorm.DB, cfg *config.Config) *Seeder {
	return &Seeder{
		db:  db,
		cfg: cfg,
	}
}

// SeedAll seeds all tables in the correct order
func (s *Seeder) SeedAll() error {
	log.Println("Seeding colleges...")
	if err := s.SeedColleges(); err != nil {
		return err
	}

	log.Println("Seeding departments...")
	if err := s.SeedDepartments(); err != nil {
		return err
	}

	log.Println("Seeding programs...")
	if err := s.SeedPrograms(); err != nil {
		return err
	}

	log.Println("Seeding semesters...")
	if err := s.SeedSemesters(); err != nil {
		return err
	}

	log.Println("Seeding venues...")
	if err := s.SeedVenues(); err != nil {
		return err
	}

	log.Println("Seeding OAuth client...")
	if err := s.SeedOAuthClient(); err != nil {
		return err
	}

	log.Println("Seeding users...")
	if err := s.SeedUsers(); err != nil {
		return err
	}

	log.Println("Seeding courses...")
	if err := s.SeedCourses(); err != nil {
		return err
	}

	log.Println("Seeding lectures...")
	if err := s.SeedLectures(); err != nil {
		return err
	}

	log.Println("Seeding enrollments...")
	if err := s.SeedEnrollments(); err != nil {
		return err
	}

	log.Println("Seeding grades...")
	if err := s.SeedGrades(); err != nil {
		return err
	}

	log.Println("Seeding payments...")
	if err := s.SeedPayments(); err != nil {
		return err
	}

	return nil
}

// SeedColleges seeds the 7 MUST colleges
func (s *Seeder) SeedColleges() error {
	colleges := []models.College{
		{
			Code:      "01",
			Name:      "College of Information and Communication Technology",
			ShortName: "CoICT",
			Dean:      "Prof. Dr. Joseph Mkunda",
		},
		{
			Code:      "02",
			Name:      "College of Business Education and Technology",
			ShortName: "CoBET",
			Dean:      "Dr. Augustino Mwagike",
		},
		{
			Code:      "03",
			Name:      "College of Engineering and Technology",
			ShortName: "CoET",
			Dean:      "Prof. Dr. Richard Mcharo",
		},
		{
			Code:      "04",
			Name:      "College of Earth Sciences",
			ShortName: "CoES",
			Dean:      "Dr. Emmanuel Mutakyahwa",
		},
		{
			Code:      "05",
			Name:      "College of Veterinary Medicine and Biomedical Sciences",
			ShortName: "CoVMBS",
			Dean:      "Dr. Emmanuel Mellau",
		},
		{
			Code:      "06",
			Name:      "College of Life Sciences and Bioengineering",
			ShortName: "CoLSB",
			Dean:      "Dr. Ezekiel Mmbaga",
		},
		{
			Code:      "07",
			Name:      "College of Social Sciences and Humanities",
			ShortName: "CoSSH",
			Dean:      "Dr. Magreth Bushesha",
		},
	}

	for _, college := range colleges {
		if err := s.db.FirstOrCreate(&college, models.College{Code: college.Code}).Error; err != nil {
			return err
		}
	}

	return nil
}

// SeedDepartments seeds departments for each college
func (s *Seeder) SeedDepartments() error {
	// Get colleges
	var colleges []models.College
	s.db.Find(&colleges)

	// Map college codes to IDs
	collegeMap := make(map[string]uint)
	for _, c := range colleges {
		collegeMap[c.Code] = c.ID
	}

	departments := []models.Department{
		// CoICT (01)
		{CollegeID: collegeMap["01"], Code: "CS", Name: "Computer Science and Engineering", Head: "Dr. Mussa Ally Dida"},
		{CollegeID: collegeMap["01"], Code: "ICT", Name: "Information and Communication Technology", Head: "Dr. Devotha Nyambo"},
		{CollegeID: collegeMap["01"], Code: "EE", Name: "Electronics and Telecommunications Engineering", Head: "Dr. Joseph Mbelwa"},

		// CoBET (02)
		{CollegeID: collegeMap["02"], Code: "BAF", Name: "Business Administration and Finance", Head: "Dr. Honest Kimario"},
		{CollegeID: collegeMap["02"], Code: "ACC", Name: "Accounting", Head: "Mr. Japhet Mwemezi"},
		{CollegeID: collegeMap["02"], Code: "ECO", Name: "Economics", Head: "Dr. Honest Kimario"},

		// CoET (03)
		{CollegeID: collegeMap["03"], Code: "ME", Name: "Mechanical Engineering", Head: "Dr. Isack Kamonde"},
		{CollegeID: collegeMap["03"], Code: "CE", Name: "Civil Engineering", Head: "Dr. Japhet Kashaigili"},
		{CollegeID: collegeMap["03"], Code: "CHE", Name: "Chemical and Process Engineering", Head: "Dr. Yusufu Abeid Chande Jande"},

		// CoES (04)
		{CollegeID: collegeMap["04"], Code: "GEO", Name: "Geology", Head: "Dr. Emmanuel Mutakyahwa"},
		{CollegeID: collegeMap["04"], Code: "MIN", Name: "Mining and Mineral Processing Engineering", Head: "Dr. Simon Makundi"},

		// CoVMBS (05)
		{CollegeID: collegeMap["05"], Code: "VET", Name: "Veterinary Medicine", Head: "Dr. Emmanuel Mellau"},
		{CollegeID: collegeMap["05"], Code: "BMS", Name: "Biomedical Sciences", Head: "Dr. Fred Mfinanga"},

		// CoLSB (06)
		{CollegeID: collegeMap["06"], Code: "BIO", Name: "Biology", Head: "Dr. Ezekiel Mmbaga"},
		{CollegeID: collegeMap["06"], Code: "BCH", Name: "Biochemistry and Molecular Biology", Head: "Dr. Sylvester Leonard Lyantagaye"},
		{CollegeID: collegeMap["06"], Code: "BE", Name: "Bioengineering", Head: "Dr. Victor Mkupasi"},

		// CoSSH (07)
		{CollegeID: collegeMap["07"], Code: "ENG", Name: "English Language and Literature", Head: "Dr. Magreth Bushesha"},
		{CollegeID: collegeMap["07"], Code: "SOC", Name: "Sociology and Anthropology", Head: "Dr. Felix Kumamoto"},
	}

	for _, dept := range departments {
		if err := s.db.FirstOrCreate(&dept, models.Department{Code: dept.Code}).Error; err != nil {
			return err
		}
	}

	return nil
}

// SeedPrograms seeds academic programs
func (s *Seeder) SeedPrograms() error {
	// Get departments
	var departments []models.Department
	s.db.Find(&departments)

	// Map department codes to IDs
	deptMap := make(map[string]uint)
	for _, d := range departments {
		deptMap[d.Code] = d.ID
	}

	programs := []models.Program{
		// Computer Science and Engineering
		{DepartmentID: deptMap["CS"], Code: "MB011", Name: "Bachelor of Science in Computer Science and Engineering", DegreeLevel: "Bachelor", NTALevel: 8, Duration: 3, TuitionFees: 1500000},
		{DepartmentID: deptMap["CS"], Code: "MM007", Name: "Master of Science in Computer Science", DegreeLevel: "Masters", NTALevel: 9, Duration: 2, TuitionFees: 2500000},

		// ICT
		{DepartmentID: deptMap["ICT"], Code: "MB006", Name: "Bachelor of Science in Information and Communication Technology", DegreeLevel: "Bachelor", NTALevel: 8, Duration: 3, TuitionFees: 1500000},
		{DepartmentID: deptMap["ICT"], Code: "MD005", Name: "Diploma in Information Technology", DegreeLevel: "Diploma", NTALevel: 6, Duration: 2, TuitionFees: 1000000},
		{DepartmentID: deptMap["ICT"], Code: "MD010", Name: "Diploma in Computer Engineering", DegreeLevel: "Diploma", NTALevel: 6, Duration: 3, TuitionFees: 1200000},

		// Electronics and Telecommunications
		{DepartmentID: deptMap["EE"], Code: "MB012", Name: "Bachelor of Science in Electronics and Telecommunications Engineering", DegreeLevel: "Bachelor", NTALevel: 8, Duration: 3, TuitionFees: 1600000},

		// Business Administration and Finance
		{DepartmentID: deptMap["BAF"], Code: "MB015", Name: "Bachelor of Business Administration", DegreeLevel: "Bachelor", NTALevel: 8, Duration: 3, TuitionFees: 1400000},
		{DepartmentID: deptMap["BAF"], Code: "MB016", Name: "Bachelor of Science in Finance", DegreeLevel: "Bachelor", NTALevel: 8, Duration: 3, TuitionFees: 1400000},

		// Accounting
		{DepartmentID: deptMap["ACC"], Code: "MB017", Name: "Bachelor of Science in Accounting", DegreeLevel: "Bachelor", NTALevel: 8, Duration: 3, TuitionFees: 1400000},

		// Mechanical Engineering
		{DepartmentID: deptMap["ME"], Code: "MB020", Name: "Bachelor of Science in Mechanical Engineering", DegreeLevel: "Bachelor", NTALevel: 8, Duration: 3, TuitionFees: 1600000},

		// Civil Engineering
		{DepartmentID: deptMap["CE"], Code: "MB021", Name: "Bachelor of Science in Civil Engineering", DegreeLevel: "Bachelor", NTALevel: 8, Duration: 3, TuitionFees: 1600000},

		// Geology
		{DepartmentID: deptMap["GEO"], Code: "MB025", Name: "Bachelor of Science in Geology", DegreeLevel: "Bachelor", NTALevel: 8, Duration: 3, TuitionFees: 1500000},

		// Veterinary Medicine
		{DepartmentID: deptMap["VET"], Code: "MB030", Name: "Doctor of Veterinary Medicine", DegreeLevel: "Bachelor", NTALevel: 8, Duration: 5, TuitionFees: 2000000},

		// Biology
		{DepartmentID: deptMap["BIO"], Code: "MB035", Name: "Bachelor of Science in Biology", DegreeLevel: "Bachelor", NTALevel: 8, Duration: 3, TuitionFees: 1400000},

		// English Language and Literature
		{DepartmentID: deptMap["ENG"], Code: "MB040", Name: "Bachelor of Arts in English Language and Literature", DegreeLevel: "Bachelor", NTALevel: 8, Duration: 3, TuitionFees: 1300000},
	}

	for _, program := range programs {
		if err := s.db.FirstOrCreate(&program, models.Program{Code: program.Code}).Error; err != nil {
			return err
		}
	}

	return nil
}

// SeedSemesters creates current and past semesters
func (s *Seeder) SeedSemesters() error {
	semesters := []models.Semester{
		{
			Name:         "2023/2024 - Semester I",
			AcademicYear: "2023/2024",
			SemesterNum:  1,
			StartDate:    time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
			EndDate:      time.Date(2024, 2, 28, 0, 0, 0, 0, time.UTC),
			IsCurrent:    false,
		},
		{
			Name:         "2023/2024 - Semester II",
			AcademicYear: "2023/2024",
			SemesterNum:  2,
			StartDate:    time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC),
			EndDate:      time.Date(2024, 8, 31, 0, 0, 0, 0, time.UTC),
			IsCurrent:    false,
		},
		{
			Name:         "2024/2025 - Semester I",
			AcademicYear: "2024/2025",
			SemesterNum:  1,
			StartDate:    time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC),
			EndDate:      time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC),
			IsCurrent:    false,
		},
		{
			Name:         "2024/2025 - Semester II",
			AcademicYear: "2024/2025",
			SemesterNum:  2,
			StartDate:    time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC),
			EndDate:      time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC),
			IsCurrent:    true,
		},
	}

	for _, semester := range semesters {
		if err := s.db.FirstOrCreate(&semester, models.Semester{Name: semester.Name}).Error; err != nil {
			return err
		}
	}

	return nil
}

// SeedVenues creates lecture venues
func (s *Seeder) SeedVenues() error {
	venues := []models.Venue{
		{Building: "CoICT Building", RoomNumber: "101", Capacity: 50, VenueType: "Classroom"},
		{Building: "CoICT Building", RoomNumber: "102", Capacity: 100, VenueType: "Lecture Hall"},
		{Building: "CoICT Building", RoomNumber: "Lab A", Capacity: 30, VenueType: "Lab"},
		{Building: "Main Building", RoomNumber: "201", Capacity: 80, VenueType: "Classroom"},
		{Building: "Main Building", RoomNumber: "Auditorium", Capacity: 300, VenueType: "Lecture Hall"},
		{Building: "Engineering Block", RoomNumber: "E101", Capacity: 60, VenueType: "Classroom"},
		{Building: "Engineering Block", RoomNumber: "Workshop 1", Capacity: 40, VenueType: "Lab"},
	}

	for _, venue := range venues {
		if err := s.db.FirstOrCreate(&venue, models.Venue{
			Building:   venue.Building,
			RoomNumber: venue.RoomNumber,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}

// SeedOAuthClient creates LMS OAuth client
func (s *Seeder) SeedOAuthClient() error {
	// Hash the client secret
	hashedSecret, err := utils.HashPassword("lms-client-secret-change-in-production")
	if err != nil {
		return err
	}

	client := models.OAuthClient{
		ClientID:     "lms-client-id",
		ClientSecret: hashedSecret,
		Name:         "MUST Learning Management System",
		RedirectURIs: "http://localhost:8080/auth/callback,http://192.168.1.20:8080/auth/callback",
		Scopes:       "read_profile,read_courses,read_grades,write_grades",
		IsActive:     true,
	}

	return s.db.FirstOrCreate(&client, models.OAuthClient{ClientID: client.ClientID}).Error
}

// Continue in next part...
