// services/friend_service.go

package services

// import (
// 	"github.com/dakong-yi/im-go-server/models"
// 	"github.com/dakong-yi/im-go-server/repositories"

// 	"github.com/pkg/errors"
// )

// type FriendService struct {
// 	friendRepo *repositories.FriendRepository
// }

// func NewFriendService(friendRepo *repositories.FriendRepository) *FriendService {
// 	return &FriendService{friendRepo: friendRepo}
// }

// func (s *FriendService) Create(friend *models.Friend) error {
// 	err := s.friendRepo.Create(friend)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to create friend")
// 	}

// 	return nil
// }

// func (s *FriendService) Delete(friend *models.Friend) error {
// 	err := s.friendRepo.Delete(friend)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to delete friend")
// 	}

// 	return nil
// }

// func (s *FriendService) FindFriendsByUserID(userID uint) ([]*models.Friend, error) {
// 	friends, err := s.friendRepo.FindByUserID(userID)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to find friends by user ID")
// 	}

// 	return friends, nil
// }

// func (s *FriendService) FindFriendsByFriendID(friendID uint) ([]*models.Friend, error) {
// 	friends, err := s.friendRepo.FindByFriendID(friendID)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to find friends by friend ID")
// 	}

// 	return friends, nil
// }

// func (s *FriendService) FindByID(id uint) (*models.Friend, error) {
// 	friend, err := s.friendRepo.FindByID(id)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to find friend by ID")
// 	}

// 	return friend, nil
// }
// func (s *FriendService) FindUsersByFriendID(friendID uint) ([]uint, error) {
// 	friend, err := s.friendRepo.FindByID(friendID)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to find friend")
// 	}
// 	if friend == nil {
// 		return nil, errors.New("friend not found")
// 	}

// 	users, err := s.friendRepo.FindUsersByFriend(friend)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to find users")
// 	}

// 	return users, nil
// }
