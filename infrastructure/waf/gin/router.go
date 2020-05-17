package gin

import (
	"gin-todo-sample/domain"
	"gin-todo-sample/interface/controller"
	"gin-todo-sample/usecase/port"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin            *gin.Engine
	TodoController *controller.TodoController
}

func NewServer(todoController *controller.TodoController) *Server {
	g := gin.Default()
	g.Use(errorMiddleware())
	return &Server{
		Gin:            g,
		TodoController: todoController,
	}
}

func (s *Server) SetRoute() {
	s.Gin.GET("/:id", s.GetTodo)
}

func (s *Server) Run() {
	s.Gin.Run()
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

	ctx.JSON(http.StatusOK, gin.H{"todo": res.Todo})
}
