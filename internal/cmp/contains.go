package cmp

import (
	"fmt"
	"reflect"
	"strings"
)

// ContainsError is specific to Contains. The attribute Hint should help the
// user understand the reason behind the error. This is used to differentiate
// with any other runtime error outside our control.
type ContainsError struct {
	Hint string
}

func (e ContainsError) Error() string {
	return fmt.Sprintf("ContainsError: %s", e.Hint)
}

func newContainsError(fmtstr string, args ...interface{}) ContainsError {
	return ContainsError{
		Hint: fmt.Sprintf(fmtstr, args...),
	}
}

// Contains returns ok=true if `elem` is contained in `container`. The
// container type is introspected at runtime. The type of elem must match the
// type of elements in the container.
// Supported container types: string, map (keys), array, slice.
func Contains(elem interface{}, container interface{}) (ok bool, err error) {
	containerVal := reflect.ValueOf(container)
	elemVal := reflect.ValueOf(elem)

	defer func() {
		if r := recover(); r != nil {
			if recErr, recOK := r.(error); !recOK {
				err = fmt.Errorf("%v", r)
			} else {
				err = recErr
			}
		}
	}()

	if container == nil {
		return false, newContainsError("container is nil")
	}

	switch reflect.TypeOf(container).Kind() {
	case reflect.String:
		if elemVal.Type() != containerVal.Type() {
			return false, newContainsError(
				"container is a (string) instead of a collection of (%T)", elem)
		}
		return strings.Contains(containerVal.String(), elemVal.String()), nil
	case reflect.Map:
		if containerVal.Type().Key() != elemVal.Type() {
			return false, newContainsError(
				"element type (%T) != container key type (%v)",
				elem, containerVal.Type().Key())
		}
		return containerVal.MapIndex(elemVal).IsValid(), nil
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		for i := 0; i < containerVal.Len(); i++ {
			if containerVal.Index(i).Interface() == elem {
				return true, nil
			}
		}
		return false, nil
	}

	return false, newContainsError("(%s) is not a container",
		containerVal.Type())
}
