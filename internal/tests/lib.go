package tests

//go:generate ../codegen/gen.sh --tmpl=../assert.tmpl --pkg=tests

import (
	stdtesting "testing"

	"github.com/rubrikinc/testwell/internal/fail"
	"github.com/rubrikinc/testwell/testing"
)

type tWrapper struct {
	*stdtesting.T

	Failures []fail.TestFailure
}

func (t *tWrapper) LastFailure() *fail.TestFailure {
	if len(t.Failures) == 0 {
		return nil
	}
	return &t.Failures[len(t.Failures)-1]
}

func failTest(t testing.T, tf fail.TestFailure) bool {
	if h, ok := t.(testing.Helper); ok {
		h.Helper() // Go 1.9 compatibility
	}
	t.Log(tf.Format("test"))
	wt := t.(*tWrapper)
	wt.Failures = append(wt.Failures, tf)
	return false
}
