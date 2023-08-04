package socket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dakong-yi/im-go-server/internal/dto"
	"github.com/goinggo/mapstructure"
	"github.com/gorilla/websocket"
)

// SocketServer WebSocket 服务器
type SocketServer struct {
	upgrader          websocket.Upgrader
	register          chan *websocket.Conn
	unregister        chan *websocket.Conn
	connectionManager *ConnectionManager
}

// NewSocketServer 创建 SocketServer 实例
func NewSocketServer() *SocketServer {
	return &SocketServer{
		upgrader:          websocket.Upgrader{},
		register:          make(chan *websocket.Conn),
		unregister:        make(chan *websocket.Conn),
		connectionManager: NewConnectionManager(),
	}
}

// HandleConnection 处理客户端连接请求
func (s *SocketServer) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	go s.handleMessages(conn)
}

// handleMessages 处理客户端发送的消息
func (s *SocketServer) handleMessages(conn *websocket.Conn) {
	defer func() {
		s.unregister <- conn
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		// 处理消息
		s.handleMessage(conn, msg)
	}
}

// handleMessage 处理消息
func (s *SocketServer) handleMessage(conn *websocket.Conn, message []byte) {
	msg, err := Deserialize(message)
	if err != nil {
		return
	}
	switch msg.Action {
	case LoginAction:
		s.HandleUserLogin(conn, msg.Data)
	case ChatAction:
		s.handleChatMessage(conn, msg.Data)
	}
}

// HandleUserLogin 处理用户登录消息
func (s *SocketServer) HandleUserLogin(conn *websocket.Conn, data map[string]interface{}) {
	// 解析消息数据
	var loginMessage UserLoginMessage
	if err := mapstructure.Decode(data, &loginMessage); err != nil {
		// 处理解析错误
		fmt.Println("Failed to decode user login message:", err)
		return
	}

	// 获取用户ID
	userID := loginMessage.UserID

	// 将连接和用户ID绑定
	s.connectionManager.BindConnection(conn, userID)

	fmt.Println("User", userID, "logged in.")
}

func (s *SocketServer) handleChatMessage(conn *websocket.Conn, message map[string]interface{}) {
	// 解析消息
	// var chatMessage request.CreateMessageRequest
	// if err := mapstructure.Decode(message, &chatMessage); err != nil {
	// 	// 消息解析失败，处理错误逻辑
	// 	log.Println("Failed to parse chat message:", err)
	// 	return
	// }
	// // 保存消息到数据库 TODO
	// var msg *model.Message
	// var err error
	// if msg, err = s.messageService.CreateMessage(chatMessage); err != nil {
	// 	// 消息保存失败，处理错误逻辑
	// 	log.Println("Failed to save chat message:", err)
	// 	return
	// }
	// // 通过conversation获取接受用户id
	// userIDs, err := s.conversationService.GetUsersByConversationID(chatMessage.ConversationID)
	// if err != nil {
	// 	return
	// }
	userIDs := []string{"user1", "user2"}
	// 广播消息给所有在线用户
	for _, uid := range userIDs {
		// if uid == chatMessage.SenderID {
		// 	continue
		// }
		conn, ok := s.connectionManager.GetConn(uid)
		fmt.Println(ok)
		if ok {
			// content := response.ToMessageResponse(msg)
			// err := conn.WriteJSON(content)
			err := conn.WriteJSON(message)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (s *SocketServer) SendMessage(message *dto.V2TimMessage) {
	// 解析消息
	// var chatMessage request.CreateMessageRequest
	// if err := mapstructure.Decode(message, &chatMessage); err != nil {
	// 	// 消息解析失败，处理错误逻辑
	// 	log.Println("Failed to parse chat message:", err)
	// 	return
	// }
	// // 保存消息到数据库 TODO
	// var msg *model.Message
	// var err error
	// if msg, err = s.messageService.CreateMessage(chatMessage); err != nil {
	// 	// 消息保存失败，处理错误逻辑
	// 	log.Println("Failed to save chat message:", err)
	// 	return
	// }
	// // 通过conversation获取接受用户id
	// userIDs, err := s.conversationService.GetUsersByConversationID(chatMessage.ConversationID)
	// if err != nil {
	// 	return
	// }
	newMsg := *message
	conn, ok := s.connectionManager.GetConn(newMsg.UserID)
	fmt.Println(ok)
	if ok {
		newMsg.UserID = newMsg.Sender
		newMsg.IsSelf = false
		content := map[string]interface{}{
			"action": "RecvNewMessage",
			"data":   newMsg,
		}
		err := conn.WriteJSON(content)
		if err != nil {
			log.Println(err)
		}
	}
}

func (s *SocketServer) OnNewConversation(userID string, conversations []*dto.V2TimConversation) {
	conn, ok := s.connectionManager.GetConn(userID)
	if ok {
		content := map[string]interface{}{
			"action": "NewConversation",
			"data":   conversations,
		}
		err := conn.WriteJSON(content)
		if err != nil {
			log.Println(err)
		}
	}
}
func (s *SocketServer) OnConversationChanged(userID string, conversations []*dto.V2TimConversation) {
	conn, ok := s.connectionManager.GetConn(userID)
	if ok {
		content := map[string]interface{}{
			"action": "ConversationChanged",
			"data":   conversations,
		}
		err := conn.WriteJSON(content)
		if err != nil {
			log.Println(err)
		}
	}
}

func (s *SocketServer) OnTotalUnreadMessageCountChanged(userID string, count int) {
	conn, ok := s.connectionManager.GetConn(userID)
	if ok {
		content := map[string]interface{}{
			"action": "TotalUnreadMessageCountChanged",
			"data":   count,
		}
		err := conn.WriteJSON(content)
		if err != nil {
			log.Println(err)
		}
	}
}

// 启动 Socket 服务
func (s *SocketServer) Start() error {
	// 监听 WebSocket 连接
	http.HandleFunc("/ws", s.HandleConnection)
	log.Println("WebSocket server started.:8081")
	// 启动 WebSocket 服务器
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Failed to start WebSocket server:", err)
		return err
	}
	return nil
}
