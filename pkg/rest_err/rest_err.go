package rest_err

import "net/http"

const (
	ErrConflict            = "Resource already exists"
	ErrNotFound            = "Resource does not exist"
	ErrValidation          = "Validation error"
	ErrInternalServerError = "An internal error occurred"
)

type RestErr struct {
	StatusCode int
	Message    string
	Details    string
}

func NewBadRequest(message string, details string) *RestErr {
	return &RestErr{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Details:    details,
	}
}

func NewNotFound(message string) *RestErr {
	return &RestErr{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func NewInternalError(details string) *RestErr {
	return &RestErr{
		StatusCode: http.StatusInternalServerError,
		Message:    ErrInternalServerError,
		Details:    details,
	}
}

func NewUnprocessable(message string, details string) *RestErr {
	return &RestErr{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    message,
		Details:    details,
	}
}

func NewConflict(message string) *RestErr {
	return &RestErr{
		StatusCode: http.StatusConflict,
		Message:    message,
	}
}
