package models

import (
	"time"

	"gorm.io/gorm"
)

// ============================================================================
// ORGANIZATIONAL STRUCTURE
// ============================================================================

// College represents a faculty/college at MUST
type College struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Code      string         `gorm:"uniqueIndex;size:10;not null" json:"code"` // 01, 02, 03...
	Name      string         `gorm:"size:200;not null" json:"name"`
	ShortName string         `gorm:"size:20" json:"short_name"` // CoICT, CoET, etc.
	Dean      string         `gorm:"size:100" json:"dean"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Departments []Department `gorm:"foreignKey:CollegeID" json:"departments,omitempty"`
}

// Department represents a department within a college
type Department struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CollegeID uint           `gorm:"not null;index" json:"college_id"`
	Code      string         `gorm:"uniqueIndex;size:20;not null" json:"code"`
	Name      string         `gorm:"size:200;not null" json:"name"`
	Head      string         `gorm:"size:100" json:"head"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	College  College   `gorm:"foreignKey:CollegeID" json:"college,omitempty"`
	Programs []Program `gorm:"foreignKey:DepartmentID" json:"programs,omitempty"`
}

// Program represents an academic program (e.g., BSc Computer Science)
type Program struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	DepartmentID uint           `gorm:"not null;index" json:"department_id"`
	Code         string         `gorm:"uniqueIndex;size:20;not null" json:"code"` // MB011, MB006, etc.
	Name         string         `gorm:"size:200;not null" json:"name"`
	DegreeLevel  string         `gorm:"size:50;not null" json:"degree_level"` // Certificate, Diploma, Bachelor, Masters, PhD
	NTALevel     int            `gorm:"not null" json:"nta_level"`            // 5, 6, 8, 9, 10
	Duration     int            `gorm:"not null" json:"duration"`             // Duration in years
	TuitionFees  int            `gorm:"not null" json:"tuition_fees"`         // Annual fees in TZS
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	Department Department  `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	Students   []Student   `gorm:"foreignKey:ProgramID" json:"students,omitempty"`
	Courses    []Course    `gorm:"many2many:program_courses" json:"courses,omitempty"`
}

// Semester represents an academic semester
type Semester struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"size:50;not null" json:"name"`           // 2024/2025 Semester I
	AcademicYear  string         `gorm:"size:20;not null" json:"academic_year"`  // 2024/2025
	SemesterNum   int            `gorm:"not null" json:"semester_num"`           // 1 or 2
	StartDate     time.Time      `gorm:"not null" json:"start_date"`
	EndDate       time.Time      `gorm:"not null" json:"end_date"`
	IsCurrent     bool           `gorm:"default:false" json:"is_current"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	Enrollments []Enrollment `gorm:"foreignKey:SemesterID" json:"enrollments,omitempty"`
}

// Venue represents a classroom or lecture hall
type Venue struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Building  string         `gorm:"size:100;not null" json:"building"`
	RoomNumber string        `gorm:"size:20;not null" json:"room_number"`
	Capacity  int            `gorm:"not null" json:"capacity"`
	VenueType string         `gorm:"size:50" json:"venue_type"` // Classroom, Lab, Lecture Hall
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Lectures []Lecture `gorm:"foreignKey:VenueID" json:"lectures,omitempty"`
}

// ============================================================================
// PEOPLE (USERS)
// ============================================================================

// User is the base user model (polymorphic)
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"-"` // Bcrypt hashed
	UserType  string         `gorm:"size:20;not null;index" json:"user_type"` // student, faculty, admin
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Student represents a student
type Student struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	UserID          uint           `gorm:"uniqueIndex;not null" json:"user_id"`
	RegNumber       string         `gorm:"uniqueIndex;size:50;not null" json:"reg_number"` // 2024/01/04567
	FirstName       string         `gorm:"size:100;not null" json:"first_name"`
	MiddleName      string         `gorm:"size:100" json:"middle_name"`
	LastName        string         `gorm:"size:100;not null" json:"last_name"`
	ProgramID       uint           `gorm:"not null;index" json:"program_id"`
	YearOfStudy     int            `gorm:"not null" json:"year_of_study"` // 1, 2, 3, 4
	GPA             float64        `gorm:"type:decimal(3,2)" json:"gpa"`
	EnrollmentStatus string        `gorm:"size:20;default:'active'" json:"enrollment_status"` // active, suspended, graduated
	PaymentStatus   string         `gorm:"size:20;default:'pending'" json:"payment_status"`   // paid, pending, partial
	AdmissionYear   int            `gorm:"not null" json:"admission_year"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	User        User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Program     Program      `gorm:"foreignKey:ProgramID" json:"program,omitempty"`
	Enrollments []Enrollment `gorm:"foreignKey:StudentID" json:"enrollments,omitempty"`
	Grades      []Grade      `gorm:"foreignKey:StudentID" json:"grades,omitempty"`
}

// Faculty represents a lecturer/professor
type Faculty struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	UserID         uint           `gorm:"uniqueIndex;not null" json:"user_id"`
	StaffID        string         `gorm:"uniqueIndex;size:50;not null" json:"staff_id"`
	FirstName      string         `gorm:"size:100;not null" json:"first_name"`
	MiddleName     string         `gorm:"size:100" json:"middle_name"`
	LastName       string         `gorm:"size:100;not null" json:"last_name"`
	DepartmentID   uint           `gorm:"not null;index" json:"department_id"`
	Rank           string         `gorm:"size:50" json:"rank"` // Professor, Senior Lecturer, Lecturer, Assistant Lecturer
	Specialization string         `gorm:"size:200" json:"specialization"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	User               User               `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Department         Department         `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	CourseAssignments  []CourseAssignment `gorm:"foreignKey:FacultyID" json:"course_assignments,omitempty"`
}

// Admin represents system administrators
type Admin struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"uniqueIndex;not null" json:"user_id"`
	FirstName string         `gorm:"size:100;not null" json:"first_name"`
	LastName  string         `gorm:"size:100;not null" json:"last_name"`
	Role      string         `gorm:"size:50" json:"role"` // Registrar, Dean, Director, etc.
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// ============================================================================
// ACADEMIC ENTITIES
// ============================================================================

// Course represents a course/subject
type Course struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Code        string         `gorm:"uniqueIndex;size:20;not null" json:"code"` // CSC301, ENG201
	Name        string         `gorm:"size:200;not null" json:"name"`
	Credits     int            `gorm:"not null" json:"credits"` // 3, 4, etc.
	Level       int            `gorm:"not null" json:"level"`   // 100, 200, 300, 400
	Description string         `gorm:"type:text" json:"description"`
	DepartmentID uint          `gorm:"not null;index" json:"department_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	Department        Department         `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	Lectures          []Lecture          `gorm:"foreignKey:CourseID" json:"lectures,omitempty"`
	Enrollments       []Enrollment       `gorm:"foreignKey:CourseID" json:"enrollments,omitempty"`
	CourseAssignments []CourseAssignment `gorm:"foreignKey:CourseID" json:"course_assignments,omitempty"`
	Programs          []Program          `gorm:"many2many:program_courses" json:"programs,omitempty"`
}

// Lecture represents a scheduled class/lecture
type Lecture struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CourseID   uint           `gorm:"not null;index" json:"course_id"`
	FacultyID  uint           `gorm:"not null;index" json:"faculty_id"`
	SemesterID uint           `gorm:"not null;index" json:"semester_id"`
	VenueID    uint           `gorm:"index" json:"venue_id"`
	DayOfWeek  string         `gorm:"size:20;not null" json:"day_of_week"` // Monday, Tuesday, etc.
	StartTime  string         `gorm:"size:10;not null" json:"start_time"`  // 10:00
	EndTime    string         `gorm:"size:10;not null" json:"end_time"`    // 12:00
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	Course   Course   `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	Faculty  Faculty  `gorm:"foreignKey:FacultyID" json:"faculty,omitempty"`
	Semester Semester `gorm:"foreignKey:SemesterID" json:"semester,omitempty"`
	Venue    Venue    `gorm:"foreignKey:VenueID" json:"venue,omitempty"`
}

// Enrollment represents student enrollment in a course
type Enrollment struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	StudentID  uint           `gorm:"not null;index" json:"student_id"`
	CourseID   uint           `gorm:"not null;index" json:"course_id"`
	SemesterID uint           `gorm:"not null;index" json:"semester_id"`
	Status     string         `gorm:"size:20;default:'active'" json:"status"` // active, dropped, completed
	EnrolledAt time.Time      `gorm:"not null" json:"enrolled_at"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	Student  Student  `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	Course   Course   `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	Semester Semester `gorm:"foreignKey:SemesterID" json:"semester,omitempty"`
	Grade    *Grade   `gorm:"foreignKey:EnrollmentID" json:"grade,omitempty"`
}

// Grade represents a student's grade for a course
type Grade struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	EnrollmentID uint           `gorm:"uniqueIndex;not null" json:"enrollment_id"`
	StudentID    uint           `gorm:"not null;index" json:"student_id"`
	CourseID     uint           `gorm:"not null;index" json:"course_id"`
	CAMarks      float64        `gorm:"type:decimal(5,2)" json:"ca_marks"`         // Continuous Assessment (0-40)
	FinalExam    float64        `gorm:"type:decimal(5,2)" json:"final_exam"`       // Final Exam (0-60)
	TotalMarks   float64        `gorm:"type:decimal(5,2)" json:"total_marks"`      // Total (0-100)
	LetterGrade  string         `gorm:"size:5" json:"letter_grade"`                // A, B+, B, C, D, F
	GradePoint   float64        `gorm:"type:decimal(3,2)" json:"grade_point"`      // 5.0, 4.0, 3.5, etc.
	Remarks      string         `gorm:"type:text" json:"remarks"`
	SubmittedAt  *time.Time     `json:"submitted_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	Enrollment Enrollment `gorm:"foreignKey:EnrollmentID" json:"enrollment,omitempty"`
	Student    Student    `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	Course     Course     `gorm:"foreignKey:CourseID" json:"course,omitempty"`
}

// CourseAssignment represents faculty assignment to teach a course
type CourseAssignment struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CourseID   uint           `gorm:"not null;index" json:"course_id"`
	FacultyID  uint           `gorm:"not null;index" json:"faculty_id"`
	SemesterID uint           `gorm:"not null;index" json:"semester_id"`
	Role       string         `gorm:"size:50" json:"role"` // Lecturer, Teaching Assistant
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	Course   Course   `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	Faculty  Faculty  `gorm:"foreignKey:FacultyID" json:"faculty,omitempty"`
	Semester Semester `gorm:"foreignKey:SemesterID" json:"semester,omitempty"`
}

// ============================================================================
// OAUTH 2.0
// ============================================================================

// OAuthClient represents an OAuth client application (LMS)
type OAuthClient struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	ClientID     string         `gorm:"uniqueIndex;size:100;not null" json:"client_id"`
	ClientSecret string         `gorm:"size:255;not null" json:"-"`
	Name         string         `gorm:"size:100;not null" json:"name"`
	RedirectURIs string         `gorm:"type:text;not null" json:"redirect_uris"` // Comma-separated
	Scopes       string         `gorm:"type:text" json:"scopes"`                 // Comma-separated
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// OAuthAuthorizationCode represents a temporary authorization code
type OAuthAuthorizationCode struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Code        string         `gorm:"uniqueIndex;size:255;not null" json:"code"`
	ClientID    string         `gorm:"size:100;not null;index" json:"client_id"`
	UserID      uint           `gorm:"not null;index" json:"user_id"`
	RedirectURI string         `gorm:"size:500;not null" json:"redirect_uri"`
	Scopes      string         `gorm:"type:text" json:"scopes"`
	ExpiresAt   time.Time      `gorm:"not null;index" json:"expires_at"`
	Used        bool           `gorm:"default:false" json:"used"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// OAuthAccessToken represents an access token
type OAuthAccessToken struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Token        string         `gorm:"uniqueIndex;size:500;not null" json:"token"`
	ClientID     string         `gorm:"size:100;not null;index" json:"client_id"`
	UserID       uint           `gorm:"not null;index" json:"user_id"`
	Scopes       string         `gorm:"type:text" json:"scopes"`
	ExpiresAt    time.Time      `gorm:"not null;index" json:"expires_at"`
	RefreshToken string         `gorm:"size:500" json:"refresh_token"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// ============================================================================
// WEBHOOKS & PAYMENTS
// ============================================================================

// WebhookLog represents a webhook delivery attempt
type WebhookLog struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Event      string         `gorm:"size:100;not null;index" json:"event"` // enrollment.created, grade.submitted, etc.
	URL        string         `gorm:"size:500;not null" json:"url"`
	StatusCode int            `gorm:"index" json:"status_code"`
	Error      *string        `gorm:"type:text" json:"error"`
	SentAt     time.Time      `gorm:"not null;index" json:"sent_at"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

// Payment represents a student's tuition payment
type Payment struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	StudentID     uint           `gorm:"not null;index" json:"student_id"`
	InvoiceNumber string         `gorm:"uniqueIndex;size:50;not null" json:"invoice_number"` // MB1010231181
	Amount        float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	PaymentMethod string         `gorm:"size:50;not null" json:"payment_method"` // Bank, M-Pesa, Cash
	PaymentDate   time.Time      `gorm:"not null;index" json:"payment_date"`
	AcademicYear  string         `gorm:"size:20;not null" json:"academic_year"` // 2024/2025
	Semester      string         `gorm:"size:20;not null" json:"semester"`      // Semester I
	Status        string         `gorm:"size:20;default:'completed'" json:"status"` // completed, pending, failed
	ReceiptNumber string         `gorm:"size:50" json:"receipt_number"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	Student Student `gorm:"foreignKey:StudentID" json:"student,omitempty"`
}
