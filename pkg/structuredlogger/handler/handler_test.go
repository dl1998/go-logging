// Package handler_test has tests for handler package.
package handler

import (
	"bytes"
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/structuredlogger/formatter"
	"io"
	"os"
	"testing"
)

const (
	testFile  = "/tmp/test_file.log"
	fromLevel = level.Warning
	toLevel   = level.Null
	pretty    = false
)

var template = map[string]string{
	"level": "%(level)",
	"name":  "%(name)",
}

// TestNew tests that New creates a new Handler instance.
func TestNew(t *testing.T) {
	newFormatter := formatter.NewJSON(template, pretty)

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
	newFormatter := formatter.NewJSON(template, pretty)

	for index := 0; index < b.N; index++ {
		New(fromLevel, toLevel, newFormatter, os.Stdout)
	}
}

// TestNewConsoleHandler tests that NewConsoleHandler creates a new Handler
// instance that writes on the console.
func TestNewConsoleHandler(t *testing.T) {
	writer := os.Stdout

	newFormatter := formatter.NewJSON(template, pretty)

	newHandler := NewConsoleHandler(fromLevel, toLevel, newFormatter)

	testutils.AssertEquals(t, fromLevel, newHandler.FromLevel())
	testutils.AssertEquals(t, toLevel, newHandler.ToLevel())

	if newHandler.Writer() != writer {
		t.Fatalf("writer is not the same. expected: %v, actual: %v", writer, newHandler.Writer())
	}
}

// BenchmarkNewConsoleHandler performs benchmarking of the NewConsoleHandler().
func BenchmarkNewConsoleHandler(b *testing.B) {
	newFormatter := formatter.NewJSON(template, pretty)

	for index := 0; index < b.N; index++ {
		NewConsoleHandler(fromLevel, toLevel, newFormatter)
	}
}

// TestNewConsoleErrorHandler tests that NewConsoleErrorHandler creates a new
// Handler instance that writes on the console.
func TestNewConsoleErrorHandler(t *testing.T) {
	writer := os.Stderr

	newFormatter := formatter.NewJSON(template, pretty)

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
	newFormatter := formatter.NewJSON(template, pretty)

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
	newFormatter := formatter.NewJSON(template, pretty)

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
	newFormatter := formatter.NewJSON(template, pretty)

	osOpenFile = func(_ string, _ int, _ os.FileMode) (*os.File, error) {
		return nil, fmt.Errorf("error")
	}

	newHandler := NewFileHandler(fromLevel, toLevel, newFormatter, testFile)

	testutils.AssertEquals(t, nil, newHandler)
}

// BenchmarkNewFileHandler performs benchmarking of the NewFileHandler().
func BenchmarkNewFileHandler(b *testing.B) {
	newFormatter := formatter.NewJSON(template, pretty)

	osOpenFile = mockOpenFile

	for index := 0; index < b.N; index++ {
		NewFileHandler(fromLevel, toLevel, newFormatter, testFile)
	}
}

// TestHandler_Formatter test that Handler.Formatter() returns assigned
// Formatter.
func TestHandler_Formatter(t *testing.T) {
	var newFormatter formatter.Interface = formatter.NewJSON(template, pretty)

	newHandler := NewConsoleHandler(fromLevel, toLevel, newFormatter)

	testutils.AssertEquals(t, newFormatter, newHandler.Formatter())
}

// BenchmarkHandler_Formatter performs benchmarking of the Handler.Formatter().
func BenchmarkHandler_Formatter(b *testing.B) {
	var newFormatter formatter.Interface = formatter.NewJSON(template, pretty)

	newHandler := NewConsoleHandler(fromLevel, toLevel, newFormatter)

	for index := 0; index < b.N; index++ {
		newHandler.Formatter()
	}
}

// setupHandler is a helper function to setup a new handler for testing purposes.
func setupHandler(fromLevel, toLevel level.Level, supportsANSI bool, formatterTemplate map[string]string) *Handler {
	newFormatter := formatter.NewJSON(formatterTemplate, pretty)
	newHandler := NewConsoleHandler(fromLevel, toLevel, newFormatter)
	newHandler.ConsoleSupportsANSIColors = func() bool {
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

	handler := setupHandler(logLevel, toLevel, true, template)

	var bufferStdout bytes.Buffer

	handler.Write(logName, logLevel, message)

	_ = writerStdout.Close()

	_, _ = io.Copy(&bufferStdout, readerStdout)

	osStdout = originalStdout

	testutils.AssertEquals(t, fmt.Sprintf("\033[36m{\"level\":%q,\"name\":%q}\033[0m\n", logLevel.String(), logName), bufferStdout.String())
}

// TestHandler_WriteError tests that Handler.Write() returns error if writer fails.
func TestHandler_WriteError(t *testing.T) {
	logName := "test"
	logLevel := level.Debug
	message := "Test message."

	originalStdout := osStdout

	readerStdout, writerStdout, _ := os.Pipe()

	osStdout = writerStdout

	handler := setupHandler(fromLevel, toLevel, false, template)

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
	logLevel := level.Warning
	message := "Test message."

	newFormatter := formatter.NewJSON(template, pretty)

	newHandler := NewConsoleHandler(fromLevel, toLevel, newFormatter)

	newHandler.Handler.SetWriter(io.Discard)

	for index := 0; index < b.N; index++ {
		newHandler.Write(logName, logLevel, message)
	}
}
