package repository

import (
	"github.com/dakong-yi/im-go-server/internal/db"
	"github.com/dakong-yi/im-go-server/internal/model"
)

type ConversationRepoImpl struct {
}

func NewConversationRepoImpl() ConversationRepository {
	return &ConversationRepoImpl{}
}

func (r *ConversationRepoImpl) CreateConversation(conversation *model.Conversation) error {
	return db.DB.Create(conversation).Error
}

func (r *ConversationRepoImpl) GetConversationByID(id int) (*model.Conversation, error) {
	var conversation model.Conversation
	if err := db.DB.Where("id = ?", id).First(&conversation).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *ConversationRepoImpl) GetConversationByIdentifier(identifier string) (*model.Conversation, error) {
	var conversation model.Conversation
	if err := db.DB.Where("identifier = ?", identifier).First(&conversation).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *ConversationRepoImpl) GetConversationsByIDList(ids []int) ([]*model.Conversation, error) {
	var conversations []*model.Conversation
	if err := db.DB.Where("id IN (?)", ids).Find(&conversations).Error; err != nil {
		return nil, err
	}
	return conversations, nil
}
