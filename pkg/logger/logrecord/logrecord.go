// Package logrecord provides a structured log record implementation.
package logrecord

import (
	"fmt"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/common/logrecord"
)

// Interface represents interface that shall be satisfied by structured LogRecord.
type Interface interface {
	Name() string
	Time() string
	Timestamp() int64
	Level() level.Level
	FileName() string
	FileLine() int
	Message() string
}

// LogRecord struct represents a structured log record.
type LogRecord struct {
	// LogRecord is the base log record.
	*logrecord.LogRecord
	// message is the message of the log record.
	message string
}

// New creates a new instance of the structured LogRecord.
func New(name string, level level.Level, timeFormat string, message string, parameters []any, skipCaller int) *LogRecord {
	return &LogRecord{
		LogRecord: logrecord.New(name, level, timeFormat, skipCaller),
		message:   fmt.Sprintf(message, parameters...),
	}
}

// Message returns the message of the log record.
func (record *LogRecord) Message() string {
	return record.message
}
