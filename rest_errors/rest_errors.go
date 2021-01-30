package rest_errors

import (
	"errors"
	"net/http"
)

// RestErr : Rest Error
type RestErr struct {
	Message    string        `json:"messgae"`
	StatusCode int           `json:"code"`
	Error      string        `json:"error"`
	Causes     []interface{} `json:"causes"`
}

// NewError : Return Error Message
func NewError(msg string) error {
	return errors.New(msg)
}

// NewBadRequestError : Generate Error Response
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "Bad Request",
	}
}

// NewNotFoundError : Generate Error Response
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      "not_found",
	}
}

// NewInternalServerError : Internal Server Error
func NewInternalServerError(message string, err error) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Error:      "internal_server_error",
	}
}
