package controller

import (
	"net/http"
	"strconv"

	"github.com/dakong-yi/im-go-server/internal/dto/request"
	"github.com/dakong-yi/im-go-server/internal/service"
	"github.com/gin-gonic/gin"
)

type ConversationController struct {
	conversationService *service.ConversationService
}

func NewConversationController(conversationService *service.ConversationService) *ConversationController {
	return &ConversationController{
		conversationService: conversationService,
	}
}

// 创建对话
func (ctrl *ConversationController) CreateConversation(c *gin.Context) {
	// 从请求中获取创建对话的信息
	var request request.CreateConversationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 调用服务层方法创建对话
	conversation, err := ctrl.conversationService.CreateConversation(request.Name, request.OwnerID, request.UserIDs, request.ConversationType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create conversation"})
		return
	}

	// 返回创建成功的响应
	c.JSON(http.StatusOK, gin.H{"message": "Conversation created successfully", "conversation": conversation})
}

// 获取对话信息
func (ctrl *ConversationController) GetConversation(c *gin.Context) {
	// 从请求中获取对话ID
	conversationIDStr := c.Param("id")

	// 将对话ID转换为整数类型
	conversationID, err := strconv.Atoi(conversationIDStr)

	// 调用服务层方法获取对话信息
	conversation, err := ctrl.conversationService.GetConversationByID(conversationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get conversation"})
		return
	}

	// 返回对话信息的响应
	c.JSON(http.StatusOK, gin.H{"conversation": conversation})
}
func (ctrl *ConversationController) GetConversationsByUserID(c *gin.Context) {
	userID := c.Param("userID")

	conversations, err := ctrl.conversationService.GetConversationsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve conversations",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"conversations": conversations,
	})
}

// 其他控制器方法...
