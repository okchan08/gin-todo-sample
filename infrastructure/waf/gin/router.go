package gin

import (
	"gin-todo-sample/domain"
	"gin-todo-sample/interface/controller"
	"gin-todo-sample/usecase/port"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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

func (s *Server) GetIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.tmpl", gin.H{})
}

func (s *Server) GetTodo(ctx *gin.Context) {
	todoIDstr := ctx.Param("id")

	todoID, err := strconv.Atoi(todoIDstr)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	param := port.GetTodoRequest{TodoID: domain.TodoID(todoID)}
	res, err := s.TodoController.GetTodo(&param)

	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.HTML(http.StatusOK, "todo_index.tmpl", gin.H{"todo": res.Todo})
}

func (s *Server) GetUser(ctx *gin.Context) {
	userIDStr := ctx.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	param := port.GetUserRequest{UserID: domain.UserID(userID)}
	res, err := s.UserController.GetUser(&param)

	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"User": res.User})
}

func (s *Server) GetSignup(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.tmpl", gin.H{})
}

func (s *Server) CreateUser(ctx *gin.Context) {
	var form domain.User

	if err := ctx.Bind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"form": form})
		ctx.Abort()
	}

	// TODO ここ直したい
	form.CreatedAt = time.Now()

	request := port.CreateUserRequest{
		User: &form,
	}

	if _, err := s.UserController.CreateUser(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	ctx.Redirect(302, "/")
}
