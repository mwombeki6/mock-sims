package seeder

// Predefined data sets used by the seeder to generate rich mock data.

var studentFirstNames = []string{
	"Asha", "Erick", "Neema", "Juma", "Maria", "Kelvin", "Gloria", "Samuel", "Linda", "Abel",
	"Diana", "Peter", "Joyce", "Elias", "Rebecca", "Baraka", "Agnes", "Felix", "Zainab", "Daniel",
	"Salma", "Sofia", "Dennis", "Frank", "Mariam", "Tereza", "Catherine", "John", "Prisca", "Godfrey",
	"Happy", "George", "Magreth", "Furaha", "Bruno", "Imani", "Ester", "Laurian", "Esther", "Richard",
}

var studentMiddleNames = []string{
	"Ally", "Mwombek", "Joseph", "Grace", "Moses", "Innocent", "Jackson", "Lighton", "Matata", "Issack",
	"Gabriel", "Cuthbert", "Salim", "Vicent", "Ananilea", "Venance", "Paul", "Ibrahim", "Saidi", "Hassan",
	"Daniel", "Timothy", "Amos", "Prosper", "Steven", "Lucas", "James", "Raphael", "Filbert", "Cyril",
}

var studentLastNames = []string{
	"Lubere", "Mushi", "Mwakabungu", "Ngwale", "Shayo", "Mmbaga", "Katambala", "Mhando", "Lusajo", "Mkude",
	"Bujiku", "Kaaya", "Mgaya", "Nyoni", "Kavishe", "Mchome", "Hosiana", "Mligo", "Mrema", "Msele",
	"Mwaikendo", "Mrema", "Komba", "Kibona", "Mshindo", "Samson", "Sangoma", "Mdegela", "Mhagama", "Mwakalinga",
	"Msuya", "Mbilinyi", "Kanyeka", "Msigwa", "Chacha", "Mwangosi", "Mtui", "Maziku", "Mwansasu", "Nyangasa",
}

type facultySeed struct {
	Email          string
	FirstName      string
	MiddleName     string
	LastName       string
	StaffID        string
	DepartmentCode string
	Rank           string
	Specialization string
}

var coreFacultySeeds = []facultySeed{
	{"joseph.mkunda@must.ac.tz", "Joseph", "T", "Mkunda", "MUST-F-001", "CS", "Professor", "Computer Networks"},
	{"devotha.nyambo@must.ac.tz", "Devotha", "G", "Nyambo", "MUST-F-002", "CS", "Senior Lecturer", "Software Engineering"},
	{"mussa.dida@must.ac.tz", "Mussa", "Ally", "Dida", "MUST-F-003", "CS", "Lecturer", "Artificial Intelligence"},
}

var additionalFacultySeeds = []facultySeed{
	{"janeth.ndossi@must.ac.tz", "Janeth", "P", "Ndossi", "MUST-F-010", "CS", "Senior Lecturer", "Information Systems"},
	{"alex.katende@must.ac.tz", "Alex", "G", "Katende", "MUST-F-011", "CS", "Lecturer", "Data Communication"},
	{"rehema.mdee@must.ac.tz", "Rehema", "J", "Mdee", "MUST-F-012", "ICT", "Assistant Lecturer", "Database Systems"},
	{"geofrey.nyusi@must.ac.tz", "Geofrey", "K", "Nyusi", "MUST-F-013", "ICT", "Lecturer", "ICT Governance"},
	{"hellen.komba@must.ac.tz", "Hellen", "S", "Komba", "MUST-F-014", "EE", "Senior Lecturer", "Telecommunications"},
	{"patrick.mshomba@must.ac.tz", "Patrick", "C", "Mshomba", "MUST-F-015", "EE", "Lecturer", "Embedded Systems"},
	{"flavian.kweka@must.ac.tz", "Flavian", "M", "Kweka", "MUST-F-016", "BAF", "Lecturer", "Corporate Finance"},
	{"mercy.shayo@must.ac.tz", "Mercy", "I", "Shayo", "MUST-F-017", "ACC", "Lecturer", "Taxation"},
	{"john.muro@must.ac.tz", "John", "L", "Muro", "MUST-F-018", "ME", "Senior Lecturer", "Thermodynamics"},
	{"hulda.mollel@must.ac.tz", "Hulda", "A", "Mollel", "MUST-F-019", "CE", "Lecturer", "Structural Engineering"},
	{"arnold.sanga@must.ac.tz", "Arnold", "N", "Sanga", "MUST-F-020", "CHE", "Lecturer", "Process Control"},
	{"sabina.mwita@must.ac.tz", "Sabina", "K", "Mwita", "MUST-F-021", "VET", "Lecturer", "Clinical Studies"},
	{"augustine.msigwa@must.ac.tz", "Augustine", "D", "Msigwa", "MUST-F-022", "BIO", "Senior Lecturer", "Biotechnology"},
	{"hassan.madiba@must.ac.tz", "Hassan", "F", "Madiba", "MUST-F-023", "ENG", "Lecturer", "Literature"},
}

type programCohort struct {
	AdmissionYear int
	YearOfStudy   int
	StudentCount  int
	EntryCategory string
}

var programCohorts = map[string][]programCohort{
	"MB011": {
		{AdmissionYear: 2022, YearOfStudy: 3, StudentCount: 45, EntryCategory: "DIRECT"},
		{AdmissionYear: 2023, YearOfStudy: 2, StudentCount: 60, EntryCategory: "DIRECT"},
		{AdmissionYear: 2024, YearOfStudy: 1, StudentCount: 65, EntryCategory: "DIRECT"},
	},
	"MB006": {
		{AdmissionYear: 2023, YearOfStudy: 2, StudentCount: 55, EntryCategory: "DIRECT"},
		{AdmissionYear: 2024, YearOfStudy: 1, StudentCount: 60, EntryCategory: "DIRECT"},
	},
	"MD010": {
		{AdmissionYear: 2023, YearOfStudy: 2, StudentCount: 70, EntryCategory: "DIRECT"},
		{AdmissionYear: 2024, YearOfStudy: 1, StudentCount: 75, EntryCategory: "DIRECT"},
	},
	"MB020": {
		{AdmissionYear: 2022, YearOfStudy: 3, StudentCount: 40, EntryCategory: "DIRECT"},
		{AdmissionYear: 2023, YearOfStudy: 2, StudentCount: 45, EntryCategory: "DIRECT"},
	},
	"MB021": {
		{AdmissionYear: 2023, YearOfStudy: 2, StudentCount: 40, EntryCategory: "DIRECT"},
		{AdmissionYear: 2024, YearOfStudy: 1, StudentCount: 50, EntryCategory: "DIRECT"},
	},
	"MB015": {
		{AdmissionYear: 2023, YearOfStudy: 2, StudentCount: 50, EntryCategory: "DIRECT"},
		{AdmissionYear: 2024, YearOfStudy: 1, StudentCount: 60, EntryCategory: "DIRECT"},
	},
	"MB040": {
		{AdmissionYear: 2023, YearOfStudy: 2, StudentCount: 25, EntryCategory: "DIRECT"},
		{AdmissionYear: 2024, YearOfStudy: 1, StudentCount: 30, EntryCategory: "DIRECT"},
	},
}

type courseSeed struct {
	Code           string
	Name           string
	Credits        int
	Level          int
	DepartmentCode string
	Description    string
	ProgramCodes   []string
}

var courseSeeds = []courseSeed{
	{"CS6201", "System Software", 9, 200, "CS", "Operating system concepts and system level programming.", []string{"MD010"}},
	{"CS6202", "Object Oriented Programming", 7, 200, "CS", "Advanced object oriented principles using Java.", []string{"MB011", "MD010"}},
	{"CS6203", "Data Communication", 6, 200, "CS", "Computer networks and secure data transmission.", []string{"MB011", "MD010"}},
	{"CS6204", "Advanced Electronics", 6, 200, "EE", "Analog electronics and amplification techniques.", []string{"MD010"}},
	{"CS6205", "Digital Electronics", 6, 200, "EE", "Digital logic, combinational and sequential circuits.", []string{"MD010"}},
	{"CS6206", "System Analysis and Design", 6, 200, "ICT", "Structured and object oriented system analysis methods.", []string{"MB011", "MB006", "MD010"}},
	{"CS6207", "Database System Design and Management", 6, 200, "ICT", "Relational database theory and administration.", []string{"MB011", "MB006", "MD010"}},
	{"CS6208", "Introduction to Software Engineering", 6, 200, "CS", "Software lifecycle, processes and documentation.", []string{"MB011", "MB006"}},
	{"CS6209", "Field Practical Training II", 10, 200, "ICT", "Industrial attachment focusing on ICT operations.", []string{"MD010"}},
	{"CS6210", "Data Structure and File Handling", 9, 200, "CS", "Algorithms, data structures and file systems.", []string{"MD010"}},
	{"CS6211", "Basics of Telecommunication", 6, 200, "EE", "Telecommunication technologies and standards.", []string{"MD010"}},
	{"CS6212", "Website Design and Hosting", 9, 200, "ICT", "Web technologies, hosting and security.", []string{"MD010"}},
	{"CS6213", "Sequential Circuits", 6, 200, "EE", "Design of sequential logic circuits.", []string{"MD010"}},
	{"CS6214", "Microprocessor Technology", 6, 200, "EE", "Microprocessor architecture and interfacing.", []string{"MB011", "MD010"}},
	{"CS6215", "Information Management", 6, 200, "ICT", "Information systems for organisations.", []string{"MB006", "MB011"}},
	{"CS6216", "Introduction to Embedded Systems", 6, 200, "EE", "Embedded systems design and programming.", []string{"MB011", "MD010"}},
	{"MS6227", "Discrete Mathematics and Complex Numbers", 6, 200, "ICT", "Discrete math concepts and complex analysis.", []string{"MB011", "MB006", "MD010"}},
	{"CS6103", "Programming Fundamentals", 6, 100, "ICT", "Foundations of programming using Python.", []string{"MB011", "MB006", "MD010"}},
	{"CS6104", "Computer Systems", 6, 100, "CS", "Introduction to computer architecture.", []string{"MB011", "MD010"}},
	{"CS6105", "Mathematics for Computing", 7, 100, "ICT", "Calculus and algebra for computing.", []string{"MB011", "MB006"}},
	{"CS6106", "Professional Communication", 5, 100, "ENG", "Communication skills for technical professionals.", []string{"MB011", "MB006", "MD010", "MB015"}},
	{"CS6301", "Distributed Systems", 6, 300, "CS", "Distributed architectures and microservices.", []string{"MB011"}},
	{"CS6302", "Cloud Infrastructure", 6, 300, "ICT", "Cloud platforms, deployment and operations.", []string{"MB011"}},
	{"CS6303", "Machine Learning", 6, 300, "CS", "Machine learning algorithms and applications.", []string{"MB011"}},
	{"BA6101", "Principles of Accounting", 6, 100, "ACC", "Financial accounting fundamentals.", []string{"MB015", "MB017"}},
	{"BA6102", "Business Mathematics", 6, 100, "BAF", "Mathematics for business decision making.", []string{"MB015", "MB016"}},
	{"BA6204", "Corporate Finance", 6, 200, "BAF", "Working capital management and capital budgeting.", []string{"MB015"}},
	{"BA6205", "Investment Analysis", 6, 200, "BAF", "Portfolio theory and security valuation.", []string{"MB015"}},
	{"BA6306", "International Business", 6, 300, "BAF", "Global trade environment and strategies.", []string{"MB015"}},
	{"ACC6201", "Intermediate Accounting", 7, 200, "ACC", "Financial reporting and standards.", []string{"MB017"}},
	{"ACC6202", "Taxation Principles", 6, 200, "ACC", "Income tax computations and planning.", []string{"MB017"}},
	{"ME6102", "Engineering Mechanics", 6, 100, "ME", "Statics and dynamics for engineers.", []string{"MB020"}},
	{"ME6204", "Thermodynamics", 6, 200, "ME", "Thermodynamic systems and cycles.", []string{"MB020"}},
	{"ME6306", "Machine Design", 6, 300, "ME", "Design and analysis of mechanical systems.", []string{"MB020"}},
	{"CE6101", "Engineering Drawing", 6, 100, "CE", "Technical drawing for civil engineers.", []string{"MB021"}},
	{"CE6203", "Structural Analysis", 6, 200, "CE", "Analysis of structures and loads.", []string{"MB021"}},
	{"CE6304", "Transportation Engineering", 6, 300, "CE", "Planning and design of transportation systems.", []string{"MB021"}},
	{"BIO6201", "Molecular Biology", 6, 200, "BIO", "Cellular and molecular processes.", []string{"MB035"}},
	{"BIO6303", "Bioprocess Engineering", 6, 300, "BIO", "Bioprocess design and optimisation.", []string{"MB035"}},
	{"ENG6105", "Academic Writing", 6, 100, "ENG", "Academic writing and research skills.", []string{"MB040"}},
	{"ENG6202", "African Literature", 6, 200, "ENG", "Study of African literary works.", []string{"MB040"}},
	{"ENG6304", "Language and Society", 6, 300, "ENG", "Sociolinguistics and discourse analysis.", []string{"MB040"}},
}

var lectureStartSlots = []string{"08:00", "10:00", "14:00", "16:00"}
var lectureEndSlots = []string{"10:00", "12:00", "16:00", "18:00"}
var lectureDays = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

var paymentDescriptions = []string{
	"Tuition Fee", "Administrative Fee", "Accommodation Fee", "Meals Fee", "Library Fee", "ICT Service Fee",
}

var paymentMethods = []string{"Bank", "M-Pesa", "Tigo Pesa", "Airtel Money"}
var paymentStatuses = []string{"completed", "partial", "pending"}
