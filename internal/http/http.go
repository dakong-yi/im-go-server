package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		Router: gin.Default(),
	}
}
func (s *Server) Start() error {
	err := s.Router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	return nil
}
