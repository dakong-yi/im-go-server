package chat

import (
	"errors"
	"sync"
	"time"
)

// Room 表示聊天室
type Room struct {
	ID       string
	Name     string
	Members  map[string]bool
	Messages []Message
	// 其他聊天室属性和方法
}

// Message 表示聊天消息
type Message struct {
	ID        string
	RoomID    string
	SenderID  string
	Content   string
	Timestamp int64
	// 其他消息属性
}

// ChatService 表示聊天服务
type ChatService struct {
	Rooms map[string]*Room
	// 其他服务属性和方法
}

var (
	chatService     *ChatService
	chatServiceOnce sync.Once
)

// GetChatService 返回 ChatService 实例
func GetChatService() *ChatService {
	chatServiceOnce.Do(func() {
		chatService = &ChatService{
			Rooms: make(map[string]*Room),
		}
	})
	return chatService
}

// CreateRoom 创建聊天室
func (cs *ChatService) CreateRoom(name string) *Room {
	room := &Room{
		ID:       generateRoomID(),
		Name:     name,
		Members:  make(map[string]bool),
		Messages: make([]Message, 0),
		// 其他聊天室属性的初始化
	}

	cs.Rooms[room.ID] = room

	return room
}

// JoinRoom 加入聊天室
func (cs *ChatService) JoinRoom(roomID, userID string) error {
	room, ok := cs.Rooms[roomID]
	if !ok {
		return errors.New("room not found")
	}

	room.Members[userID] = true

	return nil
}

// LeaveRoom 离开聊天室
func (cs *ChatService) LeaveRoom(roomID, userID string) error {
	room, ok := cs.Rooms[roomID]
	if !ok {
		return errors.New("room not found")
	}

	delete(room.Members, userID)

	return nil
}

// SendMessage 发送消息
func (cs *ChatService) SendMessage(roomID, senderID, content string) error {
	room, ok := cs.Rooms[roomID]
	if !ok {
		return errors.New("room not found")
	}

	message := Message{
		ID:        generateMessageID(),
		RoomID:    roomID,
		SenderID:  senderID,
		Content:   content,
		Timestamp: time.Now().Unix(),
		// 其他消息属性的设置
	}

	room.Messages = append(room.Messages, message)

	// TODO: 实现消息的广播

	return nil
}

// GetRoomMessages 返回指定聊天室的消息列表
func (cs *ChatService) GetRoomMessages(roomID string) ([]Message, error) {
	room, ok := cs.Rooms[roomID]
	if !ok {
		return nil, errors.New("room not found")
	}

	return room.Messages, nil
}

// 辅助函数

func generateRoomID() string {
	// TODO: 实现聊天室ID的生成
	return ""
}

func generateMessageID() string {
	// TODO: 实现消息ID的生成
	return ""
}
