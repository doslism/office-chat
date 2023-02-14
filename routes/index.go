package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func CreateServer() *Server {
	router := gin.Default()
	server := Server{
		router: router,
	}
	server.init()
	return &server
}

func (s *Server) LisenOnPort(port int) {
	s.router.Run(fmt.Sprintf("0.0.0.0:%d", port))
}