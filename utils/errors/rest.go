package errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

func BadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func Error(message string) error {
	return errors.New(message)
}

func NotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "not_found",
	}
}

func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
