package errors

import (
	"errors"
	"net/http"
)

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func BadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Invalid Request",
	}
}

func NotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Resource not found",
	}
}

// Errors for internal use, ex. passing concrete error from repo to service
var (
	ErrNotFound = errors.New("not found")
)

func New(msg string) error {
	return errors.New(msg)
}
