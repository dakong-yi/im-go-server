package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8081/ws?user_id=user2", nil)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket server:", err)
	}
	defer conn.Close()

	go readMessages(conn)

	for {
		var message string
		fmt.Print("Enter message: ")
		fmt.Scanln(&message)
		//login
		// {"action":"login","data":{"user_id":"user2"}}
		// chat
		// {"action":"chat","data":{"conversation_id":1,"sender_id":"user2","content":"iamuser2","type":"private"}}
		// {"action":"chat","data":{"msgID":"1","timestamp":1635993600,"userID":"user2","sender":"user2","nickName":"John","faceUrl":"https://www.bugela.com/cjpic/frombd/1/253/1943132031/773911012.jpg","elemType":1,"textElem":{"text":"Hello,thisisamocktextmessage"}}}
		fmt.Println(message)
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Failed to send message:", err)
			break
		}
	}
}

func readMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}
		log.Println("Received message:", string(message))
	}
}
