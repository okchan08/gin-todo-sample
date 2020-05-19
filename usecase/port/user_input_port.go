package port

type UserInputPort interface {
	Get(*GetUserRequest) (*GetUserResponse, Error)
	Create(*CreateUserRequest) (*CreateUserResponse, Error)
}
