package rest_errors

import "net/http"

type RestErr struct {
	Status  int
	Message string
	Error   string
}

func NewBadRequest(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusNotFound,
		Message: message,
		Error:   "not_found",
	}
}
