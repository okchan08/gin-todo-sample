package interactor

import (
	"gin-todo-sample/domain"
	"gin-todo-sample/usecase/port"
)

type TodoOutputPortImpl struct{}

func (s *TodoOutputPortImpl) Get(todo *domain.Todo) (*port.GetTodoResponse, port.Error) {
	return &port.GetTodoResponse{
		Todo: todo,
	}, nil
}

func NewTodoOutputPortImpl() *TodoOutputPortImpl {
	return &TodoOutputPortImpl{}
}
