package main

import (
	"fmt"

	"gin-todo-sample/infrastructure/waf/gin"
	"gin-todo-sample/registry"
)

func main() {
	fmt.Println("gin todo list sample")

	todoController := registry.NewTodoController()
	userController := registry.NewUserController()
	server := gin.NewServer(todoController, userController)
	server.SetRoute()

	server.Run()
}
