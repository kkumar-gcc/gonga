package Models

import (
    "gorm.io/gorm"
)

type Post struct {
    gorm.Model
    UserID    uint       `json:"-"`
    User      User       `json:"user" gorm:"foreignKey:UserID"`
    Text      string     `json:"text"`
    ImageURL  string     `json:"image_url"`
    Likes     []Like     `json:"-" gorm:"foreignKey:PostID"`
    Comments  []Comment  `json:"-" gorm:"foreignKey:PostID"`
}

func (Post) TableName() string {
    return "posts"
}