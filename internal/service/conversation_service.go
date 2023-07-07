package service

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"github.com/dakong-yi/im-go-server/internal/db"
	"github.com/dakong-yi/im-go-server/internal/model"
	"github.com/dakong-yi/im-go-server/internal/repository"
	"github.com/google/wire"
)

type ConversationService struct {
	conversationRepo     repository.ConversationRepository
	userConversationRepo repository.UserConversationRepository
}

func NewConversationService(conversationRepo repository.ConversationRepository,
	userConversationRepo repository.UserConversationRepository) *ConversationService {
	return &ConversationService{
		conversationRepo:     conversationRepo,
		userConversationRepo: userConversationRepo,
	}
}
func InitConversationService() *ConversationService {
	wire.Build(NewConversationService)
	return nil
}
func calculateIdentifier(value string) string {
	// 使用哈希算法计算 value 的哈希值作为 identifier
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}

func (service *ConversationService) CreateConversation(name string, ownerID string, userIDs []string, conversationType string) (*model.Conversation, error) {
	// 如果是私聊对话，则判断 identifier 是否已存在
	var identifier string
	if conversationType == model.ConversationTypePrivate {
		// 将 userIDs 拼接在一起
		concatenatedIDs := strings.Join(userIDs, "")

		// 使用加密算法计算 identifier
		identifier := calculateIdentifier(concatenatedIDs)

		// 检查 identifier 是否已存在
		existingConversation, _ := service.conversationRepo.GetConversationByIdentifier(identifier)

		// 如果存在，则直接返回对话
		if existingConversation != nil {
			return existingConversation, nil
		}
	}

	// 开始事务
	tx := db.DB.Begin()

	// 创建对话
	conversation := &model.Conversation{
		Type:    conversationType,
		OwnerID: ownerID,
	}

	// 设置对话的 identifier
	if conversationType == model.ConversationTypePrivate {
		conversation.Identifier = identifier
	}

	if err := tx.Create(conversation).Error; err != nil {
		tx.Rollback() // 回滚事务
		return nil, err
	}

	// 创建用户对话关联关系
	for _, uid := range userIDs {
		userConversation := &model.UserConversation{
			ConversationID: int(conversation.ID),
			UserID:         uid,
			// 设置其他属性
		}

		if err := tx.Create(userConversation).Error; err != nil {
			tx.Rollback() // 回滚事务
			return nil, err
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // 回滚事务
		return nil, err
	}

	return conversation, nil
}

func (service *ConversationService) GetConversationByID(conversationID int) (*model.Conversation, error) {
	conversation, err := service.conversationRepo.GetConversationByID(conversationID)
	if err != nil {
		return nil, err
	}
	return conversation, nil
}

func (service *ConversationService) GetUsersByConversationID(conversationID int) ([]string, error) {
	userConversations, err := service.userConversationRepo.GetByConversationID(conversationID)
	if err != nil {
		return nil, err
	}

	userIDs := make([]string, len(userConversations))
	for i, uc := range userConversations {
		userIDs[i] = uc.UserID
	}

	return userIDs, nil
}

func (service *ConversationService) GetConversationsByUserID(userID string) ([]*model.Conversation, error) {
	conversationIDs, err := service.userConversationRepo.GetConversationsByUserID(userID)
	if err != nil {
		return nil, err
	}
	conversations, err := service.conversationRepo.GetConversationsByIDList(conversationIDs)
	if err != nil {
		return nil, err
	}
	return conversations, nil
}
