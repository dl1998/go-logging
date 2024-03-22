// Package handler_test has tests for handler package.
package handler

import (
	"bytes"
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/logrecord"
	"io"
	"os"
	"testing"
)

const (
	template   = "%(level):%(name):%(message)"
	testFile   = "/tmp/test_file.log"
	fromLevel  = level.Warning
	toLevel    = level.Null
	loggerName = "test"
	message    = "Test message."
)

var emptyParameters = make([]any, 0)

// TestNew tests that New creates a new Handler instance.
func TestNew(t *testing.T) {
	newFormatter := formatter.New(template)

	writer := os.Stdout

	newHandler := New(fromLevel, toLevel, newFormatter, writer)

	testutils.AssertEquals(t, fromLevel, newHandler.FromLevel())
	testutils.AssertEquals(t, toLevel, newHandler.ToLevel())

	if newHandler.Writer() != writer {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", writer, newHandler.Writer())
	}
}

// BenchmarkNew performs benchmarking of the New().
func BenchmarkNew(b *testing.B) {
	newFormatter := formatter.New(template)

	for index := 0; index < b.N; index++ {
		New(fromLevel, toLevel, newFormatter, os.Stdout)
	}
}

// TestNewConsoleHandler tests that NewConsoleHandler creates a new Handler
// instance that writes on the console.
func TestNewConsoleHandler(t *testing.T) {
	writer := os.Stdout

	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(fromLevel, toLevel, newFormatter)

	testutils.AssertEquals(t, fromLevel, newHandler.FromLevel())
	testutils.AssertEquals(t, toLevel, newHandler.ToLevel())

	if newHandler.Writer() != writer {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", writer, newHandler.Writer())
	}
}

// BenchmarkNewConsoleHandler performs benchmarking of the NewConsoleHandler().
func BenchmarkNewConsoleHandler(b *testing.B) {
	newFormatter := formatter.New(template)

	for index := 0; index < b.N; index++ {
		NewConsoleHandler(fromLevel, toLevel, newFormatter)
	}
}

// TestNewConsoleErrorHandler tests that NewConsoleErrorHandler creates a new
// Handler instance that writes on the console.
func TestNewConsoleErrorHandler(t *testing.T) {
	writer := os.Stderr

	newFormatter := formatter.New(template)

	newHandler := NewConsoleErrorHandler(fromLevel, toLevel, newFormatter)

	testutils.AssertEquals(t, fromLevel, newHandler.FromLevel())
	testutils.AssertEquals(t, toLevel, newHandler.ToLevel())

	if newHandler.Writer() != writer {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", writer, newHandler.Writer())
	}
}

// BenchmarkNewConsoleErrorHandler performs benchmarking of the
// NewConsoleErrorHandler().
func BenchmarkNewConsoleErrorHandler(b *testing.B) {
	newFormatter := formatter.New(template)

	for index := 0; index < b.N; index++ {
		NewConsoleErrorHandler(fromLevel, toLevel, newFormatter)
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

	newHandler := NewFileHandler(fromLevel, toLevel, newFormatter, testFile)

	testutils.AssertEquals(t, fromLevel, newHandler.FromLevel())
	testutils.AssertEquals(t, toLevel, newHandler.ToLevel())

	if newHandler.Writer() != os.Stdout {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", os.Stdout, newHandler.Writer())
	}
}

// TestNewFileHandlerError test that NewFileHandler returns error if file cannot
// be opened.
func TestNewFileHandlerError(t *testing.T) {
	newFormatter := formatter.New(template)

	osOpenFile = func(_ string, _ int, _ os.FileMode) (*os.File, error) {
		return nil, fmt.Errorf("error")
	}

	newHandler := NewFileHandler(fromLevel, toLevel, newFormatter, testFile)

	testutils.AssertEquals(t, nil, newHandler)
}

// BenchmarkNewFileHandler performs benchmarking of the NewFileHandler().
func BenchmarkNewFileHandler(b *testing.B) {
	newFormatter := formatter.New(template)

	osOpenFile = mockOpenFile

	for index := 0; index < b.N; index++ {
		NewFileHandler(fromLevel, toLevel, newFormatter, testFile)
	}
}

// TestHandler_Formatter test that Handler.Formatter() returns assigned
// Formatter.
func TestHandler_Formatter(t *testing.T) {
	var newFormatter formatter.Interface = formatter.New(template)

	newHandler := NewConsoleHandler(fromLevel, toLevel, newFormatter)

	testutils.AssertEquals(t, newFormatter, newHandler.Formatter())
}

// BenchmarkHandler_Formatter performs benchmarking of the Handler.Formatter().
func BenchmarkHandler_Formatter(b *testing.B) {
	var newFormatter formatter.Interface = formatter.New(template)

	newHandler := NewConsoleHandler(fromLevel, toLevel, newFormatter)

	for index := 0; index < b.N; index++ {
		newHandler.Formatter()
	}
}

// setupHandler is a helper function to set up a new handler for testing purposes.
func setupHandler(fromLevel, toLevel level.Level, supportsANSI bool, formatterTemplate string) *Handler {
	newFormatter := formatter.New(formatterTemplate)
	newHandler := NewConsoleHandler(fromLevel, toLevel, newFormatter)
	newHandler.ConsoleSupportsANSIColors = func() bool {
		return supportsANSI
	}
	return newHandler
}

// TestHandler_Write tests that Handler.Write() writes formatted log to the
// correct writer.
func TestHandler_Write(t *testing.T) {
	logLevel := level.Debug

	originalStdout := osStdout

	readerStdout, writerStdout, _ := os.Pipe()

	osStdout = writerStdout

	handler := setupHandler(logLevel, toLevel, true, template)

	var bufferStdout bytes.Buffer

	record := logrecord.New(loggerName, logLevel, "", message, emptyParameters, 1)

	handler.Write(record)

	_ = writerStdout.Close()

	_, _ = io.Copy(&bufferStdout, readerStdout)

	osStdout = originalStdout

	testutils.AssertEquals(t, "\033[36mdebug:test:Test message.\033[0m\n", bufferStdout.String())
}

// TestHandler_WriteError tests that Handler.Write() returns error if writer fails.
func TestHandler_WriteError(t *testing.T) {
	logLevel := level.Debug

	originalStdout := osStdout

	readerStdout, writerStdout, _ := os.Pipe()

	osStdout = writerStdout

	handler := setupHandler(fromLevel, toLevel, false, template)

	var bufferStdout bytes.Buffer

	record := logrecord.New(loggerName, logLevel, "", message, emptyParameters, 1)

	handler.Write(record)

	_ = writerStdout.Close()

	_, _ = io.Copy(&bufferStdout, readerStdout)

	osStdout = originalStdout

	testutils.AssertEquals(t, "", bufferStdout.String())
}

// BenchmarkHandler_Write performs benchmarking of the Handler.Write().
func BenchmarkHandler_Write(b *testing.B) {
	logLevel := level.Warning

	newFormatter := formatter.New(template)

	newHandler := NewConsoleHandler(fromLevel, toLevel, newFormatter)

	newHandler.Handler.SetWriter(io.Discard)

	record := logrecord.New(loggerName, logLevel, "", message, emptyParameters, 1)

	b.ResetTimer()

	for index := 0; index < b.N; index++ {
		newHandler.Write(record)
	}
}
