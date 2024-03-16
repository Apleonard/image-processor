package errorutil

import (
	"errors"
	"fmt"
	"strings"
)

const (
	Default ErrorType = iota
	TypeBadRequest
	TypeUnauthorized
	TypeNotFound
	TypeConflict
	TypeInternalServerError
	TypeUnprocessableEntity
	TypeTooManyRequests
)

type errorLog struct {
	pkg       string
	throwerFn string
	causersFn string
}

func (e *Error) Error() string {
	return e.getMessage(e.Errs)
}

func (e *Error) getMessage(errs []error) string {
	var errMsg []string
	for _, err := range errs {
		errMsg = append(errMsg, err.Error())
	}
	return strings.Join(errMsg, ", ")
}

func HandlerError(throwerFn, causerFn string, err ...error) *Error {
	var appErr *Error
	switch {
	case errors.As(err[0], &appErr):
		appErr = newAppErr("handler", throwerFn, causerFn, err[0].(*Error).Type, err...)
	default:
		appErr = newAppErr("handler", throwerFn, causerFn, Default, err...)
	}
	return appErr
}

func (e *Error) ErrorWithStackTrace() string {
	return e.getMessageStackTrace(e.Errs)
}

func (e *Error) getMessageStackTrace(errs []error) string {
	var errMsg []string
	for _, err := range errs {
		errMsg = append(errMsg, fmt.Sprintf("%+v", err))
	}

	return strings.Join(errMsg, "\n")
}

func newAppErr(pkg, throwerFn, causerFn string, errType ErrorType, err ...error) *Error {
	errLog := &errorLog{pkg, throwerFn, causerFn}
	return &Error{errLog, errType, err}
}

type Error struct {
	*errorLog

	Type ErrorType
	Errs []error
}

type ErrorType uint8
