package domain

import (
	"fmt"

	"github.com/hizzuu/plate-backend/conf"
	"github.com/pkg/errors"
)

const (
	UnauthorizedError   = "UNAUTHORIZED_ERROR"
	ForbiddenError      = "FORBIDDEN_ERROR"
	UnavailableError    = "UNAVAILABLE_ERROR"
	DBError             = "DB_ERROR"
	NotFoundError       = "NOT_FOUND_ERROR"
	ValidationError     = "VALIDATION_ERROR"
	BadRequestError     = "BAD_REQUEST_ERROR"
	InternalServerError = "INTERNAL_SERVER_ERROR"
	UnknownError        = "UNKNOWN_ERROR"
)

type StackTrace interface {
	StackTrace() errors.StackTrace
}

type Error interface {
	Error() string
	Code() string
	Extensions() map[string]interface{}
	Unwrap() error
}

func NewUnauthorizedError(e error) error {
	return newError(
		DBError,
		e.Error(),
		map[string]interface{}{
			"code": UnauthorizedError,
		},
		e,
	)
}

func NewForbiddenError(e error) error {
	return newError(
		DBError,
		e.Error(),
		map[string]interface{}{
			"code": ForbiddenError,
		},
		e,
	)
}

func NewUnavailableError(e error) error {
	return newError(
		DBError,
		e.Error(),
		map[string]interface{}{
			"code": UnauthorizedError,
		},
		e,
	)
}

func NewDBError(e error) error {
	return newError(
		DBError,
		e.Error(),
		map[string]interface{}{
			"code": DBError,
		},
		e,
	)
}

func NewNotFoundError(e error, value interface{}) error {
	return newError(
		NotFoundError,
		e.Error(),
		map[string]interface{}{
			"code":  NotFoundError,
			"value": value,
		},
		e,
	)
}

func NewInvalidParamError(e error, value interface{}) error {
	return newError(
		BadRequestError,
		e.Error(),
		map[string]interface{}{
			"code":  BadRequestError,
			"value": value,
		},
		e,
	)
}

func NewValidationError(e error) error {
	return newError(
		ValidationError,
		e.Error(),
		map[string]interface{}{
			"code": ValidationError,
		},
		e,
	)
}

func NewInternalServerError(e error) error {
	return newError(
		InternalServerError,
		e.Error(),
		map[string]interface{}{
			"code": InternalServerError,
		},
		e,
	)
}

type err struct {
	err        error
	code       string
	message    string
	extensions map[string]interface{}
}

func (e *err) Error() string                      { return e.message }
func (e *err) Code() string                       { return e.code }
func (e *err) Extensions() map[string]interface{} { return e.extensions }
func (e *err) Unwrap() error                      { return e.err }

func IsError(e error) bool {
	_, ok := e.(Error)
	return ok
}

func IsStackTrace(e error) bool {
	_, ok := e.(StackTrace)
	return ok
}

func newError(code string, message string, extensions map[string]interface{}, e error) error {
	newErr := &err{
		err:        e,
		code:       code,
		message:    message,
		extensions: extensions,
	}
	if IsStackTrace(e) {
		return newErr
	}

	return withStackTrace(newErr)
}

func withStackTrace(e error) error {
	ews := errors.WithStack(e)

	if conf.C.App.Debug {
		fmt.Printf("%+v\n", ews)
	}

	return ews
}
