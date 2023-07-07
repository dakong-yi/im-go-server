package model

import (
	"gorm.io/gorm"
)

const (
	ConversationTypeGroup   = "group"
	ConversationTypePrivate = "private"
)

// Conversation 表示对话的模型
type Conversation struct {
	gorm.Model
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	OwnerID    string `json:"owner_id"`
	Type       string `json:"type"` // 对话类型，可以是 "group" 或 "private" 等
}
