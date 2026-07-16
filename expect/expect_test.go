package expect_test

import (
	"errors"
	"testing"

	"github.com/rubrikinc/testwell/expect"
)

func TestAnError(t *testing.T) {
	if expect.AnError == nil {
		t.Fatal("AnError must not be nil")
	}
	if expect.AnError.Error() == "" {
		t.Error("AnError must have a non-empty message")
	}
	if !errors.Is(expect.AnError, expect.AnError) {
		t.Error("errors.Is(AnError, AnError) must be true")
	}
	if errors.Is(errors.New("other"), expect.AnError) {
		t.Error("AnError must be distinct from other errors")
	}
}
