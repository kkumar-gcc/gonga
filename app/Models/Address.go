package Models

import (
    "gorm.io/gorm"
)

type Address struct {
    gorm.Model
    Street  string `gorm:"not null"`
    City    string `gorm:"not null"`
    State   string `gorm:"not null"`
    ZipCode string `gorm:"not null"`
}

func (Address) TableName() string {
    return "addresses"
}
