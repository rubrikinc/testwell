package expect

//go:generate ../internal/codegen/gen.sh --pkg=expect

import (
	"github.com/rubrikinc/testwell/internal/fail"
	"github.com/rubrikinc/testwell/testing"
)

func failTest(t testing.T, tf fail.TestFailure) bool {
	t.Helper()
	t.Log(tf.Format("check"))
	t.Fail()
	return false
}
