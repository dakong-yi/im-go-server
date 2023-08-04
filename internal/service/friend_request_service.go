package service

import (
	"github.com/dakong-yi/im-go-server/internal/model"
	"github.com/dakong-yi/im-go-server/internal/repository"
)

type FriendRequestService struct {
	friendRequestRepo repository.FriendRequestRepository
}

func NewFriendRequestService(friendRequestRepo repository.FriendRequestRepository) *FriendRequestService {
	return &FriendRequestService{
		friendRequestRepo: friendRequestRepo,
	}
}

func (service *FriendRequestService) SendFriendRequest(request *model.FriendRequest) error {
	return service.friendRequestRepo.CreateFriendRequest(request)
}

func (service *FriendRequestService) AcceptFriendRequest(requestID int) error {
	return service.friendRequestRepo.UpdateFriendRequestStatus(requestID, model.FriendRequestStatusAccepted)
}

func (service *FriendRequestService) RejectFriendRequest(requestID int) error {
	return service.friendRequestRepo.UpdateFriendRequestStatus(requestID, model.FriendRequestStatusRejected)
}

func (service *FriendRequestService) GetPendingFriendRequests(userID string) ([]*model.FriendRequest, error) {
	return service.friendRequestRepo.GetPendingFriendRequestsByUserID(userID)
}

func (service *FriendRequestService) GetFriendRequests(userID string) ([]*model.FriendRequest, error) {
	return service.friendRequestRepo.GetFriendRequestsByUserID(userID)
}

func (service *FriendRequestService) GetReceivedFriendRequests(friendID string) ([]*model.FriendRequest, error) {
	return service.friendRequestRepo.GetFriendRequestsByFriendID(friendID)
}

func (service *FriendRequestService) CancelFriendRequest(requestID int) error {
	return service.friendRequestRepo.DeleteFriendRequest(requestID)
}
