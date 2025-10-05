package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/services"
	"gorm.io/gorm"
)

type OAuthHandler struct {
	db           *gorm.DB
	cfg          *config.Config
	oauthService *services.OAuthService
}

func NewOAuthHandler(db *gorm.DB, cfg *config.Config) *OAuthHandler {
	return &OAuthHandler{
		db:           db,
		cfg:          cfg,
		oauthService: services.NewOAuthService(db, cfg),
	}
}

// Authorize handles OAuth authorization endpoint
// GET /oauth/authorize?client_id=xxx&redirect_uri=xxx&response_type=code&state=xxx
func (h *OAuthHandler) Authorize(c *fiber.Ctx) error {
	clientID := c.Query("client_id")
	redirectURI := c.Query("redirect_uri")
	responseType := c.Query("response_type")
	state := c.Query("state")

	// Validate parameters
	if clientID == "" || redirectURI == "" || responseType != "code" {
		return c.Status(400).SendString("Invalid OAuth parameters")
	}

	// Validate client
	_, err := h.oauthService.ValidateClient(clientID, redirectURI)
	if err != nil {
		return c.Status(400).SendString("Invalid client_id or redirect_uri")
	}

	// Check if this is a POST (login form submission)
	if c.Method() == "POST" {
		return h.handleLogin(c, clientID, redirectURI, state)
	}

	// Render login page
	return c.Type("html").SendString(h.getLoginPageHTML(clientID, redirectURI, state))
}

// handleLogin processes the login form submission
func (h *OAuthHandler) handleLogin(c *fiber.Ctx, clientID, redirectURI, state string) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Authenticate user
	user, _, err := h.oauthService.AuthenticateUser(username, password)
	if err != nil {
		// Return login page with error
		return c.Type("html").SendString(h.getLoginPageHTML(clientID, redirectURI, state) +
			`<script>alert('Invalid username or password');</script>`)
	}

	// Create authorization code
	code, err := h.oauthService.CreateAuthorizationCode(clientID, user.ID, redirectURI, "student.read courses.read")
	if err != nil {
		return c.Status(500).SendString("Failed to create authorization code")
	}

	// Redirect back to LMS with code
	redirectURL := redirectURI + "?code=" + code
	if state != "" {
		redirectURL += "&state=" + state
	}

	return c.Redirect(redirectURL)
}

// Token handles OAuth token exchange endpoint
// POST /oauth/token
func (h *OAuthHandler) Token(c *fiber.Ctx) error {
	grantType := c.FormValue("grant_type")

	switch grantType {
	case "authorization_code":
		return h.handleAuthorizationCodeGrant(c)
	case "refresh_token":
		return h.handleRefreshTokenGrant(c)
	default:
		return c.Status(400).JSON(fiber.Map{
			"error": "unsupported_grant_type",
		})
	}
}

// handleAuthorizationCodeGrant exchanges authorization code for access token
func (h *OAuthHandler) handleAuthorizationCodeGrant(c *fiber.Ctx) error {
	code := c.FormValue("code")
	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")
	redirectURI := c.FormValue("redirect_uri")

	// Validate parameters
	if code == "" || clientID == "" || clientSecret == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid_request",
		})
	}

	// Exchange code for token
	token, err := h.oauthService.ExchangeCodeForToken(code, clientID, clientSecret, redirectURI)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid_grant",
			"error_description": err.Error(),
		})
	}

	// Return token response
	return c.JSON(fiber.Map{
		"access_token":  token.Token,
		"token_type":    "Bearer",
		"expires_in":    int(time.Until(token.ExpiresAt).Seconds()),
		"refresh_token": token.RefreshToken,
		"scope":         token.Scopes,
	})
}

// handleRefreshTokenGrant exchanges refresh token for new access token
func (h *OAuthHandler) handleRefreshTokenGrant(c *fiber.Ctx) error {
	refreshToken := c.FormValue("refresh_token")

	if refreshToken == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid_request",
		})
	}

	// Get new access token
	token, err := h.oauthService.RefreshAccessToken(refreshToken)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid_grant",
		})
	}

	return c.JSON(fiber.Map{
		"access_token":  token.Token,
		"token_type":    "Bearer",
		"expires_in":    int(time.Until(token.ExpiresAt).Seconds()),
		"refresh_token": token.RefreshToken,
	})
}

// getLoginPageHTML returns the SIMS-style login page HTML
func (h *OAuthHandler) getLoginPageHTML(clientID, redirectURI, state string) string {
	return `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>MUST SIMS - Login</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: Arial, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .container {
            background: white;
            border-radius: 10px;
            box-shadow: 0 10px 40px rgba(0,0,0,0.2);
            overflow: hidden;
            max-width: 900px;
            width: 90%;
            display: flex;
        }
        .header {
            background: #4a6fa5;
            color: white;
            padding: 30px;
            flex: 1;
        }
        .header h1 {
            font-size: 24px;
            margin-bottom: 10px;
        }
        .header p {
            font-size: 14px;
            opacity: 0.9;
        }
        .header .logo {
            margin-bottom: 20px;
            font-size: 48px;
        }
        .info {
            margin-top: 30px;
            font-size: 13px;
            line-height: 1.6;
        }
        .form-container {
            padding: 40px;
            flex: 1;
        }
        .form-container h2 {
            margin-bottom: 30px;
            color: #333;
        }
        .form-group {
            margin-bottom: 20px;
        }
        .form-group label {
            display: block;
            margin-bottom: 8px;
            color: #555;
            font-size: 14px;
            font-weight: 500;
        }
        .form-group input {
            width: 100%;
            padding: 12px;
            border: 1px solid #ddd;
            border-radius: 5px;
            font-size: 14px;
        }
        .form-group input:focus {
            outline: none;
            border-color: #4a6fa5;
        }
        .btn {
            background: #4a6fa5;
            color: white;
            padding: 12px 30px;
            border: none;
            border-radius: 5px;
            font-size: 16px;
            cursor: pointer;
            width: 100%;
            margin-top: 10px;
        }
        .btn:hover {
            background: #3a5a85;
        }
        .year-info {
            margin-top: 20px;
            font-size: 12px;
            color: #888;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <div class="logo">🎓</div>
            <h1>MBEYA UNIVERSITY OF SCIENCE AND TECHNOLOGY</h1>
            <p>STUDENT INFORMATION MANAGEMENT SYSTEM { SIMS }</p>
            <div class="info">
                <p><strong>Welcome to SIMS</strong></p>
                <p>The Student Information Management System (SIMS) holds all the information relating to students.</p>
                <br>
                <p><strong>Students</strong></p>
                <ul style="margin-left: 20px; margin-top: 5px;">
                    <li>Register for Courses online</li>
                    <li>View Course Progress and Results</li>
                    <li>Forums</li>
                </ul>
            </div>
        </div>
        <div class="form-container">
            <h2>Login</h2>
            <form method="POST" action="/oauth/authorize?client_id=` + clientID + `&redirect_uri=` + redirectURI + `&response_type=code&state=` + state + `">
                <div class="form-group">
                    <label>Username</label>
                    <input type="text" name="username" placeholder="Registration Number or Email" required>
                </div>
                <div class="form-group">
                    <label>Password</label>
                    <input type="password" name="password" placeholder="Password" required>
                </div>
                <button type="submit" class="btn">Login to Your Account</button>
            </form>
            <div class="year-info">
                Academic Year: 2024/2025
            </div>
        </div>
    </div>
</body>
</html>
	`
}
