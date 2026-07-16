package cmp

import (
	"fmt"
	"reflect"
)

// NotComparableError is returned by Equal when a value's type is not comparable.
type NotComparableError struct {
	Msg string
}

func (e NotComparableError) Error() string {
	return e.Msg
}

// TypeMismatchError is returned by Equal when the two values have different types.
type TypeMismatchError struct {
	Msg string
}

func (e TypeMismatchError) Error() string {
	return e.Msg
}

func newNotComparableError(
	fmtstr string,
	args ...any,
) NotComparableError {
	return NotComparableError{
		Msg: fmt.Sprintf(fmtstr, args...),
	}
}

// Equal returns ok=true if `left` == `right` using the Go `==` operator.
// Panics are recovered and returned as errors.
func Equal(left any, right any) (ok bool, err error) {

	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	if leftType != nil && !leftType.Comparable() {
		return false, newNotComparableError(
			"type (%T) is not comparable",
			left,
		)
	}

	if rightType != nil && !rightType.Comparable() {
		return false, newNotComparableError(
			"type (%T) is not comparable",
			right,
		)
	}

	if leftType != rightType {
		return false, TypeMismatchError{
			Msg: fmt.Sprintf("type (%T) != type (%T)", left, right),
		}
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

	return left == right, nil
}
