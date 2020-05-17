package port

type TodoInputPort interface {
	Get(*GetTodoRequest) (*GetTodoResponse, Error)
}
