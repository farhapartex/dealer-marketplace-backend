package dtos

import "github.com/google/uuid"

// CreateShopRequest represents the payload for creating a shop
type CreateShopRequest struct {
	Name          string  `json:"name" binding:"required"`
	ContactNumber *string `json:"contact_number"`
	Address       *string `json:"address"`
	HouseNo       *string `json:"house_no"`
	Street        *string `json:"street"`
	Town          *string `json:"town"`
	State         *string `json:"state"`
	Country       *string `json:"country"`
}

// CreateShopResponse represents the response after creating a shop
type CreateShopResponse struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Name          string    `json:"name"`
	ContactNumber *string   `json:"contact_number"`
	Address       *string   `json:"address"`
	HouseNo       *string   `json:"house_no"`
	Street        *string   `json:"street"`
	Town          *string   `json:"town"`
	State         *string   `json:"state"`
	Country       *string   `json:"country"`
	Message       string    `json:"message"`
}
