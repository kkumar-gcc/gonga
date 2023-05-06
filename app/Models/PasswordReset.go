package Models

import (
    "gorm.io/gorm"
)

type PasswordReset struct {
    gorm.Model
    Email    string `gorm:"unique;not null"`
    Token    string `gorm:"not null"`
    Expiry   int64  `gorm:"not null"`
}

func (PasswordReset) TableName() string {
    return "password_resets"
}