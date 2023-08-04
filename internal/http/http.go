package http

import (
	"log"

	"github.com/dakong-yi/im-go-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	router.Use(middleware.RequestLoggerMiddleware())
	return &Server{
		Router: router,
	}
}
func (s *Server) Start() error {
	err := s.Router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	return nil
}
