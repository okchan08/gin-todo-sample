package errors

import "fmt"

type ErrorUnknown struct {
	Message    string
	StatusCode int
}

func (e *ErrorUnknown) Error() string {
	return fmt.Sprintf("Unknown error!: %s", e.Message)
}

func (e *ErrorUnknown) StatusCode() int {
	return e.StatusCode
}
