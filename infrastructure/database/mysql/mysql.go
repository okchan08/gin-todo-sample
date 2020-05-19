package mysql

import (
	"fmt"
	"gin-todo-sample/interface/database/rdb"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MySQLHandler struct {
	Conn *gorm.DB
}

func NewMySQLHadler() rdb.SQLHandler {
	//var (
	//	host     = "localhost"
	//	port     = "3306"
	//	dbName   = "TODO_SAMPLE"
	//	user     = "test"
	//	password = "password"
	//)

	driverName := "mysql"
	//dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbName)
	dataSourceName := "test:password@tcp(localhost:3306)/TODO_SAMPLE?parseTime=true"

	conn, err := gorm.Open(driverName, dataSourceName)

	if err != nil {
		fmt.Printf("%v", err)
		log.Fatal("Failed to connect MySQL database!")
	}

	return &MySQLHandler{
		Conn: conn,
	}
}

func (handler *MySQLHandler) First(out interface{}, where ...interface{}) rdb.SQLHandler {
	newConn := handler.Conn.First(out, where...)
	return &MySQLHandler{
		Conn: newConn,
	}
}

func (handler *MySQLHandler) Find(out interface{}, where ...interface{}) rdb.SQLHandler {
	newConn := handler.Conn.Find(out, where...)
	return &MySQLHandler{
		Conn: newConn,
	}
}

func (handler *MySQLHandler) Where(query interface{}, args ...interface{}) rdb.SQLHandler {
	newConn := handler.Conn.Where(query, args...)
	return &MySQLHandler{
		Conn: newConn,
	}
}

func (handler *MySQLHandler) Create(value interface{}) rdb.SQLHandler {
	newConn := handler.Conn.Create(value)
	return &MySQLHandler{
		Conn: newConn,
	}
}

func (handler *MySQLHandler) Error() error {
	return handler.Conn.Error
}
