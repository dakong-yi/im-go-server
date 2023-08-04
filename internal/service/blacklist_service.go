package service

import (
	"github.com/dakong-yi/im-go-server/internal/model"
	"github.com/dakong-yi/im-go-server/internal/repository"
)

type BlacklistService struct {
	blacklistRepo repository.BlacklistRepository
}

func NewBlacklistService(blacklistRepo repository.BlacklistRepository) *BlacklistService {
	return &BlacklistService{
		blacklistRepo: blacklistRepo,
	}
}

func (s *BlacklistService) AddToBlacklist(userID string, blockedID string) error {
	return s.blacklistRepo.AddToBlacklist(userID, blockedID)
}

func (s *BlacklistService) RemoveFromBlacklist(userID string, blockedID string) error {
	return s.blacklistRepo.RemoveFromBlacklist(userID, blockedID)
}

func (s *BlacklistService) GetBlacklistByUserID(userID string) ([]*model.Blacklist, error) {
	return s.blacklistRepo.GetBlacklistByUserID(userID)
}

func (s *BlacklistService) IsBlocked(userID string, blockedID string) (bool, error) {
	return s.blacklistRepo.IsBlocked(userID, blockedID)
}
