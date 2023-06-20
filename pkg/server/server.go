package server

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/dakong-yi/im-go-server/pkg/conn"
	"github.com/dakong-yi/im-go-server/pkg/jsonutil"
	"github.com/dakong-yi/im-go-server/pkg/message"
)

type Server struct {
	Ip        string
	Port      int
	OnlineMap map[string]*conn.User
	mapLock   *sync.RWMutex
	Message   chan string
}

func NewServer(ip string, port int, messageBufferSize int) *Server {
	return &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: map[string]*conn.User{},
		mapLock:   &sync.RWMutex{},
		Message:   make(chan string, messageBufferSize),
	}
}

// Start 启动服务器的接口
func (s *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Printf("Error listening: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Listening on %s:%d\n", s.Ip, s.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		fmt.Printf("accepting connection: %v\n", "success")
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(connect net.Conn) {
	defer connect.Close()

	var user *conn.User

	// set up timer for heartbeat
	timeout := time.Second * 3000
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	go func() {
		<-timer.C
		// timeout, close the connection
		log.Printf("Connection timed out, closing...\n")
		user.Conn.Close()
	}()
	scanner := bufio.NewScanner(connect)
	for scanner.Scan() {

		// Read client input
		var req message.MessageRequest
		err := json.Unmarshal(scanner.Bytes(), &req)
		// if err := jsonutil.ReadJSONMessage(connect, &req); err != nil {
		if errors.Is(err, io.EOF) {
			log.Printf("Connection closed by client %s\n", req.RequestID)
		} else if err != nil {
			log.Printf("Error reading from connection: %v\n", err)
		}
		// }
		fmt.Println(req)
		switch req.Type {
		case message.RequestTypeChat:
			if user == nil {
				break
			}
			var chat message.ChatMsg
			if err := json.Unmarshal([]byte(req.Data), &chat); err != nil {
				log.Printf("Error unmarshaling chat message: %v\n", err)
				break
			}
			fmt.Println(chat.Content)
			if target, ok := s.OnlineMap[chat.Target]; ok {
				target.Message <- chat.Content
			}

		case message.RequestTypeSignIn:
			var signIn message.SignInMsg
			if err := json.Unmarshal([]byte(req.Data), &signIn); err != nil {
				log.Printf("Error unmarshaling sign-in message: %v\n", err)
				break
			}
			user = &conn.User{
				ID:      signIn.UserID,
				Addr:    connect.RemoteAddr().String(),
				Message: make(chan string, 10),
				Conn:    connect,
			}
			fmt.Println(user.ID, user.Addr)
			s.OnlineMap[user.ID] = user
			go user.ListenMessage()
		case message.RequestTypeHeartbeat:
			timer.Reset(timeout)
		default:
			log.Printf("Invalid message type: %d\n", req.Type)
		}

		resp := message.ResponseMessage{
			Code:    0,
			Message: "ok",
		}
		respBytes, _ := json.Marshal(resp)
		response := message.MessageRequest{
			Type:      message.RequestTypeResponse,
			Data:      string(respBytes),
			RequestID: "0",
		}
		if err := jsonutil.WriteJSONMessage(user.Conn, response); err != nil {
			log.Printf("Error sending response message: %v\n", err)
		}
	}
}
