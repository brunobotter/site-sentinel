package validator

import (
	"errors"
	"fmt"
	"reflect"
)

type minNumberValidator struct {
	min int64
}

type maxNumberValidator struct {
	max int64
}

func MinNumber(min int64) Validator {
	return &minNumberValidator{min: min}
}

func MaxNumber(max int64) Validator {
	return &maxNumberValidator{max: max}
}

func (v *minNumberValidator) Validate(field any) error {
	value, ok := toInt64(field)
	if !ok {
		return errors.New("invalid number type")
	}

	if value < v.min {
		return errors.New(fmt.Sprintf("must be greater than or equal to %d", v.min))
	}

	return nil
}

func (v *maxNumberValidator) Validate(field any) error {
	value, ok := toInt64(field)
	if !ok {
		return errors.New("invalid number type")
	}

	if value > v.max {
		return errors.New(fmt.Sprintf("must be less than or equal to %d", v.max))
	}

	return nil
}

func toInt64(field any) (int64, bool) {
	value := reflect.ValueOf(field)

	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(value.Uint()), true
	}

	return 0, false
}
