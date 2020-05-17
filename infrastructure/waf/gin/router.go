package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin *gin.Engine
}

func NewServer() *Server {
	return &Server{
		Gin: gin.Default(),
	}
}

func (s *Server) SetRoute() {
	s.Gin.GET("/", func(ctx *gin.Context) {
		//ctx.JSON(message, "Hello! gin")
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
	})
}

func (s *Server) Run() {
	s.Gin.Run()
}
