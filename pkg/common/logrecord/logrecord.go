// Package logrecord contains LogRecord struct that represents a log record.
package logrecord

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	"runtime"
	"time"
)

// LogRecord struct represents a log record.
type LogRecord struct {
	// Name of the logger.
	name string
	// Time format of the log record.
	timeFormat string
	// Time of the log record.
	timestamp time.Time
	// Level of the log record.
	level level.Level
	// File name of the log record.
	fileName string
	// Line number of the log record.
	fileLine int
}

// New creates a new instance of the LogRecord.
func New(name string, level level.Level, timeFormat string, skipCaller int) *LogRecord {
	_, fileName, fileLine, _ := runtime.Caller(skipCaller)
	if timeFormat == "" {
		timeFormat = time.RFC3339
	}
	return &LogRecord{
		name:       name,
		timeFormat: timeFormat,
		timestamp:  time.Now(),
		level:      level,
		fileName:   fileName,
		fileLine:   fileLine,
	}
}

// Name returns the name of the log record.
func (record *LogRecord) Name() string {
	return record.name
}

// Time returns the time of the log record formatted according to the time format.
func (record *LogRecord) Time() string {
	return record.timestamp.Format(record.timeFormat)
}

// Timestamp returns the time of the log record as a Unix timestamp.
func (record *LogRecord) Timestamp() int64 {
	return record.timestamp.Unix()
}

// Level returns the level of the log record.
func (record *LogRecord) Level() level.Level {
	return record.level
}

// FileName returns the file name of the log record.
func (record *LogRecord) FileName() string {
	return record.fileName
}

// FileLine returns the line number of the log record.
func (record *LogRecord) FileLine() int {
	return record.fileLine
}
