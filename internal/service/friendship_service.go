package service

import (
	"github.com/dakong-yi/im-go-server/internal/dto"
	"github.com/dakong-yi/im-go-server/internal/model"
	"github.com/dakong-yi/im-go-server/internal/repository"
)

type FriendshipService struct {
	friendshipRepo repository.FriendshipRepository
	userRepo       repository.UserRepository
}

func NewFriendshipService(friendshipRepo repository.FriendshipRepository, userRepo repository.UserRepository) *FriendshipService {
	return &FriendshipService{
		friendshipRepo: friendshipRepo,
		userRepo:       userRepo,
	}
}

func (service *FriendshipService) GetFriends(userID string) ([]*dto.V2TimFriendInfo, error) {
	friends, err := service.friendshipRepo.GetFriends(userID)
	if err != nil {
		return nil, err
	}
	return service.FriendToV2TimFriendInfo(friends)
}

func (service *FriendshipService) GetFriendsInfo(sender string, userID []string) ([]*dto.V2TimFriendInfo, error) {
	friends, err := service.friendshipRepo.GetFriendsInfo(sender, userID)
	if err != nil {
		return nil, err
	}
	return service.FriendToV2TimFriendInfo(friends)
}

func (service *FriendshipService) FriendToV2TimFriendInfo(friends []*model.Friendship) ([]*dto.V2TimFriendInfo, error) {
	userIDs := make([]string, 0)
	for _, friend := range friends {
		userIDs = append(userIDs, friend.FriendID)
	}
	users, err := service.userRepo.GetUserInfo(userIDs)
	if err != nil {
		return nil, err
	}
	userMap := make(map[string]*model.User)
	for _, user := range users {
		userMap[user.UserID] = user
	}
	vfriendInfos := make([]*dto.V2TimFriendInfo, 0)
	for _, friend := range friends {
		if user, ok := userMap[friend.FriendID]; ok {
			vfriendInfo := &dto.V2TimFriendInfo{
				UserID:       friend.FriendID,
				FriendRemark: "",
				FriendGroups: []string{},
				UserProfile: dto.V2TimUserFullInfo{
					UserID:        user.UserID,
					NickName:      user.NickName,
					FaceURL:       user.FaceURL,
					SelfSignature: user.SelfSignature,
					Gender:        user.Gender,
					AllowType:     user.AllowType,
					Role:          user.Role,
					Level:         user.Level,
					Birthday:      user.Birthday,
				},
			}
			vfriendInfos = append(vfriendInfos, vfriendInfo)
		}
	}
	return vfriendInfos, nil

}

func (service *FriendshipService) CreateFriendship(userID, friendID string) error {
	return service.friendshipRepo.CreateFriendship(userID, friendID)
}

func (service *FriendshipService) DeleteFriendship(userID, friendID string) error {
	return service.friendshipRepo.DeleteFriendship(userID, friendID)
}

func (service *FriendshipService) GetFriendship(userID, friendID string) (*model.Friendship, error) {
	return service.friendshipRepo.GetFriendship(userID, friendID)
}

func (service *FriendshipService) UpdateFriendRemark(userID, friendID string, remark string) error {
	return service.friendshipRepo.UpdateFriendRemark(userID, friendID, remark)
}

func (service *FriendshipService) GetFriendRemark(userID, friendID string) (string, error) {
	return service.friendshipRepo.GetFriendRemark(userID, friendID)
}
