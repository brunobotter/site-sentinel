package validator

import "errors"

type Validator interface {
	Validate(field any) error
}

type ValidatorError struct {
	error
}

func NewValidatorError(err string) *ValidatorError {
	if err == "" {
		return nil
	}
	return &ValidatorError{error: errors.New(err)}
}
