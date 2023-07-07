package repository

import (
	"errors"

	"github.com/dakong-yi/im-go-server/internal/db"
	"github.com/dakong-yi/im-go-server/internal/model"
	"gorm.io/gorm"
)

type FriendshipRepoImpl struct {
}

func NewFriendshipRepoImpl() FriendshipRepository {
	return &FriendshipRepoImpl{}
}

func (r *FriendshipRepoImpl) GetFriends(userID string) ([]*model.User, error) {
	var friends []*model.User
	err := db.DB.Model(&model.Friendship{}).Where("user_id = ? OR friend_id = ?", userID, userID).Find(&friends).Error
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (r *FriendshipRepoImpl) CreateFriendship(userID, friendID string) error {
	friendship := &model.Friendship{
		UserID:   userID,
		FriendID: friendID,
	}
	err := db.DB.Create(friendship).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *FriendshipRepoImpl) DeleteFriendship(userID, friendID string) error {
	err := db.DB.Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, friendID, friendID, userID).Delete(&model.Friendship{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *FriendshipRepoImpl) GetFriendship(userID, friendID string) (*model.Friendship, error) {
	friendship := &model.Friendship{}
	err := db.DB.Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, friendID, friendID, userID).First(friendship).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return friendship, nil
}

func (r *FriendshipRepoImpl) UpdateFriendRemark(userID, friendID, remark string) error {
	err := db.DB.Model(&model.Friendship{}).Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, friendID, friendID, userID).Update("remark", remark).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *FriendshipRepoImpl) GetFriendRemark(userID, friendID string) (string, error) {
	var remark string
	err := db.DB.Model(&model.Friendship{}).Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, friendID, friendID, userID).Select("remark").Row().Scan(&remark)
	if err != nil {
		return "", err
	}
	return remark, nil
}
