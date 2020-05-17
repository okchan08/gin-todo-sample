package controller

import (
	"gin-todo-sample/usecase/port"
)

type TodoController struct {
	TodoInputPort port.TodoInputPort
}

func NewTodoController(todoInputPort port.TodoInputPort) *TodoController {
	return &TodoController{
		TodoInputPort: todoInputPort,
	}
}

func (c *TodoController) GetTodo(params *port.GetTodoRequest) (*port.GetTodoResponse, port.Error) {
	return c.TodoInputPort.Get(params)
}
