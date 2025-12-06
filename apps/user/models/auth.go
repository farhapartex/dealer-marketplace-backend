package models

import (
	"time"

	"github.com/google/uuid"
)

type AuthVerification struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Code      string    `json:"code" gorm:"not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	IsUsed    bool      `json:"is_used" gorm:"default:false"`
	ExpireAt  time.Time `json:"expire_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
