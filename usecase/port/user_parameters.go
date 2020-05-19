package port

import "gin-todo-sample/domain"

type GetUserRequest struct {
	UserID domain.UserID
}

type GetUserResponse struct {
	User *domain.User
}

type CreateUserRequest struct {
	User *domain.User
}

type CreateUserResponse struct {
	UserID   domain.UserID
	UserName string
	Email    string
}
