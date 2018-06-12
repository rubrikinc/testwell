package expect

//go:generate ../internal/codegen/gen.sh --pkg=expect

import (
	"github.com/rubrikinc/testwell/internal/fail"
	"github.com/rubrikinc/testwell/testing"
)

func failTest(t testing.T, tf fail.TestFailure) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}
	t.Log(tf.Format("check"))
	t.Fail()
	return false
}
