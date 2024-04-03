// Package loglevel_test has tests for level package.
package level

import (
	"github.com/dl1998/go-logging/internal/testutils"
	"testing"
)

// TestParseLevel tests that ParseLevel correctly converts string to level.
func TestParseLevel(t *testing.T) {
	parameters := []struct {
		input    string
		expected Level
	}{
		{"all", All},
		{"trace", Trace},
		{"debug", Debug},
		{"verbose", Verbose},
		{"info", Info},
		{"notice", Notice},
		{"warning", Warning},
		{"severe", Severe},
		{"error", Error},
		{"alert", Alert},
		{"critical", Critical},
		{"emergency", Emergency},
		{"null", Null},
		{"", Null},
	}

	for index := range parameters {
		actual := ParseLevel(parameters[index].input)
		testutils.AssertEquals(t, parameters[index].expected, actual)
	}
}

// BenchmarkParseLevel performs benchmarking of the ParseLevel().
func BenchmarkParseLevel(b *testing.B) {
	for index := 0; index < b.N; index++ {
		ParseLevel("debug")
	}
}

// TestLogLevel_String tests that Level correctly converts value to string.
func TestLogLevel_String(t *testing.T) {
	parameters := []struct {
		input    Level
		expected string
	}{
		{All, "all"},
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
		{Null, "null"},
	}

	for index := range parameters {
		testutils.AssertEquals(t, parameters[index].expected, parameters[index].input.String())
	}
}

// BenchmarkLogLevel_String performs benchmarking of the Level.String().
func BenchmarkLogLevel_String(b *testing.B) {
	level := Debug

	for index := 0; index < b.N; index++ {
		_ = level.String()
	}
}

// TestLogLevel_String tests that Level returns correct digital representation
// of the value.
func TestLogLevel_DigitRepresentation(t *testing.T) {
	parameters := []struct {
		input    Level
		expected int
	}{
		{All, 0},
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
		testutils.AssertEquals(t, parameters[index].expected, actual)
	}
}

// BenchmarkLogLevel_DigitRepresentation performs benchmarking of the
// Level.DigitRepresentation().
func BenchmarkLogLevel_DigitRepresentation(b *testing.B) {
	level := Debug

	for index := 0; index < b.N; index++ {
		level.DigitRepresentation()
	}
}

// TestLogLevel_Next tests that Level returns next Level.
func TestLogLevel_Next(t *testing.T) {
	parameters := []struct {
		input    Level
		expected Level
	}{
		{All, Trace},
		{Trace, Debug},
		{Debug, Verbose},
		{Verbose, Info},
		{Info, Notice},
		{Notice, Warning},
		{Warning, Severe},
		{Severe, Error},
		{Error, Alert},
		{Alert, Critical},
		{Critical, Emergency},
		{Emergency, Null},
		{Null, Null},
	}

	for index := range parameters {
		actual := parameters[index].input.Next()
		testutils.AssertEquals(t, parameters[index].expected, actual)
	}
}

// BenchmarkLogLevel_Next performs benchmarking of the Level.Next().
func BenchmarkLogLevel_Next(b *testing.B) {
	level := Debug

	for index := 0; index < b.N; index++ {
		level.Next()
	}
}

// TestLogLevel_Previous tests that Level returns previous Level.
func TestLogLevel_Previous(t *testing.T) {
	parameters := []struct {
		input    Level
		expected Level
	}{
		{All, All},
		{Trace, All},
		{Debug, Trace},
		{Verbose, Debug},
		{Info, Verbose},
		{Notice, Info},
		{Warning, Notice},
		{Severe, Warning},
		{Error, Severe},
		{Alert, Error},
		{Critical, Alert},
		{Emergency, Critical},
		{Null, Emergency},
	}

	for index := range parameters {
		actual := parameters[index].input.Previous()
		testutils.AssertEquals(t, parameters[index].expected, actual)
	}
}

// BenchmarkLogLevel_Previous performs benchmarking of the Level.Previous().
func BenchmarkLogLevel_Previous(b *testing.B) {
	level := Debug

	for index := 0; index < b.N; index++ {
		level.Previous()
	}
}
