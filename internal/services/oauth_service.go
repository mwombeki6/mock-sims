package services

import (
	"errors"
	"strings"
	"time"

	"github.com/mwombeki6/mock-sims/internal/config"
	"github.com/mwombeki6/mock-sims/internal/models"
	"github.com/mwombeki6/mock-sims/internal/utils"
	"gorm.io/gorm"
)

type OAuthService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewOAuthService(db *gorm.DB, cfg *config.Config) *OAuthService {
	return &OAuthService{
		db:  db,
		cfg: cfg,
	}
}

// ValidateClient checks if client_id and redirect_uri are valid
func (s *OAuthService) ValidateClient(clientID, redirectURI string) (*models.OAuthClient, error) {
	var client models.OAuthClient
	if err := s.db.Where("client_id = ? AND is_active = ?", clientID, true).First(&client).Error; err != nil {
		return nil, errors.New("invalid client_id")
	}

	// Check if redirect_uri is in allowed list
	if redirectURI == "" {
		return nil, errors.New("redirect_uri is required")
	}

	// Validate redirect URI is in client's allowed list
	allowedURIs := strings.Split(client.RedirectURIs, ",")
	isValid := false
	for _, allowedURI := range allowedURIs {
		if strings.TrimSpace(allowedURI) == redirectURI {
			isValid = true
			break
		}
	}

	if !isValid {
		return nil, errors.New("redirect_uri not allowed for this client")
	}

	return &client, nil
}

// AuthenticateUser validates user credentials and returns user
func (s *OAuthService) AuthenticateUser(username, password string) (*models.User, *models.Student, error) {
	// Username can be registration number or email
	var user models.User

	// Try to find by email first
	if err := s.db.Where("email = ?", username).First(&user).Error; err != nil {
		// If not found by email, try finding student by reg_number
		var student models.Student
		if err := s.db.Preload("User").Where("reg_number = ?", username).First(&student).Error; err != nil {
			return nil, nil, errors.New("invalid credentials")
		}
		user = student.User
	}

	// Verify password
	if !utils.CheckPassword(user.Password, password) {
		return nil, nil, errors.New("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, nil, errors.New("account is inactive")
	}

	// Get student profile if user is a student
	var student *models.Student
	if user.UserType == "student" {
		var st models.Student
		if err := s.db.Where("user_id = ?", user.ID).First(&st).Error; err == nil {
			student = &st
		}
	}

	return &user, student, nil
}

// CreateAuthorizationCode creates a new authorization code
func (s *OAuthService) CreateAuthorizationCode(clientID string, userID uint, redirectURI, scopes string) (string, error) {
	code, err := utils.GenerateAuthCode()
	if err != nil {
		return "", err
	}

	authCode := models.OAuthAuthorizationCode{
		Code:        code,
		ClientID:    clientID,
		UserID:      userID,
		RedirectURI: redirectURI,
		Scopes:      scopes,
		ExpiresAt:   time.Now().Add(10 * time.Minute), // 10 minutes validity
		Used:        false,
	}

	if err := s.db.Create(&authCode).Error; err != nil {
		return "", err
	}

	return code, nil
}

// ExchangeCodeForToken validates authorization code and creates access token
func (s *OAuthService) ExchangeCodeForToken(code, clientID, clientSecret, redirectURI string) (*models.OAuthAccessToken, error) {
	// Validate client credentials
	var client models.OAuthClient
	if err := s.db.Where("client_id = ? AND is_active = ?", clientID, true).First(&client).Error; err != nil {
		return nil, errors.New("invalid client credentials")
	}

	// Verify client secret using bcrypt
	if !utils.CheckPassword(client.ClientSecret, clientSecret) {
		return nil, errors.New("invalid client credentials")
	}

	// Find authorization code
	var authCode models.OAuthAuthorizationCode
	if err := s.db.Where("code = ? AND client_id = ? AND used = ? AND redirect_uri = ?", code, clientID, false, redirectURI).First(&authCode).Error; err != nil {
		return nil, errors.New("invalid or expired authorization code")
	}

	// Check if code is expired
	if authCode.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("authorization code has expired")
	}

	// Mark code as used
	authCode.Used = true
	s.db.Save(&authCode)

	// Generate access token
	accessToken, err := utils.GenerateRandomAccessToken()
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken, err := utils.GenerateRefreshTokenString()
	if err != nil {
		return nil, err
	}

	// Create access token record
	token := models.OAuthAccessToken{
		Token:        accessToken,
		ClientID:     clientID,
		UserID:       authCode.UserID,
		Scopes:       authCode.Scopes,
		ExpiresAt:    time.Now().Add(1 * time.Hour), // 1 hour validity
		RefreshToken: refreshToken,
	}

	if err := s.db.Create(&token).Error; err != nil {
		return nil, err
	}

	return &token, nil
}

// ValidateAccessToken checks if access token is valid
func (s *OAuthService) ValidateAccessToken(token string) (*models.OAuthAccessToken, *models.User, error) {
	var accessToken models.OAuthAccessToken
	if err := s.db.Where("token = ?", token).First(&accessToken).Error; err != nil {
		return nil, nil, errors.New("invalid access token")
	}

	// Check if token is expired
	if accessToken.ExpiresAt.Before(time.Now()) {
		return nil, nil, errors.New("access token has expired")
	}

	// Get user
	var user models.User
	if err := s.db.Where("id = ?", accessToken.UserID).First(&user).Error; err != nil {
		return nil, nil, errors.New("user not found")
	}

	return &accessToken, &user, nil
}

// RefreshAccessToken generates new access token from refresh token
func (s *OAuthService) RefreshAccessToken(refreshToken string) (*models.OAuthAccessToken, error) {
	// Find existing token by refresh token
	var existingToken models.OAuthAccessToken
	if err := s.db.Where("refresh_token = ?", refreshToken).First(&existingToken).Error; err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Generate new access token
	accessToken, err := utils.GenerateRandomAccessToken()
	if err != nil {
		return nil, err
	}

	// Create new access token record
	newToken := models.OAuthAccessToken{
		Token:        accessToken,
		ClientID:     existingToken.ClientID,
		UserID:       existingToken.UserID,
		Scopes:       existingToken.Scopes,
		ExpiresAt:    time.Now().Add(1 * time.Hour),
		RefreshToken: refreshToken, // Keep same refresh token
	}

	if err := s.db.Create(&newToken).Error; err != nil {
		return nil, err
	}

	// Optionally delete old token
	s.db.Delete(&existingToken)

	return &newToken, nil
}
