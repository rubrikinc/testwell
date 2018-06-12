package cmp

import (
	"fmt"
	"reflect"
)

// NotComparableError is specific to Equal.
type NotComparableError struct {
	Msg string
}

func (e NotComparableError) Error() string {
	return e.Msg
}

func newNotComparableError(
	fmtstr string,
	args ...interface{},
) NotComparableError {
	return NotComparableError{
		Msg: fmt.Sprintf(fmtstr, args...),
	}
}

// Equal returns ok=true if `tval` is equal to `expected` using the Go `==`
// operator. Panics are recovered and returned as errors.
func Equal(expected interface{}, tval interface{}) (ok bool, err error) {

	expectedType := reflect.TypeOf(expected)
	testType := reflect.TypeOf(tval)

	if expectedType != nil && !expectedType.Comparable() {
		return false, newNotComparableError(
			"expected value type (%T) is not comparable",
			expected,
		)
	}

	if testType != nil && !testType.Comparable() {
		return false, newNotComparableError(
			"test value type (%T) is not comparable",
			tval,
		)
	}

	if expectedType != testType {
		return false, newNotComparableError(
			"expected value type (%T) and test value type (%T) differ",
			expected,
			tval,
		)
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

	return tval == expected, nil
}
