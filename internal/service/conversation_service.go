package service

import (
	"fmt"
	"time"

	"github.com/dakong-yi/im-go-server/internal/dto"
	"github.com/dakong-yi/im-go-server/internal/dto/response"
	"github.com/dakong-yi/im-go-server/internal/model"
	"github.com/dakong-yi/im-go-server/internal/repository"
	"github.com/dakong-yi/im-go-server/internal/socket"
)

type ConversationService struct {
	conversationRepo    repository.ConversationRepository
	userRepo            repository.UserRepository
	msgConversationRepo repository.MessageRepository
	socketService       *socket.SocketServer
}

func NewConversationService(conversationRepo repository.ConversationRepository,
	userRepo repository.UserRepository, msgConversationRepo repository.MessageRepository, socketService *socket.SocketServer) *ConversationService {
	return &ConversationService{
		conversationRepo:    conversationRepo,
		userRepo:            userRepo,
		msgConversationRepo: msgConversationRepo,
		socketService:       socketService,
	}
}

// 发送消息需要消息双方创建对话
func (service *ConversationService) CreateConversation(message *dto.V2TimMessage) (*model.Conversation, error) {
	conversationType := model.CONVERSATION_TYPE_C2C
	if message.UserID == "" {
		conversationType = model.CONVERSATION_TYPE_GROUP
	}
	service.createConversation(conversationType, message.Sender, message.UserID, message.GroupID, message.ID)
	service.createConversation(conversationType, message.UserID, message.Sender, message.GroupID, message.ID)
	return nil, nil
}
func (service *ConversationService) createConversation(conversationType, sender, userID, groupID, ID string) (*model.Conversation, error) {
	var conversation *model.Conversation
	// 判断是否已存在
	if conversationType == model.CONVERSATION_TYPE_C2C {
		conversation, _ = service.conversationRepo.GetConversationByUserID(sender, userID)
	} else {
		conversation, _ = service.conversationRepo.GetConversationByGroupID(sender, groupID)
	}
	// 如果存在，则直接返回对话
	if conversation != nil {
		conversation.LastMessageID = ID
		conversation.UnreadCount = conversation.UnreadCount + 1
		err := service.conversationRepo.UpdateLastMessage(conversation)
		if err != nil {
			return nil, err
		}
		resp, _ := service.GetConversationsByOwnerID(sender)
		service.socketService.OnConversationChanged(sender, resp.ConversationList)
		return conversation, nil
	}
	// 创建对话
	conversation = &model.Conversation{
		Type:          conversationType,
		OwnerID:       sender,
		UserID:        userID,
		GroupID:       groupID,
		LastMessageID: ID,
		Orderkey:      time.Now().Unix(),
	}
	err := service.conversationRepo.CreateConversation(conversation)
	if err != nil {
		return nil, err
	}
	resp, _ := service.GetConversationsByOwnerID(sender)
	service.socketService.OnNewConversation(sender, resp.ConversationList)
	return conversation, nil
}

func (service *ConversationService) GetConversationByID(ownerID, userID, groupID, conversationType string) (*dto.V2TimConversation, error) {
	var conversation *model.Conversation
	var err error
	// 如果是私聊对话，
	if conversationType == model.CONVERSATION_TYPE_C2C {
		// 检查  是否已存在
		conversation, err = service.conversationRepo.GetConversationByUserID(ownerID, userID)
	} else {
		// 检查  是否已存在
		conversation, err = service.conversationRepo.GetConversationByGroupID(ownerID, groupID)
	}
	if err != nil {
		return nil, err
	}
	return service.ConversationToV2TimConversation(conversation)
}

func (service *ConversationService) ConversationToV2TimConversation(conversation *model.Conversation) (*dto.V2TimConversation, error) {
	var groupIDs []string
	var userIDs []string
	var messageIDs []string
	if conversation.Type == model.CONVERSATION_TYPE_GROUP {
		groupIDs = append(groupIDs, conversation.GroupID)
	} else {
		userIDs = append(userIDs, conversation.UserID)
	}
	if conversation.LastMessageID != "" {
		messageIDs = append(messageIDs, conversation.LastMessageID)
	}
	// TODO: 获取群组信息
	if len(groupIDs) > 0 {
		// group_infos := service.conversationRepo.GetGroupInfo(group_ids)
	}
	userInfoMap := make(map[string]*model.User)
	if len(userIDs) > 0 {
		userInfos, err := service.userRepo.GetUserInfo(userIDs)
		if err != nil {
			return nil, err
		}
		for _, user := range userInfos {
			userInfoMap[user.UserID] = user
		}
	}
	messageMap := make(map[string]*model.Message)
	messageInfos, err := service.msgConversationRepo.GetLatestMessageByIDs(messageIDs)
	if err != nil {
		return nil, err
	}
	for _, message := range messageInfos {
		messageMap[message.ID] = message
	}
	conversationID := fmt.Sprintf("%s_%s", conversation.Type, conversation.UserID)
	if conversation.Type == model.CONVERSATION_TYPE_GROUP {
		conversationID = fmt.Sprintf("%s_%s", conversation.Type, conversation.GroupID)
	}
	v2Message := &dto.V2TimMessage{}
	if m, ok := messageMap[conversation.LastMessageID]; ok {
		v2Message, err = service.MessageToV2TimMessage(m)
		if err != nil {
			return nil, err
		}
	}
	conversationType := model.ConversationTypes[conversation.Type]
	res := &dto.V2TimConversation{
		ConversationID:        conversationID,
		Type:                  conversationType,
		UserID:                conversation.UserID,
		GroupID:               conversation.GroupID,
		GroupType:             conversation.GroupType,
		DraftText:             "",
		DraftTimestamp:        0,
		GroupAtInfoList:       []interface{}{},
		IsPinned:              conversation.IsPinned,
		RecvOpt:               conversation.RecvOpt,
		Orderkey:              conversation.Orderkey,
		MarkList:              []int{},
		CustomData:            "",
		ConversationGroupList: []string{},
		UnreadCount:           conversation.UnreadCount,
		LastMessage:           v2Message,
	}
	if u, ok := userInfoMap[conversation.UserID]; ok {
		res.ShowName = u.NickName
		res.FaceUrl = u.FaceURL
	}
	return res, nil
}

func (service *ConversationService) GetConversationsByOwnerID(ownerID string) (*response.ConversationResponse, error) {
	conversations, err := service.conversationRepo.GetConversationsByOwnerID(ownerID)
	if err != nil {
		return nil, err
	}
	var groupIDs []string
	var userIDs []string
	var messageIDs []string
	for _, conversation := range conversations {
		if conversation.Type == model.CONVERSATION_TYPE_GROUP {
			groupIDs = append(groupIDs, conversation.GroupID)
		} else {
			userIDs = append(userIDs, conversation.UserID)
		}
		if conversation.LastMessageID != "" {
			messageIDs = append(messageIDs, conversation.LastMessageID)
		}
	}
	// TODO: 获取群组信息
	if len(groupIDs) > 0 {
		// group_infos := service.conversationRepo.GetGroupInfo(group_ids)
	}
	userInfoMap := make(map[string]*model.User)
	if len(userIDs) > 0 {
		userInfos, err := service.userRepo.GetUserInfo(userIDs)
		if err != nil {
			return nil, err
		}
		for _, user := range userInfos {
			userInfoMap[user.UserID] = user
		}
	}
	messageMap := make(map[string]*model.Message)
	messageInfos, err := service.msgConversationRepo.GetLatestMessageByIDs(messageIDs)
	if err != nil {
		return nil, err
	}
	for _, message := range messageInfos {
		messageMap[message.ID] = message
	}
	result := make([]*dto.V2TimConversation, 0)
	for _, conversation := range conversations {
		conversationID := fmt.Sprintf("%s_%s", conversation.Type, conversation.UserID)
		if conversation.Type == model.CONVERSATION_TYPE_GROUP {
			conversationID = fmt.Sprintf("%s_%s", conversation.Type, conversation.GroupID)
		}
		v2Message := &dto.V2TimMessage{}
		if m, ok := messageMap[conversation.LastMessageID]; ok {
			v2Message, err = service.MessageToV2TimMessage(m)
			if err != nil {
				return nil, err
			}
		}
		conversationType := model.ConversationTypes[conversation.Type]
		res := &dto.V2TimConversation{
			ConversationID:        conversationID,
			Type:                  conversationType,
			UserID:                conversation.UserID,
			GroupID:               conversation.GroupID,
			GroupType:             conversation.GroupType,
			DraftText:             "",
			DraftTimestamp:        0,
			GroupAtInfoList:       []interface{}{},
			IsPinned:              conversation.IsPinned,
			RecvOpt:               conversation.RecvOpt,
			Orderkey:              conversation.Orderkey,
			MarkList:              []int{},
			CustomData:            "",
			ConversationGroupList: []string{},
			UnreadCount:           conversation.UnreadCount,
			LastMessage:           v2Message,
		}
		if u, ok := userInfoMap[conversation.UserID]; ok {
			res.ShowName = u.NickName
			res.FaceUrl = u.FaceURL
		}
		result = append(result, res)
	}
	r := &response.ConversationResponse{NextSeq: "0", IsFinished: true, ConversationList: result}
	return r, nil
}

func (service *ConversationService) MessageToV2TimMessage(message *model.Message) (*dto.V2TimMessage, error) {
	userInfoMap := make(map[string]*model.User)
	userInfos, err := service.userRepo.GetUserInfo([]string{message.UserID})
	if err != nil {
		return nil, err
	}
	for _, user := range userInfos {
		userInfoMap[user.UserID] = user
	}

	vmessage := &dto.V2TimMessage{
		MsgID:     message.ID,
		Timestamp: message.Timestamp,
		Progress:  message.Progress,
		Sender:    message.Sender,
		NameCard:  "",
		GroupID:   message.GroupID,
		UserID:    message.UserID,
		Status:    int(message.Status),
		ElemType:  int(message.ElemType),
		TextElem: &dto.V2TimTextElem{
			Text: message.Text,
		},
		IsSelf:          message.IsSelf,
		IsRead:          message.IsRead,
		Seq:             message.Seq,
		Random:          message.Random,
		ID:              message.ID,
		NeedReadReceipt: message.NeedReadReceipt,
	}
	if u, ok := userInfoMap[message.UserID]; ok {
		vmessage.NickName = u.NickName
		vmessage.FaceUrl = u.FaceURL
	}
	return vmessage, nil
}

func (service *ConversationService) GetAllUnreadMessageCount(ownerID string) (int, error) {
	// TODO : 获取群组信息
	groupIDs := []string{}
	count, err := service.conversationRepo.GetAllUnreadMessageCount(ownerID, groupIDs)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (service *ConversationService) OnConversationChanged(userID string) {
	resp, _ := service.GetConversationsByOwnerID(userID)
	service.socketService.OnConversationChanged(userID, resp.ConversationList)
}
func (service *ConversationService) OnTotalUnreadMessageCountChanged(userID string) {
	count, _ := service.GetAllUnreadMessageCount(userID)
	service.socketService.OnTotalUnreadMessageCountChanged(userID, count)
}
