// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/port/todo_output_port.go

// Package mock_port is a generated GoMock package.
package mock_port

import (
	domain "gin-todo-sample/domain"
	port "gin-todo-sample/usecase/port"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTodoOutputPort is a mock of TodoOutputPort interface
type MockTodoOutputPort struct {
	ctrl     *gomock.Controller
	recorder *MockTodoOutputPortMockRecorder
}

// MockTodoOutputPortMockRecorder is the mock recorder for MockTodoOutputPort
type MockTodoOutputPortMockRecorder struct {
	mock *MockTodoOutputPort
}

// NewMockTodoOutputPort creates a new mock instance
func NewMockTodoOutputPort(ctrl *gomock.Controller) *MockTodoOutputPort {
	mock := &MockTodoOutputPort{ctrl: ctrl}
	mock.recorder = &MockTodoOutputPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTodoOutputPort) EXPECT() *MockTodoOutputPortMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockTodoOutputPort) Get(arg0 *domain.Todo) (*port.GetTodoResponse, port.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*port.GetTodoResponse)
	ret1, _ := ret[1].(port.Error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTodoOutputPortMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTodoOutputPort)(nil).Get), arg0)
}