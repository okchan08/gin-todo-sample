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

	if err != nil {
		return &port.GetUserResponse{
			User: nil,
		}, err
	}

	return &port.GetUserResponse{
		User: &user,
	}, nil
}
