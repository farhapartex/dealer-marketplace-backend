package services

import (
	"fmt"

	"github.com/farhapartex/dealer-marketplace-be/apps/shop/models"
	"github.com/farhapartex/dealer-marketplace-be/apps/user/dtos"
	userModels "github.com/farhapartex/dealer-marketplace-be/apps/user/models"
	"github.com/farhapartex/dealer-marketplace-be/pkg/database"
	"github.com/google/uuid"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUserByID(userID uuid.UUID) (*dtos.UserMeResponse, error) {
	var user userModels.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	response := &dtos.UserMeResponse{
		UserID:            user.ID.String(),
		Name:              user.Name,
		Email:             user.Email,
		Phone:             user.Mobile,
		UserType:          user.UserType,
		IsOnboardComplete: user.IsOnboardComplete,
	}

	if user.UserType != nil && (*user.UserType == "dealer" || *user.UserType == "sub_dealer") {
		var shop models.Shop
		if err := database.DB.Where("user_id = ?", userID).First(&shop).Error; err == nil {
			response.Shop = &dtos.ShopInfo{
				ID:   shop.ID.String(),
				Name: shop.Name,
			}
		}
	}

	return response, nil
}
