package http

import "net/http"

type StatusCode string

const (
	CodeSuccess             StatusCode = "00"
	CodeConflict            StatusCode = "02"
	CodeUnprocessableEntity StatusCode = "03"
	CodeGeneralError        StatusCode = "99"
)

type StatusText string

const (
	StatusSuccess             StatusText = "SUCCESS"
	StatusNotFound            StatusText = "NOT_FOUND"
	StatusBadRequest          StatusText = "BAD_REQUEST"
	StatusUnauthorized        StatusText = "UNAUTHORIZED"
	StatusConflict            StatusText = "CONFLICT"
	StatusUnprocessableEntity StatusText = "UNPROCESSABLE_ENTITY"
	StatusInternalServerError StatusText = "INTERNAL_SERVER_ERROR"
)

func statusCode(code int) StatusCode {
	switch code {
	case http.StatusConflict:
		return CodeConflict
	case http.StatusUnprocessableEntity:
		return CodeUnprocessableEntity
	default:
		return CodeGeneralError
	}
}

func statusText(code int) StatusText {
	switch code {
	case http.StatusBadRequest:
		return StatusBadRequest
	case http.StatusUnauthorized:
		return StatusUnauthorized
	case http.StatusNotFound:
		return StatusNotFound
	case http.StatusConflict:
		return StatusConflict
	case http.StatusUnprocessableEntity:
		return StatusUnprocessableEntity
	case http.StatusInternalServerError:
		return StatusInternalServerError
	default:
		return StatusInternalServerError
	}
}
