package dto

type Error struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

func NewError(code int, message string) Error {
	return Error{
		StatusCode: code,
		Message:    message,
	}
}
