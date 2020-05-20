package gin

import (
	"gin-todo-sample/interface/controller"

	"github.com/gin-gonic/gin"
)

// This file set up gin routings
// The gin callback functions are implemented in the file
// 'infrastructure/waf/gin/server.go'

type Server struct {
	Gin            *gin.Engine
	TodoController *controller.TodoController
	UserController *controller.UserController
}

func NewServer(
	todoController *controller.TodoController,
	userController *controller.UserController,
) *Server {
	g := gin.Default()
	g.Use(errorMiddleware())
	g.LoadHTMLGlob("templates/*.tmpl")
	return &Server{
		Gin:            g,
		TodoController: todoController,
		UserController: userController,
	}
}

func (s *Server) SetRoute() {
	s.Gin.GET("/", s.GetIndex)
	s.Gin.GET("/todo/:id", s.GetTodo)
	s.Gin.GET("/user/:id", s.GetUser)
	s.Gin.GET("/signup", s.GetSignup)
	s.Gin.POST("/signup", s.CreateUser)
}

func (s *Server) Run() {
	s.Gin.Run()
}
