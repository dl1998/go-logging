// Package formatter_test has tests for formatter package.
package formatter

import (
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger/logrecord"
	"strconv"
	"strings"
	"testing"
)

var (
	template        = "%(level):%(name):%(message)"
	message         = "Test message."
	loggerName      = "test"
	loggingLevel    = level.Debug
	timeFormat      = ""
	emptyParameters = make([]any, 0)
	skipCaller      = 1
)

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

	tests := map[string]struct {
		colored  bool
		expected string
	}{
		"not colored": {false, fmt.Sprintf("%s:%s:%s\n", loggingLevel.String(), loggerName, message)},
		"colored":     {true, fmt.Sprintf("\033[36m%s:%s:%s\033[0m\n", loggingLevel.String(), loggerName, message)},
	}

	for name, parameters := range tests {
		record := logrecord.New(loggerName, loggingLevel, timeFormat, message, emptyParameters, skipCaller)

		t.Run(name, func(t *testing.T) {
			actual := newFormatter.Format(record, parameters.colored)

			testutils.AssertEquals(t, parameters.expected, actual)
		})
	}
}

// BenchmarkFormatter_Format performs benchmarking of the Formatter.Format().
func BenchmarkFormatter_Format(b *testing.B) {
	newFormatter := New(template)

	record := logrecord.New(loggerName, loggingLevel, timeFormat, message, emptyParameters, skipCaller)

	b.ResetTimer()

	for index := 0; index < b.N; index++ {
		newFormatter.Format(record, true)
	}
}

// TestParseTemplate tests that ParseTemplate correctly replaces keys with values.
func TestParseTemplate(t *testing.T) {
	format := ParseTemplate(template, logrecord.New(loggerName, loggingLevel, timeFormat, message, emptyParameters, skipCaller))

	expected := fmt.Sprintf("%s:%s:%s", loggingLevel.String(), loggerName, message)

	testutils.AssertEquals(t, expected, format)
}

// BenchmarkParseTemplate performs benchmarking of the ParseTemplate().
func BenchmarkParseTemplate(b *testing.B) {
	for index := 0; index < b.N; index++ {
		ParseTemplate(template, logrecord.New(loggerName, loggingLevel, timeFormat, message, emptyParameters, skipCaller))
	}
}

// TestReplaceKey tests that ReplaceKey correctly replaces key with value.
func TestReplaceKey(t *testing.T) {
	record := logrecord.New(loggerName, loggingLevel, timeFormat, message, emptyParameters, skipCaller)

	tests := map[string]struct {
		key      string
		expected string
	}{
		"Name":      {key: "%(name)", expected: strings.ReplaceAll(template, "%(name)", loggerName)},
		"Not a key": {key: "not a key", expected: template},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := ReplaceKey(template, test.key, record)

			testutils.AssertEquals(t, test.expected, value)
		})
	}
}

// BenchmarkReplaceKey performs benchmarking of the ReplaceKey().
func BenchmarkReplaceKey(b *testing.B) {
	record := logrecord.New(loggerName, loggingLevel, timeFormat, message, emptyParameters, skipCaller)

	benchmarks := map[string]struct {
		key string
	}{
		"Name":      {key: "%(name)"},
		"Not a key": {key: "not a key"},
	}

	for name, benchmark := range benchmarks {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for index := 0; index < b.N; index++ {
				ReplaceKey(template, benchmark.key, record)
			}
		})
	}
}

// TestParseKey tests that ParseKey returns correct value for the key.
func TestParseKey(t *testing.T) {
	record := logrecord.New(loggerName, loggingLevel, timeFormat, message, emptyParameters, skipCaller)

	tests := map[string]struct {
		key      string
		expected string
	}{
		"Name":          {key: "%(name)", expected: loggerName},
		"Level name":    {key: "%(level)", expected: loggingLevel.String()},
		"Level number":  {key: "%(levelnr)", expected: strconv.Itoa(loggingLevel.DigitRepresentation())},
		"Date time":     {key: "%(datetime)", expected: record.Time()},
		"Timestamp":     {key: "%(timestamp)", expected: strconv.FormatInt(record.Timestamp(), 10)},
		"Function name": {key: "%(fname)", expected: record.FileName()},
		"Function line": {key: "%(fline)", expected: strconv.Itoa(record.FileLine())},
		"Not a key":     {key: "not a key", expected: "not a key"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			value := ParseKey(test.key, record)

			testutils.AssertEquals(t, test.expected, value)
		})
	}
}

// BenchmarkParseKey performs benchmarking of the ParseKey().
func BenchmarkParseKey(b *testing.B) {
	record := logrecord.New(loggerName, loggingLevel, timeFormat, message, emptyParameters, skipCaller)

	benchmarks := map[string]struct {
		key string
	}{
		"Name":          {key: "%(name)"},
		"Level name":    {key: "%(level)"},
		"Level number":  {key: "%(levelnr)"},
		"Date time":     {key: "%(datetime)"},
		"Timestamp":     {key: "%(timestamp)"},
		"Function name": {key: "%(fname)"},
		"Function line": {key: "%(fline)"},
		"Not a key":     {key: "not a key"},
	}

	for name, benchmark := range benchmarks {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for index := 0; index < b.N; index++ {
				ParseKey(benchmark.key, record)
			}
		})
	}
}
