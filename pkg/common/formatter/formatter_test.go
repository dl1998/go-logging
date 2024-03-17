// Package formatter_test has tests for formatter package.
package formatter

import (
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"testing"
)

var loggerName = "test"
var loggingLevel = level.Debug

// TestFormatter_EvaluatePreset tests that Formatter.EvaluatePreset correctly
// evaluates tags.
func TestEvaluatePreset(t *testing.T) {
	preset := EvaluatePreset(loggerName, loggingLevel, 1)

	testutils.AssertEquals(t, loggerName, preset["%(name)"].(string))
	testutils.AssertEquals(t, loggingLevel.String(), preset["%(level)"].(string))
}

// BenchmarkFormatter_EvaluatePreset performs benchmarking of the Formatter.EvaluatePreset().
func BenchmarkEvaluatePreset(b *testing.B) {
	for index := 0; index < b.N; index++ {
		EvaluatePreset(loggerName, loggingLevel, 1)
	}
}
