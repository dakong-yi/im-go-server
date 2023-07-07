// services/message_service.go

package services

// import (
// 	"github.com/dakong-yi/im-go-server/models"
// 	"github.com/dakong-yi/im-go-server/repositories"

// 	"github.com/pkg/errors"
// )

// type MessageService struct {
// 	messageRepo *repositories.MessageRepository
// }

// func NewMessageService(messageRepo *repositories.MessageRepository) *MessageService {
// 	return &MessageService{messageRepo: messageRepo}
// }

// func (s *MessageService) Create(message *models.Message) error {
// 	err := s.messageRepo.Create(message)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to create message")
// 	}

// 	return nil
// }

// func (s *MessageService) UpdateStatus(message *models.Message, status models.MessageStatus) error {
// 	err := s.messageRepo.UpdateStatus(message, status)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to update message status")
// 	}

// 	return nil
// }

// func (s *MessageService) FindByID(id uint) (*models.Message, error) {
// 	message, err := s.messageRepo.FindByID(id)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to find message by ID")
// 	}

// 	return message, nil
// }

// func (s *MessageService) FindByConversationID(conversationID uint) ([]*models.Message, error) {
// 	messages, err := s.messageRepo.FindByConversationID(conversationID)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to find messages by conversation ID")
// 	}

// 	return messages, nil
// }

// func (s *MessageService) FindByUserID(userID uint) ([]*models.Message, error) {
// 	messages, err := s.messageRepo.FindByUserID(userID)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to find messages by user ID")
// 	}

// 	return messages, nil
// }
