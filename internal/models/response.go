package models

import (
	"net/http"
)

func ResponseBadGateway(errors []string) Response {
	return Response{
		Status:  http.StatusBadGateway,
		Message: "Bad gateway",
		Errors:  errors,
		Data:    nil,
	}
}

func ResponseOK(data interface{}) Response {
	return Response{
		Status:  http.StatusOK,
		Message: "Ok",
		Errors:  []string{},
		Data:    data,
	}
}

func ResponseCreated(data interface{}) Response {
	return Response{
		Status:  http.StatusCreated,
		Message: "Created",
		Errors:  []string{},
		Data:    data,
	}
}

func ResponseBadRequest(errors []string) Response {
	return Response{
		Status:  http.StatusBadRequest,
		Message: "Bad request",
		Errors:  errors,
		Data:    nil,
	}
}
