package cmp

import (
	"fmt"
	"reflect"
)

// Nil returns ok=true if `nullable` is a nullable type or the zero value type.
func Nil(nullable interface{}) (ok bool, err error) {
	nullableVal := reflect.ValueOf(nullable)

	if !nullableVal.IsValid() {
		return true, nil
	}

	switch nullableVal.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
	default:
		return false, nil
	}

	defer func() {
		if r := recover(); r != nil {
			if recErr, recOK := r.(error); !recOK {
				err = fmt.Errorf("%v", r)
			} else {
				err = recErr
			}
		}
	}()

	return nullableVal.IsNil(), nil
}
