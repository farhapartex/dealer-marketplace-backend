package dtos

import "github.com/google/uuid"

// CreateShopRequest represents the payload for creating a shop
type CreateShopRequest struct {
	Name          string  `json:"name" binding:"required"`
	ContactNumber *string `json:"contact_number"`
	Address       *string `json:"address"`
}

// CreateShopResponse represents the response after creating a shop
type CreateShopResponse struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Name          string    `json:"name"`
	ContactNumber *string   `json:"contact_number"`
	Address       *string   `json:"address"`
	Message       string    `json:"message"`
}
