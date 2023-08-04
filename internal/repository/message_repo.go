package repository

import "github.com/dakong-yi/im-go-server/internal/model"

// MessageRepository 定义消息存储库的接口
type MessageRepository interface {
	CreateMessage(message *model.Message) error
	GetMessage(id string) (*model.Message, error)
	UpdateMessage(id string, message *model.Message) (*model.Message, error)
	GetMessagesByC2C(sender, userID string) ([]*model.Message, error)
	GetMessagesByGroup(sender, GroupID string) ([]*model.Message, error)
	GetHistoryMessageListByC2C(sender, userID string, getType int, lastMsgID string, lastMsgSeqint, count int, messageTypeList []int) ([]*model.Message, error)
	GetHistoryMessageListByGroup(sender, groupID string, getType int, lastMsgID string, lastMsgSeq int, count int, messageTypeList []int) ([]*model.Message, error)
	GetLatestMessageByIDs(IDs []string) ([]*model.Message, error)
	UpdateC2CMessageRead(sender, userID string) error
	UpdateGroupMessageRead(sender, groupID string) error
}
