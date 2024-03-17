// Package handler_test has tests for handler package.
package handler

import (
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"os"
	"testing"
)

const (
	template  = "%(level):%(name):%(message)"
	fromLevel = level.Warning
	toLevel   = level.Null
)

// TestHandler_Writer test that Handler.Writer() returns writer for the Handler.
func TestHandler_Writer(t *testing.T) {
	newHandler := New(fromLevel, toLevel, os.Stdout)

	if newHandler.Writer() != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.Writer())
	}
}

// BenchmarkHandler_Writer performs benchmarking of the Handler.Writer().
func BenchmarkHandler_Writer(b *testing.B) {
	newHandler := New(fromLevel, toLevel, os.Stdout)

	for index := 0; index < b.N; index++ {
		newHandler.Writer()
	}
}

// TestHandler_SetWriter test that Handler.SetWriter() set a new writer for the Handler.
func TestHandler_SetWriter(t *testing.T) {
	newHandler := New(fromLevel, toLevel, os.Stdout)

	writer := os.Stderr

	newHandler.SetWriter(writer)

	if newHandler.Writer() != writer {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", writer, newHandler.Writer())
	}
}

// BenchmarkHandler_SetWriter performs benchmarking of the Handler.SetWriter().
func BenchmarkHandler_SetWriter(b *testing.B) {
	newHandler := New(fromLevel, toLevel, os.Stdout)

	writer := os.Stderr

	for index := 0; index < b.N; index++ {
		newHandler.SetWriter(writer)
	}
}

// TestHandler_FromLevel test that Handler.FromLevel() returns log fromLevel for
// the Handler.
func TestHandler_FromLevel(t *testing.T) {
	newHandler := New(fromLevel, toLevel, os.Stdout)

	testutils.AssertEquals(t, fromLevel, newHandler.FromLevel())
}

// BenchmarkHandler_FromLevel performs benchmarking of the Handler.FromLevel().
func BenchmarkHandler_FromLevel(b *testing.B) {
	newHandler := New(fromLevel, toLevel, os.Stdout)

	for index := 0; index < b.N; index++ {
		newHandler.FromLevel()
	}
}

// TestHandler_SetFromLevel test that Handler.SetFromLevel() set a new log toLevel for
// the Handler.
func TestHandler_SetFromLevel(t *testing.T) {
	newHandler := New(fromLevel, toLevel, os.Stdout)

	newLevel := level.Info

	newHandler.SetFromLevel(newLevel)

	testutils.AssertEquals(t, newLevel, newHandler.FromLevel())
}

// BenchmarkHandler_SetFromLevel performs benchmarking of the
// Handler.SetFromLevel().
func BenchmarkHandler_SetFromLevel(b *testing.B) {
	newHandler := New(fromLevel, toLevel, os.Stdout)

	newLevel := level.Info

	for index := 0; index < b.N; index++ {
		newHandler.SetFromLevel(newLevel)
	}
}

// TestHandler_ToLevel test that Handler.ToLevel() returns log toLevel for the Handler.
func TestHandler_ToLevel(t *testing.T) {
	newHandler := New(fromLevel, toLevel, os.Stdout)

	testutils.AssertEquals(t, toLevel, newHandler.ToLevel())
}

// BenchmarkHandler_ToLevel performs benchmarking of the Handler.ToLevel().
func BenchmarkHandler_ToLevel(b *testing.B) {
	newHandler := New(fromLevel, toLevel, os.Stdout)

	for index := 0; index < b.N; index++ {
		newHandler.ToLevel()
	}
}

// TestHandler_SetToLevel test that Handler.SetToLevel() set a new log toLevel
// for the Handler.
func TestHandler_SetToLevel(t *testing.T) {
	newHandler := New(fromLevel, toLevel, os.Stdout)

	newLevel := level.Info

	newHandler.SetToLevel(newLevel)

	testutils.AssertEquals(t, newLevel, newHandler.ToLevel())
}

// BenchmarkHandler_SetToLevel performs benchmarking of the Handler.SetToLevel().
func BenchmarkHandler_SetToLevel(b *testing.B) {
	newHandler := New(fromLevel, toLevel, os.Stdout)

	newLevel := level.Info

	for index := 0; index < b.N; index++ {
		newHandler.SetToLevel(newLevel)
	}
}

// TestConsoleSupportsANSIColors tests that consoleSupportsANSIColors returns
// true if console supports ANSI colors, or otherwise it returns false.
func TestConsoleSupportsANSIColors(t *testing.T) {
	parameters := map[string]struct {
		Expected bool
		Term     string
	}{
		"Console supports ANSI colors": {
			true, "xterm-256color",
		},
		"Console don't support ANSI colors": {
			false, "",
		},
	}
	for name, parameter := range parameters {
		t.Run(name, func(t *testing.T) {
			_ = os.Setenv("TERM", parameter.Term)
			result := consoleSupportsANSIColors()
			testutils.AssertEquals(t, parameter.Expected, result)
		})
	}
	_ = os.Setenv("TERM", "xterm-256color")
}

// BenchmarkConsoleSupportsANSIColors performs benchmarking of the
// consoleSupportsANSIColors().
func BenchmarkConsoleSupportsANSIColors(b *testing.B) {
	for index := 0; index < b.N; index++ {
		consoleSupportsANSIColors()
	}
}
