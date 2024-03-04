// Package handler_test has tests for handler package.
package handler

import (
	"bytes"
	"fmt"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/loglevel"
	"io"
	"os"
	"testing"
)

var format = "%(level):%(name):%(message)"
var testFile = "/tmp/test_file.log"

// TestNew tests that New creates a new Handler instance.
func TestNew(t *testing.T) {
	newFormatter := formatter.New(format)

	newHandler := New(loglevel.Debug, *newFormatter, os.Stdout, os.Stderr)

	if newHandler.level != loglevel.Debug {
		t.Fatalf("log level is not the same. expected: %v, actual: %v", loglevel.Debug, newHandler.level)
	}

	if !newHandler.formatter.IsEqual(newFormatter) {
		t.Fatalf("formatter is not the same. expected: %v, actual: %v", newFormatter, newHandler.formatter)
	}

	if newHandler.writer != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.writer)
	}

	if newHandler.errorWriter != os.Stderr {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stderr, newHandler.errorWriter)
	}
}

// BenchmarkNew performs benchmarking of the New().
func BenchmarkNew(b *testing.B) {
	newFormatter := formatter.New(format)

	for index := 0; index < b.N; index++ {
		New(loglevel.Debug, *newFormatter, os.Stdout, os.Stderr)
	}
}

// TestNewConsoleHandler tests that NewConsoleHandler creates a new Handler
// instance that writes on the console.
func TestNewConsoleHandler(t *testing.T) {
	newFormatter := formatter.New(format)

	newHandler := NewConsoleHandler(loglevel.Debug, *newFormatter)

	if newHandler.level != loglevel.Debug {
		t.Fatalf("log level is not the same. expected: %v, actual: %v", loglevel.Debug, newHandler.level)
	}

	if !newHandler.formatter.IsEqual(newFormatter) {
		t.Fatalf("formatter is not the same. expected: %v, actual: %v", newFormatter, newHandler.formatter)
	}

	if newHandler.writer != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.writer)
	}

	if newHandler.errorWriter != os.Stderr {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stderr, newHandler.errorWriter)
	}
}

// BenchmarkNewConsoleHandler performs benchmarking of the NewConsoleHandler().
func BenchmarkNewConsoleHandler(b *testing.B) {
	newFormatter := formatter.New(format)

	for index := 0; index < b.N; index++ {
		NewConsoleHandler(loglevel.Debug, *newFormatter)
	}
}

// mockOpenFile mocks os.OpenFile method and returns writer to os.Stdout.
func mockOpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.Stdout, nil
}

// TestNewFileHandler test that NewFileHandler creates a new Handler instance
// that writes to the file.
func TestNewFileHandler(t *testing.T) {
	newFormatter := formatter.New(format)

	osOpenFile = mockOpenFile

	newHandler := NewFileHandler(loglevel.Debug, *newFormatter, testFile)

	if newHandler.level != loglevel.Debug {
		t.Fatalf("log level is not the same. expected: %v, actual: %v", loglevel.Debug, newHandler.level)
	}

	if !newHandler.formatter.IsEqual(newFormatter) {
		t.Fatalf("formatter is not the same. expected: %v, actual: %v", newFormatter, newHandler.formatter)
	}

	if newHandler.writer != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.writer)
	}

	if newHandler.errorWriter != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.errorWriter)
	}
}

// BenchmarkNewFileHandler performs benchmarking of the NewFileHandler().
func BenchmarkNewFileHandler(b *testing.B) {
	newFormatter := formatter.New(format)

	osOpenFile = mockOpenFile

	for index := 0; index < b.N; index++ {
		NewFileHandler(loglevel.Debug, *newFormatter, testFile)
	}
}

// TestHandler_Level test that Handler.Level() returns log level for the Handler.
func TestHandler_Level(t *testing.T) {
	newFormatter := formatter.New(format)

	newHandler := NewConsoleHandler(loglevel.Debug, *newFormatter)

	if newHandler.Level() != newHandler.level {
		t.Fatalf("expected: %v, actual: %v", newHandler.level, newHandler.Level())
	}
}

// BenchmarkHandler_Level performs benchmarking of the Handler.Level().
func BenchmarkHandler_Level(b *testing.B) {
	newFormatter := formatter.New(format)

	newHandler := NewConsoleHandler(loglevel.Debug, *newFormatter)

	for index := 0; index < b.N; index++ {
		newHandler.Level()
	}
}

// TestHandler_Level test that Handler.SetLevel() set a new log level for the
// Handler.
func TestHandler_SetLevel(t *testing.T) {
	newFormatter := formatter.New(format)

	newHandler := NewConsoleHandler(loglevel.Debug, *newFormatter)

	newLevel := loglevel.LogLevel(loglevel.Info)

	newHandler.SetLevel(newLevel)

	if newHandler.level != newLevel {
		t.Fatalf("expected: %v, actual: %v", newLevel, newHandler.level)
	}
}

// BenchmarkHandler_Level performs benchmarking of the Handler.SetLevel().
func BenchmarkHandler_SetLevel(b *testing.B) {
	newFormatter := formatter.New(format)

	newHandler := NewConsoleHandler(loglevel.Debug, *newFormatter)

	newLevel := loglevel.LogLevel(loglevel.Info)

	for index := 0; index < b.N; index++ {
		newHandler.SetLevel(newLevel)
	}
}

// TestHandler_Write tests that Handler.Write() writes formatted log to the
// correct writer.
func TestHandler_Write(t *testing.T) {
	logName := "test"
	logLevel := loglevel.LogLevel(loglevel.Debug)
	logLevelError := loglevel.LogLevel(loglevel.Error)
	message := "Test message."

	formattedStdoutMessage := fmt.Sprintf("%s%s:%s:%s%s\n", "\033[36m", logLevel.String(), logName, message, "\033[0m")
	formattedStderrMessage := fmt.Sprintf("%s%s:%s:%s%s\n", "\033[31m", logLevelError.String(), logName, message, "\033[0m")

	newFormatter := formatter.New(format)

	var bufferStdout bytes.Buffer
	var bufferStderr bytes.Buffer

	originalStdout := osStdout
	originalStderr := osStderr

	readerStdout, writerStdout, _ := os.Pipe()
	readerStderr, writerStderr, _ := os.Pipe()

	osStdout = writerStdout
	osStderr = writerStderr

	newHandler := NewConsoleHandler(loglevel.Debug, *newFormatter)

	newHandler.Write(logName, logLevel, message)
	newHandler.Write(logName, logLevelError, message)

	writerStdout.Close()
	writerStderr.Close()

	io.Copy(&bufferStdout, readerStdout)
	io.Copy(&bufferStderr, readerStderr)

	osStdout = originalStdout
	osStderr = originalStderr

	if bufferStdout.String() != formattedStdoutMessage {
		t.Fatalf("expected: %s, actual: %s", formattedStdoutMessage, bufferStdout.String())
	}

	if bufferStderr.String() != formattedStderrMessage {
		t.Fatalf("expected: %s, actual: %s", formattedStderrMessage, bufferStderr.String())
	}
}

// BenchmarkHandler_Write performs benchmarking of the Handler.Write().
func BenchmarkHandler_Write(b *testing.B) {
	logName := "test"
	logLevel := loglevel.LogLevel(loglevel.Debug)
	message := "Test message."

	newFormatter := formatter.New(format)

	newHandler := NewConsoleHandler(loglevel.Debug, *newFormatter)

	newHandler.writer = io.Discard

	for index := 0; index < b.N; index++ {
		newHandler.Write(logName, logLevel, message)
	}
}