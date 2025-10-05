package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type DocsHandler struct{}

func NewDocsHandler() *DocsHandler {
	return &DocsHandler{}
}

// ServeAPIDocs serves a beautiful custom API documentation with dark mode and orange accents
func (h *DocsHandler) ServeAPIDocs(c *fiber.Ctx) error {
	html := h.generateHTML()
	return c.Type("html").SendString(html)
}

// ServeSwaggerUI redirects to the new API docs
func (h *DocsHandler) ServeSwaggerUI(c *fiber.Ctx) error {
	return c.Redirect("/docs")
}

// ServeReDoc redirects to the new API docs
func (h *DocsHandler) ServeReDoc(c *fiber.Ctx) error {
	return c.Redirect("/docs")
}

// ServeSwaggerJSON serves a minimal OpenAPI spec (for compatibility)
func (h *DocsHandler) ServeSwaggerJSON(c *fiber.Ctx) error {
	spec := map[string]interface{}{
		"openapi": "3.0.3",
		"info": map[string]interface{}{
			"title":       "Mock SIMS API",
			"description": "MUST Student Information Management System - OAuth 2.0 & REST API",
			"version":     "1.0.0",
		},
		"servers": []map[string]string{
			{"url": "http://localhost:8000", "description": "Local Development"},
		},
		"paths": map[string]interface{}{
			"/health": map[string]interface{}{
				"get": map[string]interface{}{
					"summary": "Health check",
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Service is healthy",
						},
					},
				},
			},
		},
	}
	return c.JSON(spec)
}

func (h *DocsHandler) generateHTML() string {
	var html strings.Builder

	html.WriteString(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Mock SIMS API Documentation</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
    <style>`)

	html.WriteString(h.getCSS())

	html.WriteString(`</style>
</head>
<body>`)

	html.WriteString(h.getHeader())
	html.WriteString(h.getContainer())
	html.WriteString(h.getFooter())
	html.WriteString(h.getJavaScript())

	html.WriteString(`</body>
</html>`)

	return html.String()
}

func (h *DocsHandler) getCSS() string {
	return `
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        :root {
            --bg-primary: #0f0f0f;
            --bg-secondary: #1a1a1a;
            --bg-tertiary: #242424;
            --orange-primary: #ff6b35;
            --orange-secondary: #ff8c61;
            --orange-dark: #e55a2b;
            --text-primary: #ffffff;
            --text-secondary: #b0b0b0;
            --text-muted: #6b6b6b;
            --border-color: #333333;
            --success: #4ade80;
            --warning: #fbbf24;
            --error: #ef4444;
            --info: #3b82f6;
            --get-color: #61affe;
            --post-color: #49cc90;
            --put-color: #fca130;
            --delete-color: #f93e3e;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Helvetica', 'Arial', sans-serif;
            background: var(--bg-primary);
            color: var(--text-primary);
            line-height: 1.6;
        }

        .header {
            background: linear-gradient(135deg, #1a1a1a 0%, #242424 100%);
            border-bottom: 2px solid var(--orange-primary);
            padding: 2rem 0;
            position: sticky;
            top: 0;
            z-index: 100;
            box-shadow: 0 4px 20px rgba(0, 0, 0, 0.5);
        }

        .header-content {
            max-width: 1400px;
            margin: 0 auto;
            padding: 0 2rem;
            display: flex;
            align-items: center;
            justify-content: space-between;
        }

        .logo {
            display: flex;
            align-items: center;
            gap: 1rem;
        }

        .logo i {
            font-size: 2.5rem;
            color: var(--orange-primary);
        }

        .logo-text h1 {
            font-size: 1.8rem;
            font-weight: 700;
            background: linear-gradient(135deg, var(--orange-primary), var(--orange-secondary));
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .logo-text p {
            color: var(--text-secondary);
            font-size: 0.9rem;
            margin-top: 0.2rem;
        }

        .header-actions {
            display: flex;
            gap: 1rem;
        }

        .btn {
            padding: 0.6rem 1.5rem;
            border-radius: 8px;
            border: none;
            font-size: 0.9rem;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.3s ease;
            text-decoration: none;
            display: inline-flex;
            align-items: center;
            gap: 0.5rem;
        }

        .btn-primary {
            background: var(--orange-primary);
            color: white;
        }

        .btn-primary:hover {
            background: var(--orange-secondary);
            transform: translateY(-2px);
            box-shadow: 0 4px 12px rgba(255, 107, 53, 0.4);
        }

        .btn-secondary {
            background: var(--bg-tertiary);
            color: var(--text-primary);
            border: 1px solid var(--border-color);
        }

        .btn-secondary:hover {
            background: var(--bg-secondary);
            border-color: var(--orange-primary);
        }

        .container {
            max-width: 1400px;
            margin: 0 auto;
            padding: 3rem 2rem;
        }

        .hero {
            background: linear-gradient(135deg, var(--bg-secondary) 0%, var(--bg-tertiary) 100%);
            border-radius: 16px;
            padding: 3rem;
            margin-bottom: 3rem;
            border: 1px solid var(--border-color);
            position: relative;
            overflow: hidden;
        }

        .hero::before {
            content: '';
            position: absolute;
            top: 0;
            right: 0;
            width: 300px;
            height: 300px;
            background: radial-gradient(circle, rgba(255, 107, 53, 0.1) 0%, transparent 70%);
            border-radius: 50%;
        }

        .hero-content {
            position: relative;
            z-index: 1;
        }

        .hero h2 {
            font-size: 2.5rem;
            margin-bottom: 1rem;
            background: linear-gradient(135deg, var(--text-primary), var(--text-secondary));
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .hero p {
            color: var(--text-secondary);
            font-size: 1.1rem;
            max-width: 800px;
        }

        .stats {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 1.5rem;
            margin-top: 2rem;
        }

        .stat-card {
            background: var(--bg-primary);
            padding: 1.5rem;
            border-radius: 12px;
            border: 1px solid var(--border-color);
            text-align: center;
            transition: all 0.3s ease;
        }

        .stat-card:hover {
            border-color: var(--orange-primary);
            transform: translateY(-2px);
        }

        .stat-card i {
            font-size: 2rem;
            color: var(--orange-primary);
            margin-bottom: 0.5rem;
        }

        .stat-card h3 {
            font-size: 2rem;
            margin-bottom: 0.3rem;
        }

        .stat-card p {
            color: var(--text-secondary);
            font-size: 0.9rem;
        }

        .nav-tabs {
            display: flex;
            gap: 1rem;
            border-bottom: 2px solid var(--border-color);
            margin-bottom: 2rem;
            overflow-x: auto;
        }

        .nav-tab {
            padding: 1rem 1.5rem;
            background: transparent;
            color: var(--text-secondary);
            border: none;
            cursor: pointer;
            font-size: 1rem;
            font-weight: 600;
            border-bottom: 3px solid transparent;
            transition: all 0.3s ease;
            white-space: nowrap;
        }

        .nav-tab:hover {
            color: var(--orange-primary);
        }

        .nav-tab.active {
            color: var(--orange-primary);
            border-bottom-color: var(--orange-primary);
        }

        .endpoint-group {
            margin-bottom: 2rem;
        }

        .group-title {
            font-size: 1.5rem;
            margin-bottom: 1rem;
            color: var(--orange-primary);
            display: flex;
            align-items: center;
            gap: 0.5rem;
        }

        .endpoint-card {
            background: var(--bg-secondary);
            border: 1px solid var(--border-color);
            border-radius: 12px;
            margin-bottom: 1rem;
            overflow: hidden;
            transition: all 0.3s ease;
        }

        .endpoint-card:hover {
            border-color: var(--orange-primary);
            box-shadow: 0 4px 20px rgba(255, 107, 53, 0.2);
        }

        .endpoint-header {
            display: flex;
            align-items: center;
            padding: 1.5rem;
            cursor: pointer;
            gap: 1rem;
        }

        .method-badge {
            padding: 0.4rem 1rem;
            border-radius: 6px;
            font-weight: 700;
            font-size: 0.8rem;
            text-transform: uppercase;
            letter-spacing: 0.5px;
            min-width: 70px;
            text-align: center;
        }

        .method-get { background: var(--get-color); color: #000; }
        .method-post { background: var(--post-color); color: #000; }
        .method-put { background: var(--put-color); color: #000; }
        .method-delete { background: var(--delete-color); color: #fff; }

        .endpoint-path {
            flex: 1;
            font-family: 'Courier New', monospace;
            font-size: 1.1rem;
            color: var(--text-primary);
        }

        .endpoint-description {
            color: var(--text-secondary);
            font-size: 0.9rem;
        }

        .expand-icon {
            color: var(--orange-primary);
            transition: transform 0.3s ease;
        }

        .endpoint-header.expanded .expand-icon {
            transform: rotate(180deg);
        }

        .endpoint-body {
            padding: 0 1.5rem 1.5rem;
            display: none;
        }

        .endpoint-body.show {
            display: block;
        }

        .endpoint-section {
            margin-bottom: 1.5rem;
        }

        .section-title {
            font-size: 1.1rem;
            color: var(--orange-primary);
            margin-bottom: 0.8rem;
            font-weight: 600;
        }

        .code-block {
            background: var(--bg-primary);
            border: 1px solid var(--border-color);
            border-radius: 8px;
            padding: 1rem;
            font-family: 'Courier New', monospace;
            font-size: 0.9rem;
            overflow-x: auto;
            position: relative;
        }

        .code-block pre {
            margin: 0;
            color: var(--text-secondary);
        }

        .copy-btn {
            position: absolute;
            top: 0.5rem;
            right: 0.5rem;
            background: var(--bg-tertiary);
            border: 1px solid var(--border-color);
            color: var(--text-secondary);
            padding: 0.4rem 0.8rem;
            border-radius: 6px;
            cursor: pointer;
            font-size: 0.8rem;
            transition: all 0.3s ease;
        }

        .copy-btn:hover {
            background: var(--orange-primary);
            color: white;
            border-color: var(--orange-primary);
        }

        .param-table {
            width: 100%;
            border-collapse: collapse;
        }

        .param-table th,
        .param-table td {
            padding: 0.8rem;
            text-align: left;
            border-bottom: 1px solid var(--border-color);
        }

        .param-table th {
            background: var(--bg-tertiary);
            color: var(--orange-primary);
            font-weight: 600;
        }

        .param-table td {
            color: var(--text-secondary);
        }

        .param-name {
            color: var(--text-primary);
            font-family: 'Courier New', monospace;
        }

        .param-type {
            color: var(--info);
            font-size: 0.85rem;
        }

        .param-required {
            color: var(--error);
            font-size: 0.75rem;
            text-transform: uppercase;
            font-weight: 600;
        }

        .tab-pane {
            display: none;
        }

        .tab-pane.active {
            display: block;
        }

        .footer {
            background: var(--bg-secondary);
            border-top: 1px solid var(--border-color);
            padding: 2rem 0;
            margin-top: 4rem;
        }

        .footer-content {
            max-width: 1400px;
            margin: 0 auto;
            padding: 0 2rem;
            text-align: center;
        }

        .footer p {
            color: var(--text-secondary);
        }

        ::-webkit-scrollbar {
            width: 10px;
            height: 10px;
        }

        ::-webkit-scrollbar-track {
            background: var(--bg-primary);
        }

        ::-webkit-scrollbar-thumb {
            background: var(--border-color);
            border-radius: 5px;
        }

        ::-webkit-scrollbar-thumb:hover {
            background: var(--orange-primary);
        }

        @media (max-width: 768px) {
            .header-content {
                flex-direction: column;
                gap: 1rem;
            }
            .hero h2 {
                font-size: 2rem;
            }
            .stats {
                grid-template-columns: 1fr;
            }
            .endpoint-header {
                flex-wrap: wrap;
            }
        }
    `
}

func (h *DocsHandler) getHeader() string {
	return `
    <header class="header">
        <div class="header-content">
            <div class="logo">
                <i class="fas fa-graduation-cap"></i>
                <div class="logo-text">
                    <h1>Mock SIMS API</h1>
                    <p>Mbeya University of Science and Technology</p>
                </div>
            </div>
            <div class="header-actions">
                <a href="#quick-start" class="btn btn-secondary">
                    <i class="fas fa-rocket"></i> Quick Start
                </a>
                <a href="/health" class="btn btn-primary" target="_blank">
                    <i class="fas fa-heartbeat"></i> Health Check
                </a>
            </div>
        </div>
    </header>`
}

func (h *DocsHandler) getContainer() string {
	return fmt.Sprintf(`
    <div class="container">
        <div class="hero">
            <div class="hero-content">
                <h2>Complete OAuth 2.0 & RESTful API</h2>
                <p>Mock SIMS provides a fully functional OAuth 2.0 authorization server and comprehensive REST API for MUST Learning Management System development and testing.</p>
                <div class="stats">
                    <div class="stat-card">
                        <i class="fas fa-plug"></i>
                        <h3>22</h3>
                        <p>API Endpoints</p>
                    </div>
                    <div class="stat-card">
                        <i class="fas fa-database"></i>
                        <h3>17</h3>
                        <p>Database Models</p>
                    </div>
                    <div class="stat-card">
                        <i class="fas fa-building"></i>
                        <h3>7</h3>
                        <p>MUST Colleges</p>
                    </div>
                    <div class="stat-card">
                        <i class="fas fa-shield-alt"></i>
                        <h3>OAuth 2.0</h3>
                        <p>Authorization</p>
                    </div>
                </div>
            </div>
        </div>

        <div class="nav-tabs">
            <button class="nav-tab active" onclick="showTab('overview')">
                <i class="fas fa-info-circle"></i> Overview
            </button>
            <button class="nav-tab" onclick="showTab('oauth')">
                <i class="fas fa-key"></i> OAuth 2.0
            </button>
            <button class="nav-tab" onclick="showTab('students')">
                <i class="fas fa-user-graduate"></i> Students
            </button>
            <button class="nav-tab" onclick="showTab('faculty')">
                <i class="fas fa-chalkboard-teacher"></i> Faculty
            </button>
            <button class="nav-tab" onclick="showTab('courses')">
                <i class="fas fa-book"></i> Courses
            </button>
            <button class="nav-tab" onclick="showTab('admin')">
                <i class="fas fa-user-shield"></i> Admin
            </button>
        </div>

        <div id="tab-content">
            %s
            %s
            %s
            %s
            %s
            %s
        </div>
    </div>`,
		h.getOverviewTab(),
		h.getOAuthTab(),
		h.getStudentsTab(),
		h.getFacultyTab(),
		h.getCoursesTab(),
		h.getAdminTab(),
	)
}

func (h *DocsHandler) getOverviewTab() string {
	return `
    <div id="overview-tab" class="tab-pane active">
        <div class="endpoint-group" id="quick-start">
            <h2 class="group-title"><i class="fas fa-rocket"></i> Quick Start</h2>
            <div class="endpoint-section">
                <h3 class="section-title">Base URL</h3>
                <div class="code-block">
                    <button class="copy-btn" onclick="copyCode(this)">Copy</button>
                    <pre>http://localhost:8000</pre>
                </div>
            </div>
            <div class="endpoint-section">
                <h3 class="section-title">Authentication Flow</h3>
                <div class="code-block">
                    <button class="copy-btn" onclick="copyCode(this)">Copy</button>
                    <pre>1. Get authorization code: GET /oauth/authorize?client_id=...
2. Exchange for token: POST /oauth/token
3. Use Bearer token: Authorization: Bearer {token}</pre>
                </div>
            </div>
            <div class="endpoint-section">
                <h3 class="section-title">Test Users</h3>
                <table class="param-table">
                    <thead>
                        <tr>
                            <th>Email</th>
                            <th>Password</th>
                            <th>Type</th>
                            <th>Details</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td class="param-name">john.doe@must.ac.tz</td>
                            <td>password123</td>
                            <td class="param-type">Student</td>
                            <td>Year 2, Computer Science</td>
                        </tr>
                        <tr>
                            <td class="param-name">jane.smith@must.ac.tz</td>
                            <td>password123</td>
                            <td class="param-type">Student</td>
                            <td>Year 2</td>
                        </tr>
                        <tr>
                            <td class="param-name">mussa.dida@must.ac.tz</td>
                            <td>password123</td>
                            <td class="param-type">Faculty</td>
                            <td>Lecturer, CS Department</td>
                        </tr>
                        <tr>
                            <td class="param-name">admin@must.ac.tz</td>
                            <td>password123</td>
                            <td class="param-type">Admin</td>
                            <td>Registrar</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>`
}

func (h *DocsHandler) getOAuthTab() string {
	return fmt.Sprintf(`
    <div id="oauth-tab" class="tab-pane">
        <div class="endpoint-group">
            <h2 class="group-title"><i class="fas fa-key"></i> OAuth 2.0 Endpoints</h2>
            %s
            %s
        </div>
    </div>`,
		h.createEndpoint("GET", "/oauth/authorize", "Authorization page", "Displays SIMS login page for user authentication",
			`<table class="param-table">
                <thead>
                    <tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr>
                </thead>
                <tbody>
                    <tr><td class="param-name">client_id</td><td class="param-type">string</td><td class="param-required">required</td><td>OAuth client ID (lms-client-id)</td></tr>
                    <tr><td class="param-name">redirect_uri</td><td class="param-type">string</td><td class="param-required">required</td><td>Callback URL</td></tr>
                    <tr><td class="param-name">response_type</td><td class="param-type">string</td><td class="param-required">required</td><td>Must be 'code'</td></tr>
                    <tr><td class="param-name">state</td><td class="param-type">string</td><td style="color: var(--text-muted)">optional</td><td>Optional state parameter</td></tr>
                </tbody>
            </table>`,
			"Redirects to redirect_uri with authorization code",
			"GET /oauth/authorize?client_id=lms-client-id&redirect_uri=http://localhost:8080/callback&response_type=code"),
		h.createEndpoint("POST", "/oauth/token", "Exchange authorization code for access token", "Exchanges authorization code or refresh token for access token",
			`<table class="param-table">
                <thead>
                    <tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr>
                </thead>
                <tbody>
                    <tr><td class="param-name">grant_type</td><td class="param-type">string</td><td class="param-required">required</td><td>authorization_code or refresh_token</td></tr>
                    <tr><td class="param-name">code</td><td class="param-type">string</td><td class="param-required">required</td><td>Authorization code (for grant_type=authorization_code)</td></tr>
                    <tr><td class="param-name">client_id</td><td class="param-type">string</td><td class="param-required">required</td><td>OAuth client ID</td></tr>
                    <tr><td class="param-name">client_secret</td><td class="param-type">string</td><td class="param-required">required</td><td>OAuth client secret</td></tr>
                    <tr><td class="param-name">redirect_uri</td><td class="param-type">string</td><td class="param-required">required</td><td>Must match authorization redirect_uri</td></tr>
                </tbody>
            </table>`,
			`{
  "access_token": "eyJhbGci...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "eyJhbGci..."
}`,
			`curl -X POST http://localhost:8000/oauth/token \
  -d "grant_type=authorization_code" \
  -d "code=ABC123" \
  -d "client_id=lms-client-id" \
  -d "client_secret=secret"`),
	)
}

func (h *DocsHandler) getStudentsTab() string {
	return fmt.Sprintf(`
    <div id="students-tab" class="tab-pane">
        <div class="endpoint-group">
            <h2 class="group-title"><i class="fas fa-user-graduate"></i> Student Endpoints</h2>
            %s
            %s
            %s
            %s
        </div>
    </div>`,
		h.createEndpoint("GET", "/api/students/me", "Get authenticated student profile", "Returns complete profile of the authenticated student", "",
			`{
  "student_id": 1,
  "reg_number": "23100523050032",
  "name": "John Doe Mwamba",
  "email": "john.doe@must.ac.tz",
  "program": {
    "code": "MB011",
    "name": "Bachelor of Science in Computer Science"
  },
  "year_of_study": 2,
  "gpa": 3.45
}`,
			`curl http://localhost:8000/api/students/me \
  -H "Authorization: Bearer {token}"`),
		h.createEndpoint("GET", "/api/students/:id/courses", "Get student's enrolled courses", "Returns all courses the student is enrolled in",
			`<table class="param-table">
                <thead><tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr></thead>
                <tbody>
                    <tr><td class="param-name">id</td><td class="param-type">integer</td><td class="param-required">required</td><td>Student ID</td></tr>
                </tbody>
            </table>`,
			`{
  "courses": [
    {
      "code": "CS 1101",
      "name": "Introduction to Programming",
      "credits": 3,
      "lecturer": "Dr. Mussa Dida"
    }
  ]
}`,
			`curl http://localhost:8000/api/students/1/courses \
  -H "Authorization: Bearer {token}"`),
		h.createEndpoint("GET", "/api/students/:id/grades", "Get student's grades", "Returns all grades with CA marks and final exam scores",
			`<table class="param-table">
                <thead><tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr></thead>
                <tbody>
                    <tr><td class="param-name">id</td><td class="param-type">integer</td><td class="param-required">required</td><td>Student ID</td></tr>
                </tbody>
            </table>`,
			`{
  "grades": [
    {
      "course_code": "CS 1101",
      "ca_marks": 35.5,
      "final_exam": 52.0,
      "total_marks": 87.5,
      "letter_grade": "A",
      "grade_point": 5.0
    }
  ]
}`,
			`curl http://localhost:8000/api/students/1/grades \
  -H "Authorization: Bearer {token}"`),
		h.createEndpoint("GET", "/api/students/:id/timetable", "Get student's class timetable", "Returns weekly class schedule",
			`<table class="param-table">
                <thead><tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr></thead>
                <tbody>
                    <tr><td class="param-name">id</td><td class="param-type">integer</td><td class="param-required">required</td><td>Student ID</td></tr>
                </tbody>
            </table>`,
			`{
  "timetable": [
    {
      "day": "Monday",
      "time": "10:00-12:00",
      "course": "CS 1101",
      "venue": "CoICT 101",
      "lecturer": "Dr. Mussa Dida"
    }
  ]
}`,
			`curl http://localhost:8000/api/students/1/timetable \
  -H "Authorization: Bearer {token}"`),
	)
}

func (h *DocsHandler) getFacultyTab() string {
	return fmt.Sprintf(`
    <div id="faculty-tab" class="tab-pane">
        <div class="endpoint-group">
            <h2 class="group-title"><i class="fas fa-chalkboard-teacher"></i> Faculty Endpoints</h2>
            %s
            %s
            %s
        </div>
    </div>`,
		h.createEndpoint("GET", "/api/faculty/me", "Get authenticated faculty profile", "Returns complete profile of the authenticated faculty member", "",
			`{
  "faculty_id": 1,
  "staff_id": "MUST001",
  "name": "Dr. Mussa Ally Dida",
  "department": "Computer Science",
  "rank": "Lecturer"
}`,
			`curl http://localhost:8000/api/faculty/me \
  -H "Authorization: Bearer {token}"`),
		h.createEndpoint("GET", "/api/faculty/:id/courses", "Get faculty teaching assignments", "Returns all courses assigned to the faculty member",
			`<table class="param-table">
                <thead><tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr></thead>
                <tbody>
                    <tr><td class="param-name">id</td><td class="param-type">integer</td><td class="param-required">required</td><td>Faculty ID</td></tr>
                </tbody>
            </table>`,
			`{
  "courses": [
    {
      "code": "CS 1101",
      "name": "Introduction to Programming",
      "students_count": 45,
      "semester": "2024/2025 - Semester II"
    }
  ]
}`,
			`curl http://localhost:8000/api/faculty/1/courses \
  -H "Authorization: Bearer {token}"`),
		h.createEndpoint("POST", "/api/faculty/courses/:id/ca-marks", "Submit Continuous Assessment marks", "Submit CA marks for students in a course (0-40)",
			`<table class="param-table">
                <thead><tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr></thead>
                <tbody>
                    <tr><td class="param-name">id</td><td class="param-type">integer</td><td class="param-required">required</td><td>Course ID</td></tr>
                </tbody>
            </table>`,
			`{
  "message": "CA marks submitted successfully",
  "course_id": 1,
  "marks_count": 2
}`,
			`curl -X POST http://localhost:8000/api/faculty/courses/1/ca-marks \
  -H "Authorization: Bearer {token}" \
  -H "Content-Type: application/json" \
  -d '{
    "marks": [
      {"student_id": 1, "ca_marks": 35},
      {"student_id": 2, "ca_marks": 32}
    ]
  }'`),
	)
}

func (h *DocsHandler) getCoursesTab() string {
	return fmt.Sprintf(`
    <div id="courses-tab" class="tab-pane">
        <div class="endpoint-group">
            <h2 class="group-title"><i class="fas fa-book"></i> Course Endpoints</h2>
            %s
            %s
            %s
            %s
        </div>
    </div>`,
		h.createEndpoint("GET", "/api/courses", "List all courses (paginated)", "Returns paginated list of all courses",
			`<table class="param-table">
                <thead><tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr></thead>
                <tbody>
                    <tr><td class="param-name">page</td><td class="param-type">integer</td><td style="color: var(--text-muted)">optional</td><td>Page number (default: 1)</td></tr>
                    <tr><td class="param-name">limit</td><td class="param-type">integer</td><td style="color: var(--text-muted)">optional</td><td>Items per page (default: 20)</td></tr>
                </tbody>
            </table>`,
			`{
  "courses": [
    {
      "code": "CS 1101",
      "name": "Introduction to Programming",
      "credits": 3,
      "level": 100
    }
  ],
  "total": 50,
  "page": 1,
  "limit": 20
}`,
			`curl "http://localhost:8000/api/courses?page=1&limit=20" \
  -H "Authorization: Bearer {token}"`),
		h.createEndpoint("GET", "/api/courses/:code", "Get course details", "Returns detailed information about a specific course",
			`<table class="param-table">
                <thead><tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr></thead>
                <tbody>
                    <tr><td class="param-name">code</td><td class="param-type">string</td><td class="param-required">required</td><td>Course code (e.g., CS 1101)</td></tr>
                </tbody>
            </table>`,
			`{
  "code": "CS 1101",
  "name": "Introduction to Programming",
  "credits": 3,
  "level": 100,
  "description": "Introduction to programming concepts",
  "department": "Computer Science"
}`,
			`curl "http://localhost:8000/api/courses/CS%201101" \
  -H "Authorization: Bearer {token}"`),
		h.createEndpoint("GET", "/api/courses/:code/lectures", "Get course lecture schedule", "Returns weekly lecture schedule for a course",
			`<table class="param-table">
                <thead><tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr></thead>
                <tbody>
                    <tr><td class="param-name">code</td><td class="param-type">string</td><td class="param-required">required</td><td>Course code</td></tr>
                </tbody>
            </table>`,
			`{
  "lectures": [
    {
      "day": "Monday",
      "start_time": "10:00",
      "end_time": "12:00",
      "venue": "CoICT 101",
      "lecturer": "Dr. Mussa Dida"
    }
  ]
}`,
			`curl "http://localhost:8000/api/courses/CS%201101/lectures" \
  -H "Authorization: Bearer {token}"`),
		h.createEndpoint("GET", "/api/courses/:code/students", "Get enrolled students", "Returns all students enrolled in a course",
			`<table class="param-table">
                <thead><tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr></thead>
                <tbody>
                    <tr><td class="param-name">code</td><td class="param-type">string</td><td class="param-required">required</td><td>Course code</td></tr>
                </tbody>
            </table>`,
			`{
  "students": [
    {
      "reg_number": "23100523050032",
      "name": "John Doe",
      "program": "BSc Computer Science",
      "year": 2
    }
  ]
}`,
			`curl "http://localhost:8000/api/courses/CS%201101/students" \
  -H "Authorization: Bearer {token}"`),
	)
}

func (h *DocsHandler) getAdminTab() string {
	return fmt.Sprintf(`
    <div id="admin-tab" class="tab-pane">
        <div class="endpoint-group">
            <h2 class="group-title"><i class="fas fa-user-shield"></i> Admin Endpoints</h2>
            %s
            %s
            %s
            %s
        </div>
    </div>`,
		h.createEndpoint("GET", "/api/colleges", "Get all MUST colleges", "Returns list of all 7 MUST colleges", "",
			`{
  "colleges": [
    {
      "id": 1,
      "code": "01",
      "name": "College of Information and Communication Technology",
      "short_name": "CoICT",
      "dean": "Prof. Dr. Joseph Mkunda"
    }
  ]
}`,
			`curl http://localhost:8000/api/colleges \
  -H "Authorization: Bearer {token}"`),
		h.createEndpoint("GET", "/api/departments", "Get departments", "Returns departments, optionally filtered by college",
			`<table class="param-table">
                <thead><tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr></thead>
                <tbody>
                    <tr><td class="param-name">college_id</td><td class="param-type">integer</td><td style="color: var(--text-muted)">optional</td><td>Filter by college ID</td></tr>
                </tbody>
            </table>`,
			`{
  "departments": [
    {
      "id": 1,
      "code": "CS",
      "name": "Computer Science and Engineering",
      "head": "Dr. Mussa Ally Dida",
      "college": "CoICT"
    }
  ]
}`,
			`curl "http://localhost:8000/api/departments?college_id=1" \
  -H "Authorization: Bearer {token}"`),
		h.createEndpoint("GET", "/api/programs", "Get academic programs", "Returns programs, filterable by department or level",
			`<table class="param-table">
                <thead><tr><th>Name</th><th>Type</th><th>Required</th><th>Description</th></tr></thead>
                <tbody>
                    <tr><td class="param-name">department_id</td><td class="param-type">integer</td><td style="color: var(--text-muted)">optional</td><td>Filter by department</td></tr>
                    <tr><td class="param-name">level</td><td class="param-type">string</td><td style="color: var(--text-muted)">optional</td><td>Filter by level (Bachelor, Diploma, Masters)</td></tr>
                </tbody>
            </table>`,
			`{
  "programs": [
    {
      "code": "MB011",
      "name": "Bachelor of Science in Computer Science",
      "level": "Bachelor",
      "duration": 3,
      "tuition_fees": 1500000
    }
  ]
}`,
			`curl "http://localhost:8000/api/programs?level=Bachelor" \
  -H "Authorization: Bearer {token}"`),
		h.createEndpoint("POST", "/api/enrollments", "Create bulk enrollments", "Enroll multiple students in courses", "",
			`{
  "message": "Enrollments created successfully",
  "count": 5
}`,
			`curl -X POST http://localhost:8000/api/enrollments \
  -H "Authorization: Bearer {token}" \
  -H "Content-Type: application/json" \
  -d '{
    "enrollments": [
      {
        "student_id": 1,
        "course_id": 5,
        "semester_id": 4
      }
    ]
  }'`),
	)
}

func (h *DocsHandler) createEndpoint(method, path, title, description, params, response, example string) string {
	paramsSection := ""
	if params != "" {
		paramsSection = fmt.Sprintf(`
            <div class="endpoint-section">
                <h3 class="section-title">Parameters</h3>
                %s
            </div>`, params)
	}

	responseSection := ""
	if response != "" {
		responseSection = fmt.Sprintf(`
            <div class="endpoint-section">
                <h3 class="section-title">Response Example</h3>
                <div class="code-block">
                    <button class="copy-btn" onclick="copyCode(this)">Copy</button>
                    <pre>%s</pre>
                </div>
            </div>`, response)
	}

	exampleSection := ""
	if example != "" {
		exampleSection = fmt.Sprintf(`
            <div class="endpoint-section">
                <h3 class="section-title">Example Request</h3>
                <div class="code-block">
                    <button class="copy-btn" onclick="copyCode(this)">Copy</button>
                    <pre>%s</pre>
                </div>
            </div>`, example)
	}

	return fmt.Sprintf(`
    <div class="endpoint-card">
        <div class="endpoint-header" onclick="toggleEndpoint(this)">
            <span class="method-badge method-%s">%s</span>
            <div style="flex: 1;">
                <div class="endpoint-path">%s</div>
                <div class="endpoint-description">%s</div>
            </div>
            <i class="fas fa-chevron-down expand-icon"></i>
        </div>
        <div class="endpoint-body">
            %s
            %s
            %s
        </div>
    </div>`,
		strings.ToLower(method), method, path, description,
		paramsSection, responseSection, exampleSection,
	)
}

func (h *DocsHandler) getFooter() string {
	return `
    <footer class="footer">
        <div class="footer-content">
            <p>&copy; 2025 Mbeya University of Science and Technology - Mock SIMS API v1.0</p>
            <p style="margin-top: 0.5rem; color: var(--text-muted);">
                Built with <i class="fas fa-heart" style="color: var(--orange-primary);"></i> for LMS Development
            </p>
        </div>
    </footer>`
}

func (h *DocsHandler) getJavaScript() string {
	return `
    <script>
        function showTab(tabName) {
            document.querySelectorAll('.tab-pane').forEach(tab => {
                tab.classList.remove('active');
            });
            document.querySelectorAll('.nav-tab').forEach(tab => {
                tab.classList.remove('active');
            });
            document.getElementById(tabName + '-tab').classList.add('active');
            event.target.classList.add('active');
        }

        function toggleEndpoint(element) {
            const body = element.nextElementSibling;
            element.classList.toggle('expanded');
            body.classList.toggle('show');
        }

        function copyCode(button) {
            const codeBlock = button.nextElementSibling;
            const code = codeBlock.textContent;
            navigator.clipboard.writeText(code).then(() => {
                const originalText = button.textContent;
                button.textContent = 'Copied!';
                button.style.background = 'var(--success)';
                button.style.color = 'white';
                setTimeout(() => {
                    button.textContent = originalText;
                    button.style.background = '';
                    button.style.color = '';
                }, 2000);
            });
        }
    </script>`
}
