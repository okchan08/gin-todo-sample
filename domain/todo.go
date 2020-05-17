package domain

type TodoID int64

type Todo struct {
	TodoID  TodoID
	UserID  int64
	Content string
}
