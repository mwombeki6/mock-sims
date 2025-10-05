package seeder

import (
	"time"

	"github.com/mwombeki6/mock-sims/internal/models"
)

// SeedCourses creates sample courses
func (s *Seeder) SeedCourses() error {
	// Get first department (CS)
	var departments []models.Department
	s.db.Limit(3).Find(&departments)

	if len(departments) == 0 {
		return nil
	}

	courses := []models.Course{
		// Year 1 Courses
		{
			Code:         "CS 1101",
			Name:         "Introduction to Computer Science",
			Credits:      4,
			Level:        100,
			Description:  "Fundamentals of computer science including algorithms, data structures, and programming basics.",
			DepartmentID: departments[0].ID,
		},
		{
			Code:         "CS 1102",
			Name:         "Programming with Python",
			Credits:      4,
			Level:        100,
			Description:  "Introduction to programming using Python language.",
			DepartmentID: departments[0].ID,
		},
		{
			Code:         "MS 1101",
			Name:         "Calculus I",
			Credits:      3,
			Level:        100,
			Description:  "Differential and integral calculus.",
			DepartmentID: departments[0].ID,
		},
		{
			Code:         "MS 1102",
			Name:         "Linear Algebra",
			Credits:      3,
			Level:        100,
			Description:  "Vectors, matrices, and linear transformations.",
			DepartmentID: departments[0].ID,
		},

		// Year 2 Courses
		{
			Code:         "CS 2201",
			Name:         "Data Structures and Algorithms",
			Credits:      4,
			Level:        200,
			Description:  "Advanced data structures, algorithm design and analysis.",
			DepartmentID: departments[0].ID,
		},
		{
			Code:         "CS 2202",
			Name:         "Object-Oriented Programming",
			Credits:      4,
			Level:        200,
			Description:  "Object-oriented programming concepts using Java.",
			DepartmentID: departments[0].ID,
		},
		{
			Code:         "CS 2203",
			Name:         "Database Systems",
			Credits:      4,
			Level:        200,
			Description:  "Relational database design, SQL, normalization.",
			DepartmentID: departments[0].ID,
		},
		{
			Code:         "CS 2204",
			Name:         "Computer Networks",
			Credits:      3,
			Level:        200,
			Description:  "Network protocols, TCP/IP, network security.",
			DepartmentID: departments[0].ID,
		},
		{
			Code:         "MS 2201",
			Name:         "Discrete Mathematics",
			Credits:      3,
			Level:        200,
			Description:  "Logic, sets, graph theory, combinatorics.",
			DepartmentID: departments[0].ID,
		},

		// Year 3 Courses
		{
			Code:         "CS 3301",
			Name:         "Software Engineering",
			Credits:      4,
			Level:        300,
			Description:  "Software development methodologies, design patterns, testing.",
			DepartmentID: departments[0].ID,
		},
		{
			Code:         "CS 3302",
			Name:         "Web Technologies",
			Credits:      4,
			Level:        300,
			Description:  "HTML, CSS, JavaScript, web frameworks.",
			DepartmentID: departments[0].ID,
		},
		{
			Code:         "CS 3303",
			Name:         "Artificial Intelligence",
			Credits:      4,
			Level:        300,
			Description:  "Search algorithms, machine learning, neural networks.",
			DepartmentID: departments[0].ID,
		},
		{
			Code:         "CS 3304",
			Name:         "Operating Systems",
			Credits:      4,
			Level:        300,
			Description:  "Process management, memory management, file systems.",
			DepartmentID: departments[0].ID,
		},
	}

	for _, course := range courses {
		if err := s.db.FirstOrCreate(&course, models.Course{Code: course.Code}).Error; err != nil {
			return err
		}
	}

	return nil
}

// SeedLectures creates lecture schedules
func (s *Seeder) SeedLectures() error {
	// Get current semester
	var semester models.Semester
	if err := s.db.Where("is_current = ?", true).First(&semester).Error; err != nil {
		return nil
	}

	// Get some courses
	var courses []models.Course
	s.db.Limit(6).Find(&courses)

	// Get faculty members
	var faculty []models.Faculty
	s.db.Limit(3).Find(&faculty)

	// Get venues
	var venues []models.Venue
	s.db.Limit(5).Find(&venues)

	if len(courses) == 0 || len(faculty) == 0 || len(venues) == 0 {
		return nil
	}

	// Create course assignments first
	assignments := []models.CourseAssignment{
		{CourseID: courses[0].ID, FacultyID: faculty[0].ID, SemesterID: semester.ID, Role: "Lecturer"},
		{CourseID: courses[1].ID, FacultyID: faculty[1].ID, SemesterID: semester.ID, Role: "Lecturer"},
		{CourseID: courses[2].ID, FacultyID: faculty[2].ID, SemesterID: semester.ID, Role: "Lecturer"},
		{CourseID: courses[3].ID, FacultyID: faculty[0].ID, SemesterID: semester.ID, Role: "Lecturer"},
		{CourseID: courses[4].ID, FacultyID: faculty[1].ID, SemesterID: semester.ID, Role: "Lecturer"},
		{CourseID: courses[5].ID, FacultyID: faculty[2].ID, SemesterID: semester.ID, Role: "Lecturer"},
	}

	for _, assignment := range assignments {
		if err := s.db.FirstOrCreate(&assignment, models.CourseAssignment{
			CourseID:   assignment.CourseID,
			FacultyID:  assignment.FacultyID,
			SemesterID: assignment.SemesterID,
		}).Error; err != nil {
			return err
		}
	}

	// Create lecture schedules
	lectures := []models.Lecture{
		// Monday
		{CourseID: courses[0].ID, FacultyID: faculty[0].ID, SemesterID: semester.ID, VenueID: venues[0].ID, DayOfWeek: "Monday", StartTime: "08:00", EndTime: "10:00"},
		{CourseID: courses[1].ID, FacultyID: faculty[1].ID, SemesterID: semester.ID, VenueID: venues[1].ID, DayOfWeek: "Monday", StartTime: "10:00", EndTime: "12:00"},

		// Tuesday
		{CourseID: courses[2].ID, FacultyID: faculty[2].ID, SemesterID: semester.ID, VenueID: venues[0].ID, DayOfWeek: "Tuesday", StartTime: "08:00", EndTime: "10:00"},
		{CourseID: courses[3].ID, FacultyID: faculty[0].ID, SemesterID: semester.ID, VenueID: venues[2].ID, DayOfWeek: "Tuesday", StartTime: "14:00", EndTime: "16:00"},

		// Wednesday
		{CourseID: courses[4].ID, FacultyID: faculty[1].ID, SemesterID: semester.ID, VenueID: venues[1].ID, DayOfWeek: "Wednesday", StartTime: "08:00", EndTime: "10:00"},
		{CourseID: courses[5].ID, FacultyID: faculty[2].ID, SemesterID: semester.ID, VenueID: venues[0].ID, DayOfWeek: "Wednesday", StartTime: "10:00", EndTime: "12:00"},

		// Thursday
		{CourseID: courses[0].ID, FacultyID: faculty[0].ID, SemesterID: semester.ID, VenueID: venues[3].ID, DayOfWeek: "Thursday", StartTime: "08:00", EndTime: "10:00"},

		// Friday
		{CourseID: courses[1].ID, FacultyID: faculty[1].ID, SemesterID: semester.ID, VenueID: venues[0].ID, DayOfWeek: "Friday", StartTime: "10:00", EndTime: "12:00"},
	}

	for _, lecture := range lectures {
		if err := s.db.FirstOrCreate(&lecture, models.Lecture{
			CourseID:   lecture.CourseID,
			SemesterID: lecture.SemesterID,
			DayOfWeek:  lecture.DayOfWeek,
			StartTime:  lecture.StartTime,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}

// SeedEnrollments creates student course enrollments
func (s *Seeder) SeedEnrollments() error {
	// Get current semester
	var semester models.Semester
	if err := s.db.Where("is_current = ?", true).First(&semester).Error; err != nil {
		return nil
	}

	// Get students
	var students []models.Student
	s.db.Limit(5).Find(&students)

	// Get courses (first 6 for testing)
	var courses []models.Course
	s.db.Limit(6).Find(&courses)

	if len(students) == 0 || len(courses) == 0 {
		return nil
	}

	// Enroll each student in 4-5 courses
	for _, student := range students {
		coursesToEnroll := courses[:4] // First 4 courses for each student

		for _, course := range coursesToEnroll {
			enrollment := models.Enrollment{
				StudentID:  student.ID,
				CourseID:   course.ID,
				SemesterID: semester.ID,
				Status:     "active",
				EnrolledAt: time.Now().Add(-30 * 24 * time.Hour), // 30 days ago
			}

			if err := s.db.FirstOrCreate(&enrollment, models.Enrollment{
				StudentID:  enrollment.StudentID,
				CourseID:   enrollment.CourseID,
				SemesterID: enrollment.SemesterID,
			}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
