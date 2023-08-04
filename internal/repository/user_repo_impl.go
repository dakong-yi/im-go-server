package repository

import (
	"github.com/dakong-yi/im-go-server/internal/db"
	"github.com/dakong-yi/im-go-server/internal/model"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) CreateUser(user *model.User) error {
	return db.DB.Create(user).Error
}

func (r *UserRepositoryImpl) GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserRepositoryImpl) GetUserByUserID(userID string) (*model.User, error) {
	var user model.User
	err := db.DB.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserRepositoryImpl) GetUserInfo(userIDs []string) ([]*model.User, error) {
	var user []*model.User
	err := db.DB.Where("user_id in ?", userIDs).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user *model.User) error {
	return db.DB.Save(user).Error
}

func (r *UserRepositoryImpl) DeleteUser(user *model.User) error {
	return db.DB.Delete(user).Error
}

func (r *UserRepositoryImpl) GetUserByIdentifier(identifier string) (*model.User, error) {
	var user model.User
	err := db.DB.Where("identifier = ?", identifier).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
