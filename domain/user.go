package domain

import "time"

type UserID int64

type User struct {
	UserID    UserID `form:"username" binding:"required"`
	Email     string `form:"email" binding:"required"`
	UserName  string `form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	CreatedAt time.Time
}

func (User) TableName() string {
	return "USERS"
}
