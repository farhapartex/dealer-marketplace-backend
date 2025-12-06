package models

import (
	"time"

	"github.com/google/uuid"
)

type Shop struct {
	ID            uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name          string     `json:"name" gorm:"not null"`
	ContactNumber *string    `json:"contact_number"`
	Logo          *string    `json:"logo"`
	Address       *string    `json:"address"`
	CreatedAt     time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (Shop) TableName() string {
	return "shop"
}
