package socket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ConnectionManager struct {
	connUserMap sync.Map
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connUserMap: sync.Map{},
	}
}
func (cm *ConnectionManager) BindConnection(conn *websocket.Conn, userID string) {
	cm.connUserMap.Store(userID, conn)
}

func (cm *ConnectionManager) UnbindConnection(userID string) {
	cm.connUserMap.Delete(userID)
}

func (cm *ConnectionManager) GetConn(userID string) (*websocket.Conn, bool) {
	conn, ok := cm.connUserMap.Load(userID)
	if ok {
		return conn.(*websocket.Conn), true
	}
	return nil, false
}
