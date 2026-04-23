package fail

import (
	"fmt"
	"path"
	"runtime"
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
	Value  interface{}
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
func (tf TestFailure) LeftValue(left interface{}) TestFailure {
	tf.Left = valueSlot{Format: valueAsValue, Value: left}
	return tf
}

// RightValue used for the test.
func (tf TestFailure) RightValue(right interface{}) TestFailure {
	tf.Right = valueSlot{Format: valueAsValue, Value: right}
	return tf
}

// LeftType used for the test.
func (tf TestFailure) LeftType(left interface{}) TestFailure {
	tf.Left = valueSlot{Format: valueAsType, Value: left}
	return tf
}

// RightType used for the test.
func (tf TestFailure) RightType(right interface{}) TestFailure {
	tf.Right = valueSlot{Format: valueAsType, Value: right}
	return tf
}

// Reason returns a new TestFailure with a Reason msg attached to it. Any Error
// attached will override the reason.
func (tf TestFailure) Reason(msg string, args ...interface{}) TestFailure {
	tf.ReasonStr = fmt.Sprintf(msg, args...)
	return tf
}

// Hint returns a new TestFailure with a Hint msg attached to it. Any Error
// attached will override the hint.
func (tf TestFailure) Hint(msg string, args ...interface{}) TestFailure {
	tf.HintStr = fmt.Sprintf(msg, args...)
	return tf
}

// ExtraMsg returns a new TestFailure with an extra message attached to it.
// Extra messages are expected to come from the user.
func (tf TestFailure) ExtraMsg(args ...interface{}) TestFailure {
	if len(args) == 1 {
		tf.ExtraMsgStr = args[0].(string)
	} else if len(args) > 1 {
		tf.ExtraMsgStr = fmt.Sprintf(args[0].(string), args[1:]...)
	}
	return tf
}

// Error returns a new TestFailure with an error attached to it. It overrides
// the reason msg.
func (tf TestFailure) Error(err error) TestFailure {
	tf.Err = err
	return tf
}

// Errorf returns a new TestFailure with a newly formatted error attached to
// it. It overrides the reason msg.
func (tf TestFailure) Errorf(msg string, args ...interface{}) TestFailure {
	tf.Err = fmt.Errorf(msg, args...)
	return tf
}

func reverseSliceOfStrings(s []string) {
	first, last := 0, len(s)-1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func (tf TestFailure) formattedFrames() []string {
	if len(tf.Stack) == 0 {
		return nil
	}
	r := make([]string, 0, len(tf.Stack))
	frames := runtime.CallersFrames(tf.Stack)
	for {
		frame, more := frames.Next()
		if strings.Contains(frame.File, "src/testing/") {
			break
		}
		file := path.Base(frame.File)
		r = append(r, fmt.Sprintf("%s:%d (%s)", file, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	reverseSliceOfStrings(r)
	return r
}

// Format returns a properly formatted test failure message with a stacktrace.
func (tf TestFailure) Format(failType string) string {
	leftStr := tf.Left.String()
	rightStr := tf.Right.String()

	r := fmt.Sprintf("%s %s failed:\nTrace (most recent last):\n  %s",
		tf.Name, failType, strings.Join(tf.formattedFrames(), "\n  "))

	if tf.Err != nil {
		if leftStr != "" && rightStr != "" {
			r = fmt.Sprintf("%s\n Left: %s\nRight: %s", r, leftStr, rightStr)
		}
		r = fmt.Sprintf("%s\nError: %s", r, tf.Err.Error())
		if tf.HintStr != "" {
			r = fmt.Sprintf("%s\n Hint: %s", r, tf.HintStr)
		}
	} else {
		if leftStr != "" && rightStr != "" {
			r = fmt.Sprintf("%s\n  Left: %s\n Right: %s", r, leftStr, rightStr)
		}
		r = fmt.Sprintf("%s\nReason: %s", r, tf.ReasonStr)
		if tf.HintStr != "" {
			r = fmt.Sprintf("%s\n  Hint: %s", r, tf.HintStr)
		}
	}
	if tf.ExtraMsgStr != "" {
		r = fmt.Sprintf("%s\n%s", r, tf.ExtraMsgStr)
	}
	return r
}
