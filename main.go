/*
Package main is to run the server
*/
package main

import (
	"doslism/office-chat/routes"
)

func main() {
	server := routes.CreateServer()
	server.LisenOnPort(14000)
}
