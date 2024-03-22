// Package formatter_test has tests for formatter package.
package formatter

import (
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/common/logrecord"
	"testing"
)

var (
	loggerName   = "test"
	loggingLevel = level.Debug
	timeFormat   = ""
	skipCaller   = 1
)

// TestParseKey tests that ParseKey returns correct value for the key.
func TestParseKey(t *testing.T) {
	record := logrecord.New(loggerName, loggingLevel, timeFormat, skipCaller)

	tests := map[string]struct {
		key      string
		expected interface{}
	}{
		"Name":          {key: "%(name)", expected: loggerName},
		"Level name":    {key: "%(level)", expected: loggingLevel.String()},
		"Level number":  {key: "%(levelnr)", expected: loggingLevel.DigitRepresentation()},
		"Date time":     {key: "%(datetime)", expected: record.Time()},
		"Timestamp":     {key: "%(timestamp)", expected: record.Timestamp()},
		"Function name": {key: "%(fname)", expected: record.FileName()},
		"Function line": {key: "%(fline)", expected: record.FileLine()},
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
	record := logrecord.New(loggerName, loggingLevel, timeFormat, skipCaller)

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
