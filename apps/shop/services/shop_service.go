package services

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/farhapartex/dealer-marketplace-be/apps/shop/dtos"
	"github.com/farhapartex/dealer-marketplace-be/apps/shop/models"
	userModels "github.com/farhapartex/dealer-marketplace-be/apps/user/models"
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

	// Generate full address from individual fields
	var addressParts []string

	if req.HouseNo != nil && *req.HouseNo != "" {
		addressParts = append(addressParts, *req.HouseNo)
	}
	if req.Street != nil && *req.Street != "" {
		addressParts = append(addressParts, *req.Street)
	}
	if req.Town != nil && *req.Town != "" {
		addressParts = append(addressParts, *req.Town)
	}
	if req.State != nil && *req.State != "" {
		addressParts = append(addressParts, *req.State)
	}
	if req.Country != nil && *req.Country != "" {
		addressParts = append(addressParts, *req.Country)
	}

	var addressPtr *string
	if len(addressParts) > 0 {
		fullAddress := strings.Join(addressParts, ", ")
		addressPtr = &fullAddress
	}

	// Create new shop
	shop := models.Shop{
		UserID:        userID,
		Name:          req.Name,
		ContactNumber: req.ContactNumber,
		Address:       addressPtr,
		HouseNo:       req.HouseNo,
		Street:        req.Street,
		Town:          req.Town,
		State:         req.State,
		Country:       req.Country,
	}

	if err := database.DB.Create(&shop).Error; err != nil {
		return nil, err
	}

	// Update user's IsOnboardComplete to true
	if err := database.DB.Model(&userModels.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"is_onboard_complete": true,
	}).Error; err != nil {
		return nil, err
	}

	// Build response
	response := &dtos.CreateShopResponse{
		ID:            shop.ID,
		UserID:        shop.UserID,
		Name:          shop.Name,
		ContactNumber: shop.ContactNumber,
		HouseNo:       shop.HouseNo,
		Street:        shop.Street,
		Town:          shop.Town,
		State:         shop.State,
		Country:       shop.Country,
		Message:       "Shop created successfully",
	}

	return response, nil
}
