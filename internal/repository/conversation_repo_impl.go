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

func (r *ConversationRepoImpl) UpdateLastMessage(conversation *model.Conversation) error {
	return db.DB.Model(conversation).Updates(conversation).Error
}

func (r *ConversationRepoImpl) ClearC2CUnreadCount(ownerID, userID string) error {
	return db.DB.Model(&model.Conversation{}).Where("owner_id = ? AND user_id = ?", ownerID, userID).Update("unread_count", 0).Error
}
func (r *ConversationRepoImpl) ClearGroupUnreadCount(ownerID, groupID string) error {
	return db.DB.Model(&model.Conversation{}).Where("owner_id = ? AND group_id = ?", ownerID, groupID).Update("unread_count", 0).Error
}

func (r *ConversationRepoImpl) GetConversationByID(id int) (*model.Conversation, error) {
	var conversation model.Conversation
	if err := db.DB.Where("id = ?", id).First(&conversation).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *ConversationRepoImpl) GetConversationByUserID(owner_id, user_id string) (*model.Conversation, error) {
	var conversation model.Conversation
	if err := db.DB.Where("owner_id = ? AND user_id = ?", owner_id, user_id).First(&conversation).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}
func (r *ConversationRepoImpl) GetConversationByGroupID(owner_id, group_id string) (*model.Conversation, error) {
	var conversation model.Conversation
	if err := db.DB.Table("conversations").Where("owner_id = ? AND group_id = ?", owner_id, group_id).First(&conversation).Error; err != nil {
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
func (r *ConversationRepoImpl) GetConversationsByOwnerID(ownerID string) ([]*model.Conversation, error) {
	var conversations []*model.Conversation
	if err := db.DB.Where("owner_id = ?", ownerID).Find(&conversations).Error; err != nil {
		return nil, err
	}
	return conversations, nil
}

func (r *ConversationRepoImpl) GetAllUnreadMessageCount(ownerID string, groupIDs []string) (int, error) {
	var count int64
	tx := db.DB.Model(&model.Conversation{})
	if len(groupIDs) > 0 {
		tx = tx.Where("owner_id = ? or group_id in ?", ownerID, groupIDs)
	} else {
		tx = tx.Where("owner_id = ?", ownerID)
	}
	err := tx.Select("SUM(unread_count)").Scan(&count).Error
	return int(count), err
}
