package main

import (
	"fmt"

	"gin-todo-sample/infrastructure/waf/gin"
)

func main() {
	fmt.Println("gin todo list sample")

	server := gin.NewServer()
	server.SetRoute()

	server.Run()
}
