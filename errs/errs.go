package errs

import (
	"net/http"

	"github.com/pkg/errors"
)

type ErrorType uint

const (
	Unknown ErrorType = iota
	Invalidated
	Forbidden
	NotFound
	Conflict
	Failed
)

func StatusCode(err error) int {
	switch GetType(err) {
	case Invalidated:
		return http.StatusBadRequest
	case Forbidden:
		return http.StatusForbidden
	case NotFound:
		return http.StatusNotFound
	case Conflict:
		return http.StatusConflict
	case Failed:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}

type typeGetter interface {
	Type() ErrorType
}

type customError struct {
	errorType     ErrorType
	originalError error
}

func (et ErrorType) New(message string) error {
	return customError{errorType: et, originalError: errors.New(message)}
}

func (et ErrorType) Wrap(err error, message string) error {
	return customError{errorType: et, originalError: errors.Wrap(err, message)}
}

func (e customError) Error() string {
	return e.originalError.Error()
}

func (e customError) Type() ErrorType {
	return e.errorType
}

func Wrap(err error, message string) error {
	we := errors.Wrap(err, message)
	if ce, ok := err.(typeGetter); ok {
		return customError{errorType: ce.Type(), originalError: we}
	}
	return customError{errorType: Unknown, originalError: we}
}

func Cause(err error) error {
	return errors.Cause(err)
}

func GetType(err error) ErrorType {
	for {
		if e, ok := err.(typeGetter); ok {
			return e.Type()
		}
		break
	}
	return Unknown
}
