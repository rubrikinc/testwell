package assert_test

import (
	"errors"
	"testing"

	"github.com/rubrikinc/testwell/assert"
)

func TestAnError(t *testing.T) {
	if assert.AnError == nil {
		t.Fatal("AnError must not be nil")
	}
	if assert.AnError.Error() == "" {
		t.Error("AnError must have a non-empty message")
	}
	if !errors.Is(assert.AnError, assert.AnError) {
		t.Error("errors.Is(AnError, AnError) must be true")
	}
	if errors.Is(errors.New("other"), assert.AnError) {
		t.Error("AnError must be distinct from other errors")
	}
}
