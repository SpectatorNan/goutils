package errors

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestGetStackTrace(t *testing.T) {
	// Test with nil error
	st := GetStackTrace(nil)
	if st != nil {
		t.Errorf("Expected nil stack trace for nil error, got %v", st)
	}

	// Test with normal error (no stack trace)
	st = GetStackTrace(fmt.Errorf("simple error"))
	if st != nil {
		t.Errorf("Expected nil stack trace for simple error, got %v", st)
	}

	// Test with error that has stack trace
	err := WithStack(fmt.Errorf("error with stack"))
	st = GetStackTrace(err)
	if st == nil {
		t.Fatal("Expected stack trace for WithStack error, got nil")
	}
	if len(st) == 0 {
		t.Error("Got empty stack trace for WithStack error")
	}

	// Test with nested errors
	nested := Wrap(WithStack(io.EOF), "wrapped")
	st = GetStackTrace(nested)
	if st == nil {
		t.Fatal("Expected stack trace for nested error, got nil")
	}
	if len(st) == 0 {
		t.Error("Got empty stack trace for nested error")
	}
}

func TestWithStackFrom(t *testing.T) {
	// Test with nil errors
	err := WithStackFrom(nil, nil)
	if err != nil {
		t.Errorf("Expected nil result for WithStackFrom(nil, nil), got %v", err)
	}

	// Target error is nil
	err = WithStackFrom(nil, WithStack(io.EOF))
	if err != nil {
		t.Errorf("Expected nil result when target is nil, got %v", err)
	}

	// Source error is nil
	err = WithStackFrom(io.EOF, nil)
	if err != io.EOF {
		t.Errorf("Expected unchanged error when source is nil, got %v", err)
	}

	// Source has no stack
	err = WithStackFrom(io.EOF, fmt.Errorf("no stack"))
	if err != io.EOF {
		t.Errorf("Expected unchanged error when source has no stack, got %v", err)
	}

	// Normal case - transferring stack
	sourceErr := WithStack(fmt.Errorf("source"))
	targetErr := fmt.Errorf("target")
	result := WithStackFrom(targetErr, sourceErr)

	// Check if result has stack trace
	st := GetStackTrace(result)
	if st == nil {
		t.Fatal("Expected stack trace for result, got nil")
	}
	if len(st) == 0 {
		t.Error("Got empty stack trace for result")
	}

	// Check if error message is preserved
	if result.Error() != "target" {
		t.Errorf("Expected error message 'target', got '%s'", result.Error())
	}
}

func TestWithExtractedStackFormatting(t *testing.T) {
	// Create an error with stack
	sourceErr := WithStack(fmt.Errorf("source"))
	targetErr := fmt.Errorf("target")
	result := WithStackFrom(targetErr, sourceErr)

	// Test standard formatting
	if s := fmt.Sprint(result); s != "target" {
		t.Errorf("Expected 'target', got '%s'", s)
	}

	// Test %q formatting
	if s := fmt.Sprintf("%q", result); s != `"target"` {
		t.Errorf("Expected '\"target\"', got '%s'", s)
	}

	// Test %+v formatting
	s := fmt.Sprintf("%+v", result)
	if !strings.Contains(s, "target") {
		t.Errorf("Expected output to contain 'target', got '%s'", s)
	}

	// Stack trace should be included
	if !strings.Contains(s, "TestWithExtractedStackFormatting") {
		t.Errorf("Expected stack trace to contain function name, got '%s'", s)
	}
}

func TestStackTraceInterface(t *testing.T) {
	// Test withStack
	err1 := WithStack(io.EOF)
	st1 := GetStackTrace(err1)
	if st1 == nil || len(st1) == 0 {
		t.Error("withStack doesn't properly implement stackTracer interface")
	}

	// Test withExtractedStack
	err2 := WithStackFrom(fmt.Errorf("new error"), err1)
	st2 := GetStackTrace(err2)
	if st2 == nil || len(st2) == 0 {
		t.Error("withExtractedStack doesn't properly implement stackTracer interface")
	}
}

func TestStackTracePropagation(t *testing.T) {
	// Create a chain of errors with stack
	original := WithStack(io.EOF)
	wrapped := Wrap(original, "layer 1")
	wrapped = WithMessage(wrapped, "layer 2")

	// Stack should be retrievable from any level
	if st := GetStackTrace(wrapped); st == nil {
		t.Error("Stack trace lost during error wrapping")
	}

	// Create a new error with transferred stack
	transferred := WithStackFrom(fmt.Errorf("new error"), wrapped)

	// Format both errors
	wrappedStr := fmt.Sprintf("%+v", wrapped)
	transferredStr := fmt.Sprintf("%+v", transferred)

	// Both should have stack traces
	if !strings.Contains(wrappedStr, "TestStackTracePropagation") {
		t.Errorf("Original error is missing stack trace: %s", wrappedStr)
	}

	if !strings.Contains(transferredStr, "TestStackTracePropagation") {
		t.Errorf("Transferred error is missing stack trace: %s", transferredStr)
	}
}
