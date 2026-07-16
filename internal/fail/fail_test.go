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
