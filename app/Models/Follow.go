package Models

import (
    "gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	FollowerID uint `json:"follower_id" gorm:"not null"`
	FollowingID uint `json:"following_id" gorm:"not null"`
}

func (Follow) TableName() string {
	return "follows"
}