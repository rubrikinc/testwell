// Package expect provides type-safe test assertion functions that continue
// execution after a failure.
//
// On failure, each function logs the failure and marks the test as failed via
// [testing.T.Fail], but allows the test to keep running. This lets a single
// test surface multiple independent failures in one pass rather than stopping
// at the first.
//
// Use the assert package when a failure means subsequent assertions cannot
// safely execute — for example, after checking a pointer is non-nil before
// dereferencing it, or after checking an error is nil before using its result.
// Use expect when the assertions are independent of each other.
//
// # Usage
//
//	import "github.com/rubrikinc/testwell/expect"
//
//	func TestUser(t *testing.T) {
//	    // All three are checked even if earlier ones fail.
//	    expect.Equal(t, "Alice", user.Name)
//	    expect.Equal(t, 30, user.Age)
//	    expect.Nil(t, user.DeletedAt)
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
package expect
