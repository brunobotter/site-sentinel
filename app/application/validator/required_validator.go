package validator

import (
	"errors"
	"reflect"
)

type requiredValidator struct{}

func Required() Validator {
	return &requiredValidator{}
}

func (v *requiredValidator) Validate(field any) error {
	t := reflect.TypeOf(field)
	value := reflect.ValueOf(field)
	switch t.Kind() {
	case reflect.Slice, reflect.Array:
		if value.Len() == 0 {
			return errors.New(Field_Is_Required)
		}
	case reflect.String:
		if value.String() == "" {
			return errors.New(Field_Is_Required)
		}
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		if value.IsZero() {
			return errors.New(Field_Is_Required)
		}
	}
	return nil
}
