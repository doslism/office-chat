/*
Package routes is to config the route
*/
package routes

import (
	"doslism/office-chat/types"
	"doslism/office-chat/utils"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) init() {
	fmt.Println("init")
	s.router.Use(cors.Default())

	// todo: user register
	s.router.POST("/", func(ctx *gin.Context) {
		var body types.RegisterData
		err := ctx.BindJSON(&body)
		if err != nil {
			log.Fatalln("Failed to parse")
			return
		}
		uid := utils.GenerateNewID()
		fmt.Printf("username: %s \n", body.Username)
		fmt.Printf("password: %s \n", body.Password)
		fmt.Printf("uid: %s \n", uid)
	})
	s.router.GET("/", func(ctx *gin.Context) {
		fmt.Println(ctx.Request.Host)
		fmt.Println(2)
	})

	// todo: user login

	// todo: get current user info

	// todo: get user info by ID
}
