package gin

import (
	"fmt"
	"gin-todo-sample/domain"
	password "gin-todo-sample/infrastructure/security"
	"gin-todo-sample/usecase/port"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetIndex is called to return top page for the service
func (s *Server) GetIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.tmpl", gin.H{})
}

// GetTodo will return each todo detail page
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

// GetUser will return user page
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

// GetSignup will show signup page for new users
func (s *Server) GetSignup(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.tmpl", gin.H{})
}

// CreateUser will POST data in the form
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

	hashedPass, err := password.HashPassword(form.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	form.Password = hashedPass

	if _, err := s.UserController.CreateUser(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	ctx.Redirect(http.StatusFound, "/")
}

func (s *Server) GetSignIn(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signin.tmpl", gin.H{})
}

func (s *Server) PostSignIn(ctx *gin.Context) {
	email := ctx.PostForm("email")
	pass := ctx.PostForm("password")

	request := &port.GetUserRequest{Email: email}
	userResponse, err := s.UserController.GetUser(request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		ctx.Abort()
	}

	user := userResponse.User
	if !password.ComparePasswordAndHash(pass, user.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		ctx.Abort()
	}

	// create session
	session := sessions.Default(ctx)
	session.Set("UserID", user.UserID)
	session.Save()

	url := fmt.Sprintf("/user/%d", user.UserID)
	ctx.Redirect(http.StatusFound, url)
}

func (s *Server) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{"message": "logout"})
}
