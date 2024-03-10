// Package handler_test has tests for handler package.
package handler

import (
	"bytes"
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/loglevel"
	"io"
	"os"
	"testing"
)

var template = "%(level):%(name):%(message)"
var testFile = "/tmp/test_file.log"

// TestNew tests that New creates a new Handler instance.
func TestNew(t *testing.T) {
	newFormatter := formatter.New(template)

	newHandler := New(loglevel.Debug, newFormatter, os.Stdout, os.Stderr)

	testutils.AssertEquals(t, loglevel.Debug, newHandler.level)

	if newHandler.writer != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.writer)
	}

	if newHandler.errorWriter != os.Stderr {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stderr, newHandler.errorWriter)
	}
}

// BenchmarkNew performs benchmarking of the New().
func BenchmarkNew(b *testing.B) {
	newFormatter := formatter.New(template)

	for index := 0; index < b.N; index++ {
		New(loglevel.Debug, newFormatter, os.Stdout, os.Stderr)
	}
}

// TestNewConsoleHandler tests that NewConsoleHandler creates a new Handler
// instance that writes on the console.
func TestNewConsoleHandler(t *testing.T) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(loglevel.Debug, newFormatter)

	testutils.AssertEquals(t, loglevel.Debug, newHandler.level)

	if newHandler.writer != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.writer)
	}

	if newHandler.errorWriter != os.Stderr {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stderr, newHandler.errorWriter)
	}
}

// BenchmarkNewConsoleHandler performs benchmarking of the NewConsoleHandler().
func BenchmarkNewConsoleHandler(b *testing.B) {
	newFormatter := formatter.New(template)

	for index := 0; index < b.N; index++ {
		NewConsoleHandler(loglevel.Debug, newFormatter)
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

	newHandler := NewFileHandler(loglevel.Debug, newFormatter, testFile)

	testutils.AssertEquals(t, loglevel.Debug, newHandler.level)

	if newHandler.writer != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.writer)
	}

	if newHandler.errorWriter != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.errorWriter)
	}
}

// BenchmarkNewFileHandler performs benchmarking of the NewFileHandler().
func BenchmarkNewFileHandler(b *testing.B) {
	newFormatter := formatter.New(template)

	osOpenFile = mockOpenFile

	for index := 0; index < b.N; index++ {
		NewFileHandler(loglevel.Debug, newFormatter, testFile)
	}
}

// TestHandler_Level test that Handler.Level() returns log level for the Handler.
func TestHandler_Level(t *testing.T) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(loglevel.Debug, newFormatter)

	testutils.AssertEquals(t, newHandler.level, newHandler.Level())
}

// BenchmarkHandler_Level performs benchmarking of the Handler.Level().
func BenchmarkHandler_Level(b *testing.B) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(loglevel.Debug, newFormatter)

	for index := 0; index < b.N; index++ {
		newHandler.Level()
	}
}

// TestHandler_Level test that Handler.SetLevel() set a new log level for the
// Handler.
func TestHandler_SetLevel(t *testing.T) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(loglevel.Debug, newFormatter)

	newLevel := loglevel.Info

	newHandler.SetLevel(newLevel)

	testutils.AssertEquals(t, newLevel, newHandler.level)
}

// BenchmarkHandler_Level performs benchmarking of the Handler.SetLevel().
func BenchmarkHandler_SetLevel(b *testing.B) {
	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(loglevel.Debug, newFormatter)

	newLevel := loglevel.Info

	for index := 0; index < b.N; index++ {
		newHandler.SetLevel(newLevel)
	}
}

// TestHandler_Formatter test that Handler.Formatter() returns assigned
// Formatter.
func TestHandler_Formatter(t *testing.T) {
	var newFormatter formatter.Interface = formatter.New(template)

	newHandler := NewConsoleHandler(loglevel.Debug, newFormatter)

	testutils.AssertEquals(t, newFormatter, newHandler.Formatter())
}

// BenchmarkHandler_Formatter performs benchmarking of the Handler.Formatter().
func BenchmarkHandler_Formatter(b *testing.B) {
	var newFormatter formatter.Interface = formatter.New(template)

	newHandler := NewConsoleHandler(loglevel.Debug, newFormatter)

	for index := 0; index < b.N; index++ {
		newHandler.Formatter()
	}
}

// TestHandler_Write tests that Handler.Write() writes formatted log to the
// correct writer.
func TestHandler_Write(t *testing.T) {
	logName := "test"
	logLevel := loglevel.Debug
	logLevelError := loglevel.Error
	message := "Test message."

	formattedStdoutMessage := fmt.Sprintf("%s%s:%s:%s%s\n", "\033[36m", logLevel.String(), logName, message, "\033[0m")
	formattedStderrMessage := fmt.Sprintf("%s%s:%s:%s%s\n", "\033[31m", logLevelError.String(), logName, message, "\033[0m")

	newFormatter := formatter.New(template)

	var bufferStdout bytes.Buffer
	var bufferStderr bytes.Buffer

	originalStdout := osStdout
	originalStderr := osStderr

	readerStdout, writerStdout, _ := os.Pipe()
	readerStderr, writerStderr, _ := os.Pipe()

	osStdout = writerStdout
	osStderr = writerStderr

	newHandler := NewConsoleHandler(loglevel.Debug, newFormatter)

	newHandler.consoleSupportsANSIColors = func() bool {
		return true
	}

	newHandler.Write(logName, logLevel, message)
	newHandler.Write(logName, logLevelError, message)

	_ = writerStdout.Close()
	_ = writerStderr.Close()

	_, _ = io.Copy(&bufferStdout, readerStdout)
	_, _ = io.Copy(&bufferStderr, readerStderr)

	osStdout = originalStdout
	osStderr = originalStderr

	testutils.AssertEquals(t, formattedStdoutMessage, bufferStdout.String())
	testutils.AssertEquals(t, formattedStderrMessage, bufferStderr.String())
}

// BenchmarkHandler_Write performs benchmarking of the Handler.Write().
func BenchmarkHandler_Write(b *testing.B) {
	logName := "test"
	logLevel := loglevel.Debug
	message := "Test message."

	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(loglevel.Debug, newFormatter)

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
