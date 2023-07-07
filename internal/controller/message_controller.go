package controller

import (
	"net/http"
	"strconv"

	"github.com/dakong-yi/im-go-server/internal/dto/request"
	"github.com/dakong-yi/im-go-server/internal/repository"
	"github.com/dakong-yi/im-go-server/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type MessageController struct {
	messageService *service.MessageService
}

func NewMessageController(messageService *service.MessageService) *MessageController {
	return &MessageController{
		messageService: messageService,
	}
}

//go:generate wire
func initializeMessageController() (*MessageController, error) {
	wire.Build(
		NewMessageController,
		service.NewMessageService,
		repository.NewMessageRepoImpl,
	)
	return nil, nil
}

// 创建消息
func (controller *MessageController) CreateMessage(c *gin.Context) {
	var request request.CreateMessageRequest
	// var request struct {
	// 	ConversationID int    `json:"conversation_id"`
	// 	SenderID       string `json:"sender_id"`
	// 	Content        string `json:"content"`
	// 	Type           string `json:"type"`
	// }

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := controller.messageService.CreateMessage(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

// 获取对话消息列表
func (controller *MessageController) GetMessagesByConversationID(c *gin.Context) {
	conversationIDStr := c.Param("conversationID")

	conversationID, err := strconv.Atoi(conversationIDStr)
	messages, err := controller.messageService.GetMessagesByConversationID(conversationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
