package gin

import (
	"encoding/gob"
	"gin-todo-sample/domain"
	"gin-todo-sample/interface/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

// Register some types for gob
// We need this because sessions use gob
// to encode some types before storing them to sessions
func initSessionSetting() {
	gob.Register(SessionInfo{})
	gob.Register(domain.UserID(0))
}

func (s *Server) SetRoute() {
	initSessionSetting()
	store := cookie.NewStore([]byte("secret"))
	s.Gin.Use(sessions.Sessions("mysession", store))

	s.Gin.GET("/", s.GetIndex)
	s.Gin.GET("/todo/:id", s.GetTodo)
	s.Gin.GET("/signup", s.GetSignup)
	s.Gin.POST("/signup", s.CreateUser)
	s.Gin.GET("/signin", s.GetSignIn)
	s.Gin.POST("/signin", s.PostSignIn)

	userPage := s.Gin.Group("/user")
	userPage.Use(IsAuthenticated())
	{
		userPage.GET("/:id", s.GetUser)
	}

	s.Gin.GET("/logout", s.Logout)
}

func (s *Server) Run() {
	s.Gin.Run()
}
