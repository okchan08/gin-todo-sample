package interactor

import (
	"gin-todo-sample/usecase/port"
	"gin-todo-sample/usecase/repository"
)

type TodoInputPortImpl struct {
	TodoRepository repository.TodoRepository
	TodoOutputPort port.TodoOutputPort
}

func NewTodoInputPortImpl(repository repository.TodoRepository, outputPort port.TodoOutputPort) *TodoInputPortImpl {
	return &TodoInputPortImpl{
		TodoRepository: repository,
		TodoOutputPort: outputPort,
	}
}

func (i *TodoInputPortImpl) Get(param *port.GetTodoRequest) (*port.GetTodoResponse, port.Error) {
	todo := i.TodoRepository.Find(param.TodoID)
	return i.TodoOutputPort.Get(&todo)
}
