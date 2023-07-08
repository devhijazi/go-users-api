package errors

import "net/http"

func UserNotFoundError() *Error {
	return &Error{
		Code:   2000,
		Status: http.StatusNotFound,
	}
}
