// Package logrecord provides a structured log record implementation.
package logrecord

import (
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
	Parameters() map[string]interface{}
}

// LogRecord struct represents a structured log record.
type LogRecord struct {
	// LogRecord is the base log record.
	*logrecord.LogRecord
	// parameters are map of the parameters of the log record.
	parameters map[string]interface{}
}

// New creates a new instance of the structured LogRecord.
func New(name string, level level.Level, timeFormat string, parameters map[string]interface{}, skipCaller int) *LogRecord {
	return &LogRecord{
		LogRecord:  logrecord.New(name, level, timeFormat, skipCaller),
		parameters: parameters,
	}
}

// Parameters returns the parameters of the log record.
func (record *LogRecord) Parameters() map[string]interface{} {
	return record.parameters
}
