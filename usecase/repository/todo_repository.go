package repository

import "gin-todo-sample/domain"

type TodoRepository interface {
	Find(domain.TodoID) domain.Todo
}
