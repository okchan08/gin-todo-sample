package database

import (
	"gin-todo-sample/domain"
	"gin-todo-sample/interface/database/rdb"
	"gin-todo-sample/usecase/port"
)

type UserRepositoryImpl struct {
	SQLHandler rdb.SQLHandler
}

func NewUserRepositoryImpl(handler rdb.SQLHandler) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		SQLHandler: handler,
	}
}

type UserRepositoryError struct{}

func (e *UserRepositoryError) Error() string {
	return "user repository error"
}

func (e *UserRepositoryError) StatusCode() int {
	return 407
}

func (repo *UserRepositoryImpl) FindOne(id domain.UserID) (domain.User, port.Error) {
	user := domain.User{}
	user.UserID = id
	if err := repo.SQLHandler.Where("user_id = ?", id).Find(&user).GetErrors(); len(err) > 0 {
		return user, &UserRepositoryError{}
	}

	return user, nil
}

func (repo *UserRepositoryImpl) Create(user domain.User) port.Error {
	return nil
}

func (repo *UserRepositoryImpl) Update(user domain.User) port.Error {
	return nil
}
