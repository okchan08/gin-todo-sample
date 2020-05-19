package port

import "gin-todo-sample/domain"

type UserOutputPort interface {
	Get(*domain.User) (*GetUserResponse, Error)
	Create(*domain.User) (*CreateUserResponse, Error)
}
