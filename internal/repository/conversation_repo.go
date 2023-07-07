package repository

import "github.com/dakong-yi/im-go-server/internal/model"

type ConversationRepository interface {
	CreateConversation(conversation *model.Conversation) error
	GetConversationByID(id int) (*model.Conversation, error)
	GetConversationsByIDList(id []int) ([]*model.Conversation, error)
	GetConversationByIdentifier(identifier string) (*model.Conversation, error)
	// 其他方法...
}
