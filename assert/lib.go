package assert

//go:generate ../internal/codegen/gen.sh --pkg=assert

import (
	"github.com/rubrikinc/testwell/internal/fail"
	"github.com/rubrikinc/testwell/testing"
)

func failTest(t testing.T, tf fail.TestFailure) bool {
	t.Helper()
	t.Log(tf.Format("assertion"))
	t.FailNow()
	return false
}
