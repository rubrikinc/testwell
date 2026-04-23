// Package assert provides type-safe test assertion functions.
//
// On failure, each function logs the failure and immediately stops the test
// via [testing.T.FailNow]. Use the [expect] package instead if you want
// execution to continue after a failure.
//
// # Usage
//
//	import "github.com/rubrikinc/testwell/assert"
//
//	func TestExample(t *testing.T) {
//	    assert.True(t, ready, "service must be ready")
//	    assert.Equal(t, 42, answer)
//	    assert.Nil(t, err)
//	    assert.Contains(t, "needle", haystack)
//	}
//
// # Type safety
//
// Equality and contains functions are type-safe: comparing values of different
// types (e.g. int32 and int64) is a failure, not a silent coercion. Each
// function is available in a generic form (e.g. [Equal]) that accepts
// interface{} and in type-specific forms (e.g. [EqualInt32]) for additional
// compile-time checking.
//
// # Parameter order
//
// All functions follow a consistent parameter order: the testing.T first,
// then the expected or reference value, then the actual value, then an
// optional printf-style message.
//
// [Contains] and [NotContains] take the element before the container, which
// is the reverse of Testify.
package assert
