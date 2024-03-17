// Package formatter_test has tests for formatter package.
package formatter

import (
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"math"
	"testing"
)

const (
	loggerName        = "test"
	loggingLevel      = level.Debug
	pretty            = false
	keyValueDelimiter = "="
	pairSeparator     = ", "
	message           = "Test message."
	static            = "test"
)

var template = map[string]string{
	"level":  "%(level)",
	"name":   "%(name)",
	"static": static,
}

// TestNewJSON tests that NewJSON create correct Formatter instance.
func TestNewJSON(t *testing.T) {
	newFormatter := NewJSON(template, pretty)

	testutils.AssertEquals(t, template, newFormatter.baseFormatter.template)
	testutils.AssertEquals(t, pretty, newFormatter.pretty)
}

// BenchmarkNewJSON performs benchmarking of the NewJSON().
func BenchmarkNewJSON(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NewJSON(template, false)
	}
}

// TestJSONFormatter_Format tests that JSONFormatter.Format correctly formats string.
func TestJSONFormatter_Format(t *testing.T) {
	color := logLevelColors[loggingLevel]
	tests := map[string]struct {
		template map[string]string
		pretty   bool
		colored  bool
		expected string
	}{
		"One Line JSON Not Colored": {
			template: template,
			pretty:   false,
			colored:  false,
			expected: fmt.Sprintf("{\"level\":\"%s\",\"message\":\"%s\",\"name\":\"%s\",\"static\":\"%s\"}\n", loggingLevel.String(), message, loggerName, static),
		},
		"One Line JSON Colored": {
			template: template,
			pretty:   false,
			colored:  true,
			expected: fmt.Sprintf("%s{\"level\":\"%s\",\"message\":\"%s\",\"name\":\"%s\",\"static\":\"%s\"}%s\n", color, loggingLevel.String(), message, loggerName, static, resetColor),
		},
		"Pretty JSON Not Colored": {
			template: template,
			pretty:   true,
			colored:  false,
			expected: fmt.Sprintf("{\n  \"level\": \"%s\",\n  \"message\": \"%s\",\n  \"name\": \"%s\",\n  \"static\": \"%s\"\n}\n", loggingLevel.String(), message, loggerName, static),
		},
		"Pretty JSON Colored": {
			template: template,
			pretty:   true,
			colored:  true,
			expected: fmt.Sprintf("%s{%s\n%s  \"level\": \"%s\",%s\n%s  \"message\": \"%s\",%s\n%s  \"name\": \"%s\",%s\n%s  \"static\": \"%s\"%s\n%s}%s\n", color, resetColor, color, loggingLevel.String(), resetColor, color, message, resetColor, color, loggerName, resetColor, color, static, resetColor, color, resetColor),
		},
	}

	for testName, parameters := range tests {
		t.Run(testName, func(t *testing.T) {
			newFormatter := NewJSON(parameters.template, parameters.pretty)

			testutils.AssertEquals(t, parameters.expected, newFormatter.Format(loggerName, loggingLevel, parameters.colored, "message", message))
		})
	}
}

// TestJSONFormatter_FormatError tests that JSONFormatter.Format returns empty string if error occurred.
func TestJSONFormatter_FormatError(t *testing.T) {
	newFormatter := NewJSON(template, pretty)

	testutils.AssertEquals(t, "", newFormatter.Format(loggerName, loggingLevel, false, "key", math.Inf(1)))
}

// BenchmarkJSONFormatter_Format performs benchmarking of the JSONFormatter.Format().
func BenchmarkJSONFormatter_Format(b *testing.B) {
	benchmarks := map[string]struct {
		template map[string]string
		pretty   bool
		colored  bool
	}{
		"One Line JSON Not Colored": {
			template: template,
			pretty:   false,
			colored:  false,
		},
		"One Line JSON Colored": {
			template: template,
			pretty:   false,
			colored:  true,
		},
		"Pretty JSON Not Colored": {
			template: template,
			pretty:   true,
			colored:  false,
		},
		"Pretty JSON Colored": {
			template: template,
			pretty:   true,
			colored:  true,
		},
	}

	for testName, parameters := range benchmarks {
		b.Run(testName, func(b *testing.B) {
			newFormatter := NewJSON(parameters.template, parameters.pretty)

			for index := 0; index < b.N; index++ {
				newFormatter.Format(loggerName, loggingLevel, parameters.colored, "message", message)
			}
		})
	}
}

// TestJSONFormatter_Template tests that JSONFormatter.Template return assigned template.
func TestJSONFormatter_Template(t *testing.T) {
	newFormatter := NewJSON(template, pretty)

	testutils.AssertEquals(t, template, newFormatter.Template())
}

// BenchmarkJSONFormatter_Template performs benchmarking of the JSONFormatter.Template().
func BenchmarkJSONFormatter_Template(b *testing.B) {
	newFormatter := NewJSON(template, pretty)

	for index := 0; index < b.N; index++ {
		newFormatter.Template()
	}
}

// TestNewKeyValue tests that NewKeyValue create correct Formatter instance.
func TestNewKeyValue(t *testing.T) {
	newFormatter := NewKeyValue(template, keyValueDelimiter, pairSeparator)

	testutils.AssertEquals(t, template, newFormatter.baseFormatter.template)
}

// BenchmarkNewKeyValue performs benchmarking of the NewKeyValue().
func BenchmarkNewKeyValue(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NewKeyValue(template, keyValueDelimiter, pairSeparator)
	}
}

// TestKeyValueFormatter_Format tests that KeyValueFormatter.Format correctly formats string.
func TestKeyValueFormatter_Format(t *testing.T) {
	boolValue := true
	intValue := 1
	int64Value := int64(1)
	float64Value := 1.0
	float32Value := float32(1.0)
	color := logLevelColors[loggingLevel]
	tests := map[string]struct {
		template  map[string]string
		delimiter string
		separator string
		colored   bool
		expected  string
	}{
		"Key Value Not Colored": {
			template:  template,
			delimiter: keyValueDelimiter,
			separator: pairSeparator,
			colored:   false,
			expected:  fmt.Sprintf("bool=%t%sfloat32=%g%sfloat64=%g%sint=%d%sint64=%d%slevel=%q%smessage=%q%sname=%q%sstatic=%q\n", boolValue, pairSeparator, float32Value, pairSeparator, float64Value, pairSeparator, intValue, pairSeparator, int64Value, pairSeparator, loggingLevel.String(), pairSeparator, message, pairSeparator, loggerName, pairSeparator, static),
		},
		"Key Value Colored": {
			template:  template,
			delimiter: keyValueDelimiter,
			separator: pairSeparator,
			colored:   true,
			expected:  fmt.Sprintf("%sbool=%t%sfloat32=%g%sfloat64=%g%sint=%d%sint64=%d%slevel=%q%smessage=%q%sname=%q%sstatic=%q%s\n", color, boolValue, pairSeparator, float32Value, pairSeparator, float64Value, pairSeparator, intValue, pairSeparator, int64Value, pairSeparator, loggingLevel.String(), pairSeparator, message, pairSeparator, loggerName, pairSeparator, static, resetColor),
		},
	}

	for testName, parameters := range tests {
		t.Run(testName, func(t *testing.T) {
			newFormatter := NewKeyValue(parameters.template, parameters.delimiter, parameters.separator)

			testutils.AssertEquals(t, parameters.expected, newFormatter.Format(loggerName, loggingLevel, parameters.colored, "message", message, "bool", boolValue, "int", intValue, "int64", int64Value, "float64", float64Value, "float32", float32Value))
		})
	}
}

// BenchmarkKeyValueFormatter_Format performs benchmarking of the KeyValueFormatter.Format().
func BenchmarkKeyValueFormatter_Format(b *testing.B) {
	benchmarks := map[string]struct {
		template  map[string]string
		delimiter string
		separator string
		colored   bool
	}{
		"Key Value Not Colored": {
			template:  template,
			delimiter: keyValueDelimiter,
			separator: pairSeparator,
			colored:   false,
		},
		"Key Value Colored": {
			template:  template,
			delimiter: keyValueDelimiter,
			separator: pairSeparator,
			colored:   true,
		},
	}

	for testName, parameters := range benchmarks {
		b.Run(testName, func(b *testing.B) {
			newFormatter := NewKeyValue(parameters.template, parameters.delimiter, parameters.separator)

			for index := 0; index < b.N; index++ {
				newFormatter.Format(loggerName, loggingLevel, parameters.colored, "message", message)
			}
		})
	}
}

// TestKeyValueFormatter_Template tests that KeyValueFormatter.Template return assigned template.
func TestKeyValueFormatter_Template(t *testing.T) {
	newFormatter := NewKeyValue(template, keyValueDelimiter, pairSeparator)

	testutils.AssertEquals(t, template, newFormatter.Template())
}

// BenchmarkKeyValueFormatter_Template performs benchmarking of the KeyValueFormatter.Template().
func BenchmarkKeyValueFormatter_Template(b *testing.B) {
	newFormatter := NewKeyValue(template, keyValueDelimiter, pairSeparator)

	for index := 0; index < b.N; index++ {
		newFormatter.Template()
	}
}
