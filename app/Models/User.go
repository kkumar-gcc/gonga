package Models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string     `gorm:"uniqueIndex:idx_username_length;not null;unique"`
	Email     string     `gorm:"unique;not null"`
	Password  string     `gorm:"not null"`
	FirstName string     `gorm:"not null"`
	LastName  string     `gorm:"not null"`
	AvatarURL string     `json:"avatar_url"`
	Bio       string     `json:"bio"`
	Gender    string     `json:"gender"`
	Birthday  *time.Time `json:"birthday"`
	Country   string     `json:"country"`
	City      string     `json:"city"`
	Posts     []Post     `json:"-" gorm:"foreignKey:UserID"`
	Comments  []Comment  `json:"-" gorm:"foreignKey:UserID"`
	Followers []User     `json:"-" gorm:"many2many:user_followers;"`
	Following []User     `json:"-" gorm:"many2many:user_following;"`
}

func (User) TableName() string {
	return "users"
}
