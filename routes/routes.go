/*
Package routes is to config the route
*/
package routes

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) init() {
	fmt.Println("init")
	s.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4900"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// todo: user register
	s.router.POST("/users/new", func(ctx *gin.Context) {
		fmt.Println(ctx)
	})

	// todo: user login

	// todo: get current user info

	// todo: get user info by ID
}
