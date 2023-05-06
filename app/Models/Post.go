package Models

import (
    "gorm.io/gorm"
)

type Post struct {
    gorm.Model
    UserID    uint       `json:"-"`
    User      User       `json:"user" gorm:"foreignKey:UserID"`
    Body      string     `json:"body"`
    MediaURL  string     `json:"media_url"`
    Likes     []Like     `json:"-" gorm:"foreignKey:PostID"`
    Comments  []Comment  `json:"-" gorm:"foreignKey:PostID"`
}

func (Post) TableName() string {
    return "posts"
}