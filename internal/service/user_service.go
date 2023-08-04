package service

import (
	"errors"

	"github.com/dakong-yi/im-go-server/internal/dto/request"
	"github.com/dakong-yi/im-go-server/internal/dto/response"
	"github.com/dakong-yi/im-go-server/internal/model"
	"github.com/dakong-yi/im-go-server/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*model.User, error) {
	return s.userRepo.GetUserByID(id)
}
func (s *UserService) GetUserByUserID(userID string) (*model.User, error) {
	return s.userRepo.GetUserByUserID(userID)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(user *model.User) error {
	return s.userRepo.DeleteUser(user)
}

func (s *UserService) GetUserInfo(userIDs []string) ([]*response.UserResponse, error) {
	users, err := s.userRepo.GetUserInfo(userIDs)
	if err != nil {
		return nil, err
	}
	res := make([]*response.UserResponse, 0)
	for _, v := range users {
		res = append(res, response.ToUserResponse(v))
	}
	return res, nil
}

func (service *UserService) Register(request request.CreateUserRequest) (*model.User, error) {
	// 检查用户名是否已存在
	existingUser, err := service.userRepo.GetUserByIdentifier(request.Username)
	if err != gorm.ErrRecordNotFound && err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// 创建用户对象
	user := &model.User{
		UserID:        request.Username,
		Identifier:    request.Username,
		Email:         request.Email,
		NickName:      request.NickName,
		FaceURL:       request.FaceURL,
		SelfSignature: request.SelfSignature,
		Gender:        request.Gender,
		AllowType:     request.AllowType,
		Role:          request.Role,
		Level:         request.Level,
		Birthday:      request.Birthday,
	}

	// 对密码进行加密处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// 创建用户
	err = service.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) Login(username, password string) (*response.UserResponse, error) {
	// 根据用户名查询用户
	user, err := service.userRepo.GetUserByIdentifier(username)
	if err != nil {
		return nil, err
	}

	// 检查用户密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return response.ToUserResponse(user), nil
}
