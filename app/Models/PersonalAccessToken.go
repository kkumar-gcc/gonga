package Models

import (
	"time"

	"gorm.io/gorm"
)

type PersonalAccessToken struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Token string `gorm:"unique;not null"`
    LastUsedAt time.Time
	ExpiresAt time.Time
}

func (PersonalAccessToken) TableName() string {
	return "personal_access_tokens"
}
