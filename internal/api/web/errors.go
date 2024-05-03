package web

import (
	"net/http"

	"github.com/quinntas/go-rest-template/internal/api/utils"
)

type HttpError struct {
	status  int
	message string
	body    utils.Map[interface{}]
}

func (e *HttpError) Error() string {
	return e.message
}

func (e *HttpError) ToJsonResponse(response http.ResponseWriter) {
	e.body["message"] = e.message
	JsonResponse(
		response,
		e.status,
		&e.body,
	)
}

func NewHttpError(status int, message string, body utils.Map[interface{}]) *HttpError {
	return &HttpError{
		status:  status,
		body:    body,
		message: message,
	}
}

func UnprocessableEntity() *HttpError {
	return NewHttpError(
		http.StatusUnprocessableEntity,
		"unprocessable entity",
		utils.Map[interface{}]{
			"hint": "please check the request's body",
		},
	)
}

func InternalError() *HttpError {
	return NewHttpError(
		http.StatusInternalServerError,
		"internal server error",
		utils.Map[interface{}]{},
	)
}

func BadRequest() *HttpError {
	return NewHttpError(
		http.StatusBadRequest,
		"bad request",
		utils.Map[interface{}]{},
	)
}

func NotFound() *HttpError {
	return NewHttpError(
		http.StatusNotFound,
		"not found",
		utils.Map[interface{}]{},
	)
}
