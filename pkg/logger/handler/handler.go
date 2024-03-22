// Package handler provides handlers for the logger, it contains logic for the
// logging messages.
package handler

import (
	"fmt"
	"github.com/dl1998/go-logging/pkg/common/handler"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/logrecord"
	"io"
	"os"
)

var osOpenFile = os.OpenFile
var osStdout = os.Stdout
var osStderr = os.Stderr

// Interface represents interface that shall be satisfied by Handler.
type Interface interface {
	Writer() io.Writer
	FromLevel() level.Level
	SetFromLevel(fromLevel level.Level)
	ToLevel() level.Level
	SetToLevel(toLevel level.Level)
	Formatter() formatter.Interface
	Write(record logrecord.Interface)
}

// Handler struct contains information where it shall write log message, how to
// format them and their log fromLevel.
type Handler struct {
	*handler.Handler
	formatter formatter.Interface
}

// New create a new instance of the Handler.
func New(fromLevel level.Level, toLevel level.Level, newFormatter formatter.Interface, writer io.Writer) *Handler {
	return &Handler{
		Handler:   handler.New(fromLevel, toLevel, writer),
		formatter: newFormatter,
	}
}

// NewConsoleHandler create a new instance of the Handler that writes log
// messages to the os.Stdout.
func NewConsoleHandler(fromLevel level.Level, toLevel level.Level, newFormatter formatter.Interface) *Handler {
	return New(fromLevel, toLevel, newFormatter, osStdout)
}

// NewConsoleErrorHandler create a new instance of the Handler that writes log
// messages to the os.Stderr.
func NewConsoleErrorHandler(fromLevel level.Level, toLevel level.Level, newFormatter formatter.Interface) *Handler {
	return New(fromLevel, toLevel, newFormatter, osStderr)
}

// NewFileHandler creates a new instance of the Handler that writes log message
// to the log file.
func NewFileHandler(fromLevel level.Level, toLevel level.Level, newFormatter formatter.Interface, file string) *Handler {
	writer, err := osOpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return New(fromLevel, toLevel, newFormatter, writer)
}

// Formatter returns formatter of the Handler.
func (handler *Handler) Formatter() formatter.Interface {
	return handler.formatter
}

// Write writes log message to the defined by the Handler writer.
func (handler *Handler) Write(record logrecord.Interface) {
	if record.Level().DigitRepresentation() < handler.FromLevel().DigitRepresentation() || record.Level().DigitRepresentation() > handler.ToLevel().DigitRepresentation() {
		return
	}

	var colored = false

	if handler.ConsoleSupportsANSIColors() && (handler.Writer() == osStdout || handler.Writer() == osStderr) {
		colored = true
	}

	log := handler.formatter.Format(record, colored)

	if _, err := handler.Writer().Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}
