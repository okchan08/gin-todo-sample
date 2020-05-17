package port

import "gin-todo-sample/domain"

type TodoOutputPort interface {
	Get(*domain.Todo) (*GetTodoResponse, Error)
}
