package repository

import (
	"gin-todo-sample/domain"
	"gin-todo-sample/usecase/port"
)

type UserRepository interface {
	FindOne(domain.UserID) (domain.User, port.Error)
	FindOneByEmail(string) (domain.User, port.Error)
	Create(domain.User) port.Error
	Update(domain.User) port.Error
}
