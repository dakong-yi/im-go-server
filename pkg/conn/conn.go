package conn

import (
	"container/list"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	TypeTcp int8 = 1 // tcp连接
	TypeWs  int8 = 2 // websocket连接
)

type Conn struct {
	ConnType int8            // 连接类型
	TCP      net.Conn        // tcp连接
	WSMutex  sync.Mutex      // WS写锁
	WS       *websocket.Conn // websocket连接
	UserId   int64           // 用户ID
	DeviceId int64           // 设备ID
	Element  *list.Element   // 链表节点
}

// Write 写入数据
func (c *Conn) Write(bytes []byte) error {
	if c.ConnType == TypeTcp {
		_, err := c.TCP.Write(bytes)
		return err
	}
	if c.ConnType == TypeWs {
		return c.WriteToWS(bytes)
	}
	return nil
}

// WriteToWS 消息写入WebSocket
func (c *Conn) WriteToWS(bytes []byte) error {
	c.WSMutex.Lock()
	defer c.WSMutex.Unlock()

	err := c.WS.SetWriteDeadline(time.Now().Add(10 * time.Millisecond))
	if err != nil {
		return err
	}
	return c.WS.WriteMessage(websocket.BinaryMessage, bytes)
}

// Send 下发消息
func (c *Conn) Send(pt int, requestId int64, message Message, err error) {
	//outputBytes, err := json.Marshal(message.Data)
	//fmt.Println(string(outputBytes))
	//if err != nil {
	//	return
	//}
	err = c.Write([]byte(message.Data))
	if err != nil {
		//c.Close()
		return
	}
}

// SignIn 登录
func (c *Conn) SignIn(input Message) {
	var signIn SignInInput
	err := json.Unmarshal([]byte(input.Data), &signIn)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Send(1, input.RequestId, Message{Data: `{"name":"xy"}`}, err)
	if err != nil {
		return
	}
	c.UserId = signIn.UserId
	c.DeviceId = signIn.DeviceId
	SetConn(signIn.UserId, c)
}

// SignInInput 设备登录,
type SignInInput struct {
	DeviceId int64  `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"` // 设备id
	UserId   int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`       // 用户id
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`                        // 秘钥
}

type Message struct {
	RequestId int64  `json:"request_id"`
	Data      string `json:"data"`
	Type      int32  `json:"type"`
}
