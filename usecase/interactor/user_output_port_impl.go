package interactor

import (
	"gin-todo-sample/domain"
	"gin-todo-sample/usecase/port"
)

type UserOutputPortImpl struct{}

func NewUserOutputPortImpl() *UserOutputPortImpl {
	return &UserOutputPortImpl{}
}

func (p *UserOutputPortImpl) Get(user *domain.User) (*port.GetUserResponse, port.Error) {
	return &port.GetUserResponse{
		User: user,
	}, nil
}
