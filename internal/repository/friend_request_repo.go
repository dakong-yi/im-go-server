package repository

import "github.com/dakong-yi/im-go-server/internal/model"

type FriendRequestRepository interface {
	CreateFriendRequest(request *model.FriendRequest) error
	GetFriendRequestByID(id uint) (*model.FriendRequest, error)
	GetFriendRequestsByUserID(userID string) ([]*model.FriendRequest, error)
	GetPendingFriendRequestsByUserID(userID string) ([]*model.FriendRequest, error)
	GetFriendRequestsByFriendID(friendID string) ([]*model.FriendRequest, error)
	UpdateFriendRequestStatus(requestID uint, status string) error
	DeleteFriendRequest(requestID uint) error
}
