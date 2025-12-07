package services

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/farhapartex/dealer-marketplace-be/apps/shop/dtos"
	"github.com/farhapartex/dealer-marketplace-be/apps/shop/models"
	"github.com/farhapartex/dealer-marketplace-be/pkg/database"
)

type ShopService struct{}

func NewShopService() *ShopService {
	return &ShopService{}
}

// CreateShop creates a new shop for the given user
func (s *ShopService) CreateShop(userID uuid.UUID, req *dtos.CreateShopRequest) (*dtos.CreateShopResponse, error) {
	// Check if user already has a shop
	var existingShop models.Shop
	err := database.DB.Where("user_id = ?", userID).First(&existingShop).Error
	if err == nil {
		return nil, errors.New("user already has a shop")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Create new shop
	shop := models.Shop{
		UserID:        userID,
		Name:          req.Name,
		ContactNumber: req.ContactNumber,
		Address:       req.Address,
	}

	if err := database.DB.Create(&shop).Error; err != nil {
		return nil, err
	}

	// Build response
	response := &dtos.CreateShopResponse{
		ID:            shop.ID,
		UserID:        shop.UserID,
		Name:          shop.Name,
		ContactNumber: shop.ContactNumber,
		Address:       shop.Address,
		Message:       "Shop created successfully",
	}

	return response, nil
}
