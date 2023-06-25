package server

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	connect "github.com/dakong-yi/im-go-server/pkg/conn"
	"github.com/dakong-yi/im-go-server/pkg/jsonutil"
	"github.com/dakong-yi/im-go-server/pkg/message"
	"github.com/dakong-yi/im-go-server/pkg/request"
	"github.com/dakong-yi/im-go-server/pkg/user"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
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

// 处理请求的主函数
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	user := &user.User{
		Conn: conn,
		Addr: conn.RemoteAddr().String(),
	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		req, err := parseRequest(scanner.Bytes())
		if errors.Is(err, io.EOF) {
			connect.DeleteConn(user.ID)
			return
		} else if err != nil {
			log.Printf("Error reading from connection: %v\n", err)
			continue
		}

		resp := (&request.RequestHandlersFacade{}).HandleRequest(req, user)

		if err := writeResponse(conn, resp); err != nil {
			log.Printf("Error sending response message: %v\n", err)
			continue
		}
	}
}

func parseRequest(data []byte) (*message.MessageRequest, error) {
	var req message.MessageRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling request: %v", err)
	}
	return &req, nil
}

func writeResponse(connect net.Conn, response message.MessageRequest) error {
	return jsonutil.WriteJSONMessage(connect, response)
}
