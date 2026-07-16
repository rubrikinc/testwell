package expect

//go:generate ../internal/codegen/gen.sh --pkg=expect

import (
	"errors"

	"github.com/rubrikinc/testwell/internal/fail"
	"github.com/rubrikinc/testwell/testing"
)

// AnError is a sentinel error for use in tests that need an error value but
// don't care about its contents.
var AnError = errors.New("testwell.AnError general error for testing")

func failTest(t testing.T, tf fail.TestFailure) bool {
	t.Helper()
	tf = tf.WithStack(2)
	t.Log(tf.Format("check"))
	t.Fail()
	return false
}
