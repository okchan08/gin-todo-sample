package controller

import "gin-todo-sample/usecase/port"

type UserController struct {
	UserInputPort port.UserInputPort
}

func NewUserController(port port.UserInputPort) *UserController {
	return &UserController{
		UserInputPort: port,
	}
}

func (c *UserController) GetUser(params *port.GetUserRequest) (*port.GetUserResponse, port.Error) {
	return c.UserInputPort.Get(params)
}

func (c *UserController) CreateUser(params *port.CreateUserRequest) (*port.CreateUserResponse, port.Error) {
	return c.UserInputPort.Create(params)
}
