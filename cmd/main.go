package main

import "github.com/dakong-yi/im-go-server/pkg/server"

func main() {
	server := server.NewServer("127.0.0.1", 1234)
	server.Start()
}
