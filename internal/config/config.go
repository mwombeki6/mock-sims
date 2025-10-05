package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Port string
	Host string
	Env  string

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// OAuth
	OAuthClientID            string
	OAuthClientSecret        string
	OAuthRedirectURI         string
	OAuthTokenExpiry         string
	OAuthRefreshTokenExpiry  string

	// JWT
	JWTSecret string
	JWTExpiry string

	// LMS Integration
	LMSWebhookURL    string
	LMSWebhookSecret string

	// CORS
	AllowedOrigins string

	// Logging
	LogLevel string
}

func Load() *Config {
	// Load .env file if exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := &Config{
		// Server
		Port: getEnv("PORT", "8000"),
		Host: getEnv("HOST", "0.0.0.0"),
		Env:  getEnv("ENV", "development"),

		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "sims_user"),
		DBPassword: getEnv("DB_PASSWORD", "sims_password"),
		DBName:     getEnv("DB_NAME", "mock_sims"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

		// OAuth
		OAuthClientID:           getEnv("OAUTH_CLIENT_ID", "lms-client-id"),
		OAuthClientSecret:       getEnv("OAUTH_CLIENT_SECRET", "lms-client-secret"),
		OAuthRedirectURI:        getEnv("OAUTH_REDIRECT_URI", "http://localhost:8080/auth/callback"),
		OAuthTokenExpiry:        getEnv("OAUTH_TOKEN_EXPIRY", "3600"),
		OAuthRefreshTokenExpiry: getEnv("OAUTH_REFRESH_TOKEN_EXPIRY", "604800"),

		// JWT
		JWTSecret: getEnv("JWT_SECRET", "change-this-secret-in-production"),
		JWTExpiry: getEnv("JWT_EXPIRY", "86400"),

		// LMS Integration
		LMSWebhookURL:    getEnv("LMS_WEBHOOK_URL", "http://localhost:50051/webhooks/sims"),
		LMSWebhookSecret: getEnv("LMS_WEBHOOK_SECRET", "webhook-secret"),

		// CORS
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:8080"),

		// Logging
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetDSN returns database connection string
func (c *Config) GetDSN() string {
	return strings.Join([]string{
		"host=" + c.DBHost,
		"port=" + c.DBPort,
		"user=" + c.DBUser,
		"password=" + c.DBPassword,
		"dbname=" + c.DBName,
		"sslmode=" + c.DBSSLMode,
	}, " ")
}
