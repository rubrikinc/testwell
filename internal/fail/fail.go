package fail

import (
	"fmt"
	"path"
	"runtime"
	"slices"
	"strings"
)

// valueFormat selects how a value attached to a TestFailure is rendered.
type valueFormat int8

const (
	valueUnset valueFormat = iota
	valueAsValue
	valueAsType
)

// valueSlot defers rendering of a left/right value until Format runs.
// Test assertions often deal with large structs/maps; building the
// string only at Format time keeps the success path allocation-free
// and isolates presentation from data capture.
type valueSlot struct {
	Format valueFormat
	Value  any
}

func (s valueSlot) String() string {
	switch s.Format {
	case valueAsValue:
		return fmt.Sprintf("`%+v` (%T)", s.Value, s.Value)
	case valueAsType:
		return fmt.Sprintf("(%T)", s.Value)
	}
	return ""
}

// TestFailure is a builder of test failure messages with stacktrace.
// Failure() returns a zero-cost builder; the stack is captured lazily
// by WithStack at the failure path (see failTest in each package).
type TestFailure struct {
	Name        string
	Stack       []uintptr
	Left        valueSlot
	Right       valueSlot
	ReasonStr   string
	HintStr     string
	ExtraMsgStr string
	Err         error
}

// Failure returns a new TestFailure for the given assertion name.
// It allocates only the struct itself; stack capture is deferred to
// WithStack so the success path of every assertion stays cheap.
func Failure(assertion string) TestFailure {
	return TestFailure{Name: assertion}
}

// WithStack records a stack trace into tf, skipping `skip` frames above
// WithStack's caller. Pass skip=0 to start the trace at WithStack's
// immediate caller.
func (tf TestFailure) WithStack(skip int) TestFailure {
	stack := make([]uintptr, 32)
	n := runtime.Callers(2+skip, stack)
	tf.Stack = stack[:n]
	return tf
}

// LeftValue used for the test.
func (tf TestFailure) LeftValue(left any) TestFailure {
	tf.Left = valueSlot{Format: valueAsValue, Value: left}
	return tf
}

// RightValue used for the test.
func (tf TestFailure) RightValue(right any) TestFailure {
	tf.Right = valueSlot{Format: valueAsValue, Value: right}
	return tf
}

// LeftType used for the test.
func (tf TestFailure) LeftType(left any) TestFailure {
	tf.Left = valueSlot{Format: valueAsType, Value: left}
	return tf
}

// RightType used for the test.
func (tf TestFailure) RightType(right any) TestFailure {
	tf.Right = valueSlot{Format: valueAsType, Value: right}
	return tf
}

// Reason returns a new TestFailure with a Reason msg attached to it. Any Error
// attached will override the reason.
func (tf TestFailure) Reason(msg string, args ...any) TestFailure {
	tf.ReasonStr = fmt.Sprintf(msg, args...)
	return tf
}

// Hint returns a new TestFailure with a Hint msg attached to it. Any Error
// attached will override the hint.
func (tf TestFailure) Hint(msg string, args ...any) TestFailure {
	tf.HintStr = fmt.Sprintf(msg, args...)
	return tf
}

// ExtraMsg returns a new TestFailure with an extra message attached to it.
// Extra messages are expected to come from the user.
func (tf TestFailure) ExtraMsg(args ...any) TestFailure {
	if len(args) == 1 {
		tf.ExtraMsgStr = fmt.Sprint(args[0])
	} else if len(args) > 1 {
		if fmtStr, ok := args[0].(string); ok {
			tf.ExtraMsgStr = fmt.Sprintf(fmtStr, args[1:]...)
		} else {
			tf.ExtraMsgStr = fmt.Sprint(args...)
		}
	}
	return tf
}

// Error returns a new TestFailure with an error attached to it. It overrides
// the reason msg.
func (tf TestFailure) Error(err error) TestFailure {
	tf.Err = err
	return tf
}

func (tf TestFailure) formattedFrames() []string {
	if len(tf.Stack) == 0 {
		return nil
	}
	r := make([]string, 0, len(tf.Stack))
	frames := runtime.CallersFrames(tf.Stack)
	for {
		frame, more := frames.Next()
		if strings.HasPrefix(frame.Function, "testing.") {
			break
		}
		file := path.Base(frame.File)
		r = append(r, fmt.Sprintf("%s:%d (%s)", file, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	slices.Reverse(r)
	return r
}

// valueDiff returns a unified line diff of the left/right values when a diff
// would be more readable than printing both blobs: both must be multi-line
// strings that actually differ, and the assertion must be an equality check
// (where left and right are expected to match). Contains/Regexp-style checks
// also carry two string values but a diff would misrepresent them, so they are
// excluded by name.
func (tf TestFailure) valueDiff() (string, bool) {
	if tf.Left.Format != valueAsValue || tf.Right.Format != valueAsValue {
		return "", false
	}
	switch {
	case strings.HasPrefix(tf.Name, "Equal"),
		strings.HasPrefix(tf.Name, "NotEqual"),
		strings.HasPrefix(tf.Name, "DeepEqual"),
		strings.HasPrefix(tf.Name, "NotDeepEqual"):
	default:
		return "", false
	}
	l, lok := tf.Left.Value.(string)
	r, rok := tf.Right.Value.(string)
	if !lok || !rok || l == r {
		return "", false
	}
	if !strings.Contains(l, "\n") && !strings.Contains(r, "\n") {
		return "", false
	}
	return unifiedDiff(l, r), true
}

// Format returns a properly formatted test failure message with a stacktrace.
func (tf TestFailure) Format(failType string) string {
	leftStr := tf.Left.String()
	rightStr := tf.Right.String()
	diff, hasDiff := tf.valueDiff()

	var b strings.Builder
	fmt.Fprintf(&b, "%s %s failed:\nTrace (most recent last):\n  %s",
		tf.Name, failType, strings.Join(tf.formattedFrames(), "\n  "))

	writeValues := func() {
		if hasDiff {
			// The diff only fires for equality assertions, where the
			// documented convention is left=expected (want), right=actual
			// (got) — so -want/+got is accurate here and matches Go's own
			// cmp.Diff idiom, unlike the neutral Left/Right used for the raw
			// display shared by all assertions.
			fmt.Fprintf(&b, "\nDiff (-want +got):\n%s", diff)
		} else if leftStr != "" && rightStr != "" {
			fmt.Fprintf(&b, "\n Left: %s\nRight: %s", leftStr, rightStr)
		}
	}

	if tf.Err != nil {
		writeValues()
		fmt.Fprintf(&b, "\nError: %s", tf.Err.Error())
		if tf.HintStr != "" {
			fmt.Fprintf(&b, "\n Hint: %s", tf.HintStr)
		}
	} else {
		writeValues()
		fmt.Fprintf(&b, "\nReason: %s", tf.ReasonStr)
		if tf.HintStr != "" {
			fmt.Fprintf(&b, "\n  Hint: %s", tf.HintStr)
		}
	}
	if tf.ExtraMsgStr != "" {
		fmt.Fprintf(&b, "\n%s", tf.ExtraMsgStr)
	}
	return b.String()
}
