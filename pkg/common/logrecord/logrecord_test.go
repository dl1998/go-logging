// Package logrecord contains tests for the LogRecord.
package logrecord

import (
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"runtime"
	"testing"
	"time"
)

var (
	skipCaller = 1
	logLevel   = level.Debug
	name       = "test"
	timeFormat = ""
)

// TestNew tests that New function returns a new LogRecord struct instance.
func TestNew(t *testing.T) {
	_, file, line, _ := runtime.Caller(0)
	record := New(name, logLevel, timeFormat, skipCaller)

	testutils.AssertEquals(t, name, record.name)
	testutils.AssertEquals(t, logLevel, record.level)
	testutils.AssertEquals(t, file, record.fileName)
	testutils.AssertEquals(t, line+1, record.fileLine)
	testutils.AssertEquals(t, time.RFC3339, record.timeFormat)
	testutils.AssertNotNil(t, record.timestamp)
}

// BenchmarkNew benchmarks the New function.
func BenchmarkNew(b *testing.B) {
	for index := 0; index < b.N; index++ {
		New(name, logLevel, "", skipCaller)
	}
}

// TestName tests that Name function returns the name of the log record.
func TestName(t *testing.T) {
	record := New(name, logLevel, "", skipCaller)

	testutils.AssertEquals(t, name, record.Name())
}

// BenchmarkName benchmarks the Name function.
func BenchmarkName(b *testing.B) {
	record := New(name, logLevel, "", skipCaller)
	for index := 0; index < b.N; index++ {
		record.Name()
	}
}

// TestTime tests that Time function returns the time of the log record.
func TestTime(t *testing.T) {
	record := New(name, logLevel, "", skipCaller)

	testutils.AssertEquals(t, record.timestamp.Format(record.timeFormat), record.Time())
}

// BenchmarkTime benchmarks the Time function.
func BenchmarkTime(b *testing.B) {
	record := New(name, logLevel, "", skipCaller)
	for index := 0; index < b.N; index++ {
		record.Time()
	}
}

// TestTimestamp tests that Timestamp function returns the timestamp of the log record.
func TestTimestamp(t *testing.T) {
	record := New(name, logLevel, "", skipCaller)

	testutils.AssertEquals(t, record.timestamp.Unix(), record.Timestamp())
}

// BenchmarkTimestamp benchmarks the Timestamp function.
func BenchmarkTimestamp(b *testing.B) {
	record := New(name, logLevel, "", skipCaller)
	for index := 0; index < b.N; index++ {
		record.Timestamp()
	}
}

// TestLevel tests that Level function returns the level of the log record.
func TestLevel(t *testing.T) {
	record := New(name, logLevel, "", skipCaller)

	testutils.AssertEquals(t, logLevel, record.Level())
}

// BenchmarkLevel benchmarks the Level function.
func BenchmarkLevel(b *testing.B) {
	record := New(name, logLevel, "", skipCaller)
	for index := 0; index < b.N; index++ {
		record.Level()
	}
}

// TestFileName tests that FileName function returns the file name of the log record.
func TestFileName(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	record := New(name, logLevel, "", skipCaller)

	testutils.AssertEquals(t, file, record.FileName())
}

// BenchmarkFileName benchmarks the FileName function.
func BenchmarkFileName(b *testing.B) {
	record := New(name, logLevel, "", skipCaller)
	for index := 0; index < b.N; index++ {
		record.FileName()
	}
}

// TestFileLine tests that FileLine function returns the file line of the log record.
func TestFileLine(t *testing.T) {
	_, _, line, _ := runtime.Caller(0)
	record := New(name, logLevel, "", skipCaller)

	testutils.AssertEquals(t, line+1, record.FileLine())
}

// BenchmarkFileLine benchmarks the FileLine function.
func BenchmarkFileLine(b *testing.B) {
	record := New(name, logLevel, "", skipCaller)
	for index := 0; index < b.N; index++ {
		record.FileLine()
	}
}
