package services

import (
	"fmt"
	"time"

	"github.com/farhapartex/dealer-marketplace-be/apps/user/dtos"
	"github.com/farhapartex/dealer-marketplace-be/apps/user/models"
	"github.com/farhapartex/dealer-marketplace-be/config"
	"github.com/farhapartex/dealer-marketplace-be/pkg/database"
	"github.com/farhapartex/dealer-marketplace-be/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateNewUser(payload dtos.CreateUserPayload) (*dtos.CreateUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Mobile:   payload.Phone,
		Password: string(hashedPassword),
		UserType: &payload.UserType,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	code := utils.GenerateVerificationCode()

	expiry, err := parseExpiry(config.AppSettings.AuthCodeExpiry)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expiry: %w", err)
	}

	authVerification := models.AuthVerification{
		Code:     code,
		UserID:   user.ID,
		IsUsed:   false,
		ExpireAt: time.Now().Add(expiry),
	}

	if err := database.DB.Create(&authVerification).Error; err != nil {
		return nil, fmt.Errorf("failed to create auth verification: %w", err)
	}

	token, err := utils.GenerateJWTWithExpiry(user.ID, expiry)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	fmt.Printf("Verification Code: %s\n", code)

	return &dtos.CreateUserResponse{
		Code:  code,
		Token: token,
	}, nil
}

func (s *AuthService) VerifyAuthCode(userID uuid.UUID, code string) error {
	var authVerification models.AuthVerification
	if err := database.DB.Where("user_id = ? AND code = ?", userID, code).First(&authVerification).Error; err != nil {
		return fmt.Errorf("invalid verification code")
	}

	if authVerification.IsUsed {
		return fmt.Errorf("verification code has already been used")
	}

	if time.Now().After(authVerification.ExpireAt) {
		return fmt.Errorf("verification code has expired")
	}

	if err := database.DB.Model(&authVerification).Update("is_used", true).Error; err != nil {
		return fmt.Errorf("failed to update verification status: %w", err)
	}

	if err := database.DB.Model(&models.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"is_active": true,
	}).Error; err != nil {
		return fmt.Errorf("failed to activate user: %w", err)
	}

	return nil
}

func (s *AuthService) Signin(payload dtos.SigninRequest) (*dtos.SigninResponse, error) {
	if payload.Email == nil && payload.Mobile == nil {
		return nil, fmt.Errorf("either email or mobile is required")
	}

	var user models.User
	query := database.DB

	if payload.Email != nil {
		query = query.Where("email = ?", *payload.Email)
	} else if payload.Mobile != nil {
		query = query.Where("mobile = ?", *payload.Mobile)
	}

	if err := query.First(&user).Error; err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if !user.IsActive {
		return nil, fmt.Errorf("account is not active. Please verify your account first")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	now := time.Now()
	if err := database.DB.Model(&user).Update("last_login_at", &now).Error; err != nil {
		return nil, fmt.Errorf("failed to update last login: %w", err)
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &dtos.SigninResponse{
		Token:   token,
		Message: "Login successful",
	}, nil
}

func parseExpiry(expiry string) (time.Duration, error) {
	switch expiry {
	case "1HR":
		return 1 * time.Hour, nil
	case "2HR":
		return 2 * time.Hour, nil
	case "24HR":
		return 24 * time.Hour, nil
	case "30MIN":
		return 30 * time.Minute, nil
	default:
		return 1 * time.Hour, nil
	}
}
