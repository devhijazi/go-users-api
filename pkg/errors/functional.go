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

/*
Adicionar mais tratativas de error
*PasswordError
*ValidationError
*PaginationError
*/
