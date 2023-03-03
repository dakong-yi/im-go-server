package conn

import (
	"sync"
)

var ManagerConn = sync.Map{}

// SetConn 存储
func SetConn(uid int64, conn *Conn) {
	ManagerConn.Store(uid, conn)
}

// GetConn 获取
func GetConn(uid int64) *Conn {
	value, ok := ManagerConn.Load(uid)
	if ok {
		return value.(*Conn)
	}
	return nil
}

// DeleteConn 删除
func DeleteConn(uid int64) {
	ManagerConn.Delete(uid)
}
