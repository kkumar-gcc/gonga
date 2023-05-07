package Models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Title       string `json:"title" gorm:"unique"`
	CoverImage  string `json:"coverImage"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Slug        string `json:"slug"`
	UserID      uint   `json:"-"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
	Posts       []Post `json:"posts" gorm:"many2many:post_hashtags;"`
	// Interests []*Interest `json:"interests" gorm:"many2many:interest_tags;"`
}

func (Tag) TableName() string {
	return "tags"
}
