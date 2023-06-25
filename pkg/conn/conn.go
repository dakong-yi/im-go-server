package conn

import (
	"sync"

	"github.com/dakong-yi/im-go-server/pkg/user"
)

var ConnsPool = sync.Map{}

// SetConn 存储
func SetConn(userID string, user *user.User) {
	ConnsPool.Store(userID, user)
}

// GetConn 获取
func GetConn(userID string) *user.User {
	value, ok := ConnsPool.Load(userID)
	if ok {
		return value.(*user.User)
	}
	return nil
}

// DeleteConn 删除
func DeleteConn(userID string) {
	ConnsPool.Delete(userID)
}
