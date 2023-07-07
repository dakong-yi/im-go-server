package socket

import (
	"encoding/json"
	"time"

	"github.com/dakong-yi/im-go-server/internal/model"
)

type Action string

const (
	ChatAction   Action = "chat"
	LoginAction  Action = "login"
	SystemAction Action = "system"
)

// MessageProtocol 定义消息协议
type MessageProtocol struct {
	Action Action                 `json:"action"` // 消息类型
	Data   map[string]interface{} `json:"data"`   // 消息数据
}

// Serialize 将消息对象序列化为字节流
func (m *MessageProtocol) Serialize() ([]byte, error) {
	return json.Marshal(m)
}

// Deserialize 将字节流反序列化为消息对象
func Deserialize(data []byte) (*MessageProtocol, error) {
	var msg MessageProtocol
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

// UserLoginMessage 定义用户登录消息
type UserLoginMessage struct {
	UserID string `json:"user_id" mapstructure:"user_id"` // 用户ID
}

// ChatMessage 定义聊天消息
type ChatMessage struct {
	ConversationID int
	SenderID       string
	Content        string
	Type           model.MessageType
	ImageURL       string
	AudioURL       string
	VideoURL       string
	FileURL        string
}

func (c *ChatMessage) ToModelMessage() model.Message {
	return model.Message{
		ConversationID: c.ConversationID,
		SenderID:       c.SenderID,
		Content:        c.Content,
		Type:           c.Type,
		ImageURL:       c.ImageURL,
		AudioURL:       c.AudioURL,
		VideoURL:       c.VideoURL,
		FileURL:        c.FileURL,
		Timestamp:      time.Now(),
	}
}
