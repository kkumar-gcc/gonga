package Models

import (
    "gorm.io/gorm"
)

type Like struct {
    gorm.Model
    UserID    uint       `json:"-"`
    User      User       `json:"user" gorm:"foreignKey:UserID"`
    PostID    uint       `json:"-"`
    Post      Post       `json:"post" gorm:"foreignKey:PostID"`
    CommentID uint       `json:"-"`
    Comment   Comment    `json:"comment" gorm:"foreignKey:CommentID"`
}

func (Like) TableName() string {
    return "likes"
}