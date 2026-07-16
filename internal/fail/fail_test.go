package fail_test

import (
	"strings"
	"testing"

	"github.com/rubrikinc/testwell/internal/fail"
)

func TestFormattedFramesTrimsTesting(t *testing.T) {
	tf := fail.Failure("X").WithStack(0)
	out := tf.Format("test")

	if !strings.Contains(out, "TestFormattedFramesTrimsTesting") {
		t.Errorf("trace should include the calling test function:\n%s", out)
	}
	if strings.Contains(out, "testing.tRunner") {
		t.Errorf("trace should not include testing.tRunner:\n%s", out)
	}
}

func TestFormatDiffForMultilineEqual(t *testing.T) {
	left := "line one\nline two\nline three"
	right := "line one\nline TWO\nline three"
	out := fail.Failure("Equal").LeftValue(left).RightValue(right).Format("test")

	if !strings.Contains(out, "Diff (-want +got):") {
		t.Errorf("expected a diff section, got:\n%s", out)
	}
	if !strings.Contains(out, "- line two") || !strings.Contains(out, "+ line TWO") {
		t.Errorf("expected changed lines in diff, got:\n%s", out)
	}
	// When a diff is shown, the raw Left:/Right: blobs are omitted.
	if strings.Contains(out, "\n Left: ") {
		t.Errorf("raw Left/Right should be replaced by the diff, got:\n%s", out)
	}
}

func TestFormatNoDiffForSingleLine(t *testing.T) {
	out := fail.Failure("Equal").LeftValue("foo").RightValue("bar").Format("test")

	if strings.Contains(out, "Diff (-want +got):") {
		t.Errorf("single-line values should not produce a diff, got:\n%s", out)
	}
	if !strings.Contains(out, "Left:") || !strings.Contains(out, "Right:") {
		t.Errorf("expected raw Left/Right for single-line values, got:\n%s", out)
	}
}

func TestFormatNoDiffForNonEqualityAssertion(t *testing.T) {
	// ErrorContains carries two string values but a diff would misrepresent a
	// containment check, so it must not be diffed even when multi-line.
	errText := "boom\nsecond line"
	out := fail.Failure("ErrorContains").LeftValue(errText).RightValue("needle").Format("test")

	if strings.Contains(out, "Diff (-want +got):") {
		t.Errorf("non-equality assertion should not be diffed, got:\n%s", out)
	}
}
