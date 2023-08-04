package model

import (
	"gorm.io/gorm"
)

const (
	CONVERSATION_TYPE_GROUP = "group"
	CONVERSATION_TYPE_C2C   = "c2c"
)

var ConversationTypes = map[string]int{
	CONVERSATION_TYPE_C2C:   1,
	CONVERSATION_TYPE_GROUP: 2,
}

// Conversation 表示对话的模型
type Conversation struct {
	gorm.Model
	OwnerID       string `json:"owner_id"`
	UserID        string `json:"user_id"`
	GroupID       string `json:"group_id"`
	FaceUrl       string `json:"face_url"`
	ShowName      string `json:"show_name"`
	GroupType     string `json:"group_type"`
	UnreadCount   int    `json:"unread_count"`
	LastMessageID string `json:"last_message_id"`
	Type          string `json:"type"`      // 对话类型，可以是 "group" 或 "c2c" 等
	IsPinned      bool   `json:"is_pinned"` // 是否置顶
	RecvOpt       int    `json:"recv_opt"`  // 接收消息的优先级 0:在线正常接收消息，离线时会进行离线推送1:不会接收到消息，离线不会有推送通知 2:在线正常接收消息，离线不会有推送通知
	Orderkey      int64  `json:"orderkey"`  // 排序
	// MarkList              []int    `json:"mark_list"`               // 标记
	// ConversationGroupList []string `json:"conversation_group_list"` // 会话组
	// GroupAtInfoList       []int    `json:"group_at_info_list"`      // 群组@信息
}
