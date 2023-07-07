package services

// import (
// 	"github.com/dakong-yi/im-go-server/models"
// 	"github.com/dakong-yi/im-go-server/repositories"
// )

// type ConversationService struct {
// 	convRepo *repositories.ConversationRepository
// }

// func NewConversationService(convRepo *repositories.ConversationRepository) *ConversationService {
// 	return &ConversationService{convRepo}
// }

// func (s *ConversationService) CreateConversation(name string) (*models.Conversation, error) {
// 	return s.convRepo.CreateConversation(name)
// }

// func (s *ConversationService) GetConversationByID(id uint) (*models.Conversation, error) {
// 	return s.convRepo.GetConversationByID(id)
// }

// func (s *ConversationService) UpdateConversation(conversation *models.Conversation) error {
// 	return s.convRepo.UpdateConversation(conversation)
// }

// func (s *ConversationService) DeleteConversation(id uint) error {
// 	return s.convRepo.DeleteConversation(id)
// }
