package main

import (
	"fmt"

	"gin-todo-sample/infrastructure/waf/gin"
	"gin-todo-sample/registry"
)

func main() {
	fmt.Println("gin todo list sample")

	todoController := registry.NewTodoController()
	server := gin.NewServer(todoController)
	server.SetRoute()

	server.Run()
}
