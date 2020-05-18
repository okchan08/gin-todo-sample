package registry

import (
	"gin-todo-sample/domain"
	"gin-todo-sample/infrastructure/database"
	"gin-todo-sample/infrastructure/database/mysql"
	"gin-todo-sample/interface/controller"
	"gin-todo-sample/usecase/interactor"
)

func NewTodoController() *controller.TodoController {
	inputPort := newInputPort()
	return controller.NewTodoController(inputPort)
}

// temporary on memory repository implementation.
// You can swicth this implementation to another such as MySQL, NoSQL, etc.

type memoryTodoRepository struct{}

func (m *memoryTodoRepository) Find(id domain.TodoID) domain.Todo {
	if id == domain.TodoID(1) {
		return domain.Todo{
			TodoID:  1,
			UserID:  1,
			Content: "This is memory repository",
		}
	} else {
		return domain.Todo{
			TodoID:  2,
			UserID:  2,
			Content: "This is memory repository 2",
		}
	}
}

func newMemoryTodoRepository() *memoryTodoRepository {
	return &memoryTodoRepository{}
}

func newInputPort() *interactor.TodoInputPortImpl {
	return interactor.NewTodoInputPortImpl(
		newMemoryTodoRepository(),
		interactor.NewTodoOutputPortImpl(),
	)
}

// User controller registry
func NewUserController() *controller.UserController {
	inputPort := newUserInputPort()
	return &controller.UserController{
		UserInputPort: inputPort,
	}
}

func newUserInputPort() *interactor.UserInputPortImpl {
	return interactor.NewUserInputPortImpl(
		database.NewUserRepositoryImpl(mysql.NewMySQLHadler()),
		interactor.NewUserOutputPortImpl())
}
