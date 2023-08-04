package repository

import "github.com/dakong-yi/im-go-server/internal/model"

type ConversationRepository interface {
	CreateConversation(conversation *model.Conversation) error
	UpdateLastMessage(conversation *model.Conversation) error
	GetConversationByID(id int) (*model.Conversation, error)
	GetConversationsByIDList(id []int) ([]*model.Conversation, error)
	GetConversationsByOwnerID(ownerID string) ([]*model.Conversation, error)
	GetConversationByUserID(ownerID string, userID string) (*model.Conversation, error)
	GetConversationByGroupID(ownerID string, GroupID string) (*model.Conversation, error)
	// 获取未读消息总数
	GetAllUnreadMessageCount(ownerID string, groupIDs []string) (int, error)
	// 标记消息为已读
	ClearC2CUnreadCount(ownerID, userID string) error
	ClearGroupUnreadCount(ownerID, groupID string) error
}
