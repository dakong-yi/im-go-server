package repository

import "github.com/dakong-yi/im-go-server/internal/model"

type FriendshipRepository interface {
	GetFriends(userID string) ([]*model.User, error)
	CreateFriendship(userID, friendID string) error
	DeleteFriendship(userID, friendID string) error
	GetFriendship(userID, friendID string) (*model.Friendship, error)
	UpdateFriendRemark(userID, friendID string, remark string) error
	GetFriendRemark(userID, friendID string) (string, error)
}
