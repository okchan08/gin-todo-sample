package domain

import "time"

type UserID int64

type User struct {
	UserID    UserID
	Email     string
	UserName  string
	Password  string
	CreatedAt time.Time
}

func (User) TableName() string {
	return "USERS"
}
