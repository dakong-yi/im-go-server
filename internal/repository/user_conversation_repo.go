package repository

import "github.com/dakong-yi/im-go-server/internal/model"

type UserConversationRepository interface {
	AddUserToConversation(userID string, conversationID uint) error
	RemoveUserFromConversation(userID string, conversationID uint) error
	GetConversationsByUserID(userID string) ([]*model.UserConversation, error)
	GetUnreadMessageCount(userID string) (int, error)
	UpdateUnreadMessageCount(userID string, conversationID, count int) error
	CreateUserConversations(userConversations []*model.UserConversation) error
	GetByConversationID(conversationID int) ([]*model.UserConversation, error)
}
