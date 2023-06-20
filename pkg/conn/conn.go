package conn

import (
	"fmt"
	"net"
)

type User struct {
	ID      string
	Addr    string
	Message chan string
	Conn    net.Conn
}

func (user *User) ListenMessage() {
	for {
		msg := <-user.Message
		_, err := fmt.Fprintf(user.Conn, "%s\n", msg)
		if err != nil {
			fmt.Printf("Error sending message to user %s: %v", user.ID, err)
			return
		}
	}
}

type Group struct {
	Users []*User
}
