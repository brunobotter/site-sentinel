package application

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	_errors "github.com/go-errors/errors"
)

var (
	NotFoundApplicationErrorType        = "NotFoundApplicationError"
	TimeoutExceededApplicationErrorType = "TimeoutExceededApplicationError"
	IntegrationApplicationErrorType     = "IntegrationApplicationError"
	ValidationApplicationErrorType      = "ValidationApplicationError"
	MaintenanceErrorType                = "MaintenanceError"
	ManyRequestsApplicationErrorType    = "ManyRequestsApplicationError"
	UnauthorizedApplicationErrorType    = "UnauthorizedApplicationError"
	ForbiddenApplicationErrorType       = "ForbiddenApplicationError"
	BadRequestErrorType                 = "BadRequestError"
)

const (
	FraudGeneric           string = "CO100"
	IntegrationGeneric     string = "CO200"
	TimeoutExceededGeneric string = "CO300"
	ValidationDomain       string = "CO400"
	ValidationRequired     string = "CO401"
	ServiceUnavailable     string = "CO600"
)

type NotFoundApplicationError struct {
	code string
	error
}

func (e NotFoundApplicationError) Code() string {
	return e.code
}

type TimeoutExceededApplicationError struct {
	code string
	error
}

func (e TimeoutExceededApplicationError) Code() string {
	return e.code
}

type IntegrationApplicationError struct {
	code string
	error
}

func (e IntegrationApplicationError) Code() string {
	return e.code
}

type BadRequestApplicationError struct {
	code string
	error
}

func (e BadRequestApplicationError) Code() string {
	return e.code
}

type ValidationApplicationError struct {
	code string
	error
}

func (e ValidationApplicationError) Code() string {
	return e.code
}

type MaintenanceError struct {
	code string
	error
}

func IsMaintenanceError(err error) bool {
	if err == nil {
		return false
	}
	var maintenanceErr MaintenanceError
	return errors.As(err, &maintenanceErr)
}

func (e MaintenanceError) Code() string {
	return e.code
}

type ManyRequestsApplicationError struct {
	code string
	error
}

func (e ManyRequestsApplicationError) Code() string {
	return e.code
}

func IsManyRequestsError(err error) bool {
	if err == nil {
		return false
	}
	var manyRequestsErr ManyRequestsApplicationError
	return errors.As(err, &manyRequestsErr)

}

type BadRequestError struct {
	code string
	error
}

func (e BadRequestError) Code() string {
	return e.code
}

type UnauthorizedApplicationError struct {
	code string
	error
}

func (e UnauthorizedApplicationError) Code() string {
	return e.code
}

type ForbiddenApplicationError struct {
	code string
	error
}

func (e ForbiddenApplicationError) Code() string {
	return e.code
}

type errorWrapper interface {
	Error() string
	GetOriginalError() error
}

type WrappedError struct {
	originalError error
	path          string
	messages      []string
}

func (err WrappedError) Error() string {
	if len(err.messages) > 0 {
		retVal := fmt.Sprintf("%s: ", err.path)

		for _, message := range err.messages {
			retVal += message + "; "
		}

		return fmt.Sprintf("%s => %v", retVal, err.originalError)
	}

	return fmt.Sprintf("%s => %v", err.path, err.originalError)
}

func (err WrappedError) GetOriginalError() error {
	if err.originalError != nil {
		originalError, ok := (err.originalError).(errorWrapper)
		if ok {
			return originalError.GetOriginalError()
		}
	}

	return err.originalError
}

func (err WrappedError) Unwrap() error {
	return err.originalError
}

func Wrap(err error, messages ...string) error {
	return &WrappedError{
		originalError: err,
		path:          caller(),
		messages:      messages,
	}
}

func GetOriginalError(err error) error {
	wrappedErr, ok := err.(errorWrapper)
	if ok {
		return wrappedErr.GetOriginalError()
	}

	return err
}

func caller() string {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	funcRef := runtime.FuncForPC(pc[0])

	pathArr := strings.Split(funcRef.Name(), "/")

	return pathArr[len(pathArr)-1]
}

func NewIntegrationApplicationError(code string, err error) error {
	return newApplicationError(code, IntegrationApplicationErrorType, err)
}

func NewTimeoutExceededApplicationError(code string, err error) error {
	return newApplicationError(code, TimeoutExceededApplicationErrorType, err)
}

func NewNotFoundApplicationError(code string, err error) error {
	return newApplicationError(code, NotFoundApplicationErrorType, err)
}

func NewBadRequestError(code string, err error) error {
	return newApplicationError(code, BadRequestErrorType, err)
}

func NewValidationApplicationError(code string, err error) error {
	return newApplicationError(code, ValidationApplicationErrorType, err)
}

func NewMaintenanceError(code string, err error) error {
	return newApplicationError(code, MaintenanceErrorType, err)
}

func NewManyRequestsApplicationError(code string, err error) error {
	return newApplicationError(code, ManyRequestsApplicationErrorType, err)
}

func NewUnauthorizedApplicationError(code string, err error) error {
	return newApplicationError(code, UnauthorizedApplicationErrorType, err)
}

func NewForbiddenApplicationError(code string, err error) error {
	return newApplicationError(code, ForbiddenApplicationErrorType, err)
}

func newApplicationError(code, errType string, err error) error {
	if err == nil {
		return err
	}
	switch errType {
	case NotFoundApplicationErrorType:
		return _errors.Wrap(NotFoundApplicationError{code, err}, 1)
	case TimeoutExceededApplicationErrorType:
		return _errors.Wrap(TimeoutExceededApplicationError{code, err}, 1)
	case IntegrationApplicationErrorType:
		return _errors.Wrap(IntegrationApplicationError{code, err}, 1)
	case ValidationApplicationErrorType:
		return _errors.Wrap(ValidationApplicationError{code, err}, 1)
	case MaintenanceErrorType:
		return _errors.Wrap(MaintenanceError{code: code, error: err}, 1)
	case ManyRequestsApplicationErrorType:
		return _errors.Wrap(ManyRequestsApplicationError{code: code, error: err}, 1)
	case UnauthorizedApplicationErrorType:
		return _errors.Wrap(UnauthorizedApplicationError{code: code, error: err}, 1)
	case ForbiddenApplicationErrorType:
		return _errors.Wrap(ForbiddenApplicationError{code: code, error: err}, 1)
	case BadRequestErrorType:
		return _errors.Wrap(BadRequestError{code: code, error: err}, 1)
	default:
		return _errors.Wrap(err, 1)
	}
}
