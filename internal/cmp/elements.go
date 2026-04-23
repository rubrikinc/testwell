package cmp

import (
	"fmt"
	"reflect"
)

// ElementsMatch returns true if listA and listB contain the same elements
// regardless of order. Both must be slices or arrays.
func ElementsMatch(listA, listB interface{}) (bool, error) {
	aVal := reflect.ValueOf(listA)
	bVal := reflect.ValueOf(listB)

	aKind := aVal.Kind()
	bKind := bVal.Kind()

	if aKind != reflect.Slice && aKind != reflect.Array {
		return false, fmt.Errorf("first argument must be a slice or array, got %T", listA)
	}
	if bKind != reflect.Slice && bKind != reflect.Array {
		return false, fmt.Errorf("second argument must be a slice or array, got %T", listB)
	}

	aLen := aVal.Len()
	bLen := bVal.Len()

	if aLen != bLen {
		return false, nil
	}
	if aLen == 0 {
		return true, nil
	}

	// O(n²) via DeepEqual — correct for all types including non-comparable ones.
	matched := make([]bool, bLen)
	for i := 0; i < aLen; i++ {
		aElem := aVal.Index(i).Interface()
		found := false
		for j := 0; j < bLen; j++ {
			if !matched[j] && reflect.DeepEqual(aElem, bVal.Index(j).Interface()) {
				matched[j] = true
				found = true
				break
			}
		}
		if !found {
			return false, nil
		}
	}
	return true, nil
}
