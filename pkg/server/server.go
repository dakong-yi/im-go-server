package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"

	ImConn "github.com/dakong-yi/im-go-server/pkg/conn"
)

type Server struct {
	Ip   string
	Port int
	//OnlineMap map[string]*User
	mapLock sync.RWMutex
	Message chan string
}

// NewServer 创建一个普通的server接口（函数）
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
		//OnlineMap: make(map[string]*User),
		Message: make(chan string),
	}
	return server
}

func (s *Server) Handler(conn net.Conn) {
	c := ImConn.Conn{}
	c.TCP = conn
	c.ConnType = 1
	//当前连接的业务
	for {
		message := ImConn.Message{}
		d := json.NewDecoder(c.TCP)
		err1 := d.Decode(&message)
		if err1 != nil {
			if err1 == io.EOF {
				err := c.TCP.Close()
				if err != nil {
					fmt.Println(err)
					return
				}
				return
			}
			fmt.Println(err1)
		}
		if message.Type == 1 {
			c.SignIn(message)
		}
	}
}

func (s *Server) OnMessage(c ImConn.Conn, buffer []byte) {

	//buffer1 := []byte(`{"request_id":1,"data":"{\"user_id\":123,\"device_id\":123}","type":1}`)
	//fmt.Println(string(buffer))
	//fmt.Println(string(buffer1))

}

// Start 启动服务器的接口
func (s *Server) Start() {
	//socket listen
	fmt.Println("server is starting...")
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("listener err:", err)
		return
	}
	//为了防止遗忘关闭连接，加上defer保证在接口结束之后close
	//close listen socket
	defer listener.Close()

	//启动监听Message的goroutine
	//go s.ListenMessage()

	//listener一直监听连接
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		//go程处理连接
		go s.Handler(conn)
	}
}
