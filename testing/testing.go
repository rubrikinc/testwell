// Package testing defines minimal interfaces over the standard library's
// testing.T. Using interfaces rather than the concrete *testing.T allows
// callers to substitute fakes (e.g. for testing assertion libraries themselves)
// and keeps the core T interface compatible with Go versions older than 1.9.
package testing

// T is the interface satisfied by *testing.T from the standard library.
// All assertion functions in this module accept T so that callers can pass
// either a real *testing.T or a test double.
type T interface {
	Log(args ...interface{})
	Fail()
	FailNow()
}

// Helper is satisfied by *testing.T starting with Go 1.9. It is kept as a
// separate interface so that T remains usable on older Go versions. Assertion
// functions call Helper() via a type assertion when available, which causes
// the testing framework to omit the assertion's own stack frames from failure
// output and point directly at the caller's line instead.
type Helper interface {
	Helper()
}
