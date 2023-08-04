package repository

import "github.com/dakong-yi/im-go-server/internal/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id int) (*model.User, error)
	GetUserByUserID(userID string) (*model.User, error)
	GetUserInfo(userIDs []string) ([]*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User) error
	GetUserByIdentifier(identifier string) (*model.User, error)
	// 其他方法...
}
