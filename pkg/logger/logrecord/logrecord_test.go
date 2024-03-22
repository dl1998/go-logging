// Package logrecord contains tests for the LogRecord.
package logrecord

import (
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"testing"
)

var (
	skipCaller = 1
	logLevel   = level.Debug
	name       = "test"
	timeFormat = ""
	message    = "Test message."
	parameters = make([]any, 0)
)

// TestNew tests that New function returns a new LogRecord instance.
func TestNew(t *testing.T) {
	record := New(name, logLevel, timeFormat, message, parameters, skipCaller)

	testutils.AssertEquals(t, message, record.message)
}

// BenchmarkNew benchmarks the New function.
func BenchmarkNew(b *testing.B) {
	for index := 0; index < b.N; index++ {
		New(name, logLevel, timeFormat, message, parameters, skipCaller)
	}
}

// TestMessage tests that Message function returns the formatted message of the
// log record.
func TestMessage(t *testing.T) {
	record := New(name, logLevel, timeFormat, message, parameters, skipCaller)

	testutils.AssertEquals(t, message, record.Message())
}

// BenchmarkMessage benchmarks the Message function.
func BenchmarkMessage(b *testing.B) {
	record := New(name, logLevel, timeFormat, message, parameters, skipCaller)
	for index := 0; index < b.N; index++ {
		record.Message()
	}
}
