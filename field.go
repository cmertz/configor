package configor

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var (
	// ErrInvalidInput signals that the input was invalid
	ErrInvalidInput = errors.New("invalid input")
)

type field struct {
	structPath []string
	path       []string
	t          reflect.Type
	v          reflect.Value
	isSet      bool
}

func fieldFor(i interface{}) field {
	return field{
		t: reflect.TypeOf(i).Elem(),
		v: reflect.ValueOf(i).Elem(),
	}
}

// nolint:exhaustive
func (f *field) set(value string) error {
	switch f.t.Kind() {
	case reflect.Bool:
		if value != "true" && value != "false" {
			return fmt.Errorf("%w: illegal bool value %s, legal value are \"true\" and \"false\"", ErrInvalidInput, value)
		}

		f.v.SetBool(value == "true")

	case reflect.Int:
		i, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("%w: atoi %v", err, i)
		}

		f.v.SetInt(int64(i))

	case reflect.String:
		f.v.SetString(value)

	default:
		return fmt.Errorf("%w: unknown type %s", ErrInvalidInput, f.t.Kind())
	}

	f.isSet = true

	return nil
}

func (f field) isStruct() bool {
	return f.t.Kind() == reflect.Struct
}

func (f field) fields() []field {
	if !f.isStruct() {
		return nil
	}

	var res []field
	for i := 0; i < f.t.NumField(); i++ {
		res = append(res, field{
			structPath: []string{f.t.Field(i).Name},
			path:       []string{f.t.Field(i).Name},
			t:          f.t.Field(i).Type,
			v:          f.v.Field(i),
		})
	}

	return res
}
