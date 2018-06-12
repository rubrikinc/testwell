package cmp

import (
	"fmt"
	"reflect"
)

// Empty returns ok=true if `container is empty. A nil pointer to a container
// type is empty. A container with zero element is empty.
func Empty(container interface{}) (ok bool, err error) {
	containerVal := reflect.ValueOf(container)

	switch containerVal.Kind() {
	case reflect.Array:
	case reflect.Chan:
	case reflect.Map:
	case reflect.Slice:
	case reflect.String:
	default:
		return false, fmt.Errorf("(%T) is not a container type", container)
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

	return containerVal.Len() == 0, nil
}
