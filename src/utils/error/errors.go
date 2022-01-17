package errors

import (
	"net/http"
)

type RestError struct {
	Message string `json:"messgae"`
	Error   string `json:"error"`
	Status  int    `json:"status"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "Bad Request",
		Status:  http.StatusBadRequest,
	}
}

func NewNotFoundtError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "Not Found",
		Status:  http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "Internal Server Error",
		Status:  http.StatusInternalServerError,
	}
}
