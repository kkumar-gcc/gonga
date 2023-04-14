package Models

import (
    "gorm.io/gorm"
)

type Follow struct {
    gorm.Model
    FollowerID uint      `json:"-"`
    Follower   User      `json:"follower" gorm:"foreignKey:FollowerID"`
    FollowingID uint     `json:"-"`
    Following   User     `json:"following" gorm:"foreignKey:FollowingID"`
}

func (Follow) TableName() string {
    return "follows"
}