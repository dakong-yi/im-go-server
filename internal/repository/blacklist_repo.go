package repository

import "github.com/dakong-yi/im-go-server/internal/model"

type BlacklistRepository interface {
	AddToBlacklist(userID string, blockedID string) error
	RemoveFromBlacklist(userID string, blockedID string) error
	GetBlacklistByUserID(userID string) ([]*model.Blacklist, error)
	IsBlocked(userID string, blockedID string) (bool, error)
}
