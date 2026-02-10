package exception

import (
	"net/http"
)

type AppError struct {
	Code     int    // HTTP status
	Message  string // aman untuk client
	Internal error  // error asli (log only)
}

func (e *AppError) Error() string {
	return e.Message
}

func BadRequest(msg string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: msg,
		Internal: nil,
	}
}

func Unauthorized(msg string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: msg,
		Internal: nil,
	}
}

func NotFound(msg string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: msg,
		Internal: nil,
	}
}

func Internal(err error) *AppError {
	return &AppError{
		Code:     http.StatusInternalServerError,
		Message:  "internal server error",
		Internal: err,
	}
}