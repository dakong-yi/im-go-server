package service

import (
	"strconv"
	"time"

	"github.com/dakong-yi/im-go-server/internal/dto"
	"github.com/dakong-yi/im-go-server/internal/dto/request"
	"github.com/dakong-yi/im-go-server/internal/model"
	"github.com/dakong-yi/im-go-server/internal/repository"
	"github.com/google/uuid"
)

type MessageService struct {
	messageRepo      repository.MessageRepository
	userService      *UserService
	conversationRepo repository.ConversationRepository
}

func NewMessageService(messageRepo repository.MessageRepository, userService *UserService, conversationRepo repository.ConversationRepository) *MessageService {
	return &MessageService{
		messageRepo:      messageRepo,
		userService:      userService,
		conversationRepo: conversationRepo,
	}
}

func (service *MessageService) CreateMessage(req request.CreateTextMessageRequest) (*dto.V2TimMessage, error) {
	message := &model.Message{
		Sender:    req.Sender,
		Text:      req.Text,
		ElemType:  model.TEXT_MESSAGE,
		ID:        uuid.New().String(),
		Timestamp: time.Now().Unix(),
	}
	err := service.messageRepo.CreateMessage(message)
	if err != nil {
		return nil, err
	}
	vMessage, err := service.messageToV2TimMessage(message)
	return vMessage, err
}

func (service *MessageService) messageToV2TimMessage(message *model.Message) (*dto.V2TimMessage, error) {
	vMessage := &dto.V2TimMessage{
		Timestamp: message.Timestamp,
		UserID:    message.UserID,
		Sender:    message.Sender,
		IsSelf:    message.IsSelf,
		IsRead:    message.IsRead,
		ElemType:  int(message.ElemType),
		TextElem: &dto.V2TimTextElem{
			Text: message.Text,
		},
		ID:    message.ID,
		MsgID: message.MsgID,
		Seq:   message.Seq,
	}
	if message.Sender != "" {
		userInfos, err := service.userService.GetUserInfo([]string{message.Sender})
		if err != nil {
			return nil, err
		}
		vMessage.NickName = userInfos[0].NickName
		vMessage.FaceUrl = userInfos[0].FaceURL
	}
	return vMessage, nil
}

func (service *MessageService) GetMessage(id string) (*dto.V2TimMessage, error) {
	message, err := service.messageRepo.GetMessage(id)
	if err != nil {
		return nil, err
	}
	vMessage, err := service.messageToV2TimMessage(message)
	return vMessage, err
}
func (service *MessageService) UpdateMessage(id string, req *request.SendMessageRequest) (*dto.V2TimMessage, error) {
	message, err := service.messageRepo.GetMessage(id)
	if err != nil {
		return nil, err
	}
	message.Sender = req.Sender
	message.UserID = req.Receiver
	message.GroupID = req.GroupID
	message.IsSelf = true
	message.Priority = model.PRIORITY[req.Priority]
	message.NeedReadReceipt = req.NeedReadReceipt
	message.IsExcludedFromLastMessage = req.IsExcludedFromLastMessage
	message.IsExcludedFromUnreadCount = req.IsExcludedFromUnreadCount
	message.CloudCustomData = req.CloudCustomData
	message.MsgID = strconv.FormatInt(time.Now().UnixNano(), 10)
	message, err = service.messageRepo.UpdateMessage(id, message)
	if err != nil {
		return nil, err
	}
	vMessage, err := service.messageToV2TimMessage(message)
	return vMessage, err
}

func (service *MessageService) GetHistoryMessageList(req request.GetHistoryMessageListRequest) ([]*dto.V2TimMessage, error) {
	var messages []*model.Message
	var err error
	if req.UserID != "" {
		messages, err = service.messageRepo.GetHistoryMessageListByC2C(req.Sender, req.UserID, req.GetType, req.LastMsgID, req.LastMsgSeq, req.Count, req.MessageTypeList)
	} else {
		messages, err = service.messageRepo.GetHistoryMessageListByGroup(req.Sender, req.GroupID, req.GetType, req.LastMsgID, req.LastMsgSeq, req.Count, req.MessageTypeList)
	}
	if err != nil {
		return nil, err
	}
	vmessages := make([]*dto.V2TimMessage, 0)
	for _, v := range messages {
		if v.Sender == req.Sender {
			v.IsSelf = true
		} else {
			v.IsSelf = false
		}
		m, _ := service.messageToV2TimMessage(v)
		vmessages = append(vmessages, m)
	}
	return vmessages, nil
}

func (service *MessageService) MarkC2CMessageAsRead(ownerID, userID string) error {
	service.conversationRepo.ClearC2CUnreadCount(ownerID, userID)
	return service.messageRepo.UpdateC2CMessageRead(ownerID, userID)
}

func (service *MessageService) MarkGroupMessageAsRead(ownerID, groupID string) error {
	return service.messageRepo.UpdateC2CMessageRead(ownerID, groupID)
}
