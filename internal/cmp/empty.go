package cmp

import (
	"fmt"
	"reflect"
)

// Empty returns ok=true if `container` is empty. A nil pointer to a container
// type is empty. A container with zero element is empty.
func Empty(container interface{}) (bool, error) {
	containerVal := reflect.ValueOf(container)

	switch containerVal.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
	default:
		return false, fmt.Errorf("(%T) is not a container type", container)
	}

	// Len is safe: the switch above only falls through for Array/Chan/Map/Slice/String,
	// all of which support Len; nil Chan/Map/Slice return 0 rather than panicking.
	return containerVal.Len() == 0, nil
}
