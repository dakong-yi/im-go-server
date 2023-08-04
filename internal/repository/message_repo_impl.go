package repository

import (
	"github.com/dakong-yi/im-go-server/internal/db"
	"github.com/dakong-yi/im-go-server/internal/model"
)

type MessageRepoImpl struct {
}

func NewMessageRepoImpl() MessageRepository {
	return &MessageRepoImpl{}
}

func (r *MessageRepoImpl) CreateMessage(message *model.Message) error {
	return db.DB.Create(message).Error
}
func (r *MessageRepoImpl) GetMessage(id string) (*model.Message, error) {
	var message model.Message
	err := db.DB.Where("id = ?", id).First(&message).Error
	return &message, err
}
func (r *MessageRepoImpl) UpdateMessage(id string, message *model.Message) (*model.Message, error) {
	err := db.DB.Model(message).Where("id = ?", id).Updates(message).Error
	return message, err
}

func (r *MessageRepoImpl) GetMessagesByC2C(sender, userID string) ([]*model.Message, error) {
	var messages []*model.Message
	err := db.DB.
		Where("sender = ? and user_id = ?", sender, userID).
		Order("created_at ASC").
		Find(&messages).
		Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}
func (r *MessageRepoImpl) GetMessagesByGroup(sender, groupID string) ([]*model.Message, error) {
	var messages []*model.Message
	err := db.DB.
		Where("sender = ? and group_id = ?", sender, groupID).
		Order("created_at ASC").
		Find(&messages).
		Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepoImpl) GetHistoryMessageListByC2C(sender, userID string, getType int, lastMsgID string, lastMsgSeq int, count int, messageTypeList []int) ([]*model.Message, error) {
	var message []*model.Message
	tx := db.DB.Where("(sender = ? and user_id = ?) or (sender = ? and user_id = ?)", sender, userID, userID, sender)
	if getType == model.V2TIM_GET_LOCAL_NEWER_MSG || getType == model.V2TIM_GET_CLOUD_NEWER_MSG {
		if lastMsgID != "" {
			tx = tx.Where("msg_id > ?", lastMsgID)
		} else {
			// tx = tx.Where("seq > ?", lastMsgSeq)
		}
		tx = tx.Order("msg_id ASC")
	}
	if getType == model.V2TIM_GET_LOCAL_OLDER_MSG || getType == model.V2TIM_GET_CLOUD_OLDER_MSG {
		if lastMsgID != "" {
			tx = tx.Where("msg_id < ?", lastMsgID)
		} else {
			// tx = tx.Where("seq < ?", lastMsgSeq)
		}
		tx = tx.Order("msg_id DESC")
	}
	tx = tx.Limit(count)
	err := tx.Find(&message).Error
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (r *MessageRepoImpl) GetHistoryMessageListByGroup(sender, groupID string, getType int, lastMsgID string, lastMsgSeq int, count int, messageTypeList []int) ([]*model.Message, error) {
	var message []*model.Message
	tx := db.DB.Where("sender = ? and group_id = ?", sender, groupID)
	if getType == model.V2TIM_GET_LOCAL_NEWER_MSG || getType == model.V2TIM_GET_CLOUD_NEWER_MSG {
		if lastMsgID != "" {
			tx = tx.Where("msg_id > ?", lastMsgID)
		} else {
			// tx = tx.Where("seq > ?", lastMsgSeq)
		}
		tx = tx.Order("msg_id ASC")
	}
	if getType == model.V2TIM_GET_LOCAL_OLDER_MSG || getType == model.V2TIM_GET_CLOUD_OLDER_MSG {
		if lastMsgID != "" {
			tx = tx.Where("msg_id < ?", lastMsgID)
		} else {
			// tx = tx.Where("seq < ?", lastMsgSeq)
		}
		tx = tx.Order("msg_id DESC")
	}
	tx = tx.Limit(count)
	err := tx.Find(&message).Error
	if err != nil {
		return nil, err
	}

	return message, nil
}
func (r *MessageRepoImpl) GetLatestMessageByIDs(IDs []string) ([]*model.Message, error) {
	var message []*model.Message
	err := db.DB.
		Where("id in ?", IDs).
		Find(&message).
		Error

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (r *MessageRepoImpl) UpdateC2CMessageRead(sender, userID string) error {
	return db.DB.Where("sender = ? and user_id = ? ", userID, sender).Updates(model.Message{IsRead: true}).Error
}
func (r *MessageRepoImpl) UpdateGroupMessageRead(sender, groupID string) error {
	return db.DB.Where("sender = ? and group_id = ?", sender, groupID).Updates(model.Message{IsRead: true}).Error
}
