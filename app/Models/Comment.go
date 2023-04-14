package Models

import (
    "gorm.io/gorm"
)

type Comment struct {
    gorm.Model
    UserID    uint      `json:"-"`
    User      User      `json:"user" gorm:"foreignKey:UserID"`
    PostID    uint      `json:"-"`
    Post      Post      `json:"post" gorm:"foreignKey:PostID"`
    Text      string    `json:"text"`
    Likes     []Like    `json:"-" gorm:"foreignKey:CommentID"`
}

func (Comment) TableName() string {
    return "comments"
}