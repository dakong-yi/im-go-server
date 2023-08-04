package controller

import (
	"net/http"

	"github.com/dakong-yi/im-go-server/internal/dto"
	"github.com/dakong-yi/im-go-server/internal/dto/request"
	"github.com/dakong-yi/im-go-server/internal/dto/response"
	"github.com/dakong-yi/im-go-server/internal/service"
	"github.com/dakong-yi/im-go-server/internal/socket"
	"github.com/gin-gonic/gin"
)

type MessageController struct {
	messageService      *service.MessageService
	socketService       *socket.SocketServer
	conversationService *service.ConversationService
}

func NewMessageController(messageService *service.MessageService, socketService *socket.SocketServer, conversationService *service.ConversationService) *MessageController {
	return &MessageController{
		messageService:      messageService,
		socketService:       socketService,
		conversationService: conversationService,
	}
}

// 创建消息
func (ctrl *MessageController) CreateMessage(c *gin.Context) {
	var request request.CreateTextMessageRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := ctrl.messageService.CreateMessage(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	v2TimMsgCreateInfoResult := response.V2TimMsgCreateInfoResult{
		ID:          message.ID,
		MessageInfo: message,
	}

	c.JSON(http.StatusOK, gin.H{"data": v2TimMsgCreateInfoResult, "code": 0, "desc": "success"})
}

// 发送消息
func (ctrl *MessageController) SendMessage(c *gin.Context) {
	var request request.SendMessageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	v2TimMessage, err := ctrl.messageService.UpdateMessage(request.ID, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctrl.socketService.SendMessage(v2TimMessage)

	ctrl.conversationService.CreateConversation(v2TimMessage)
	ctrl.conversationService.OnTotalUnreadMessageCountChanged(request.Receiver)
	c.JSON(http.StatusOK, gin.H{"data": v2TimMessage, "code": 0, "desc": "success"})
}

// 获取对话消息列表
func (ctrl *MessageController) GetHistoryMessageList(c *gin.Context) {
	var request request.GetHistoryMessageListRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messages, err := ctrl.messageService.GetHistoryMessageList(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// result := &response.V2TimMessageListResult{
	// 	IsFinished:  true,
	// 	MessageList: messages,
	// }
	c.JSON(http.StatusOK, gin.H{"code": 0, "desc": "success", "data": messages})
}

func (ctrl *MessageController) GetHistoryMessageListV2(c *gin.Context) {
	var request request.GetHistoryMessageListRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messages, err := ctrl.messageService.GetHistoryMessageList(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	isFinished := false
	if len(messages) < request.Count || len(messages) == 0 {
		isFinished = true
	}
	result := &response.V2TimMessageListResult{
		IsFinished:  isFinished,
		MessageList: messages,
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "desc": "success", "data": result})
}

func (ctrl *MessageController) GetMessageByMsgID(c *gin.Context) {
	// conversationIDStr := c.Param("conversationID")

	// conversationID, err := strconv.Atoi(conversationIDStr)

	// messages, err := ctrl.messageService.GetMessagesByConversationID(conversationID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	mockMessage := &dto.V2TimMessage{
		MsgID:     "1",
		Timestamp: 1635993602,
		UserID:    "user2",
		Sender:    "user1",
		IsSelf:    true,
		NickName:  "John",
		FaceUrl:   "https://www.bugela.com/cjpic/frombd/1/253/1943132031/773911012.jpg",
		ElemType:  1,
		TextElem: &dto.V2TimTextElem{
			Text: "Hello, this is a mock text message",
		},
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "desc": "", "data": mockMessage})
}

func (ctrl *MessageController) MarkC2CMessageAsRead(c *gin.Context) {
	var request request.MarkC2CMessageAsReadRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.messageService.MarkC2CMessageAsRead(request.OwnerID, request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctrl.conversationService.OnConversationChanged(request.OwnerID)
	ctrl.conversationService.OnTotalUnreadMessageCountChanged(request.OwnerID)
	c.JSON(http.StatusOK, gin.H{"code": 0, "desc": "", "data": ""})
}
