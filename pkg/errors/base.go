package errors

type Error struct {
	Code   int `json:"code"`
	Status int `json:"status"`
}

func (e *Error) ToObject() *Error {
	return &Error{
		Code:   e.Code,
		Status: e.Status,
	}
}

func (e *Error) GetCode() int {
	return e.Code
}

func (e *Error) GetStatus() int {
	return e.Status
}
