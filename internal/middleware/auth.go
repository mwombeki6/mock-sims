package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/services"
	"gorm.io/gorm"
)

// AuthMiddleware verifies OAuth access token
func AuthMiddleware(db *gorm.DB, cfg *config.Config) fiber.Handler {
	oauthService := services.NewOAuthService(db, cfg)

	return func(c *fiber.Ctx) error {
		// Get Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "missing authorization header",
			})
		}

		// Extract Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid authorization header format",
			})
		}

		token := parts[1]

		// Validate token
		accessToken, user, err := oauthService.ValidateAccessToken(token)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid or expired access token",
			})
		}

		// Store user and token info in context
		c.Locals("user_id", user.ID)
		c.Locals("user_email", user.Email)
		c.Locals("user_type", user.UserType)
		c.Locals("access_token", accessToken)

		return c.Next()
	}
}

// RequireUserType middleware checks if user has specific type
func RequireUserType(allowedTypes ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userType := c.Locals("user_type")
		if userType == nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		userTypeStr := userType.(string)
		for _, allowed := range allowedTypes {
			if userTypeStr == allowed {
				return c.Next()
			}
		}

		return c.Status(403).JSON(fiber.Map{
			"error": "forbidden - insufficient permissions",
		})
	}
}
