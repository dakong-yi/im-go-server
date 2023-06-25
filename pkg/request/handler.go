package request

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/dakong-yi/im-go-server/pkg/conn"
	"github.com/dakong-yi/im-go-server/pkg/message"
	"github.com/dakong-yi/im-go-server/pkg/user"
)

// Chat请求的处理策略
type chatHandler struct{}

func (h *chatHandler) handle(data string, user *user.User) message.MessageRequest {
	var chat message.ChatMsg
	if err := json.Unmarshal([]byte(data), &chat); err != nil {
		log.Printf("Error unmarshaling chat message: %v\n", err)
		return respError(err)
	}
	target := conn.GetConn(chat.Target)
	target.Message <- chat.Content
	return respOk()
}

// SignIn请求的处理策略
type signInHandler struct{}

func (h *signInHandler) handle(data string, user *user.User) message.MessageRequest {
	var signIn message.SignInMsg
	if err := json.Unmarshal([]byte(data), &signIn); err != nil {
		log.Printf("Error unmarshaling sign-in message: %v\n", err)
		return respError(err)
	}
	user.ID = signIn.UserID
	user.Message = make(chan string, 10)
	conn.SetConn(signIn.UserID, user)
	user.Timer = time.NewTimer(time.Second * 3000)
	defer user.Timer.Stop()
	go user.ListenMessage()
	return respOk()
}

// Heartbeat请求的处理策略
type heartbeatHandler struct{}

func (h *heartbeatHandler) handle(data string, user *user.User) message.MessageRequest {
	user.Timer.Reset(time.Second * 3000)
	return respOk()
}

// Invalid请求的处理策略
type invalidHandler struct{}

func (h *invalidHandler) handle(data string, user *user.User) message.MessageRequest {
	return respError(fmt.Errorf("invalid message type"))
}

func respOk() message.MessageRequest {
	resp := message.MessageResponse{
		Code:    0,
		Message: "ok",
	}
	respBytes, _ := json.Marshal(resp)
	response := message.MessageRequest{
		Type:      message.RequestTypeResponse,
		Data:      string(respBytes),
		RequestID: "0",
	}
	return response
}

func respError(err error) message.MessageRequest {
	resp := message.MessageResponse{
		Code:    0,
		Message: "ok",
	}
	respBytes, _ := json.Marshal(resp)
	response := message.MessageRequest{
		Type:      message.RequestTypeResponse,
		Data:      string(respBytes),
		RequestID: "0",
	}
	return response
}

// 处理请求的策略映射表
var requestHandlers = map[message.RequestType]requestHandler{
	message.RequestTypeChat:      &chatHandler{},
	message.RequestTypeSignIn:    &signInHandler{},
	message.RequestTypeHeartbeat: &heartbeatHandler{},
}

type RequestHandlersFacade struct{}

func (f *RequestHandlersFacade) HandleRequest(req *message.MessageRequest, user *user.User) message.MessageRequest {
	handler, found := requestHandlers[req.Type]
	if !found {
		handler = &invalidHandler{}
	}
	return handler.handle(req.Data, user)
}
