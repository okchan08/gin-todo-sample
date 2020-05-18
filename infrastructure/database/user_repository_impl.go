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

func (repo *UserRepositoryImpl) FindOne(id domain.UserID) domain.User {
	user := domain.User{}
	user.UserID = id
	repo.SQLHandler.First(user)

	return user
}

func (repo *UserRepositoryImpl) Create(user domain.User) port.Error {
	return nil
}

func (repo *UserRepositoryImpl) Update(user domain.User) port.Error {
	return nil
}
