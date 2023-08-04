package controller

import (
	"net/http"
	"strings"

	"github.com/dakong-yi/im-go-server/internal/dto/request"
	"github.com/dakong-yi/im-go-server/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	// conversation, err := ctrl.conversationService.CreateConversation(request.OwnerID, request.UserID, request.GroupID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create conversation"})
	// 	return
	// }

	// 返回创建成功的响应
	c.JSON(http.StatusOK, gin.H{"message": "Conversation created successfully", "conversation": nil})
}

// 获取对话信息
func (ctrl *ConversationController) GetConversation(c *gin.Context) {
	// 从请求中获取创建对话的信息
	conversationID := c.Param("conversationID")
	ownerID := c.Param("ownerID")
	arr := strings.Split(conversationID, "_")
	if len(arr) != 2 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get conversation"})
		return
	}
	userID := arr[1]
	groupID := arr[1]
	conversationType := arr[0]
	// 调用服务层方法获取对话信息
	conversation, err := ctrl.conversationService.GetConversationByID(ownerID, userID, groupID, conversationType)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{"code": 0, "desc": "ok", "data": conversation})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get conversation"})
		return
	}

	// 返回对话信息的响应
	c.JSON(http.StatusOK, gin.H{"code": 0, "desc": "ok", "data": conversation})
}
func (ctrl *ConversationController) GetConversationsByUserID(c *gin.Context) {
	userID := c.Param("userID")

	conversations, err := ctrl.conversationService.GetConversationsByOwnerID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve conversations",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"desc": "ok",
		"data": conversations,
	})
}
func (ctrl *ConversationController) GetTotalUnreadMessageCount(c *gin.Context) {
	ownerID := c.Param("ownerID")
	count, err := ctrl.conversationService.GetAllUnreadMessageCount(ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve GetTotalUnreadMessageCount",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"desc": "ok",
		"data": count,
	})
}

// 其他控制器方法...
