package port

import "gin-todo-sample/domain"

type GetUserRequest struct {
	UserID domain.UserID
}

type GetUserResponse struct {
	User *domain.User
}
