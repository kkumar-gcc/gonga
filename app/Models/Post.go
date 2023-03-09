package Models

import (
    "gorm.io/gorm"
)

type Post struct {
    gorm.Model
    Name    string   `gorm:"not null"`
    Email   string   `gorm:"unique;not null"`
}

func (Post) TableName() string {
    return "posts"
}