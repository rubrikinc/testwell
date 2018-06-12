package testing

// T is modeled and compatible with the standard library testing.T.
// It is used to match any testing.T looking type within the library.
type T interface {
	Log(args ...interface{})
	Fail()
	FailNow()
}

// Helper matches the new Helper() method added on testing.T by Go 1.9.
type Helper interface {
	Helper()
}
