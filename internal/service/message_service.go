package service

import (
	"time"

	"github.com/dakong-yi/im-go-server/internal/dto/request"
	"github.com/dakong-yi/im-go-server/internal/model"
	"github.com/dakong-yi/im-go-server/internal/repository"
	"github.com/google/wire"
)

type MessageService struct {
	messageRepo repository.MessageRepository
}

func NewMessageService(messageRepo repository.MessageRepository) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
	}
}
func InitMessageService() *MessageService {
	wire.Build(NewMessageService)
	return nil
}

func (service *MessageService) CreateMessage(req request.CreateMessageRequest) (*model.Message, error) {
	message := &model.Message{
		ConversationID: req.ConversationID,
		SenderID:       req.SenderID,
		Content:        req.Content,
		Type:           model.MessageType(req.Type),
		ImageURL:       req.ImageURL,
		AudioURL:       req.AudioURL,
		VideoURL:       req.VideoURL,
		FileURL:        req.FileURL,
		Timestamp:      time.Now(),
	}
	err := service.messageRepo.CreateMessage(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (service *MessageService) GetMessagesByConversationID(conversationID int) ([]*model.Message, error) {
	messages, err := service.messageRepo.GetMessagesByConversationID(conversationID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
