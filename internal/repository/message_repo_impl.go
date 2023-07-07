package repository

import (
	"github.com/dakong-yi/im-go-server/internal/db"
	"github.com/dakong-yi/im-go-server/internal/model"
)

type MessageRepoImpl struct {
}

func NewMessageRepoImpl() MessageRepository {
	return &MessageRepoImpl{}
}

func (r *MessageRepoImpl) CreateMessage(message *model.Message) error {
	return db.DB.Create(message).Error
}

func (r *MessageRepoImpl) GetMessagesByConversationID(conversationID int) ([]*model.Message, error) {
	var messages []*model.Message
	err := db.DB.
		Where("conversation_id = ?", conversationID).
		Order("created_at ASC").
		Find(&messages).
		Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepoImpl) GetLatestMessageByConversationID(conversationID int) (*model.Message, error) {
	var message model.Message
	err := db.DB.
		Where("conversation_id = ?", conversationID).
		Order("created_at DESC").
		Limit(1).
		First(&message).
		Error

	if err != nil {
		return nil, err
	}

	return &message, nil
}
