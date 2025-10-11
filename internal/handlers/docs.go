package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type DocsHandler struct{}

func NewDocsHandler() *DocsHandler {
	return &DocsHandler{}
}

// ServeSwaggerUI serves the Swagger UI with beautiful dark mode
func (h *DocsHandler) ServeSwaggerUI(c *fiber.Ctx) error {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Mock SIMS API - Swagger UI</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.10.5/swagger-ui.css">
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap" rel="stylesheet">
    <style>
        :root {
            --blue-primary: #3b82f6;
            --blue-light: #60a5fa;
            --blue-dark: #2563eb;
            --orange-primary: #f97316;
            --orange-light: #fb923c;
            --green-primary: #10b981;
            --green-light: #34d399;
            --bg-dark: #0f172a;
            --bg-secondary: #1e293b;
            --bg-tertiary: #334155;
            --text-primary: #f1f5f9;
            --text-secondary: #cbd5e1;
            --border-color: #475569;
        }

        body {
            margin: 0;
            background: var(--bg-dark);
            font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
        }

        .swagger-ui {
            max-width: 1600px;
            margin: 0 auto;
            padding: 0;
        }

        /* Top Bar */
        .swagger-ui .topbar {
            background: linear-gradient(135deg, var(--blue-primary) 0%, var(--blue-dark) 100%) !important;
            padding: 20px 0 !important;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.3);
            border: none !important;
        }

        .swagger-ui .topbar-wrapper {
            max-width: 1600px;
            margin: 0 auto;
            padding: 0 40px;
        }

        .swagger-ui .topbar .topbar-wrapper .link {
            display: flex !important;
            align-items: center;
        }

        .swagger-ui .topbar .topbar-wrapper .link::after {
            content: "🎓 Mock SIMS API";
            color: #fff !important;
            font-size: 24px;
            font-weight: 700;
            letter-spacing: -0.5px;
            line-height: 1;
        }

        .swagger-ui .topbar .topbar-wrapper .link img,
        .swagger-ui .topbar-wrapper img,
        .swagger-ui .topbar-wrapper a img {
            display: none !important;
            visibility: hidden !important;
        }

        .swagger-ui .topbar .download-url-wrapper,
        .swagger-ui .topbar-wrapper .download-url-wrapper {
            display: none !important;
        }

        /* Info Section */
        .swagger-ui .information-container {
            background: var(--bg-dark) !important;
            padding: 40px;
            margin: 0;
            border-bottom: 1px solid var(--border-color);
        }

        .swagger-ui .info {
            margin: 0 auto;
            max-width: 1520px;
        }

        .swagger-ui .info .title {
            color: var(--blue-light);
            font-size: 32px;
            font-weight: 700;
            margin-bottom: 16px;
        }

        .swagger-ui .info .description,
        .swagger-ui .info p,
        .swagger-ui .info li {
            color: var(--text-secondary);
            font-size: 15px;
            line-height: 1.7;
        }

        .swagger-ui .info a {
            color: var(--blue-light);
            text-decoration: none;
        }

        .swagger-ui .info a:hover {
            color: var(--blue-primary);
            text-decoration: underline;
        }

        /* Main Content */
        .swagger-ui .wrapper {
            padding: 0 40px 40px;
            max-width: 1600px;
            margin: 0 auto;
        }

        /* Server Selector - Fix white box */
        .swagger-ui .scheme-container {
            background: var(--bg-dark) !important;
            padding: 30px 0;
            margin: 0;
            box-shadow: none;
        }

        .swagger-ui .scheme-container .schemes {
            max-width: 1520px;
            margin: 0 auto;
            padding: 20px 30px;
            background: var(--bg-secondary) !important;
            border: 1px solid var(--border-color);
            border-radius: 8px;
        }

        .swagger-ui .scheme-container .schemes > label {
            color: var(--text-secondary) !important;
            font-weight: 600;
            font-size: 14px;
        }

        .swagger-ui .scheme-container .schemes select {
            background: var(--bg-tertiary) !important;
            color: var(--text-primary) !important;
            border: 1px solid var(--border-color) !important;
            border-radius: 6px;
            padding: 10px 14px;
        }

        /* Filter Box */
        .swagger-ui .filter-container {
            background: var(--bg-dark) !important;
            padding: 20px 0;
            margin: 0;
            border-bottom: 1px solid var(--border-color);
        }

        .swagger-ui .filter {
            max-width: 1520px;
            margin: 0 auto;
        }

        .swagger-ui .filter input {
            background: var(--bg-tertiary) !important;
            color: var(--text-primary) !important;
            border: 1px solid var(--border-color) !important;
            border-radius: 6px;
            padding: 12px 16px;
            font-size: 15px;
        }

        .swagger-ui .filter input:focus {
            border-color: var(--blue-primary) !important;
            outline: none;
            box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
        }

        /* Section Headers */
        .swagger-ui .opblock-tag-section {
            margin-top: 32px;
        }

        .swagger-ui .opblock-tag {
            background: var(--bg-secondary);
            color: var(--text-primary);
            font-size: 20px;
            font-weight: 600;
            padding: 16px 24px;
            border: none;
            border-left: 4px solid var(--blue-primary);
            margin-bottom: 16px;
            border-radius: 6px;
            transition: all 0.2s ease;
        }

        .swagger-ui .opblock-tag:hover {
            background: var(--bg-tertiary);
            border-left-color: var(--blue-light);
        }

        .swagger-ui .opblock-tag small {
            color: var(--text-secondary);
            font-weight: 400;
        }

        /* Operation Blocks */
        .swagger-ui .opblock {
            background: var(--bg-secondary);
            border: 1px solid var(--border-color);
            border-radius: 8px;
            margin-bottom: 16px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
            overflow: hidden;
        }

        .swagger-ui .opblock .opblock-summary {
            padding: 16px 20px;
            align-items: center;
        }

        .swagger-ui .opblock.opblock-get {
            border-left: 4px solid var(--blue-primary);
        }

        .swagger-ui .opblock.opblock-get .opblock-summary-method {
            background: var(--blue-primary);
        }

        .swagger-ui .opblock.opblock-post {
            border-left: 4px solid var(--green-primary);
        }

        .swagger-ui .opblock.opblock-post .opblock-summary-method {
            background: var(--green-primary);
        }

        .swagger-ui .opblock.opblock-put {
            border-left: 4px solid var(--orange-primary);
        }

        .swagger-ui .opblock.opblock-put .opblock-summary-method {
            background: var(--orange-primary);
        }

        .swagger-ui .opblock.opblock-delete {
            border-left: 4px solid #ef4444;
        }

        .swagger-ui .opblock.opblock-delete .opblock-summary-method {
            background: #ef4444;
        }

        .swagger-ui .opblock-summary-method {
            font-weight: 700;
            min-width: 80px;
            border-radius: 6px;
            padding: 8px 16px;
            text-align: center;
        }

        .swagger-ui .opblock-summary-path {
            color: var(--text-primary);
            font-family: 'Monaco', 'Courier New', monospace;
            font-size: 14px;
            font-weight: 500;
        }

        .swagger-ui .opblock-summary-description {
            color: var(--text-secondary);
        }

        /* Parameters & Request Body */
        .swagger-ui .opblock-body {
            background: var(--bg-dark);
            padding: 24px;
        }

        .swagger-ui .opblock-section-header {
            background: var(--bg-tertiary);
            padding: 12px 16px;
            margin: -24px -24px 20px -24px;
            border-bottom: 1px solid var(--border-color);
        }

        .swagger-ui .opblock-section-header h4 {
            color: var(--text-primary);
            font-size: 16px;
            font-weight: 600;
        }

        .swagger-ui .table-container,
        .swagger-ui table {
            background: transparent;
        }

        .swagger-ui table thead tr th,
        .swagger-ui table thead tr td {
            color: var(--text-secondary);
            font-size: 13px;
            font-weight: 600;
            text-transform: uppercase;
            letter-spacing: 0.5px;
            border-bottom: 2px solid var(--border-color);
            padding: 12px 16px;
        }

        .swagger-ui table tbody tr td {
            color: var(--text-primary);
            padding: 12px 16px;
            border-bottom: 1px solid var(--border-color);
        }

        .swagger-ui .parameter__name {
            color: var(--blue-light);
            font-weight: 600;
        }

        .swagger-ui .parameter__type {
            color: var(--green-light);
            font-family: 'Monaco', monospace;
        }

        /* Input Fields */
        .swagger-ui input[type=text],
        .swagger-ui input[type=password],
        .swagger-ui input[type=email],
        .swagger-ui textarea,
        .swagger-ui select {
            background: var(--bg-tertiary);
            color: var(--text-primary);
            border: 1px solid var(--border-color);
            border-radius: 6px;
            padding: 10px 14px;
            font-size: 14px;
            transition: all 0.2s ease;
        }

        .swagger-ui input[type=text]:focus,
        .swagger-ui textarea:focus,
        .swagger-ui select:focus {
            outline: none;
            border-color: var(--blue-primary);
            box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
        }

        /* Buttons */
        .swagger-ui .btn {
            background: var(--blue-primary);
            color: #fff;
            border: none;
            border-radius: 6px;
            padding: 10px 20px;
            font-weight: 600;
            font-size: 14px;
            cursor: pointer;
            transition: all 0.2s ease;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
        }

        .swagger-ui .btn:hover {
            background: var(--blue-dark);
            transform: translateY(-1px);
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
        }

        .swagger-ui .btn.execute {
            background: var(--green-primary);
        }

        .swagger-ui .btn.execute:hover {
            background: #059669;
        }

        .swagger-ui .btn.cancel {
            background: #ef4444;
        }

        .swagger-ui .btn.cancel:hover {
            background: #dc2626;
        }

        .swagger-ui .btn.authorize {
            background: var(--orange-primary);
        }

        .swagger-ui .btn.authorize:hover {
            background: #ea580c;
        }

        /* Responses */
        .swagger-ui .responses-wrapper {
            margin-top: 24px;
        }

        .swagger-ui .response {
            background: var(--bg-secondary);
            border: 1px solid var(--border-color);
            border-radius: 6px;
            padding: 16px;
            margin-bottom: 12px;
        }

        .swagger-ui .response-col_status {
            color: var(--green-light);
            font-weight: 700;
            font-size: 16px;
        }

        .swagger-ui .response-col_description {
            color: var(--text-secondary);
        }

        /* Code Blocks */
        .swagger-ui .highlight-code,
        .swagger-ui .microlight {
            background: var(--bg-dark) !important;
            border: 1px solid var(--border-color);
            border-radius: 6px;
            padding: 16px;
            color: var(--text-primary);
            font-family: 'Monaco', 'Courier New', monospace;
            font-size: 13px;
        }

        .swagger-ui .copy-to-clipboard {
            background: var(--bg-tertiary);
            border: 1px solid var(--border-color);
            right: 8px;
            top: 8px;
        }

        .swagger-ui .copy-to-clipboard button {
            color: var(--text-secondary);
        }

        /* Models */
        .swagger-ui .model-box,
        .swagger-ui .model {
            background: var(--bg-secondary);
            border-radius: 6px;
            padding: 16px;
        }

        .swagger-ui .model-title {
            color: var(--text-primary);
            font-weight: 600;
        }

        .swagger-ui .property-row {
            border-bottom: 1px solid var(--border-color);
        }

        .swagger-ui .prop-name {
            color: var(--blue-light);
        }

        .swagger-ui .prop-type {
            color: var(--green-light);
        }

        /* Scrollbar */
        ::-webkit-scrollbar {
            width: 12px;
            height: 12px;
        }

        ::-webkit-scrollbar-track {
            background: var(--bg-dark);
        }

        ::-webkit-scrollbar-thumb {
            background: var(--bg-tertiary);
            border-radius: 6px;
            border: 2px solid var(--bg-dark);
        }

        ::-webkit-scrollbar-thumb:hover {
            background: var(--border-color);
        }
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
                layout: "BaseLayout",
                defaultModelsExpandDepth: 1,
                defaultModelExpandDepth: 1,
                docExpansion: "list",
                filter: true,
                syntaxHighlight: {
                    activate: true,
                    theme: "monokai"
                }
            });
        };
    </script>
</body>
</html>`
	return c.Type("html").SendString(html)
}

// ServeReDoc serves the ReDoc documentation with beautiful dark mode
func (h *DocsHandler) ServeReDoc(c *fiber.Ctx) error {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Mock SIMS API - Documentation</title>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=JetBrains+Mono:wght@400;500&display=swap" rel="stylesheet">
    <style>
        body {
            margin: 0;
            padding: 0;
            background: #1e293b !important;
        }

        /* Force dark backgrounds everywhere */
        .api-content,
        .api-info,
        .menu-content,
        [role="navigation"],
        [role="search"],
        [class*="search-wrap"],
        [class*="SearchWrap"] {
            background-color: #1e293b !important;
        }

        /* Fix white boxes - Parameter examples, enums, etc */
        .react-tabs__tab-panel,
        .react-tabs,
        table,
        td,
        th,
        tr,
        tbody,
        thead,
        pre,
        code,
        .redoc-json,
        [class*="dropdown"],
        [class*="Dropdown"],
        [class*="example"],
        [class*="Example"],
        [class*="Table"],
        [class*="enum"],
        [class*="Enum"],
        [class*="DropdownLabel"],
        [class*="DropdownOption"] {
            background-color: #0f172a !important;
            color: #cbd5e1 !important;
            border-color: #334155 !important;
        }

        /* Search box */
        input[type="text"],
        input[type="search"],
        [class*="search"] input,
        [class*="Search"] input {
            background: #0f172a !important;
            color: #f1f5f9 !important;
            border: 1px solid #334155 !important;
            border-radius: 6px !important;
            padding: 10px 14px !important;
        }

        input:focus {
            border-color: #3b82f6 !important;
            outline: none !important;
            box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15) !important;
        }

        /* Parameter table styling */
        table {
            background: #0f172a !important;
        }

        table th {
            background: #1e293b !important;
            color: #f1f5f9 !important;
            border-bottom: 2px solid #334155 !important;
            padding: 12px 16px !important;
            font-weight: 600 !important;
        }

        table td {
            background: #0f172a !important;
            color: #cbd5e1 !important;
            border-bottom: 1px solid #334155 !important;
            padding: 12px 16px !important;
        }

        table tr:hover td {
            background: #1e293b !important;
        }

        /* Example boxes */
        [class*="example"] {
            background: #0f172a !important;
            border: 1px solid #334155 !important;
            border-radius: 6px !important;
            padding: 12px !important;
        }

        /* Dropdown/Select boxes */
        select,
        [class*="dropdown"],
        [class*="Select"] {
            background: #0f172a !important;
            color: #f1f5f9 !important;
            border: 1px solid #334155 !important;
            border-radius: 6px !important;
            padding: 8px 12px !important;
        }

        /* Code blocks */
        pre,
        code {
            background: #0f172a !important;
            color: #cbd5e1 !important;
            border: 1px solid #334155 !important;
            border-radius: 6px !important;
        }

        /* Property names */
        [class*="property"] span:first-child,
        [class*="field"] span:first-child {
            color: #60a5fa !important;
            font-weight: 600 !important;
        }

        /* Property types */
        [class*="type"] {
            color: #34d399 !important;
            font-family: 'JetBrains Mono', monospace !important;
        }

        /* Required labels */
        [class*="required"] {
            color: #ef4444 !important;
        }

        /* Headers and titles */
        h1, h2, h3, h4, h5, h6 {
            color: #f1f5f9 !important;
        }

        /* Regular text */
        p, span, div, li {
            color: #cbd5e1 !important;
        }

        /* Links */
        a {
            color: #3b82f6 !important;
        }

        a:hover {
            color: #60a5fa !important;
        }

        /* Search button */
        button[aria-label*="search"],
        [class*="search"] button {
            background: #f97316 !important;
            color: white !important;
            border: none !important;
            border-radius: 6px !important;
            padding: 10px 20px !important;
            font-weight: 600 !important;
            cursor: pointer !important;
        }

        button[aria-label*="search"]:hover,
        [class*="search"] button:hover {
            background: #ea580c !important;
        }

        /* Method badges - Color palette (Blue, Green, Orange) */
        [class*="get"],
        [class*="GET"] {
            background: #3b82f6 !important;
            color: white !important;
        }

        [class*="post"],
        [class*="POST"] {
            background: #10b981 !important;
            color: white !important;
        }

        [class*="put"],
        [class*="PUT"] {
            background: #f97316 !important;
            color: white !important;
        }

        [class*="delete"],
        [class*="DELETE"] {
            background: #ef4444 !important;
            color: white !important;
        }

        [class*="patch"],
        [class*="PATCH"] {
            background: #f59e0b !important;
            color: white !important;
        }

        /* Response tabs */
        .react-tabs__tab {
            background: #0f172a !important;
            color: #cbd5e1 !important;
            border: 1px solid #334155 !important;
            border-radius: 6px 6px 0 0 !important;
            padding: 10px 16px !important;
            margin-right: 4px !important;
        }

        .react-tabs__tab--selected {
            background: #1e293b !important;
            color: #3b82f6 !important;
            border-bottom-color: #1e293b !important;
        }

        .react-tabs__tab:hover {
            background: #1e293b !important;
        }

        /* Scrollbar */
        ::-webkit-scrollbar {
            width: 12px;
            height: 12px;
        }

        ::-webkit-scrollbar-track {
            background: #0f172a;
        }

        ::-webkit-scrollbar-thumb {
            background: #334155;
            border-radius: 6px;
            border: 2px solid #0f172a;
        }

        ::-webkit-scrollbar-thumb:hover {
            background: #475569;
        }

        /* Server dropdown */
        [class*="ServerItem"],
        [class*="server"] {
            background: #0f172a !important;
            border: 1px solid #334155 !important;
            border-radius: 6px !important;
        }

        /* Input fields */
        input {
            background: #0f172a !important;
            color: #f1f5f9 !important;
            border: 1px solid #334155 !important;
            border-radius: 6px !important;
            padding: 8px 12px !important;
        }

        input:focus {
            border-color: #3b82f6 !important;
            outline: none !important;
        }
    </style>
</head>
<body>
    <redoc
        spec-url='/swagger.json'
        hide-download-button
        required-props-first
        expand-responses="200,201"
        json-sample-expand-level="2"
        native-scrollbars
    ></redoc>
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
