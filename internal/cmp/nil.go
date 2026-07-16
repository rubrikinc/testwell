package cmp

import (
	"fmt"
	"reflect"
)

// NonNilableError is returned by Nil when the value's type can never be nil.
type NonNilableError struct {
	Msg string
}

func (e NonNilableError) Error() string {
	return e.Msg
}

// Nil returns ok=true if `nullable` is a nullable type or the zero value type.
func Nil(nullable any) (bool, error) {
	nullableVal := reflect.ValueOf(nullable)

	if !nullableVal.IsValid() {
		return true, nil
	}

	switch nullableVal.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
	default:
		return false, NonNilableError{
			Msg: fmt.Sprintf("type %T cannot be nil", nullable),
		}
	}

	// IsNil is safe: the switch above only falls through for nilable kinds;
	// reflect.Value.IsNil panics only on non-nilable kinds.
	return nullableVal.IsNil(), nil
}
