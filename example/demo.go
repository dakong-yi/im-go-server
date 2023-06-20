package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"

	"github.com/dakong-yi/im-go-server/pkg/message"
)

func client1() {
	// 创建一个新的客户端连接
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}

	// 使用 bufio 包装连接
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// 向 Server 发送一条消息
	signInMsg := message.SignInMsg{
		UserID:   "1",
		DeviceID: "1",
		Token:    "1",
	}
	signInMsgBytes, _ := json.Marshal(signInMsg)
	msgReq := message.MessageRequest{
		Type: message.RequestTypeSignIn,
		Data: string(signInMsgBytes),
	}
	msgReqBytes, _ := json.Marshal(msgReq)
	fmt.Fprint(writer, string(msgReqBytes)+"\n")
	writer.Flush()

	// 循环读取从 Server 发送过来的消息
	go func() {
		for {
			// 从 Server 读取响应消息
			response, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error receiving response from server:", err)
				os.Exit(1)
			}

			// 输出响应消息
			fmt.Println("Response from server:", response)
		}
	}()
	for {
	}
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	// 从标准输入读取一行数据
	// 	input := scanner.Text()
	// 	msgReq := message.MessageRequest{
	// 		Type: message.RequestTypeChat,
	// 		Data: input,
	// 	}
	// 	msgReqBytes, _ := json.Marshal(msgReq)
	// 	fmt.Fprint(writer, string(msgReqBytes)+"\n")
	// 	writer.Flush()
	// }

	// 手动关闭连接前刷新缓冲区
	writer.Flush()

	// 关闭连接
	conn.Close()
}
func client2() {
	// 创建一个新的客户端连接
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}

	// 使用 bufio 包装连接
	// reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// 向 Server 发送一条消息
	signInMsg := message.SignInMsg{
		UserID:   "2",
		DeviceID: "1",
		Token:    "1",
	}
	signInMsgBytes, _ := json.Marshal(signInMsg)
	msgReq := message.MessageRequest{
		Type: message.RequestTypeSignIn,
		Data: string(signInMsgBytes),
	}
	msgReqBytes, _ := json.Marshal(msgReq)
	fmt.Fprint(writer, string(msgReqBytes)+"\n")
	writer.Flush()

	// 循环读取从 Server 发送过来的消息
	// go func() {
	// 	for {
	// 		// 从 Server 读取响应消息
	// 		response, err := reader.ReadString('\n')
	// 		if err != nil {
	// 			fmt.Println("Error receiving response from server:", err)
	// 			os.Exit(1)
	// 		}

	// 		// 输出响应消息
	// 		fmt.Println("Response from server:", response)
	// 	}
	// }()
	for i := 0; i < 100; i++ {
		msg := message.ChatMsg{
			ChatType: message.ChatTypePrivate,
			Type:     message.MessageTypeText,
			Content:  strconv.Itoa(i),
			Sender:   "2",
			Target:   "1",
		}
		msgBytes, _ := json.Marshal(msg)
		msgReq := message.MessageRequest{
			Type: message.RequestTypeChat,
			Data: string(msgBytes),
		}
		fmt.Println(i)
		// time.Sleep(time.Second * 1)
		msgReqBytes, _ := json.Marshal(msgReq)
		n, e := fmt.Fprint(writer, string(msgReqBytes)+"\n")
		fmt.Println(n, e)
		writer.Flush()
	}
	for {
	}
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	// 从标准输入读取一行数据
	// 	input := scanner.Text()
	// 	msgReq := message.MessageRequest{
	// 		Type: message.RequestTypeChat,
	// 		Data: input,
	// 	}
	// 	msgReqBytes, _ := json.Marshal(msgReq)
	// 	fmt.Fprint(writer, string(msgReqBytes)+"\n")
	// 	writer.Flush()
	// }

	// 手动关闭连接前刷新缓冲区
	writer.Flush()

	// 关闭连接
	conn.Close()
}

func main() {
	go client1()
	client2()

}

func randString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
