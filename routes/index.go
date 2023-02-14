package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server is server to run
type Server struct {
	router *gin.Engine
}

// CreateServer is to create a new http server
func CreateServer() *Server {
	router := gin.Default()
	server := Server{
		router: router,
	}
	server.init()
	return &server
}

// LisenOnPort is to run the server on a given port
func (s *Server) LisenOnPort(port int) {
	s.router.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
