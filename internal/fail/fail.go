package fail

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

// TestFailure is a builder of test failure message with stacktrace.
type TestFailure struct {
	Name        string
	Stack       []uintptr
	ReasonStr   string
	HintStr     string
	ExtraMsgStr string
	Err         error
}

// Failure instantiate a new msg builder with a snapshot of call stack.
func Failure(assertion string) TestFailure {
	stack := make([]uintptr, 32)
	stackSize := runtime.Callers(3, stack)
	stack = stack[:stackSize]
	return TestFailure{Name: assertion, Stack: stack}
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
	r := fmt.Sprintf("%s %s failed:\nTrace (most recent last):\n  %s",
		tf.Name, failType, strings.Join(tf.formattedFrames(), "\n  "))

	if tf.Err != nil {
		r = fmt.Sprintf("%s\nError: %s", r, tf.Err.Error())
		if tf.HintStr != "" {
			r = fmt.Sprintf("%s\n Hint: %s", r, tf.HintStr)
		}
	} else {
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
