package interactor

import (
	"gin-todo-sample/domain"
	"gin-todo-sample/usecase/port"
	mockPort "gin-todo-sample/usecase/port/mock"
	mockRepository "gin-todo-sample/usecase/repository/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestTodoGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepository := mockRepository.NewMockTodoRepository(ctrl)
	mockOutputPort := mockPort.NewMockTodoOutputPort(ctrl)

	todo := domain.Todo{
		TodoID:  domain.TodoID(1),
		UserID:  2,
		Content: "test todo",
	}

	expect := port.GetTodoResponse{
		Todo: &todo,
	}

	mockRepository.EXPECT().Find(todo.TodoID).Return(todo).Times(1)
	mockOutputPort.EXPECT().Get(&todo).Return(&expect, nil).Times(1)

	interactor := NewTodoInputPortImpl(mockRepository, mockOutputPort)

	getTodoRequest := port.GetTodoRequest{
		TodoID: todo.TodoID,
	}
	result, err := interactor.Get(&getTodoRequest)

	if err != nil {
		t.Fail()
	}

	if diff := cmp.Diff(&expect, result); diff != "" {
		t.Fail()
	}

}
