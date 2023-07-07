package model

import (
	"gorm.io/gorm"
)

// UserConversation 表示用户与对话的关联模型
type UserConversation struct {
	gorm.Model
	UserID         string `json:"user_id"`
	ConversationID int    `json:"conversation_id"`
	UnreadCount    int    `json:"unread_count"`
}
