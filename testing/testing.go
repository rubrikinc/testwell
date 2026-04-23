// Package testing defines minimal interfaces over the standard library's
// testing.T. Using interfaces rather than the concrete *testing.T allows
// callers to substitute fakes (e.g. for testing assertion libraries themselves).
package testing

// T is the interface satisfied by *testing.T from the standard library.
// All assertion functions in this module accept T so that callers can pass
// either a real *testing.T or a test double. Implementations must include
// Helper() so the testing framework omits assertion stack frames from failure
// output and points directly at the caller's line instead.
type T interface {
	Log(args ...interface{})
	Fail()
	FailNow()
	Helper()
}
