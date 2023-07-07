package model

import (
	"time"

	"gorm.io/gorm"
)

type MessageType string

const (
	MessageTypeText  MessageType = "text"
	MessageTypeImage MessageType = "image"
	MessageTypeAudio MessageType = "audio"
	MessageTypeVideo MessageType = "video"
	MessageTypeFile  MessageType = "file"
	// 添加其他消息类型...
)

type Message struct {
	gorm.Model
	ConversationID int
	SenderID       string
	Content        string
	Type           MessageType
	ImageURL       string
	AudioURL       string
	VideoURL       string
	FileURL        string
	Timestamp      time.Time
}
