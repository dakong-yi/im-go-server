package services

// import (
// 	"context"

// 	"github.com/dakong-yi/im-go-server/repositories"
// )

// type ContactService interface {
// 	SendFriendRequest(ctx context.Context, userID, friendID uint) error
// 	ApproveFriendRequest(ctx context.Context, userID, friendID uint) error
// 	RejectFriendRequest(ctx context.Context, userID, friendID uint) error
// 	BlockFriend(ctx context.Context, userID, friendID uint) error
// 	Unfriend(ctx context.Context, userID, friendID uint) error
// 	IsFriend(ctx context.Context, userID, friendID uint) (bool, error)
// 	// 其他方法...
// }

// type contactService struct {
// 	contactRepo repositories.ContactRepository
// }

// func NewContactService(contactRepo repositories.ContactRepository) *contactService {
// 	return &contactService{contactRepo: contactRepo}
// }

// // 发起好友申请
// func (s *contactService) SendFriendRequest(ctx context.Context, userID, friendID uint) error {

// 	err := s.contactRepo.SendFriendRequest(userID, friendID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // 同意好友申请
// func (s *contactService) ApproveFriendRequest(ctx context.Context, userID, friendID uint) error {

// 	err := s.contactRepo.ApproveFriendRequest(userID, friendID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // 拒绝好友申请
// func (s *contactService) RejectFriendRequest(ctx context.Context, userID, friendID uint) error {
// 	err := s.contactRepo.RejectFriendRequest(userID, friendID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // 拉黑好友
// func (s *contactService) BlockFriend(ctx context.Context, userID, friendID uint) error {
// 	err := s.contactRepo.BlockFriend(userID, friendID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // 删除好友
// func (s *contactService) Unfriend(ctx context.Context, userID, friendID uint) error {
// 	err := s.contactRepo.DeleteFriend(userID, friendID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
