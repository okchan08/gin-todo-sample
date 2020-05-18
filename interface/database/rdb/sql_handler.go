package rdb

import "gin-todo-sample/usecase/port"

type SQLHandler interface {
	First(interface{}, ...interface{}) SQLHandler
	Find(interface{}, ...interface{}) SQLHandler
}

type Result interface {
	LastInsertId() (int64, port.Error)
	RowAffected() (int64, port.Error)
}

type Row interface {
	Scan(...interface{}) port.Error
	Next() bool
	Close() port.Error
}
