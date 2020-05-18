package port

type UserInputPort interface {
	Get(*GetUserRequest) (*GetUserResponse, Error)
}
