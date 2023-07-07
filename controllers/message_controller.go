// controllers/message_controller.go

package controllers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/dakong-yi/im-go-server/models"
// 	"github.com/dakong-yi/im-go-server/services"

// 	"github.com/gin-gonic/gin"
// 	"github.com/pkg/errors"
// )

// type MessageController struct {
// 	messageService *services.MessageService
// }

// func NewMessageController(messageService *services.MessageService) *MessageController {
// 	return &MessageController{messageService: messageService}
// }

// func (c *MessageController) CreateMessage(ctx *gin.Context) {
// 	var message models.Message
// 	err := ctx.ShouldBindJSON(&message)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
// 		return
// 	}

// 	err = c.messageService.Create(&message)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to create message").Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"message": "message created successfully"})
// }

// func (c *MessageController) UpdateMessageStatus(ctx *gin.Context) {
// 	messageID, err := strconv.ParseUint(ctx.Param("message_id"), 10, 64)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid message ID"})
// 		return
// 	}

// 	status, err := strconv.Atoi(ctx.Param("status"))
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
// 		return
// 	}

// 	message, err := c.messageService.FindByID(uint(messageID))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to find message").Error()})
// 		return
// 	}

// 	if message == nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "message not found"})
// 		return
// 	}

// 	err = c.messageService.UpdateStatus(message, models.MessageStatus(rune(status)))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to update message status").Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"message": "message status updated successfully"})
// }

// func (c *MessageController) GetMessagesByConversationID(ctx *gin.Context) {
// 	conversationID, err := strconv.ParseUint(ctx.Param("conversation_id"), 10, 64)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid conversation ID"})
// 		return
// 	}

// 	messages, err := c.messageService.FindByConversationID(uint(conversationID))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to find messages").Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"messages": messages})
// }

// func (c *MessageController) GetMessagesByUserID(ctx *gin.Context) {
// 	userID, err := strconv.ParseUint(ctx.Param("user_id"), 10, 64)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
// 		return
// 	}

// 	messages, err := c.messageService.FindByUserID(uint(userID))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to find messages").Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"messages": messages})
// }
