package services

import (
	"fmt"

	"github.com/farhapartex/dealer-marketplace-be/apps/user/dtos"
	"github.com/farhapartex/dealer-marketplace-be/apps/user/models"
	"github.com/farhapartex/dealer-marketplace-be/pkg/database"
	"github.com/google/uuid"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUserByID(userID uuid.UUID) (*dtos.UserMeResponse, error) {
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &dtos.UserMeResponse{
		UserID:            user.ID.String(),
		Name:              user.Name,
		Email:             user.Email,
		Phone:             user.Mobile,
		IsOnboardComplete: user.IsOnboardComplete,
	}, nil
}
