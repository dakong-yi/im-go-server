package repository

import (
	"github.com/dakong-yi/im-go-server/internal/db"
	"github.com/dakong-yi/im-go-server/internal/model"
)

type BlacklistRepoImpl struct {
}

func NewBlacklistRepoImpl() BlacklistRepository {
	return &BlacklistRepoImpl{}
}

func (r *BlacklistRepoImpl) AddToBlacklist(userID string, blockedID string) error {
	blacklist := &model.Blacklist{
		UserID:    userID,
		BlockedID: blockedID,
	}
	return db.DB.Create(blacklist).Error
}

func (r *BlacklistRepoImpl) RemoveFromBlacklist(userID string, blockedID string) error {
	return db.DB.Where("user_id = ? AND blocked_id = ?", userID, blockedID).Delete(&model.Blacklist{}).Error
}

func (r *BlacklistRepoImpl) GetBlacklistByUserID(userID string) ([]*model.Blacklist, error) {
	var blacklist []*model.Blacklist
	err := db.DB.Where("user_id = ?", userID).Find(&blacklist).Error
	if err != nil {
		return nil, err
	}
	return blacklist, nil
}

func (r *BlacklistRepoImpl) IsBlocked(userID string, blockedID string) (bool, error) {
	var count int64
	err := db.DB.Model(&model.Blacklist{}).Where("user_id = ? AND blocked_id = ?", userID, blockedID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
