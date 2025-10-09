package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type DocsHandler struct{}

func NewDocsHandler() *DocsHandler {
	return &DocsHandler{}
}

// ServeSwaggerUI serves the Swagger UI with dark mode
func (h *DocsHandler) ServeSwaggerUI(c *fiber.Ctx) error {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Mock SIMS API - Swagger UI</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.10.5/swagger-ui.css">
    <style>
        body {
            margin: 0;
            background: #1a1a1a;
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
        }
        .swagger-ui {
            max-width: 1400px;
            margin: 0 auto;
            padding: 20px;
        }
        .swagger-ui .topbar {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            padding: 20px;
        }
        .swagger-ui .topbar-wrapper::before {
            content: "🎓 Mock SIMS API";
            color: #fff;
            font-size: 24px;
            font-weight: 700;
        }
        .swagger-ui .topbar-wrapper img { display: none; }
        .swagger-ui .info { background: #2d2d2d; color: #e0e0e0; padding: 30px; border-radius: 8px; }
        .swagger-ui .info .title { color: #667eea; }
        .swagger-ui .opblock-tag { color: #fff; background: #2d2d2d; border-bottom: 2px solid #667eea; }
        .swagger-ui .opblock { background: #2d2d2d; border-color: #3d3d3d; }
        .swagger-ui .opblock.opblock-get .opblock-summary { background: #61affe; }
        .swagger-ui .opblock.opblock-post .opblock-summary { background: #49cc90; }
        .swagger-ui .parameters-col_description input { background: #1a1a1a; color: #e0e0e0; border: 1px solid #3d3d3d; }
        .swagger-ui .btn { background: #667eea; color: #fff; }
        .swagger-ui .btn.execute { background: #49cc90; }
        .swagger-ui .response-col_status { color: #49cc90; }
        .swagger-ui select, .swagger-ui input[type=text] { background: #1a1a1a; color: #e0e0e0; border-color: #3d3d3d; }
    </style>
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5.10.5/swagger-ui-bundle.js"></script>
    <script>
        window.onload = function() {
            SwaggerUIBundle({
                url: "/swagger.json",
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIBundle.SwaggerUIStandalonePreset
                ],
                layout: "BaseLayout"
            });
        };
    </script>
</body>
</html>`
	return c.Type("html").SendString(html)
}

// ServeReDoc serves the ReDoc documentation with dark mode
func (h *DocsHandler) ServeReDoc(c *fiber.Ctx) error {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Mock SIMS API - Documentation</title>
    <style>
        body { margin: 0; padding: 0; }
    </style>
</head>
<body>
    <redoc spec-url='/swagger.json'
    theme='{
        "colors": {
            "primary": { "main": "#667eea" },
            "success": { "main": "#49cc90" },
            "text": { "primary": "#e0e0e0" },
            "http": {
                "get": "#61affe",
                "post": "#49cc90",
                "put": "#fca130",
                "delete": "#f93e3e"
            }
        },
        "typography": {
            "fontSize": "15px",
            "fontFamily": "-apple-system, BlinkMacSystemFont, Segoe UI, sans-serif",
            "headings": { "fontWeight": "700" },
            "code": { "fontSize": "14px", "color": "#667eea" }
        },
        "sidebar": { "backgroundColor": "#1a1a1a", "textColor": "#e0e0e0" },
        "rightPanel": { "backgroundColor": "#0d0d0d", "textColor": "#e0e0e0" }
    }'></redoc>
    <script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"></script>
</body>
</html>`
	return c.Type("html").SendString(html)
}

// ServeSwaggerJSON serves the OpenAPI 3.0.3 specification
func (h *DocsHandler) ServeSwaggerJSON(c *fiber.Ctx) error {
	spec := map[string]interface{}{
		"openapi": "3.0.3",
		"info": map[string]interface{}{
			"title":       "Mock SIMS API",
			"description": "**MUST Student Information Management System**\n\nOAuth 2.0 Server and REST API for LMS Integration Testing\n\n## Features\n- OAuth 2.0 Authorization Code Flow\n- Student Management APIs\n- Faculty Management APIs\n- Course Catalog\n- Administrative Endpoints",
			"version":     "1.0.0",
			"contact": map[string]string{
				"name":  "API Support",
				"url":   "https://github.com/mwombeki6/mock-sims",
				"email": "support@must.ac.tz",
			},
			"license": map[string]string{
				"name": "MIT",
				"url":  "https://opensource.org/licenses/MIT",
			},
		},
		"servers": []map[string]string{
			{
				"url":         "http://localhost:8000",
				"description": "Development Server",
			},
			{
				"url":         "http://192.168.1.10:8000",
				"description": "4-PC Deployment (Mock SIMS)",
			},
		},
		"tags": []map[string]string{
			{"name": "OAuth", "description": "OAuth 2.0 Authorization Server"},
			{"name": "Health", "description": "Service health check"},
			{"name": "Students", "description": "Student profile and academic information"},
			{"name": "Faculty", "description": "Faculty profile and teaching assignments"},
			{"name": "Courses", "description": "Course catalog and management"},
			{"name": "Admin", "description": "Administrative endpoints (colleges, departments, programs)"},
		},
		"paths":      h.getPaths(),
		"components": h.getComponents(),
	}
	return c.JSON(spec)
}

// getPaths returns all API paths
func (h *DocsHandler) getPaths() map[string]interface{} {
	return map[string]interface{}{
		"/health": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Health"},
				"summary":     "Health check",
				"description": "Returns service health status",
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Service is healthy",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/HealthResponse",
								},
							},
						},
					},
				},
			},
		},
		"/oauth/authorize": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"OAuth"},
				"summary":     "OAuth authorization endpoint",
				"description": "Displays login page for OAuth authorization",
				"parameters": []map[string]interface{}{
					{
						"name":        "client_id",
						"in":          "query",
						"required":    true,
						"description": "OAuth client ID (e.g., lms-client-id)",
						"schema":      map[string]string{"type": "string"},
						"example":     "lms-client-id",
					},
					{
						"name":        "redirect_uri",
						"in":          "query",
						"required":    true,
						"description": "Redirect URI after authorization",
						"schema":      map[string]string{"type": "string"},
						"example":     "http://localhost:8080/auth/callback",
					},
					{
						"name":        "response_type",
						"in":          "query",
						"required":    true,
						"description": "Must be 'code'",
						"schema":      map[string]string{"type": "string", "enum": "code"},
						"example":     "code",
					},
					{
						"name":        "state",
						"in":          "query",
						"required":    false,
						"description": "Optional state parameter",
						"schema":      map[string]string{"type": "string"},
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Login page HTML",
						"content": map[string]interface{}{
							"text/html": map[string]interface{}{
								"schema": map[string]string{"type": "string"},
							},
						},
					},
					"400": map[string]interface{}{
						"description": "Invalid OAuth parameters",
					},
				},
			},
			"post": map[string]interface{}{
				"tags":        []string{"OAuth"},
				"summary":     "Login form submission",
				"description": "Processes login credentials and returns authorization code",
				"parameters": []map[string]interface{}{
					{
						"name":     "client_id",
						"in":       "query",
						"required": true,
						"schema":   map[string]string{"type": "string"},
					},
					{
						"name":     "redirect_uri",
						"in":       "query",
						"required": true,
						"schema":   map[string]string{"type": "string"},
					},
					{
						"name":     "response_type",
						"in":       "query",
						"required": true,
						"schema":   map[string]string{"type": "string"},
					},
					{
						"name":     "state",
						"in":       "query",
						"required": false,
						"schema":   map[string]string{"type": "string"},
					},
				},
				"requestBody": map[string]interface{}{
					"required": true,
					"content": map[string]interface{}{
						"application/x-www-form-urlencoded": map[string]interface{}{
							"schema": map[string]interface{}{
								"type": "object",
								"properties": map[string]interface{}{
									"username": map[string]interface{}{
										"type":        "string",
										"description": "Registration number or email",
										"example":     "john.doe@must.ac.tz",
									},
									"password": map[string]interface{}{
										"type":        "string",
										"description": "User password",
										"example":     "password123",
									},
								},
								"required": []string{"username", "password"},
							},
						},
					},
				},
				"responses": map[string]interface{}{
					"302": map[string]interface{}{
						"description": "Redirect to redirect_uri with authorization code",
					},
					"400": map[string]interface{}{
						"description": "Invalid credentials",
					},
				},
			},
		},
		"/oauth/token": map[string]interface{}{
			"post": map[string]interface{}{
				"tags":        []string{"OAuth"},
				"summary":     "Exchange authorization code for access token",
				"description": "Exchanges authorization code or refresh token for access token",
				"requestBody": map[string]interface{}{
					"required": true,
					"content": map[string]interface{}{
						"application/x-www-form-urlencoded": map[string]interface{}{
							"schema": map[string]interface{}{
								"oneOf": []map[string]interface{}{
									{
										"type": "object",
										"properties": map[string]interface{}{
											"grant_type":    map[string]interface{}{"type": "string", "enum": []string{"authorization_code"}},
											"code":          map[string]interface{}{"type": "string", "description": "Authorization code"},
											"client_id":     map[string]interface{}{"type": "string", "example": "lms-client-id"},
											"client_secret": map[string]interface{}{"type": "string", "example": "lms-client-secret-change-in-production"},
											"redirect_uri":  map[string]interface{}{"type": "string", "example": "http://localhost:8080/auth/callback"},
										},
										"required": []string{"grant_type", "code", "client_id", "client_secret"},
									},
									{
										"type": "object",
										"properties": map[string]interface{}{
											"grant_type":    map[string]interface{}{"type": "string", "enum": []string{"refresh_token"}},
											"refresh_token": map[string]interface{}{"type": "string", "description": "Refresh token"},
										},
										"required": []string{"grant_type", "refresh_token"},
									},
								},
							},
						},
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Access token issued",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/TokenResponse",
								},
							},
						},
					},
					"400": map[string]interface{}{
						"description": "Invalid grant",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/ErrorResponse",
								},
							},
						},
					},
				},
			},
		},
		"/api/students/me": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Students"},
				"summary":     "Get authenticated student profile",
				"description": "Returns the profile of the currently authenticated student",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Student profile",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/StudentProfile",
								},
							},
						},
					},
					"401": map[string]interface{}{
						"description": "Unauthorized",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/ErrorResponse",
								},
							},
						},
					},
				},
			},
		},
		"/api/students/{id}/courses": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Students"},
				"summary":     "Get student's enrolled courses",
				"description": "Returns all courses a student is enrolled in",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"parameters": []map[string]interface{}{
					{
						"name":        "id",
						"in":          "path",
						"required":    true,
						"description": "Student ID",
						"schema":      map[string]string{"type": "integer"},
						"example":     1,
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Student courses",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/StudentCoursesResponse",
								},
							},
						},
					},
					"400": map[string]interface{}{
						"description": "Invalid student ID",
					},
				},
			},
		},
		"/api/students/{id}/grades": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Students"},
				"summary":     "Get student's grades",
				"description": "Returns all grades for a student",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"parameters": []map[string]interface{}{
					{
						"name":        "id",
						"in":          "path",
						"required":    true,
						"description": "Student ID",
						"schema":      map[string]string{"type": "integer"},
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Student grades",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/StudentGradesResponse",
								},
							},
						},
					},
				},
			},
		},
		"/api/students/{id}/timetable": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Students"},
				"summary":     "Get student's timetable",
				"description": "Returns class schedule for a student",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"parameters": []map[string]interface{}{
					{
						"name":        "id",
						"in":          "path",
						"required":    true,
						"description": "Student ID",
						"schema":      map[string]string{"type": "integer"},
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Student timetable",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/StudentTimetableResponse",
								},
							},
						},
					},
				},
			},
		},
		"/api/faculty/me": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Faculty"},
				"summary":     "Get authenticated faculty profile",
				"description": "Returns the profile of the currently authenticated faculty member",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Faculty profile",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/FacultyProfile",
								},
							},
						},
					},
					"401": map[string]interface{}{
						"description": "Unauthorized",
					},
				},
			},
		},
		"/api/faculty/{id}/courses": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Faculty"},
				"summary":     "Get faculty's teaching assignments",
				"description": "Returns all courses assigned to a faculty member",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"parameters": []map[string]interface{}{
					{
						"name":        "id",
						"in":          "path",
						"required":    true,
						"description": "Faculty ID",
						"schema":      map[string]string{"type": "integer"},
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Faculty courses",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/FacultyCoursesResponse",
								},
							},
						},
					},
				},
			},
		},
		"/api/faculty/courses/{id}/ca-marks": map[string]interface{}{
			"post": map[string]interface{}{
				"tags":        []string{"Faculty"},
				"summary":     "Submit CA marks",
				"description": "Submit Continuous Assessment marks for students",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"parameters": []map[string]interface{}{
					{
						"name":        "id",
						"in":          "path",
						"required":    true,
						"description": "Course ID",
						"schema":      map[string]string{"type": "integer"},
					},
				},
				"requestBody": map[string]interface{}{
					"required": true,
					"content": map[string]interface{}{
						"application/json": map[string]interface{}{
							"schema": map[string]interface{}{
								"$ref": "#/components/schemas/CAMarksRequest",
							},
						},
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "CA marks submitted successfully",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"message":     map[string]string{"type": "string", "example": "CA marks submitted successfully"},
										"course_id":   map[string]string{"type": "integer"},
										"marks_count": map[string]string{"type": "integer"},
									},
								},
							},
						},
					},
				},
			},
		},
		"/api/courses": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Courses"},
				"summary":     "List all courses",
				"description": "Returns paginated list of courses",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"parameters": []map[string]interface{}{
					{
						"name":        "page",
						"in":          "query",
						"description": "Page number (default: 1)",
						"schema":      map[string]interface{}{"type": "integer", "default": 1},
					},
					{
						"name":        "limit",
						"in":          "query",
						"description": "Items per page (default: 20, max: 100)",
						"schema":      map[string]interface{}{"type": "integer", "default": 20},
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Course list",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/CoursesListResponse",
								},
							},
						},
					},
				},
			},
		},
		"/api/courses/{code}": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Courses"},
				"summary":     "Get course details",
				"description": "Returns detailed information about a specific course",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"parameters": []map[string]interface{}{
					{
						"name":        "code",
						"in":          "path",
						"required":    true,
						"description": "Course code (e.g., CS 1101)",
						"schema":      map[string]string{"type": "string"},
						"example":     "CS 1101",
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Course details",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/CourseDetails",
								},
							},
						},
					},
					"404": map[string]interface{}{
						"description": "Course not found",
					},
				},
			},
		},
		"/api/courses/{code}/lectures": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Courses"},
				"summary":     "Get course lectures/schedule",
				"description": "Returns lecture schedule for a course",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"parameters": []map[string]interface{}{
					{
						"name":        "code",
						"in":          "path",
						"required":    true,
						"description": "Course code",
						"schema":      map[string]string{"type": "string"},
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Course lectures",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/CourseLecturesResponse",
								},
							},
						},
					},
				},
			},
		},
		"/api/courses/{code}/students": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Courses"},
				"summary":     "Get enrolled students",
				"description": "Returns list of students enrolled in a course",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"parameters": []map[string]interface{}{
					{
						"name":        "code",
						"in":          "path",
						"required":    true,
						"description": "Course code",
						"schema":      map[string]string{"type": "string"},
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Enrolled students",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/CourseStudentsResponse",
								},
							},
						},
					},
				},
			},
		},
		"/api/colleges": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Admin"},
				"summary":     "Get all colleges",
				"description": "Returns list of all MUST colleges",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Colleges list",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/CollegesResponse",
								},
							},
						},
					},
				},
			},
		},
		"/api/departments": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Admin"},
				"summary":     "Get departments",
				"description": "Returns list of departments (filterable by college)",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"parameters": []map[string]interface{}{
					{
						"name":        "college_id",
						"in":          "query",
						"description": "Filter by college ID",
						"schema":      map[string]string{"type": "integer"},
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Departments list",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/DepartmentsResponse",
								},
							},
						},
					},
				},
			},
		},
		"/api/programs": map[string]interface{}{
			"get": map[string]interface{}{
				"tags":        []string{"Admin"},
				"summary":     "Get programs",
				"description": "Returns list of academic programs (filterable by department or level)",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"parameters": []map[string]interface{}{
					{
						"name":        "department_id",
						"in":          "query",
						"description": "Filter by department ID",
						"schema":      map[string]string{"type": "integer"},
					},
					{
						"name":        "level",
						"in":          "query",
						"description": "Filter by degree level",
						"schema":      map[string]interface{}{"type": "string", "enum": []string{"Certificate", "Diploma", "Bachelor", "Masters", "PhD"}},
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Programs list",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/ProgramsResponse",
								},
							},
						},
					},
				},
			},
		},
		"/api/enrollments": map[string]interface{}{
			"post": map[string]interface{}{
				"tags":        []string{"Admin"},
				"summary":     "Create bulk enrollments",
				"description": "Creates multiple student enrollments",
				"security":    []map[string][]string{{"BearerAuth": {}}},
				"requestBody": map[string]interface{}{
					"required": true,
					"content": map[string]interface{}{
						"application/json": map[string]interface{}{
							"schema": map[string]interface{}{
								"$ref": "#/components/schemas/BulkEnrollmentRequest",
							},
						},
					},
				},
				"responses": map[string]interface{}{
					"200": map[string]interface{}{
						"description": "Enrollments created successfully",
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"message": map[string]string{"type": "string", "example": "enrollments created successfully"},
										"count":   map[string]string{"type": "integer"},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// getComponents returns schemas and security definitions
func (h *DocsHandler) getComponents() map[string]interface{} {
	return map[string]interface{}{
		"securitySchemes": map[string]interface{}{
			"BearerAuth": map[string]string{
				"type":         "http",
				"scheme":       "bearer",
				"bearerFormat": "JWT",
				"description":  "Enter your Bearer token (obtained from /oauth/token)",
			},
		},
		"schemas": map[string]interface{}{
			"HealthResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"status":  map[string]string{"type": "string", "example": "ok"},
					"service": map[string]string{"type": "string", "example": "mock-sims"},
					"version": map[string]string{"type": "string", "example": "1.0.0"},
				},
			},
			"TokenResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"access_token":  map[string]string{"type": "string", "example": "eyJhbGci..."},
					"token_type":    map[string]string{"type": "string", "example": "Bearer"},
					"expires_in":    map[string]interface{}{"type": "integer", "example": 3600},
					"refresh_token": map[string]string{"type": "string", "example": "eyJhbGci..."},
					"scope":         map[string]string{"type": "string", "example": "student.read courses.read"},
				},
			},
			"ErrorResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"error":             map[string]string{"type": "string", "example": "unauthorized"},
					"error_description": map[string]string{"type": "string", "example": "Invalid credentials"},
				},
			},
			"StudentProfile": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"student_id":        map[string]interface{}{"type": "integer", "example": 1},
					"reg_number":        map[string]string{"type": "string", "example": "23100523050032"},
					"name":              map[string]string{"type": "string", "example": "John Doe Mwamba"},
					"email":             map[string]string{"type": "string", "example": "john.doe@must.ac.tz"},
					"year_of_study":     map[string]interface{}{"type": "integer", "example": 2},
					"gpa":               map[string]interface{}{"type": "number", "example": 3.45},
					"enrollment_status": map[string]string{"type": "string", "example": "active"},
					"payment_status":    map[string]string{"type": "string", "example": "paid"},
				},
			},
			"StudentCoursesResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"student_id": map[string]interface{}{"type": "integer"},
					"courses": map[string]interface{}{
						"type": "array",
						"items": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"course_code": map[string]string{"type": "string", "example": "CS 1101"},
								"course_name": map[string]string{"type": "string"},
								"credits":     map[string]interface{}{"type": "integer"},
							},
						},
					},
					"total": map[string]interface{}{"type": "integer"},
				},
			},
			"StudentGradesResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"student_id": map[string]interface{}{"type": "integer"},
					"grades": map[string]interface{}{
						"type": "array",
						"items": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"course_code":  map[string]string{"type": "string"},
								"ca_marks":     map[string]interface{}{"type": "number"},
								"final_exam":   map[string]interface{}{"type": "number"},
								"total_marks":  map[string]interface{}{"type": "number"},
								"letter_grade": map[string]string{"type": "string"},
							},
						},
					},
					"total": map[string]interface{}{"type": "integer"},
				},
			},
			"StudentTimetableResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"student_id": map[string]interface{}{"type": "integer"},
					"timetable": map[string]interface{}{
						"type": "array",
						"items": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"course_code": map[string]string{"type": "string"},
								"day":         map[string]string{"type": "string"},
								"start_time":  map[string]string{"type": "string"},
								"end_time":    map[string]string{"type": "string"},
								"venue":       map[string]string{"type": "string"},
							},
						},
					},
				},
			},
			"FacultyProfile": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"faculty_id":     map[string]interface{}{"type": "integer"},
					"staff_id":       map[string]string{"type": "string"},
					"name":           map[string]string{"type": "string"},
					"email":          map[string]string{"type": "string"},
					"rank":           map[string]string{"type": "string"},
					"specialization": map[string]string{"type": "string"},
				},
			},
			"FacultyCoursesResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"faculty_id": map[string]interface{}{"type": "integer"},
					"courses": map[string]interface{}{
						"type":  "array",
						"items": map[string]string{"type": "object"},
					},
					"total": map[string]interface{}{"type": "integer"},
				},
			},
			"CAMarksRequest": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"marks": map[string]interface{}{
						"type": "array",
						"items": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"student_id": map[string]interface{}{"type": "integer"},
								"ca_marks":   map[string]interface{}{"type": "number"},
							},
						},
					},
				},
			},
			"CoursesListResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"courses": map[string]interface{}{
						"type":  "array",
						"items": map[string]string{"type": "object"},
					},
					"total": map[string]interface{}{"type": "integer"},
					"page":  map[string]interface{}{"type": "integer"},
					"limit": map[string]interface{}{"type": "integer"},
				},
			},
			"CourseDetails": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"code":        map[string]string{"type": "string"},
					"name":        map[string]string{"type": "string"},
					"credits":     map[string]interface{}{"type": "integer"},
					"level":       map[string]interface{}{"type": "integer"},
					"description": map[string]string{"type": "string"},
				},
			},
			"CourseLecturesResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"course_code": map[string]string{"type": "string"},
					"lectures": map[string]interface{}{
						"type":  "array",
						"items": map[string]string{"type": "object"},
					},
					"total": map[string]interface{}{"type": "integer"},
				},
			},
			"CourseStudentsResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"course_code": map[string]string{"type": "string"},
					"students": map[string]interface{}{
						"type":  "array",
						"items": map[string]string{"type": "object"},
					},
					"total": map[string]interface{}{"type": "integer"},
				},
			},
			"CollegesResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"colleges": map[string]interface{}{
						"type":  "array",
						"items": map[string]string{"type": "object"},
					},
					"total": map[string]interface{}{"type": "integer"},
				},
			},
			"DepartmentsResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"departments": map[string]interface{}{
						"type":  "array",
						"items": map[string]string{"type": "object"},
					},
					"total": map[string]interface{}{"type": "integer"},
				},
			},
			"ProgramsResponse": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"programs": map[string]interface{}{
						"type":  "array",
						"items": map[string]string{"type": "object"},
					},
					"total": map[string]interface{}{"type": "integer"},
				},
			},
			"BulkEnrollmentRequest": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"enrollments": map[string]interface{}{
						"type":  "array",
						"items": map[string]string{"type": "object"},
					},
				},
			},
		},
	}
}
