package model

import (
	"gorm.io/gorm"
)

type Friendship struct {
	gorm.Model
	UserID   string // 用户ID
	FriendID string // 好友ID
}
