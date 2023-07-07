package repository

import (
	"github.com/dakong-yi/im-go-server/internal/db"
	"github.com/dakong-yi/im-go-server/internal/model"
	"gorm.io/gorm"
)

type UserConversationRepoImpl struct {
}

func NewUserConversationRepoImpl() UserConversationRepository {
	return &UserConversationRepoImpl{}
}

func (r *UserConversationRepoImpl) AddUserToConversation(userID string, conversationID uint) error {
	userConversation := &model.UserConversation{
		UserID:         userID,
		ConversationID: int(conversationID),
	}

	err := db.DB.Create(userConversation).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserConversationRepoImpl) RemoveUserFromConversation(userID string, conversationID uint) error {
	err := db.DB.Where("user_id = ? AND conversation_id = ?", userID, conversationID).Delete(model.UserConversation{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserConversationRepoImpl) GetConversationsByUserID(userID string) ([]int, error) {
	var conversationIDs []int

	// 查询用户对话列表的 conversation_id
	rows, err := db.DB.
		Select("user_conversations.conversation_id").
		Where("user_conversations.user_id = ?", userID).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var conversationID int
		if err := rows.Scan(&conversationID); err != nil {
			return nil, err
		}
		conversationIDs = append(conversationIDs, conversationID)
	}

	return conversationIDs, nil
}

func (r *UserConversationRepoImpl) GetUnreadMessageCount(userID string) (int, error) {
	var count int64
	err := db.DB.
		Model(&model.UserConversation{}).
		Where("user_id = ? AND unread_count > 0", userID).
		Count(&count).
		Error

	return int(count), err
}

func (r *UserConversationRepoImpl) UpdateUnreadMessageCount(userID string, conversationID int, count int) error {
	return db.DB.
		Model(&model.UserConversation{}).
		Where("user_id = ? AND conversation_id = ?", userID, conversationID).
		Update("unread_count", count).
		Error
}

func (r *UserConversationRepoImpl) CreateUserConversations(userConversations []*model.UserConversation) error {
	// 使用事务执行批量创建操作
	return db.DB.Transaction(func(tx *gorm.DB) error {
		for _, userConv := range userConversations {
			if err := tx.Create(userConv).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *UserConversationRepoImpl) GetByConversationID(conversationID int) ([]*model.UserConversation, error) {
	var userConversations []*model.UserConversation
	if err := db.DB.Where("conversation_id = ?", conversationID).Find(&userConversations).Error; err != nil {
		return nil, err
	}
	return userConversations, nil
}
