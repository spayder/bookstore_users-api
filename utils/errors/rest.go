package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Error string `json:"error"`
}

func BadRequestError(message string) *RestErr  {
	return &RestErr{
		Message: message,
		Code: http.StatusBadRequest,
		Error: "bad_request",
	}
}

func NotFoundError(message string) *RestErr  {
	return &RestErr{
		Message: message,
		Code: http.StatusNotFound,
		Error: "not_found",
	}
}