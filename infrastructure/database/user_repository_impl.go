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

type UserRepositoryError struct {
	Message string
}

func (e *UserRepositoryError) Error() string {
	return e.Message
}

func (e *UserRepositoryError) StatusCode() int {
	return 407
}

func (repo *UserRepositoryImpl) FindOne(id domain.UserID) (domain.User, port.Error) {
	user := domain.User{}
	user.UserID = id
	if err := repo.SQLHandler.Where("user_id = ?", id).Find(&user).Error(); err != nil {
		return user, &UserRepositoryError{"Resource Not Found"}
	}

	return user, nil
}

func (repo *UserRepositoryImpl) FindOneByEmail(email string) (domain.User, port.Error) {
	user := domain.User{}
	if err := repo.SQLHandler.Where("email = ?", email).Find(&user).Error(); err != nil {
		return user, &UserRepositoryError{"Resource Not Found"}
	}

	return user, nil
}

func (repo *UserRepositoryImpl) Create(user domain.User) port.Error {
	if err := repo.SQLHandler.Create(user).Error(); err != nil {
		return &UserRepositoryError{"Failed to create record"}
	}
	return nil
}

func (repo *UserRepositoryImpl) Update(user domain.User) port.Error {
	return nil
}
