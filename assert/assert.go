// @generated

package assert

import (
	"reflect"

	"github.com/rubrikinc/testwell/internal/cmp"
	"github.com/rubrikinc/testwell/internal/fail"
	"github.com/rubrikinc/testwell/testing"
)

// True tests if val is True.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func True(t testing.T, tval bool, msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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

// Equal tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two interface{}.
// This is a strict equality test. It will fail if `expected` and `tval`
// have different types. For example, int32(42) is not equal to int64(42).
// You can use the typed versions of `Equal` for more static typing.
// Note that Slice, map, and function values are not comparable.
// See also `DeepEqual`.
func Equal(t testing.T,
	expected interface{},
	tval interface{},
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("Equal")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqual tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two interface{}
// This is a strict non-equality test. It will pass only when
// `expected` and `tval` have identical type but non-equal values.
// For example, given int32(42) and int64(42), the test will fail because
// the types are not comparables. Given int32(42) and int32(42), the test
// will fail because the two values are equal.
// You can use the typed versions of `NotEqual` for more static typing.
// Note that Slice, map, and function values are not comparable.
// See also `NotDeepEqual`.
func NotEqual(t testing.T,
	expected interface{},
	tval interface{},
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqual")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualBool tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two bool.
func EqualBool(t testing.T,
	expected bool,
	tval bool,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualBool")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualBool tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two bool
func NotEqualBool(t testing.T,
	expected bool,
	tval bool,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualBool")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualByte tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two byte.
func EqualByte(t testing.T,
	expected byte,
	tval byte,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualByte")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualByte tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two byte
func NotEqualByte(t testing.T,
	expected byte,
	tval byte,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualByte")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualComplex128 tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two complex128.
func EqualComplex128(t testing.T,
	expected complex128,
	tval complex128,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualComplex128")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualComplex128 tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two complex128
func NotEqualComplex128(t testing.T,
	expected complex128,
	tval complex128,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualComplex128")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualComplex64 tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two complex64.
func EqualComplex64(t testing.T,
	expected complex64,
	tval complex64,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualComplex64")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualComplex64 tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two complex64
func NotEqualComplex64(t testing.T,
	expected complex64,
	tval complex64,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualComplex64")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualError tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two error.
func EqualError(t testing.T,
	expected error,
	tval error,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualError")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualError tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two error
func NotEqualError(t testing.T,
	expected error,
	tval error,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualError")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualFloat32 tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two float32.
func EqualFloat32(t testing.T,
	expected float32,
	tval float32,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualFloat32")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualFloat32 tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two float32
func NotEqualFloat32(t testing.T,
	expected float32,
	tval float32,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualFloat32")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualFloat64 tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two float64.
func EqualFloat64(t testing.T,
	expected float64,
	tval float64,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualFloat64")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualFloat64 tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two float64
func NotEqualFloat64(t testing.T,
	expected float64,
	tval float64,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualFloat64")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualInt tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int.
func EqualInt(t testing.T,
	expected int,
	tval int,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualInt")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualInt tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int
func NotEqualInt(t testing.T,
	expected int,
	tval int,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualInt")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualInt16 tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int16.
func EqualInt16(t testing.T,
	expected int16,
	tval int16,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualInt16")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualInt16 tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int16
func NotEqualInt16(t testing.T,
	expected int16,
	tval int16,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualInt16")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualInt32 tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int32.
func EqualInt32(t testing.T,
	expected int32,
	tval int32,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualInt32")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualInt32 tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int32
func NotEqualInt32(t testing.T,
	expected int32,
	tval int32,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualInt32")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualInt64 tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int64.
func EqualInt64(t testing.T,
	expected int64,
	tval int64,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualInt64")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualInt64 tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int64
func NotEqualInt64(t testing.T,
	expected int64,
	tval int64,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualInt64")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualInt8 tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int8.
func EqualInt8(t testing.T,
	expected int8,
	tval int8,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualInt8")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualInt8 tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two int8
func NotEqualInt8(t testing.T,
	expected int8,
	tval int8,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualInt8")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualRune tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two rune.
func EqualRune(t testing.T,
	expected rune,
	tval rune,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualRune")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualRune tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two rune
func NotEqualRune(t testing.T,
	expected rune,
	tval rune,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualRune")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualString tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two string.
func EqualString(t testing.T,
	expected string,
	tval string,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualString")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualString tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two string
func NotEqualString(t testing.T,
	expected string,
	tval string,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualString")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualUint tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint.
func EqualUint(t testing.T,
	expected uint,
	tval uint,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualUint")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualUint tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint
func NotEqualUint(t testing.T,
	expected uint,
	tval uint,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualUint")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualUint16 tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint16.
func EqualUint16(t testing.T,
	expected uint16,
	tval uint16,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualUint16")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualUint16 tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint16
func NotEqualUint16(t testing.T,
	expected uint16,
	tval uint16,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualUint16")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualUint32 tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint32.
func EqualUint32(t testing.T,
	expected uint32,
	tval uint32,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualUint32")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualUint32 tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint32
func NotEqualUint32(t testing.T,
	expected uint32,
	tval uint32,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualUint32")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualUint64 tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint64.
func EqualUint64(t testing.T,
	expected uint64,
	tval uint64,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualUint64")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualUint64 tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint64
func NotEqualUint64(t testing.T,
	expected uint64,
	tval uint64,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualUint64")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualUint8 tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint8.
func EqualUint8(t testing.T,
	expected uint8,
	tval uint8,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualUint8")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualUint8 tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uint8
func NotEqualUint8(t testing.T,
	expected uint8,
	tval uint8,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualUint8")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualUintptr tests if `tval` is equal to `expected` using the `==`
// operator. The expected value comes first, followed by the value to test
// against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uintptr.
func EqualUintptr(t testing.T,
	expected uintptr,
	tval uintptr,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("EqualUintptr")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see Nil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at DeepEqual")
			}
		}
	} else if ok {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEqualUintptr tests if `tval` is not equal to `expected` using the
// `!=` operator. The expected value comes first, followed by the value to
// test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This version takes takes two uintptr
func NotEqualUintptr(t testing.T,
	expected uintptr,
	tval uintptr,
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotEqualUintptr")

	if ok, err := cmp.Equal(expected, tval); err != nil {
		failure = failure.Error(err)
		if cerr, ok2 := err.(cmp.NotComparableError); ok2 {
			failure = failure.Error(cerr)
			if reflect.TypeOf(expected) == nil {
				failure = failure.Hint("see NotNil for <nil> checks")
			} else {
				failure = failure.Hint("take a look at NotDeepEqual")
			}
		}
	} else if !ok {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// EqualTypes tests if two values have the same types.
// The expected type comes first, followed by the value to test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// Example:
//  var a int32 = 42;
//  EqualType(t, int32(0), a)
func EqualTypes(t testing.T,
	expected interface{},
	tval interface{},
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	eType := reflect.TypeOf(expected)
	tvalType := reflect.TypeOf(tval)
	if tvalType == eType {
		return true
	}
	return failTest(t, fail.Failure("EqualTypes").
		Reason(`expected type "%v", got type "%v" instead`, eType, tvalType).
		ExtraMsg(msg...))
}

// NotEqualTypes tests if two values have distinct types.
// The expected type comes first, followed by the value to test against.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// Example:
//  var a int32 = 42;
//  EqualType(t, int32(0), a)
func NotEqualTypes(t testing.T,
	expected interface{},
	tval interface{},
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	eType := reflect.TypeOf(expected)
	tvalType := reflect.TypeOf(tval)
	if tvalType != eType {
		return true
	}
	return failTest(t, fail.Failure("NotEqualTypes").
		Reason(`expected anything but type "%v", got type "%v"`,
			eType, tvalType).
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is not contained in "%v" (%T)`,
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
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
	failure = failure.Reason(`"%v" (%T) is contained in "%v" (%T)`,
		expectedElement, expectedElement, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// Nil tests if the passed value is Nil. See also Empty.
// Makes sense solely on chan, func, interface, map, pointer, slice value and
// the zero value type (ie: the nil keyword).
// msg is an optional list of arguments following the `fmt.Printf()` format.
func Nil(t testing.T, tval interface{}, msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	isnil, err := cmp.Nil(tval)
	failure := fail.Failure("Nil")
	if err != nil {
		failure = failure.Error(err)
	} else if isnil {
		return true
	}
	failure = failure.Reason(`"%#v" is not nil`, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotNil tests if the passed value is not Nil. See also NotEmpty.
// Makes sense solely on chan, func, interface, map, pointer, slice value and
// the zero value type (ie: the nil keyword).
// msg is an optional list of arguments following the `fmt.Printf()` format.
func NotNil(t testing.T, tval interface{}, msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	isnil, err := cmp.Nil(tval)
	failure := fail.Failure("NotNil")
	if err != nil {
		failure = failure.Error(err)
	} else if !isnil {
		return true
	}
	failure = failure.Reason(`"%#v" is nil`, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// Empty tests if the passed value is Empty. A nil container or a container
// with zero elements are both empty. Uses Go len().
// Container can be any of Array, Chan, Map, Slice, or String.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// See also `Nil`.
func Empty(t testing.T, container interface{}, msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	empty, err := cmp.Empty(container)
	failure := fail.Failure("Empty")
	if err != nil {
		failure = failure.Error(err)
	} else if empty {
		return true
	}
	failure = failure.Reason(`"%v" (%T) is not empty`, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotEmpty tests if the passed value is not Empty. Only a non-nil container
// with at least one element will pass the test. Uses Go len().
// Container can be any of Array, Chan, Map, Slice, or String.
// See also `NotNil`.
// msg is an optional list of arguments following the `fmt.Printf()` format.
func NotEmpty(t testing.T, container interface{}, msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	empty, err := cmp.Empty(container)
	failure := fail.Failure("NotEmpty")
	if err != nil {
		failure = failure.Error(err)
	} else if !empty {
		return true
	}
	failure = failure.Reason(`"%v" (%T) is empty`, container, container)
	return failTest(t, failure.ExtraMsg(msg...))
}

// DeepEqual tests if `tval` is deeply equal to `expected` using
// `reflect.DeepEqual`. The expected value comes first, followed by the value
// to compare.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This is a deep, recursive equality test. See also `Equal`.
func DeepEqual(t testing.T,
	expected interface{},
	tval interface{},
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("DeepEqual{")

	if reflect.DeepEqual(expected, tval) {
		return true
	}

	failure = failure.Reason(
		`expected value "%+v" (%T), got "%+v" (%T) instead`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// NotDeepEqual tests if `tval` is deeply non-equal to `expected` using
// `reflect.DeepEqual`. The expected value comes first, followed by the value
// to compare.
// msg is an optional list of arguments following the `fmt.Printf()` format.
// This is a deep, recursive non-equality test. See also `Equal`.
func NotDeepEqual(t testing.T,
	expected interface{},
	tval interface{},
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("NotDeepEqual")

	if !reflect.DeepEqual(expected, tval) {
		return true
	}

	failure = failure.Reason(
		`expected other value than "%+v" (%T), got "%+v" (%T)`,
		expected, expected, tval, tval)
	return failTest(t, failure.ExtraMsg(msg...))
}

// Failed logs a message and fails the test.
// fmtstr and args follows the `fmt.Printf()` format.
func Failed(t testing.T, fmtstr string, args ...interface{}) {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	failure := fail.Failure("Custom")
	failure = failure.Reason(fmtstr, args...)
	failTest(t, failure)
}

// Panics asserts that the code inside the specified function f panics.
//   assert.Panics(t, func(){ GoCrazy() })
func Panics(t testing.T, f func(), msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

	if funcDidPanic, _ := cmp.Panics(f); !funcDidPanic {
		failure := fail.Failure("Panic")
		failure = failure.Reason("Expected function %#v to panic", f)
		return failTest(t, failure.ExtraMsg(msg...))
	}
	return true
}

// PanicsWith asserts that the code inside the specified function f
// panics, and that the recovered panic value equals the expected panic value
//   assert.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
func PanicsWith(t testing.T,
	expected interface{},
	f func(),
	msg ...interface{}) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}

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
