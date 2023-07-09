package errors

import "net/http"

func TokenError() *Error {
	return &Error{
		Code:   1000,
		Status: http.StatusUnauthorized,
	}
}

func AuthenticationError() *Error {
	return &Error{
		Code:   1001,
		Status: http.StatusBadRequest,
	}
}

func PasswordError() *Error {
	return &Error{
		Code:   1002,
		Status: http.StatusBadRequest,
	}
}

func ValidationError() *Error {
	return &Error{
		Code:   1003,
		Status: http.StatusBadRequest,
	}
}

func PaginationError() *Error {
	return &Error{
		Code:   1004,
		Status: http.StatusBadRequest,
	}
}
