package seeder

import (
	"strings"
	"time"

	"github.com/mwombeki6/mock-sims/internal/models"
)

// SeedCourses persists catalogue courses and attaches them to their programs.
func (s *Seeder) SeedCourses() error {
	deptMap, err := s.departmentMap()
	if err != nil {
		return err
	}
	programMap, err := s.programMap()
	if err != nil {
		return err
	}

	for _, seed := range courseSeeds {
		dept, ok := deptMap[strings.ToUpper(seed.DepartmentCode)]
		if !ok {
			continue
		}

		course := models.Course{Code: seed.Code}
		if err := s.db.FirstOrCreate(&course, models.Course{Code: seed.Code}).Error; err != nil {
			return err
		}

		course.Name = seed.Name
		course.Credits = seed.Credits
		course.Level = seed.Level
		course.Description = seed.Description
		course.DepartmentID = dept.ID
		if err := s.db.Save(&course).Error; err != nil {
			return err
		}

		if len(seed.ProgramCodes) > 0 {
			assoc := s.db.Model(&course).Association("Programs")
			if err := assoc.Clear(); err != nil {
				return err
			}
			for _, programCode := range seed.ProgramCodes {
				program, ok := programMap[strings.ToUpper(programCode)]
				if !ok {
					continue
				}
				p := program
				_ = assoc.Append(&p)
			}
		}
	}

	return nil
}

// SeedLectures generates lecture schedules for the current semester using seeded faculties and venues.
func (s *Seeder) SeedLectures() error {
	var currentSemester models.Semester
	if err := s.db.Where("is_current = ?", true).First(&currentSemester).Error; err != nil {
		return nil
	}

	var courses []models.Course
	if err := s.db.Find(&courses).Error; err != nil {
		return err
	}

	var faculties []models.Faculty
	if err := s.db.Find(&faculties).Error; err != nil {
		return err
	}

	var venues []models.Venue
	if err := s.db.Find(&venues).Error; err != nil {
		return err
	}

	if len(courses) == 0 || len(faculties) == 0 || len(venues) == 0 {
		return nil
	}

	facultyByDepartment := make(map[uint][]models.Faculty)
	for _, faculty := range faculties {
		facultyByDepartment[faculty.DepartmentID] = append(facultyByDepartment[faculty.DepartmentID], faculty)
	}

	for idx, course := range courses {
		facultyPool := facultyByDepartment[course.DepartmentID]
		if len(facultyPool) == 0 {
			facultyPool = faculties
		}
		faculty := facultyPool[idx%len(facultyPool)]
		venue := venues[idx%len(venues)]
		day := lectureDays[idx%len(lectureDays)]
		start := lectureStartSlots[idx%len(lectureStartSlots)]
		end := lectureEndSlots[idx%len(lectureEndSlots)]

		assignment := models.CourseAssignment{
			CourseID:   course.ID,
			FacultyID:  faculty.ID,
			SemesterID: currentSemester.ID,
			Role:       "Lecturer",
		}
		if err := s.db.FirstOrCreate(&assignment, models.CourseAssignment{
			CourseID:   course.ID,
			FacultyID:  faculty.ID,
			SemesterID: currentSemester.ID,
		}).Error; err != nil {
			return err
		}

		lecture := models.Lecture{
			CourseID:   course.ID,
			FacultyID:  faculty.ID,
			SemesterID: currentSemester.ID,
			VenueID:    venue.ID,
			DayOfWeek:  day,
			StartTime:  start,
			EndTime:    end,
		}
		if err := s.db.FirstOrCreate(&lecture, models.Lecture{
			CourseID:   course.ID,
			SemesterID: currentSemester.ID,
			DayOfWeek:  day,
			StartTime:  start,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}

// SeedEnrollments registers students into program courses for current and previous semesters.
func (s *Seeder) SeedEnrollments() error {
	var currentSemester models.Semester
	if err := s.db.Where("is_current = ?", true).First(&currentSemester).Error; err != nil {
		return nil
	}

	var previousSemester models.Semester
	if err := s.db.Where("id <> ?", currentSemester.ID).Order("start_date DESC").First(&previousSemester).Error; err != nil {
		// If there is no previous semester we only seed the current one
		previousSemester = models.Semester{}
	}

	var students []models.Student
	if err := s.db.Find(&students).Error; err != nil {
		return err
	}

	var courses []models.Course
	if err := s.db.Preload("Programs").Find(&courses).Error; err != nil {
		return err
	}

	programBuckets := make(map[uint]map[int][]models.Course)
	for _, course := range courses {
		for _, program := range course.Programs {
			bucket := programBuckets[program.ID]
			if bucket == nil {
				bucket = make(map[int][]models.Course)
			}
			bucket[course.Level] = append(bucket[course.Level], course)
			programBuckets[program.ID] = bucket
		}
	}

	for _, student := range students {
		bucket := programBuckets[student.ProgramID]
		if len(bucket) == 0 {
			continue
		}

		desiredLevel := student.YearOfStudy * 100
		if desiredLevel < 100 {
			desiredLevel = 100
		}

		currentCourses := selectCoursesForLevel(bucket, desiredLevel, 6)
		previousCourses := []models.Course{}
		if previousSemester.ID != 0 {
			previousCourses = selectCoursesForLevel(bucket, desiredLevel-100, 6)
		}

		for _, course := range currentCourses {
			enrollment := models.Enrollment{
				StudentID:  student.ID,
				CourseID:   course.ID,
				SemesterID: currentSemester.ID,
				Status:     "active",
				EnrolledAt: currentSemester.StartDate.Add(-24 * time.Hour * 7),
			}
			if err := s.db.FirstOrCreate(&enrollment, models.Enrollment{
				StudentID:  student.ID,
				CourseID:   course.ID,
				SemesterID: currentSemester.ID,
			}).Error; err != nil {
				return err
			}
		}

		for _, course := range previousCourses {
			if previousSemester.ID == 0 {
				break
			}
			enrollment := models.Enrollment{
				StudentID:  student.ID,
				CourseID:   course.ID,
				SemesterID: previousSemester.ID,
				Status:     "completed",
				EnrolledAt: previousSemester.StartDate.Add(-24 * time.Hour * 7),
			}
			if err := s.db.FirstOrCreate(&enrollment, models.Enrollment{
				StudentID:  student.ID,
				CourseID:   course.ID,
				SemesterID: previousSemester.ID,
			}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func selectCoursesForLevel(bucket map[int][]models.Course, desiredLevel int, maxCourses int) []models.Course {
	if maxCourses <= 0 {
		return nil
	}

	requestedLevels := []int{desiredLevel, desiredLevel - 100, desiredLevel + 100, 300, 200, 100}
	courseSet := make([]models.Course, 0, maxCourses)
	seen := make(map[uint]struct{})

	for _, level := range requestedLevels {
		if level <= 0 {
			continue
		}
		courses := bucket[level]
		for _, course := range courses {
			if _, exists := seen[course.ID]; exists {
				continue
			}
			courseSet = append(courseSet, course)
			seen[course.ID] = struct{}{}
			if len(courseSet) >= maxCourses {
				return courseSet
			}
		}
	}

	return courseSet
}
