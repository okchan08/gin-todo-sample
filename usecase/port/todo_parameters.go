package port

import "gin-todo-sample/domain"

type GetTodoRequest struct {
	TodoID domain.TodoID
}

type GetTodoResponse struct {
	Todo *domain.Todo
}
