// Package handler_test has tests for handler package.
package handler

import (
	"bytes"
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"io"
	"os"
	"testing"
)

var template = "%(level):%(name):%(message)"
var testFile = "/tmp/test_file.log"

// TestNew tests that New creates a new Handler instance.
func TestNew(t *testing.T) {
	newFormatter := formatter.New(template)

	newHandler := New(level.Debug, level.Null, newFormatter, os.Stdout)

	testutils.AssertEquals(t, level.Debug, newHandler.fromLevel)
	testutils.AssertEquals(t, level.Null, newHandler.toLevel)

	if newHandler.writer != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.writer)
	}
}

// BenchmarkNew performs benchmarking of the New().
func BenchmarkNew(b *testing.B) {
	newFormatter := formatter.New(template)

	for index := 0; index < b.N; index++ {
		New(level.Debug, level.Null, newFormatter, os.Stdout)
	}
}

// TestNewConsoleHandler tests that NewConsoleHandler creates a new Handler
// instance that writes on the console.
func TestNewConsoleHandler(t *testing.T) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	testutils.AssertEquals(t, level.Debug, newHandler.fromLevel)

	if newHandler.writer != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.writer)
	}
}

// BenchmarkNewConsoleHandler performs benchmarking of the NewConsoleHandler().
func BenchmarkNewConsoleHandler(b *testing.B) {
	newFormatter := formatter.New(template)

	for index := 0; index < b.N; index++ {
		NewConsoleHandler(level.Debug, level.Null, newFormatter)
	}
}

// mockOpenFile mocks os.OpenFile method and returns writer to os.Stdout.
func mockOpenFile(_ string, _ int, _ os.FileMode) (*os.File, error) {
	return os.Stdout, nil
}

// TestNewFileHandler test that NewFileHandler creates a new Handler instance
// that writes to the file.
func TestNewFileHandler(t *testing.T) {
	newFormatter := formatter.New(template)

	osOpenFile = mockOpenFile

	newHandler := NewFileHandler(level.Debug, level.Null, newFormatter, testFile)

	testutils.AssertEquals(t, level.Debug, newHandler.fromLevel)

	if newHandler.writer != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.writer)
	}
}

// TestNewFileHandlerError test that NewFileHandler returns error if file cannot
// be opened.
func TestNewFileHandlerError(t *testing.T) {
	newFormatter := formatter.New(template)

	osOpenFile = func(_ string, _ int, _ os.FileMode) (*os.File, error) {
		return nil, fmt.Errorf("error")
	}

	newHandler := NewFileHandler(level.Debug, level.Null, newFormatter, testFile)

	testutils.AssertEquals(t, nil, newHandler)
}

// BenchmarkNewFileHandler performs benchmarking of the NewFileHandler().
func BenchmarkNewFileHandler(b *testing.B) {
	newFormatter := formatter.New(template)

	osOpenFile = mockOpenFile

	for index := 0; index < b.N; index++ {
		NewFileHandler(level.Debug, level.Null, newFormatter, testFile)
	}
}

// TestHandler_Writer test that Handler.Writer() returns writer for the Handler.
func TestHandler_Writer(t *testing.T) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	if newHandler.Writer() != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.Writer())
	}
}

// BenchmarkHandler_Writer performs benchmarking of the Handler.Writer().
func BenchmarkHandler_Writer(b *testing.B) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	for index := 0; index < b.N; index++ {
		newHandler.Writer()
	}
}

// TestHandler_FromLevel test that Handler.FromLevel() returns log fromLevel for
// the Handler.
func TestHandler_FromLevel(t *testing.T) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	testutils.AssertEquals(t, newHandler.fromLevel, newHandler.FromLevel())
}

// BenchmarkHandler_FromLevel performs benchmarking of the Handler.FromLevel().
func BenchmarkHandler_FromLevel(b *testing.B) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	for index := 0; index < b.N; index++ {
		newHandler.FromLevel()
	}
}

// TestHandler_SetFromLevel test that Handler.SetFromLevel() set a new log toLevel for
// the Handler.
func TestHandler_SetFromLevel(t *testing.T) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	newLevel := level.Info

	newHandler.SetFromLevel(newLevel)

	testutils.AssertEquals(t, newLevel, newHandler.fromLevel)
}

// BenchmarkHandler_SetFromLevel performs benchmarking of the
// Handler.SetFromLevel().
func BenchmarkHandler_SetFromLevel(b *testing.B) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	newLevel := level.Info

	for index := 0; index < b.N; index++ {
		newHandler.SetFromLevel(newLevel)
	}
}

// TestHandler_ToLevel test that Handler.ToLevel() returns log toLevel for the Handler.
func TestHandler_ToLevel(t *testing.T) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	testutils.AssertEquals(t, newHandler.toLevel, newHandler.ToLevel())
}

// BenchmarkHandler_ToLevel performs benchmarking of the Handler.ToLevel().
func BenchmarkHandler_ToLevel(b *testing.B) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	for index := 0; index < b.N; index++ {
		newHandler.ToLevel()
	}
}

// TestHandler_SetToLevel test that Handler.SetToLevel() set a new log toLevel
// for the Handler.
func TestHandler_SetToLevel(t *testing.T) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	newLevel := level.Info

	newHandler.SetToLevel(newLevel)

	testutils.AssertEquals(t, newLevel, newHandler.toLevel)
}

// BenchmarkHandler_SetToLevel performs benchmarking of the Handler.SetToLevel().
func BenchmarkHandler_SetToLevel(b *testing.B) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	newLevel := level.Info

	for index := 0; index < b.N; index++ {
		newHandler.SetToLevel(newLevel)
	}
}

// TestHandler_Formatter test that Handler.Formatter() returns assigned
// Formatter.
func TestHandler_Formatter(t *testing.T) {
	var newFormatter formatter.Interface = formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	testutils.AssertEquals(t, newFormatter, newHandler.Formatter())
}

// BenchmarkHandler_Formatter performs benchmarking of the Handler.Formatter().
func BenchmarkHandler_Formatter(b *testing.B) {
	var newFormatter formatter.Interface = formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	for index := 0; index < b.N; index++ {
		newHandler.Formatter()
	}
}

// setupHandler is a helper function to setup a new handler for testing purposes.
func setupHandler(fromLevel, toLevel level.Level, supportsANSI bool, formatterTemplate string) *Handler {
	newFormatter := formatter.New(formatterTemplate)
	newHandler := NewConsoleHandler(fromLevel, toLevel, newFormatter)
	newHandler.consoleSupportsANSIColors = func() bool {
		return supportsANSI
	}
	return newHandler
}

// TestHandler_Write tests that Handler.Write() writes formatted log to the
// correct writer.
func TestHandler_Write(t *testing.T) {
	logName := "test"
	logLevel := level.Debug
	message := "Test message."

	originalStdout := osStdout

	readerStdout, writerStdout, _ := os.Pipe()

	osStdout = writerStdout

	handler := setupHandler(level.Debug, level.Null, true, template)

	var bufferStdout bytes.Buffer

	handler.Write(logName, logLevel, message)

	_ = writerStdout.Close()

	_, _ = io.Copy(&bufferStdout, readerStdout)

	osStdout = originalStdout

	testutils.AssertEquals(t, "\033[36mdebug:test:Test message.\033[0m\n", bufferStdout.String())
}

// TestHandler_WriteError tests that Handler.Write() returns error if writer fails.
func TestHandler_WriteError(t *testing.T) {
	logName := "test"
	logLevel := level.Debug
	message := "Test message."

	originalStdout := osStdout

	readerStdout, writerStdout, _ := os.Pipe()

	osStdout = writerStdout

	handler := setupHandler(level.Warning, level.Null, false, template)

	var bufferStdout bytes.Buffer

	handler.Write(logName, logLevel, message)

	_ = writerStdout.Close()

	_, _ = io.Copy(&bufferStdout, readerStdout)

	osStdout = originalStdout

	testutils.AssertEquals(t, "", bufferStdout.String())
}

// BenchmarkHandler_Write performs benchmarking of the Handler.Write().
func BenchmarkHandler_Write(b *testing.B) {
	logName := "test"
	logLevel := level.Debug
	message := "Test message."

	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(level.Debug, level.Null, newFormatter)

	newHandler.writer = io.Discard

	for index := 0; index < b.N; index++ {
		newHandler.Write(logName, logLevel, message)
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
