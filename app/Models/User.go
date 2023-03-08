package Models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name    string   `gorm:"not null"`
    Email   string   `gorm:"unique;not null"`
    Address Address  `gorm:"foreignKey:AddressID"`
    AddressID uint
}

func (User) TableName() string {
    return "users"
}
