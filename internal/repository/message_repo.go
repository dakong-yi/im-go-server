package repository

import "github.com/dakong-yi/im-go-server/internal/model"

// MessageRepository 定义消息存储库的接口
type MessageRepository interface {
	CreateMessage(message *model.Message) error
	GetMessagesByConversationID(conversationID int) ([]*model.Message, error)
	GetLatestMessageByConversationID(conversationID int) (*model.Message, error)
}
