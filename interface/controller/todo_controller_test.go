package controller

import (
	"gin-todo-sample/domain"
	"gin-todo-sample/usecase/port"
	mock_port "gin-todo-sample/usecase/port/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestTodoControllerGetTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInputPort := mock_port.NewMockTodoInputPort(ctrl)
	target := NewTodoController(mockInputPort)

	todo := domain.Todo{
		TodoID:  domain.TodoID(1),
		UserID:  2,
		Content: "test todo item",
	}
	request := port.GetTodoRequest{
		TodoID: domain.TodoID(todo.TodoID),
	}

	response := port.GetTodoResponse{
		Todo: &todo,
	}

	mockInputPort.EXPECT().Get(&request).Return(&response, nil).Times(1)

	result, err := target.GetTodo(&request)
	if err != nil {
		t.Log("err is not nil value!")
		t.Fail()
	}
	if diff := cmp.Diff(result, &response); diff != "" {
		t.Log(diff)
		t.Fail()
	}
}
