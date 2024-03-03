package formatter

import (
	"fmt"
	"github.com/dl1998/go-logging/pkg/logger/loglevel"
	"testing"
)

var format = "%(level):%(name):%(message)"
var message = "Test message."
var loggerName = "test"
var loggingLevel = loglevel.LogLevel(loglevel.Debug)

// TestNew tests that New create correct Formatter instance.
func TestNew(t *testing.T) {
	newFormatter := New(format)

	if newFormatter.format != format {
		t.Fatalf("expected: %s, actual: %s", format, newFormatter.format)
	}
}

// BenchmarkNew performs benchmarking of the New().
func BenchmarkNew(b *testing.B) {
	for index := 0; index < b.N; index++ {
		New(format)
	}
}

// TestFormatter_EvaluatePreset tests that Formatter.EvaluatePreset correctly
// evaluates tags.
func TestFormatter_EvaluatePreset(t *testing.T) {
	newFormatter := New(format)

	preset := newFormatter.EvaluatePreset(message, loggerName, loggingLevel)

	if preset["%(message)"] != message {
		t.Fatalf("%%(message) is incorrect. expected: %s, actual: %s", format, newFormatter.format)
	}

	if preset["%(name)"] != loggerName {
		t.Fatalf("%%(name) is incorrect. expected: %s, actual: %s", format, newFormatter.format)
	}

	if preset["%(level)"] != loggingLevel.String() {
		t.Fatalf("%%(level) is incorrect. expected: %s, actual: %s", format, newFormatter.format)
	}
}

// BenchmarkFormatter_EvaluatePreset performs benchmarking of the Formatter.EvaluatePreset().
func BenchmarkFormatter_EvaluatePreset(b *testing.B) {
	newFormatter := New(format)

	for index := 0; index < b.N; index++ {
		newFormatter.EvaluatePreset(message, loggerName, loggingLevel)
	}
}

// TestFormatter_Format tests that Formatter.Format correctly formats string.
func TestFormatter_Format(t *testing.T) {
	newFormatter := New(format)

	parameters := []struct {
		colored  bool
		expected string
	}{
		{false, fmt.Sprintf("%s:%s:%s\n", loggingLevel.String(), loggerName, message)},
		{true, fmt.Sprintf("\033[36m%s:%s:%s\033[0m\n", loggingLevel.String(), loggerName, message)},
	}

	for index := range parameters {
		actual := newFormatter.Format(message, loggerName, loggingLevel, parameters[index].colored)

		if actual != parameters[index].expected {
			t.Fatalf("expected: %s, actual: %s", parameters[index].expected, actual)
		}
	}
}

// BenchmarkFormatter_Format performs benchmarking of the Formatter.Format().
func BenchmarkFormatter_Format(b *testing.B) {
	newFormatter := New(format)

	for index := 0; index < b.N; index++ {
		newFormatter.Format(message, loggerName, loggingLevel, true)
	}
}
