// Package loglevel_test has tests for loglevel package.
package loglevel

import (
	"testing"
)

// TestLogLevel_String tests that LogLevel correctly converts value to string.
func TestLogLevel_String(t *testing.T) {
	parameters := []struct {
		input    LogLevel
		expected string
	}{
		{None, "none"},
		{Trace, "trace"},
		{Debug, "debug"},
		{Verbose, "verbose"},
		{Info, "info"},
		{Notice, "notice"},
		{Warning, "warning"},
		{Severe, "severe"},
		{Error, "error"},
		{Alert, "alert"},
		{Critical, "critical"},
		{Emergency, "emergency"},
	}

	for index := range parameters {
		actual := parameters[index].input.String()
		if actual != parameters[index].expected {
			t.Fatalf("expected: %s, actual: %s", parameters[index].expected, actual)
		}
	}
}

// BenchmarkLogLevel_String performs benchmarking of the LogLevel.String().
func BenchmarkLogLevel_String(b *testing.B) {
	level := LogLevel(Debug)

	for index := 0; index < b.N; index++ {
		level.String()
	}
}

// TestLogLevel_String tests that LogLevel returns correct digital representation
// of the value.
func TestLogLevel_DigitRepresentation(t *testing.T) {
	parameters := []struct {
		input    LogLevel
		expected int
	}{
		{None, 0},
		{Trace, 5},
		{Debug, 10},
		{Verbose, 15},
		{Info, 20},
		{Notice, 25},
		{Warning, 30},
		{Severe, 35},
		{Error, 40},
		{Alert, 45},
		{Critical, 50},
		{Emergency, 55},
	}

	for index := range parameters {
		actual := parameters[index].input.DigitRepresentation()
		if actual != parameters[index].expected {
			t.Fatalf("expected: %d, actual: %d", parameters[index].expected, actual)
		}
	}
}

// BenchmarkLogLevel_DigitRepresentation performs benchmarking of the LogLevel.DigitRepresentation().
func BenchmarkLogLevel_DigitRepresentation(b *testing.B) {
	level := LogLevel(Debug)

	for index := 0; index < b.N; index++ {
		level.DigitRepresentation()
	}
}
