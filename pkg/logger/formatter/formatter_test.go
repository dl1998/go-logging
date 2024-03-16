// Package formatter_test has tests for formatter package.
package formatter

import (
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"testing"
)

var template = "%(level):%(name):%(message)"
var message = "Test message."
var loggerName = "test"
var loggingLevel = level.Debug

// TestNew tests that New create correct Formatter instance.
func TestNew(t *testing.T) {
	newFormatter := New(template)

	testutils.AssertEquals(t, template, newFormatter.template)
}

// BenchmarkNew performs benchmarking of the New().
func BenchmarkNew(b *testing.B) {
	for index := 0; index < b.N; index++ {
		New(template)
	}
}

// TestFormatter_IsEqual tests that Formatter.IsEqual returns true, if two
// Formatter(s) are the same.
func TestFormatter_IsEqual(t *testing.T) {
	newFormatter := New(template)

	isEqual := newFormatter.IsEqual(newFormatter)

	if !isEqual {
		t.Fatalf("expected: %t, actual: %t", true, isEqual)
	}
}

// BenchmarkFormatter_IsEqual performs benchmarking of the Formatter.IsEqual().
func BenchmarkFormatter_IsEqual(b *testing.B) {
	newFormatter := New(template)

	for index := 0; index < b.N; index++ {
		newFormatter.IsEqual(newFormatter)
	}
}

// TestFormatter_EvaluatePreset tests that Formatter.EvaluatePreset correctly
// evaluates tags.
func TestFormatter_EvaluatePreset(t *testing.T) {
	newFormatter := New(template)

	preset := newFormatter.EvaluatePreset(message, loggerName, loggingLevel)

	testutils.AssertEquals(t, message, preset["%(message)"])
	testutils.AssertEquals(t, loggerName, preset["%(name)"])
	testutils.AssertEquals(t, loggingLevel.String(), preset["%(level)"])
}

// BenchmarkFormatter_EvaluatePreset performs benchmarking of the Formatter.EvaluatePreset().
func BenchmarkFormatter_EvaluatePreset(b *testing.B) {
	newFormatter := New(template)

	for index := 0; index < b.N; index++ {
		newFormatter.EvaluatePreset(message, loggerName, loggingLevel)
	}
}

// TestFormatter_Template tests that Formatter.Template return assigned template.
func TestFormatter_Template(t *testing.T) {
	newFormatter := New(template)

	testutils.AssertEquals(t, template, newFormatter.Template())
}

// BenchmarkFormatter_Template performs benchmarking of the Formatter.Template().
func BenchmarkFormatter_Template(b *testing.B) {
	newFormatter := New(template)

	for index := 0; index < b.N; index++ {
		newFormatter.Template()
	}
}

// TestFormatter_Format tests that Formatter.Format correctly formats string.
func TestFormatter_Format(t *testing.T) {
	newFormatter := New(template)

	parameters := []struct {
		colored  bool
		expected string
	}{
		{false, fmt.Sprintf("%s:%s:%s\n", loggingLevel.String(), loggerName, message)},
		{true, fmt.Sprintf("\033[36m%s:%s:%s\033[0m\n", loggingLevel.String(), loggerName, message)},
	}

	for index := range parameters {
		actual := newFormatter.Format(message, loggerName, loggingLevel, parameters[index].colored)

		testutils.AssertEquals(t, parameters[index].expected, actual)
	}
}

// BenchmarkFormatter_Format performs benchmarking of the Formatter.Format().
func BenchmarkFormatter_Format(b *testing.B) {
	newFormatter := New(template)

	for index := 0; index < b.N; index++ {
		newFormatter.Format(message, loggerName, loggingLevel, true)
	}
}
