package repository

import (
	"github.com/dakong-yi/im-go-server/internal/db"
	"github.com/dakong-yi/im-go-server/internal/model"
)

type FriendRequestRepoImpl struct {
}

func NewFriendRequestRepoImpl() FriendRequestRepository {
	return &FriendRequestRepoImpl{}
}

func (r *FriendRequestRepoImpl) CreateFriendRequest(request *model.FriendRequest) error {
	return db.DB.Create(request).Error
}

func (r *FriendRequestRepoImpl) GetFriendRequestByID(id uint) (*model.FriendRequest, error) {
	var request model.FriendRequest
	err := db.DB.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (r *FriendRequestRepoImpl) GetFriendRequestsByUserID(userID string) ([]*model.FriendRequest, error) {
	var requests []*model.FriendRequest
	err := db.DB.Where("user_id = ?", userID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (r *FriendRequestRepoImpl) GetPendingFriendRequestsByUserID(userID string) ([]*model.FriendRequest, error) {
	var requests []*model.FriendRequest
	err := db.DB.Where("user_id = ? AND status = ?", userID, "pending").Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (r *FriendRequestRepoImpl) GetFriendRequestsByFriendID(friendID string) ([]*model.FriendRequest, error) {
	var requests []*model.FriendRequest
	err := db.DB.Where("friend_id = ?", friendID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (r *FriendRequestRepoImpl) UpdateFriendRequestStatus(requestID uint, status string) error {
	return db.DB.Model(&model.FriendRequest{}).Where("id = ?", requestID).Update("status", status).Error
}

func (r *FriendRequestRepoImpl) DeleteFriendRequest(requestID uint) error {
	return db.DB.Delete(&model.FriendRequest{}, requestID).Error
}
