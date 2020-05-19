package interactor

import (
	"gin-todo-sample/usecase/port"
	"gin-todo-sample/usecase/repository"
)

type UserInputPortImpl struct {
	UserRepository repository.UserRepository
	UserOutputPort port.UserOutputPort
}

func NewUserInputPortImpl(
	repository repository.UserRepository,
	outputPort port.UserOutputPort,
) *UserInputPortImpl {
	return &UserInputPortImpl{
		UserRepository: repository,
		UserOutputPort: outputPort,
	}
}

func (p *UserInputPortImpl) Get(request *port.GetUserRequest) (*port.GetUserResponse, port.Error) {
	user, err := p.UserRepository.FindOne(request.UserID)

	// TODO: refactoring to use UserOutputPort
	if err != nil {
		return &port.GetUserResponse{
			User: nil,
		}, err
	}

	return &port.GetUserResponse{
		User: &user,
	}, nil
}

func (p *UserInputPortImpl) Create(request *port.CreateUserRequest) (*port.CreateUserResponse, port.Error) {

	err := p.UserRepository.Create(*request.User)
	if err != nil {
		return nil, err
	}

	return p.UserOutputPort.Create(request.User)
}
