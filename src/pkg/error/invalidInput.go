package error

import (
	"net/http"
)

type InvalidInputError struct {
}

func (e *InvalidInputError) Error() string {
	return "Invalid Input"
}

func (e *InvalidInputError) Message() string {
	return "Invalid Input"
}

func (e *InvalidInputError) GetErrorCode() int {
	return http.StatusUnprocessableEntity
}
