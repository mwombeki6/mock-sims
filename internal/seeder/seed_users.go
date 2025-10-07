package seeder

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/mwombeki6/mock-sims/internal/models"
	"github.com/mwombeki6/mock-sims/internal/utils"
)

var seededRand = rand.New(rand.NewSource(20241002))

// SeedUsers creates rich mock users spanning admins, faculty and large student cohorts.
func (s *Seeder) SeedUsers() error {
	hashedPassword, err := utils.HashPassword("password123")
	if err != nil {
		return err
	}

	if err := s.seedAdminUser(hashedPassword); err != nil {
		return err
	}

	deptMap, err := s.departmentMap()
	if err != nil {
		return err
	}

	if err := s.seedFacultyUsers(coreFacultySeeds, deptMap, hashedPassword); err != nil {
		return err
	}

	if err := s.seedFacultyUsers(additionalFacultySeeds, deptMap, hashedPassword); err != nil {
		return err
	}

	programMap, err := s.programMap()
	if err != nil {
		return err
	}

	if err := s.seedStudentCohorts(programMap, hashedPassword); err != nil {
		return err
	}

	return nil
}

func (s *Seeder) seedAdminUser(hashedPassword string) error {
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
	return s.db.FirstOrCreate(&admin, models.Admin{UserID: admin.UserID}).Error
}

func (s *Seeder) departmentMap() (map[string]models.Department, error) {
	var departments []models.Department
	if err := s.db.Preload("College").Find(&departments).Error; err != nil {
		return nil, err
	}

	deptMap := make(map[string]models.Department)
	for _, dept := range departments {
		deptMap[strings.ToUpper(dept.Code)] = dept
	}
	return deptMap, nil
}

func (s *Seeder) programMap() (map[string]models.Program, error) {
	var programs []models.Program
	if err := s.db.Preload("Department").Preload("Department.College").Find(&programs).Error; err != nil {
		return nil, err
	}

	programMap := make(map[string]models.Program)
	for _, program := range programs {
		programMap[strings.ToUpper(program.Code)] = program
	}
	return programMap, nil
}

func (s *Seeder) seedFacultyUsers(seeds []facultySeed, deptMap map[string]models.Department, hashedPassword string) error {
	for _, fs := range seeds {
		dept, ok := deptMap[strings.ToUpper(fs.DepartmentCode)]
		if !ok {
			continue
		}

		user := models.User{
			Email:    fs.Email,
			Password: hashedPassword,
			UserType: "faculty",
			IsActive: true,
		}
		if err := s.db.FirstOrCreate(&user, models.User{Email: user.Email}).Error; err != nil {
			return err
		}

		faculty := models.Faculty{
			UserID:         user.ID,
			StaffID:        fs.StaffID,
			FirstName:      fs.FirstName,
			MiddleName:     fs.MiddleName,
			LastName:       fs.LastName,
			DepartmentID:   dept.ID,
			Rank:           fs.Rank,
			Specialization: fs.Specialization,
		}
		if err := s.db.FirstOrCreate(&faculty, models.Faculty{UserID: faculty.UserID}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *Seeder) seedStudentCohorts(programMap map[string]models.Program, hashedPassword string) error {
	enrollmentStatuses := []string{"active", "probation", "suspended"}
	paymentStatuses := []string{"paid", "partial", "pending"}

	for programCode, cohorts := range programCohorts {
		program, ok := programMap[strings.ToUpper(programCode)]
		if !ok {
			continue
		}

		collegeCode, _ := strconv.Atoi(program.Department.College.Code)
		baseSequence := int(program.ID) * 10000

		studentIndex := 0
		for _, cohort := range cohorts {
			for i := 0; i < cohort.StudentCount; i++ {
				firstName := pickName(studentFirstNames, studentIndex+i)
				middleName := pickName(studentMiddleNames, studentIndex+i/2)
				lastName := pickName(studentLastNames, studentIndex+i/3)

				sequence := baseSequence + ((cohort.AdmissionYear % 100) * 1000) + i + 1
				regNumber := utils.GenerateRegistrationNumber(cohort.AdmissionYear, collegeCode, sequence)
				email := fmt.Sprintf("%s.%s%d@must.ac.tz",
					strings.ToLower(firstName),
					strings.ToLower(lastName),
					sequence%1000,
				)

				user := models.User{
					Email:    email,
					Password: hashedPassword,
					UserType: "student",
					IsActive: true,
				}
				if err := s.db.FirstOrCreate(&user, models.User{Email: email}).Error; err != nil {
					return err
				}

				student := models.Student{
					UserID:           user.ID,
					RegNumber:        regNumber,
					FirstName:        firstName,
					MiddleName:       middleName,
					LastName:         lastName,
					ProgramID:        program.ID,
					YearOfStudy:      cohort.YearOfStudy,
					GPA:              randomFloat(2.5, 4.9),
					EnrollmentStatus: enrollmentStatuses[sequence%len(enrollmentStatuses)],
					PaymentStatus:    paymentStatuses[(sequence/3)%len(paymentStatuses)],
					AdmissionYear:    cohort.AdmissionYear,
				}

				if err := s.db.FirstOrCreate(&student, models.Student{RegNumber: student.RegNumber}).Error; err != nil {
					return err
				}
			}
			studentIndex += cohort.StudentCount
		}
	}

	return nil
}

func pickName(list []string, index int) string {
	if len(list) == 0 {
		return ""
	}
	return list[index%len(list)]
}

func randomFloat(min, max float64) float64 {
	if max <= min {
		return min
	}
	return min + seededRand.Float64()*(max-min)
}

func init() {
	seededRand.Seed(time.Date(2024, 10, 2, 0, 0, 0, 0, time.UTC).UnixNano())
}
