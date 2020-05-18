package domain

import "time"

type UserID int64

type User struct {
	UserID    UserID
	email     string
	UserName  string
	Password  string
	createdAt time.Time
}
