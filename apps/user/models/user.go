package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                 uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name               string     `json:"name" gorm:"not null"`
	Email              *string    `json:"email" gorm:"unique"`
	Mobile             *string    `json:"mobile"`
	Password           string     `json:"-" gorm:"not null"`
	IsActive           bool       `json:"is_active" gorm:"default:false"`
	IsOnboardComplete  bool       `json:"is_onboard_complete" gorm:"default:false"`
	UserType           *string    `json:"user_type"`
	LastLoginAt        *time.Time `json:"last_login_at"`
	CreatedAt          time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
