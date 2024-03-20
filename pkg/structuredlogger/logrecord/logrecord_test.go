// Package logrecord contains tests for the structured LogRecord.
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
	parameters = map[string]interface{}{
		"message": "Test Message.",
	}
)

// TestNew tests that New function returns a new structured LogRecord instance.
func TestNew(t *testing.T) {
	record := New(name, logLevel, timeFormat, parameters, skipCaller)

	testutils.AssertEquals(t, parameters, record.parameters)
}

// BenchmarkNew benchmarks the New function.
func BenchmarkNew(b *testing.B) {
	for index := 0; index < b.N; index++ {
		New(name, logLevel, timeFormat, parameters, skipCaller)
	}
}

// TestParameters tests that Parameters function returns the parameters of the log record.
func TestParameters(t *testing.T) {
	record := New(name, logLevel, timeFormat, parameters, skipCaller)

	testutils.AssertEquals(t, parameters, record.Parameters())
}

// BenchmarkParameters benchmarks the Parameters function.
func BenchmarkParameters(b *testing.B) {
	record := New(name, logLevel, timeFormat, parameters, skipCaller)
	for index := 0; index < b.N; index++ {
		record.Parameters()
	}
}
