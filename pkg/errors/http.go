package errors

import "net/http"

func APIError() *Error {
	return &Error{
		Code:   0,
		Status: http.StatusInternalServerError,
	}
}

func BadRequestError() *Error {
	return &Error{
		Code:   1,
		Status: http.StatusBadRequest,
	}
}

func ForbiddenError() *Error {
	return &Error{
		Code:   2,
		Status: http.StatusForbidden,
	}
}
