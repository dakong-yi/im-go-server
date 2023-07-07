package model

import (
	"gorm.io/gorm"
)

type Blacklist struct {
	gorm.Model
	UserID    string // 用户ID
	BlockedID string // 被拉黑用户ID
}
