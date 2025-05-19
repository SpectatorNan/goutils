package errors

import (
	"fmt"
	"io"
)

// GetStackTrace returns the stack trace of the error if available.
// If the error doesn't implement stackTracer, nil is returned.
func GetStackTrace(err error) StackTrace {
	type stackTracer interface {
		StackTrace() StackTrace
	}

	for err != nil {
		tracer, ok := err.(stackTracer)
		if ok {
			return tracer.StackTrace()
		}

		// Try to unwrap using the standard errors.Unwrap if available
		if unwrapper, ok := err.(interface{ Unwrap() error }); ok {
			err = unwrapper.Unwrap()
			continue
		}

		// Try the pkg/errors causer interface
		if causer, ok := err.(interface{ Cause() error }); ok {
			err = causer.Cause()
			continue
		}

		break
	}
	return nil
}

// WithStackFrom adds the stack trace from sourceErr to targetErr.
// If sourceErr doesn't have a stack trace, targetErr is returned unchanged.
func WithStackFrom(targetErr, sourceErr error) error {
	if targetErr == nil || sourceErr == nil {
		return targetErr
	}

	st := GetStackTrace(sourceErr)
	if st == nil {
		return &withExtractedStack{
			error: targetErr,
			//stacktrace: st,
			stacktrace: callers().StackTrace(),
		}
	}

	return &withExtractedStack{
		error:      targetErr,
		stacktrace: st,
		//stack: callers(),
	}
}

type withExtractedStack struct {
	error
	stacktrace StackTrace
	//stack *stack
}

func (w *withExtractedStack) StackTrace() StackTrace {
	//return w.stack.StackTrace()
	return w.stacktrace
}

func (w *withExtractedStack) Cause() error  { return w.error }
func (w *withExtractedStack) Unwrap() error { return w.error }

func (w *withExtractedStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.error)
			//w.stack.Format(s, verb)
			for _, frame := range w.stacktrace {
				fmt.Fprintf(s, "\n%+v", frame)
			}
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}
