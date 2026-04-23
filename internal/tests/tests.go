// @generated

package tests

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strings"

	"github.com/rubrikinc/testwell/internal/cmp"
	"github.com/rubrikinc/testwell/internal/fail"
	"github.com/rubrikinc/testwell/testing"
)

// True tests if val is True.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func True(t testing.T, tval bool, msg ...interface{}) bool {
	t.Helper()

	if tval {
		return true
	}
	return failTest(
		t,
		fail.Failure("True").
			Reason(`expected "true", got "%v" instead`, tval).
			ExtraMsg(msg...),
	)
}

// False tests if val is False.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func False(t testing.T, tval bool, msg ...interface{}) bool {
	t.Helper()

	if !tval {
		return true
	}
	return failTest(
		t,
		fail.Failure("False").
			Reason(`expected "false", got "%v" instead`, tval).
			ExtraMsg(msg...),
	)
}

// Equal tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two interface{}.
// This is a strict equality test. It will fail if `left` and `right`
// have different types. For example, int32(42) is not equal to int64(42).
// You can use the typed versions of `Equal` for more static typing.
// Note that Slice, map, and function values are not comparable.
// See also `DeepEqual`.
func Equal(t testing.T,
	left interface{},
	right interface{},
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("Equal")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqual tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two interface{}
// This is a strict non-equality test. It will pass only when
// `left` and `right` have identical type but non-equal values.
// For example, given int32(42) and int64(42), the test will fail because
// the types are not comparable. Given int32(42) and int32(42), the test
// will fail because the two values are equal.
// You can use the typed versions of `NotEqual` for more static typing.
// Note that Slice, map, and function values are not comparable.
// See also `NotDeepEqual`.
func NotEqual(t testing.T,
	left interface{},
	right interface{},
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqual")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualBool tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two bool.
func EqualBool(t testing.T,
	left bool,
	right bool,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualBool")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualBool tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two bool
func NotEqualBool(t testing.T,
	left bool,
	right bool,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualBool")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualByte tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two byte.
func EqualByte(t testing.T,
	left byte,
	right byte,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualByte")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualByte tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two byte
func NotEqualByte(t testing.T,
	left byte,
	right byte,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualByte")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualComplex128 tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two complex128.
func EqualComplex128(t testing.T,
	left complex128,
	right complex128,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualComplex128")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualComplex128 tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two complex128
func NotEqualComplex128(t testing.T,
	left complex128,
	right complex128,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualComplex128")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualComplex64 tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two complex64.
func EqualComplex64(t testing.T,
	left complex64,
	right complex64,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualComplex64")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualComplex64 tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two complex64
func NotEqualComplex64(t testing.T,
	left complex64,
	right complex64,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualComplex64")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualError tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two error.
func EqualError(t testing.T,
	left error,
	right error,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualError")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualError tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two error
func NotEqualError(t testing.T,
	left error,
	right error,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualError")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualFloat32 tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two float32.
func EqualFloat32(t testing.T,
	left float32,
	right float32,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualFloat32")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualFloat32 tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two float32
func NotEqualFloat32(t testing.T,
	left float32,
	right float32,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualFloat32")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualFloat64 tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two float64.
func EqualFloat64(t testing.T,
	left float64,
	right float64,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualFloat64")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualFloat64 tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two float64
func NotEqualFloat64(t testing.T,
	left float64,
	right float64,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualFloat64")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualInt tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int.
func EqualInt(t testing.T,
	left int,
	right int,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualInt")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualInt tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int
func NotEqualInt(t testing.T,
	left int,
	right int,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualInt")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualInt16 tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int16.
func EqualInt16(t testing.T,
	left int16,
	right int16,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualInt16")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualInt16 tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int16
func NotEqualInt16(t testing.T,
	left int16,
	right int16,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualInt16")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualInt32 tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int32.
func EqualInt32(t testing.T,
	left int32,
	right int32,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualInt32")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualInt32 tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int32
func NotEqualInt32(t testing.T,
	left int32,
	right int32,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualInt32")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualInt64 tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int64.
func EqualInt64(t testing.T,
	left int64,
	right int64,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualInt64")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualInt64 tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int64
func NotEqualInt64(t testing.T,
	left int64,
	right int64,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualInt64")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualInt8 tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int8.
func EqualInt8(t testing.T,
	left int8,
	right int8,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualInt8")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualInt8 tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int8
func NotEqualInt8(t testing.T,
	left int8,
	right int8,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualInt8")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualRune tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two rune.
func EqualRune(t testing.T,
	left rune,
	right rune,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualRune")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualRune tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two rune
func NotEqualRune(t testing.T,
	left rune,
	right rune,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualRune")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualString tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two string.
func EqualString(t testing.T,
	left string,
	right string,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualString")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualString tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two string
func NotEqualString(t testing.T,
	left string,
	right string,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualString")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualUint tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint.
func EqualUint(t testing.T,
	left uint,
	right uint,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualUint")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualUint tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint
func NotEqualUint(t testing.T,
	left uint,
	right uint,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualUint")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualUint16 tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint16.
func EqualUint16(t testing.T,
	left uint16,
	right uint16,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualUint16")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualUint16 tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint16
func NotEqualUint16(t testing.T,
	left uint16,
	right uint16,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualUint16")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualUint32 tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint32.
func EqualUint32(t testing.T,
	left uint32,
	right uint32,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualUint32")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualUint32 tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint32
func NotEqualUint32(t testing.T,
	left uint32,
	right uint32,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualUint32")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualUint64 tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint64.
func EqualUint64(t testing.T,
	left uint64,
	right uint64,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualUint64")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualUint64 tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint64
func NotEqualUint64(t testing.T,
	left uint64,
	right uint64,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualUint64")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualUint8 tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint8.
func EqualUint8(t testing.T,
	left uint8,
	right uint8,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualUint8")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualUint8 tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint8
func NotEqualUint8(t testing.T,
	left uint8,
	right uint8,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualUint8")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualUintptr tests if `left` is equal to `right` using the `==`
// operator.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uintptr.
func EqualUintptr(t testing.T,
	left uintptr,
	right uintptr,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("EqualUintptr")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason("values should be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualUintptr tests if `left` is not equal to `right` using the
// `!=` operator. The left value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uintptr
func NotEqualUintptr(t testing.T,
	left uintptr,
	right uintptr,
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotEqualUintptr")

	if ok, err := cmp.Equal(left, right); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(left) == nil || reflect.TypeOf(right) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason("values should not be equal").
		LeftValue(left).RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// GreaterByte tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterByte(t testing.T,
	left byte,
	right byte,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterByte").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualByte tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualByte(t testing.T,
	left byte,
	right byte,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualByte").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessByte tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessByte(t testing.T,
	left byte,
	right byte,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessByte").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualByte tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualByte(t testing.T,
	left byte,
	right byte,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualByte").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterFloat32 tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterFloat32(t testing.T,
	left float32,
	right float32,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterFloat32").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualFloat32 tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualFloat32(t testing.T,
	left float32,
	right float32,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualFloat32").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessFloat32 tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessFloat32(t testing.T,
	left float32,
	right float32,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessFloat32").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualFloat32 tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualFloat32(t testing.T,
	left float32,
	right float32,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualFloat32").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterFloat64 tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterFloat64(t testing.T,
	left float64,
	right float64,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterFloat64").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualFloat64 tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualFloat64(t testing.T,
	left float64,
	right float64,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualFloat64").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessFloat64 tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessFloat64(t testing.T,
	left float64,
	right float64,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessFloat64").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualFloat64 tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualFloat64(t testing.T,
	left float64,
	right float64,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualFloat64").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterInt tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterInt(t testing.T,
	left int,
	right int,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterInt").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualInt tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualInt(t testing.T,
	left int,
	right int,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualInt").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessInt tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessInt(t testing.T,
	left int,
	right int,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessInt").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualInt tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualInt(t testing.T,
	left int,
	right int,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualInt").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterInt16 tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterInt16(t testing.T,
	left int16,
	right int16,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterInt16").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualInt16 tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualInt16(t testing.T,
	left int16,
	right int16,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualInt16").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessInt16 tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessInt16(t testing.T,
	left int16,
	right int16,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessInt16").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualInt16 tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualInt16(t testing.T,
	left int16,
	right int16,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualInt16").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterInt32 tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterInt32(t testing.T,
	left int32,
	right int32,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterInt32").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualInt32 tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualInt32(t testing.T,
	left int32,
	right int32,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualInt32").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessInt32 tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessInt32(t testing.T,
	left int32,
	right int32,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessInt32").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualInt32 tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualInt32(t testing.T,
	left int32,
	right int32,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualInt32").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterInt64 tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterInt64(t testing.T,
	left int64,
	right int64,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterInt64").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualInt64 tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualInt64(t testing.T,
	left int64,
	right int64,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualInt64").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessInt64 tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessInt64(t testing.T,
	left int64,
	right int64,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessInt64").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualInt64 tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualInt64(t testing.T,
	left int64,
	right int64,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualInt64").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterInt8 tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterInt8(t testing.T,
	left int8,
	right int8,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterInt8").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualInt8 tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualInt8(t testing.T,
	left int8,
	right int8,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualInt8").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessInt8 tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessInt8(t testing.T,
	left int8,
	right int8,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessInt8").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualInt8 tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualInt8(t testing.T,
	left int8,
	right int8,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualInt8").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterRune tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterRune(t testing.T,
	left rune,
	right rune,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterRune").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualRune tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualRune(t testing.T,
	left rune,
	right rune,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualRune").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessRune tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessRune(t testing.T,
	left rune,
	right rune,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessRune").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualRune tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualRune(t testing.T,
	left rune,
	right rune,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualRune").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterString tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterString(t testing.T,
	left string,
	right string,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterString").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualString tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualString(t testing.T,
	left string,
	right string,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualString").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessString tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessString(t testing.T,
	left string,
	right string,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessString").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualString tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualString(t testing.T,
	left string,
	right string,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualString").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterUint tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterUint(t testing.T,
	left uint,
	right uint,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterUint").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualUint tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualUint(t testing.T,
	left uint,
	right uint,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualUint").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessUint tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessUint(t testing.T,
	left uint,
	right uint,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessUint").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualUint tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualUint(t testing.T,
	left uint,
	right uint,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualUint").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterUint16 tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterUint16(t testing.T,
	left uint16,
	right uint16,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterUint16").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualUint16 tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualUint16(t testing.T,
	left uint16,
	right uint16,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualUint16").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessUint16 tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessUint16(t testing.T,
	left uint16,
	right uint16,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessUint16").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualUint16 tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualUint16(t testing.T,
	left uint16,
	right uint16,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualUint16").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterUint32 tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterUint32(t testing.T,
	left uint32,
	right uint32,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterUint32").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualUint32 tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualUint32(t testing.T,
	left uint32,
	right uint32,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualUint32").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessUint32 tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessUint32(t testing.T,
	left uint32,
	right uint32,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessUint32").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualUint32 tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualUint32(t testing.T,
	left uint32,
	right uint32,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualUint32").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterUint64 tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterUint64(t testing.T,
	left uint64,
	right uint64,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterUint64").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualUint64 tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualUint64(t testing.T,
	left uint64,
	right uint64,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualUint64").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessUint64 tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessUint64(t testing.T,
	left uint64,
	right uint64,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessUint64").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualUint64 tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualUint64(t testing.T,
	left uint64,
	right uint64,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualUint64").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterUint8 tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterUint8(t testing.T,
	left uint8,
	right uint8,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterUint8").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualUint8 tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualUint8(t testing.T,
	left uint8,
	right uint8,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualUint8").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessUint8 tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessUint8(t testing.T,
	left uint8,
	right uint8,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessUint8").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualUint8 tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualUint8(t testing.T,
	left uint8,
	right uint8,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualUint8").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterUintptr tests if `left` is strictly greater than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterUintptr(t testing.T,
	left uintptr,
	right uintptr,
	msg ...interface{}) bool {
	t.Helper()

	if left > right {
		return true
	}
	return failTest(t, fail.Failure("GreaterUintptr").
		Reason("expected left > right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// GreaterOrEqualUintptr tests if `left` is greater than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func GreaterOrEqualUintptr(t testing.T,
	left uintptr,
	right uintptr,
	msg ...interface{}) bool {
	t.Helper()

	if left >= right {
		return true
	}
	return failTest(t, fail.Failure("GreaterOrEqualUintptr").
		Reason("expected left >= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessUintptr tests if `left` is strictly less than `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessUintptr(t testing.T,
	left uintptr,
	right uintptr,
	msg ...interface{}) bool {
	t.Helper()

	if left < right {
		return true
	}
	return failTest(t, fail.Failure("LessUintptr").
		Reason("expected left < right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// LessOrEqualUintptr tests if `left` is less than or equal to `right`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func LessOrEqualUintptr(t testing.T,
	left uintptr,
	right uintptr,
	msg ...interface{}) bool {
	t.Helper()

	if left <= right {
		return true
	}
	return failTest(t, fail.Failure("LessOrEqualUintptr").
		Reason("expected left <= right").
		LeftValue(left).RightValue(right).
		ExtraMsg(msg...))
}

// EqualTypes tests if `left` and `right` have the same type.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// Example:
//
//	var a int32 = 42;
//	EqualType(t, int32(0), a)
func EqualTypes(t testing.T,
	left interface{},
	right interface{},
	msg ...interface{}) bool {
	t.Helper()

	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)
	if leftType == rightType {
		return true
	}
	return failTest(t, fail.Failure("EqualTypes").
		Reason("types should be equal").
		LeftType(left).RightType(right).
		ExtraMsg(msg...))
}

// NotEqualTypes tests if `left` and `right` have different types.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// Example:
//
//	var a int32 = 42;
//	NotEqualType(t, int64(0), a)
func NotEqualTypes(t testing.T,
	left interface{},
	right interface{},
	msg ...interface{}) bool {
	t.Helper()

	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)
	if leftType != rightType {
		return true
	}
	return failTest(t, fail.Failure("NotEqualTypes").
		Reason("types should not be equal").
		LeftType(left).RightType(right).
		ExtraMsg(msg...))
}

// Contains tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The type of elements within `container` must match the type of
// `expectedElement`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func Contains(t testing.T,
	expectedElement interface{},
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("Contains")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContains tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The type of elements within `container` must match the type of
// `expectedElement`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContains(t testing.T,
	expectedElement interface{},
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContains")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsBool tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `bool`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsBool(t testing.T,
	expectedElement bool,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsBool")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsBool tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `bool`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsBool(t testing.T,
	expectedElement bool,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsBool")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsByte tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `byte`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsByte(t testing.T,
	expectedElement byte,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsByte")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsByte tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `byte`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsByte(t testing.T,
	expectedElement byte,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsByte")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsComplex128 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `complex128`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsComplex128(t testing.T,
	expectedElement complex128,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsComplex128")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsComplex128 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `complex128`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsComplex128(t testing.T,
	expectedElement complex128,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsComplex128")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsComplex64 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `complex64`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsComplex64(t testing.T,
	expectedElement complex64,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsComplex64")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsComplex64 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `complex64`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsComplex64(t testing.T,
	expectedElement complex64,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsComplex64")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsError tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `error`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsError(t testing.T,
	expectedElement error,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsError")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsError tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `error`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsError(t testing.T,
	expectedElement error,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsError")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsFloat32 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `float32`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsFloat32(t testing.T,
	expectedElement float32,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsFloat32")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsFloat32 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `float32`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsFloat32(t testing.T,
	expectedElement float32,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsFloat32")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsFloat64 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `float64`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsFloat64(t testing.T,
	expectedElement float64,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsFloat64")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsFloat64 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `float64`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsFloat64(t testing.T,
	expectedElement float64,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsFloat64")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsInt tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `int`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsInt(t testing.T,
	expectedElement int,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsInt")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsInt tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `int`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsInt(t testing.T,
	expectedElement int,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsInt")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsInt16 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `int16`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsInt16(t testing.T,
	expectedElement int16,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsInt16")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsInt16 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `int16`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsInt16(t testing.T,
	expectedElement int16,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsInt16")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsInt32 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `int32`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsInt32(t testing.T,
	expectedElement int32,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsInt32")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsInt32 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `int32`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsInt32(t testing.T,
	expectedElement int32,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsInt32")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsInt64 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `int64`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsInt64(t testing.T,
	expectedElement int64,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsInt64")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsInt64 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `int64`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsInt64(t testing.T,
	expectedElement int64,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsInt64")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsInt8 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `int8`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsInt8(t testing.T,
	expectedElement int8,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsInt8")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsInt8 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `int8`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsInt8(t testing.T,
	expectedElement int8,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsInt8")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsRune tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `rune`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsRune(t testing.T,
	expectedElement rune,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsRune")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsRune tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `rune`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsRune(t testing.T,
	expectedElement rune,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsRune")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsString tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `string`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsString(t testing.T,
	expectedElement string,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsString")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsString tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `string`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsString(t testing.T,
	expectedElement string,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsString")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsUint tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `uint`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsUint(t testing.T,
	expectedElement uint,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsUint")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsUint tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `uint`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsUint(t testing.T,
	expectedElement uint,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsUint")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsUint16 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `uint16`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsUint16(t testing.T,
	expectedElement uint16,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsUint16")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsUint16 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `uint16`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsUint16(t testing.T,
	expectedElement uint16,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsUint16")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsUint32 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `uint32`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsUint32(t testing.T,
	expectedElement uint32,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsUint32")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsUint32 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `uint32`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsUint32(t testing.T,
	expectedElement uint32,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsUint32")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsUint64 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `uint64`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsUint64(t testing.T,
	expectedElement uint64,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsUint64")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsUint64 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `uint64`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsUint64(t testing.T,
	expectedElement uint64,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsUint64")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsUint8 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `uint8`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsUint8(t testing.T,
	expectedElement uint8,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsUint8")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsUint8 tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `uint8`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsUint8(t testing.T,
	expectedElement uint8,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsUint8")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// ContainsUintptr tests if `expectedElement` is contained within
// `container`. `container` can be a string, map (only keys are matched),
// array or slice.
// The `container` must have elements of type `uintptr`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `Contains` for more static typing.
// See also `NotContains`.
func ContainsUintptr(t testing.T,
	expectedElement uintptr,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("ContainsUintptr")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is not contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotContainsUintptr tests if `expectedElement` is contained within
// `container`. `container` can be a string, map, array or slice.
// The `container` must have elements of type `uintptr`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// You can use the typed versions of `NotContains` for more static typing.
// See also `Contains`.
func NotContainsUintptr(t testing.T,
	expectedElement uintptr,
	container interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.Contains(expectedElement, container)
	failure := fail.Failure("NotContainsUintptr")
	if err != nil {
		if ce, ok2 := err.(cmp.ContainsError); ok2 {
			failure = failure.Hint(ce.Hint)
		} else {
			failure = failure.Error(err)
		}
	} else if !ok {
		return true
	}
	failure = failure.Reason("`%v` (%T) is contained in `%v` (%T)",
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// Nil tests if the passed value is Nil. See also Empty.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func Nil(t testing.T, tval interface{}, msg ...interface{}) bool {
	t.Helper()

	isnil, err := cmp.Nil(tval)
	failure := fail.Failure("Nil")
	if err != nil {
		failure = failure.Error(err)
	} else if isnil {
		return true
	}
	failure = failure.Reason("`%v` (%T) should be nil", tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotNil tests if the passed value is not Nil. See also NotEmpty.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func NotNil(t testing.T, tval interface{}, msg ...interface{}) bool {
	t.Helper()

	isnil, err := cmp.Nil(tval)
	failure := fail.Failure("NotNil")
	if err != nil {
		failure = failure.Error(err)
	} else if !isnil {
		return true
	}
	failure = failure.Reason("`%v (%T)` should not be nil", tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NoError tests if err is nil.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// See also Error, ErrorIs, ErrorAs, ErrorContains.
func NoError(t testing.T, err error, msg ...interface{}) bool {
	t.Helper()

	if err == nil {
		return true
	}
	return failTest(t, fail.Failure("NoError").
		Reason("expected no error, got: %v", err).
		ExtraMsg(msg...))
}

// Error tests if err is not nil.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// See also NoError, ErrorIs, ErrorAs, ErrorContains.
func Error(t testing.T, err error, msg ...interface{}) bool {
	t.Helper()

	if err != nil {
		return true
	}
	return failTest(t, fail.Failure("Error").
		Reason("expected an error but got nil").
		ExtraMsg(msg...))
}

// ErrorIs tests if errors.Is(err, target) is true.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// See also NotErrorIs, NoError, Error.
func ErrorIs(t testing.T, err error, target error, msg ...interface{}) bool {
	t.Helper()

	if errors.Is(err, target) {
		return true
	}
	return failTest(t, fail.Failure("ErrorIs").
		Reason("errors.Is(err, target) returned false").
		LeftValue(err).RightValue(target).
		ExtraMsg(msg...))
}

// NotErrorIs tests if errors.Is(err, target) is false.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// See also ErrorIs, NoError, Error.
func NotErrorIs(t testing.T, err error, target error, msg ...interface{}) bool {
	t.Helper()

	if !errors.Is(err, target) {
		return true
	}
	return failTest(t, fail.Failure("NotErrorIs").
		Reason("errors.Is(err, target) returned true").
		LeftValue(err).RightValue(target).
		ExtraMsg(msg...))
}

// ErrorAs tests if errors.As(err, target) is true.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// target must be a non-nil pointer to either a type that implements error,
// or to any interface type. See also NotErrorAs.
func ErrorAs(t testing.T, err error, target interface{}, msg ...interface{}) bool {
	t.Helper()

	if errors.As(err, target) {
		return true
	}
	return failTest(t, fail.Failure("ErrorAs").
		Reason("errors.As(err, target) returned false").
		LeftValue(err).RightValue(target).
		ExtraMsg(msg...))
}

// ErrorContains tests if err is not nil and err.Error() contains substr.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// See also Error, ErrorIs.
func ErrorContains(t testing.T, err error, substr string, msg ...interface{}) bool {
	t.Helper()

	if err != nil && strings.Contains(err.Error(), substr) {
		return true
	}
	failure := fail.Failure("ErrorContains")
	if err == nil {
		failure = failure.Reason("expected an error containing %q, got nil", substr)
	} else {
		failure = failure.Reason("expected error to contain %q", substr).
			LeftValue(err.Error()).RightValue(substr)
	}
	return failTest(t, failure.ExtraMsg(msg...))
}

// Empty tests if the passed value is Empty. A nil container or a container
// with zero elements are both empty. Uses Go len().
// Container can be any of Array, Chan, Map, Slice, or String.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// See also `Nil`.
func Empty(t testing.T, container interface{}, msg ...interface{}) bool {
	t.Helper()

	empty, err := cmp.Empty(container)
	failure := fail.Failure("Empty")
	if err != nil {
		failure = failure.Error(err)
	} else if empty {
		return true
	}
	failure = failure.Reason("`%v` (%T) should be empty", container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEmpty tests if the passed value is not Empty. Only a non-nil container
// with at least one element will pass the test. Uses Go len().
// Container can be any of Array, Chan, Map, Slice, or String.
// See also `NotNil`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func NotEmpty(t testing.T, container interface{}, msg ...interface{}) bool {
	t.Helper()

	empty, err := cmp.Empty(container)
	failure := fail.Failure("NotEmpty")
	if err != nil {
		failure = failure.Error(err)
	} else if !empty {
		return true
	}
	failure = failure.Reason("`%v` (%T) should not be empty", container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// Len tests if the length of container equals expected. Uses Go len().
// Container can be any of Array, Chan, Map, Slice, or String.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func Len(t testing.T, container interface{}, expected int, msg ...interface{}) bool {
	t.Helper()

	v := reflect.ValueOf(container)
	actual := -1
	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		actual = v.Len()
	}
	if actual == expected {
		return true
	}
	failure := fail.Failure("Len")
	if actual == -1 {
		failure = failure.Reason("`%v` (%T) does not support Len", container, container)
	} else {
		failure = failure.Reason("expected length %d, got %d", expected, actual).
			LeftValue(expected).RightValue(actual)
	}
	return failTest(t, failure.ExtraMsg(msg...))
}

// ElementsMatch tests if `expected` and `actual` contain the same elements
// regardless of order. Both must be slices or arrays.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// See also Contains, DeepEqual.
func ElementsMatch(t testing.T,
	expected interface{},
	actual interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.ElementsMatch(expected, actual)
	failure := fail.Failure("ElementsMatch")
	if err != nil {
		failure = failure.Error(err)
	} else if ok {
		return true
	} else {
		failure = failure.Reason("slices contain different elements").
			LeftValue(expected).RightValue(actual)
	}
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotElementsMatch tests if `expected` and `actual` do not contain the same
// elements (regardless of order). Both must be slices or arrays.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// See also NotContains, NotDeepEqual.
func NotElementsMatch(t testing.T,
	expected interface{},
	actual interface{},
	msg ...interface{}) bool {
	t.Helper()

	ok, err := cmp.ElementsMatch(expected, actual)
	failure := fail.Failure("NotElementsMatch")
	if err != nil {
		failure = failure.Error(err)
	} else if !ok {
		return true
	} else {
		failure = failure.Reason("slices contain the same elements").
			LeftValue(expected).RightValue(actual)
	}
	return failTest(t, failure.ExtraMsg(msg...))
}

// DeepEqual tests if `left` == `right` using `reflect.DeepEqual`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This is a deep, recursive equality test. See also `Equal`.
func DeepEqual(t testing.T,
	left interface{},
	right interface{},
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("DeepEqual{")

	if reflect.DeepEqual(left, right) {
		return true
	}

	failure = failure.Reason("values should be deeply equal").
		LeftValue(left).
		RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotDeepEqual tests if `left` != `right` using `reflect.DeepEqual`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This is a deep, recursive non-equality test. See also `NotEqual`.
func NotDeepEqual(t testing.T,
	left interface{},
	right interface{},
	msg ...interface{}) bool {
	t.Helper()

	failure := fail.Failure("NotDeepEqual")

	if !reflect.DeepEqual(left, right) {
		return true
	}

	failure = failure.Reason("values should not be deeply equal").
		LeftValue(left).
		RightValue(right)
	return failTest(t, failure.ExtraMsg(msg...))
}

// Same tests if expected and actual point to the same object (pointer identity).
// msg is an optional list of arguments following the `fmt.Printf()` format.
func Same(t testing.T, expected interface{}, actual interface{}, msg ...interface{}) bool {
	t.Helper()

	ev := reflect.ValueOf(expected)
	av := reflect.ValueOf(actual)
	if ev.Kind() != reflect.Ptr || av.Kind() != reflect.Ptr {
		return failTest(t, fail.Failure("Same").
			Reason("both arguments must be pointers").
			ExtraMsg(msg...))
	}
	if ev.Pointer() == av.Pointer() {
		return true
	}
	return failTest(t, fail.Failure("Same").
		Reason("expected the same pointer").
		LeftValue(expected).RightValue(actual).
		ExtraMsg(msg...))
}

// NotSame tests if expected and actual do not point to the same object.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func NotSame(t testing.T, expected interface{}, actual interface{}, msg ...interface{}) bool {
	t.Helper()

	ev := reflect.ValueOf(expected)
	av := reflect.ValueOf(actual)
	if ev.Kind() != reflect.Ptr || av.Kind() != reflect.Ptr {
		return failTest(t, fail.Failure("NotSame").
			Reason("both arguments must be pointers").
			ExtraMsg(msg...))
	}
	if ev.Pointer() != av.Pointer() {
		return true
	}
	return failTest(t, fail.Failure("NotSame").
		Reason("expected different pointers, got the same").
		LeftValue(expected).RightValue(actual).
		ExtraMsg(msg...))
}

// Zero tests if val is the zero value for its type.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func Zero(t testing.T, val interface{}, msg ...interface{}) bool {
	t.Helper()

	if val == nil || reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface()) {
		return true
	}
	return failTest(t, fail.Failure("Zero").
		Reason("expected zero value, got `%v` (%T)", val, val).
		ExtraMsg(msg...))
}

// NotZero tests if val is not the zero value for its type.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func NotZero(t testing.T, val interface{}, msg ...interface{}) bool {
	t.Helper()

	if val != nil && !reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface()) {
		return true
	}
	return failTest(t, fail.Failure("NotZero").
		Reason("expected non-zero value, got `%v` (%T)", val, val).
		ExtraMsg(msg...))
}

// InDelta tests if the absolute difference between expected and actual is
// less than or equal to delta. Useful for floating-point comparisons.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func InDelta(t testing.T,
	expected float64,
	actual float64,
	delta float64,
	msg ...interface{}) bool {
	t.Helper()

	diff := math.Abs(actual - expected)
	if diff <= delta {
		return true
	}
	return failTest(t, fail.Failure("InDelta").
		Reason("|actual - expected| = %v, which exceeds delta %v", diff, delta).
		LeftValue(expected).RightValue(actual).
		ExtraMsg(msg...))
}

// Regexp tests if str matches the given pattern. pattern can be a string or
// a *regexp.Regexp.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// See also NotRegexp.
func Regexp(t testing.T, pattern interface{}, str string, msg ...interface{}) bool {
	t.Helper()

	rx, err := toRegexp(pattern)
	if err != nil {
		return failTest(t, fail.Failure("Regexp").Error(err).ExtraMsg(msg...))
	}
	if rx.MatchString(str) {
		return true
	}
	return failTest(t, fail.Failure("Regexp").
		Reason("expected %q to match pattern %q", str, rx.String()).
		LeftValue(str).RightValue(rx.String()).
		ExtraMsg(msg...))
}

// NotRegexp tests if str does not match the given pattern. pattern can be a
// string or a *regexp.Regexp.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// See also Regexp.
func NotRegexp(t testing.T, pattern interface{}, str string, msg ...interface{}) bool {
	t.Helper()

	rx, err := toRegexp(pattern)
	if err != nil {
		return failTest(t, fail.Failure("NotRegexp").Error(err).ExtraMsg(msg...))
	}
	if !rx.MatchString(str) {
		return true
	}
	return failTest(t, fail.Failure("NotRegexp").
		Reason("expected %q not to match pattern %q", str, rx.String()).
		LeftValue(str).RightValue(rx.String()).
		ExtraMsg(msg...))
}

func toRegexp(pattern interface{}) (*regexp.Regexp, error) {
	switch v := pattern.(type) {
	case string:
		return regexp.Compile(v)
	case *regexp.Regexp:
		return v, nil
	default:
		return nil, fmt.Errorf("pattern must be string or *regexp.Regexp, got %T", pattern)
	}
}

// Failed logs a message and fails the test.
// fmtstr and args follows the `fmt.Printf()` format.
func Failed(t testing.T, fmtstr string, args ...interface{}) {
	t.Helper()

	failure := fail.Failure("Custom")
	failure = failure.Reason(fmtstr, args...)
	failTest(t, failure)
}

// Panics asserts that the code inside the specified function f panics.
//
//	assert.Panics(t, func(){ GoCrazy() })
func Panics(t testing.T, f func(), msg ...interface{}) bool {
	t.Helper()

	if funcDidPanic, _ := cmp.Panics(f); !funcDidPanic {
		failure := fail.Failure("Panic")
		failure = failure.Reason("Expected function %#v to panic", f)
		return failTest(t, failure.ExtraMsg(msg...))
	}
	return true
}

// PanicsWith asserts that the code inside the specified function f
// panics, and that the recovered panic value equals the expected panic value
//
//	assert.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
func PanicsWith(t testing.T,
	expected interface{},
	f func(),
	msg ...interface{}) bool {
	t.Helper()

	funcDidPanic, panicValue := cmp.Panics(f)
	failure := fail.Failure("PanicWithValue")

	if !funcDidPanic {
		failure = failure.Reason("Expected %#v to panic", f)
	} else if !reflect.DeepEqual(expected, panicValue) {
		failure = failure.Reason("func %#v should panic with value:\t%v\n\r"+
			"\tPanic value:\t%v", f, expected, panicValue)
	} else {
		return true
	}
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotPanics asserts that the code inside the specified function f does not panic.
//
//	assert.NotPanics(t, func(){ Sane() })
func NotPanics(t testing.T, f func(), msg ...interface{}) bool {
	t.Helper()

	if funcDidPanic, panicValue := cmp.Panics(f); funcDidPanic {
		failure := fail.Failure("NotPanics")
		failure = failure.Reason("Expected function not to panic\n\tPanic value:\t%v", panicValue)
		return failTest(t, failure.ExtraMsg(msg...))
	}
	return true
}
