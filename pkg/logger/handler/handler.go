// Package handler provides handlers for the logger, it contains logic for the
// logging messages.
package handler

import (
	"fmt"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/loglevel"
	"io"
	"os"
	"strings"
)

var osOpenFile = os.OpenFile
var osStdout = os.Stdout
var osStderr = os.Stderr

// Interface represents interface that shall be satisfied by Handler.
type Interface interface {
	Level() loglevel.LogLevel
	SetLevel(level loglevel.LogLevel)
	Formatter() formatter.Interface
	Write(logName string, level loglevel.LogLevel, message string, parameters ...any)
}

// Handler struct contains information where it shall write log message, how to
// format them and their log level.
type Handler struct {
	level       loglevel.LogLevel
	formatter   formatter.Interface
	writer      io.Writer
	errorWriter io.Writer
}

// New create a new instance of the Handler.
func New(level loglevel.LogLevel, newFormatter formatter.Interface, writer io.Writer, errorWriter io.Writer) *Handler {
	return &Handler{
		level:       level,
		formatter:   newFormatter,
		writer:      writer,
		errorWriter: errorWriter,
	}
}

// NewConsoleHandler create a new instance of the Handler that writes log
// messages to the os.Stdout and os.Stderr respectively.
func NewConsoleHandler(level loglevel.LogLevel, newFormatter formatter.Interface) *Handler {
	return New(level, newFormatter, osStdout, osStderr)
}

// NewFileHandler creates a new instance of the Handler that writes log message
// to the log file.
func NewFileHandler(level loglevel.LogLevel, newFormatter formatter.Interface, file string) *Handler {
	writer, err := osOpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	return New(level, newFormatter, writer, writer)
}

// Level returns log level of the Handler.
func (handler *Handler) Level() loglevel.LogLevel {
	return handler.level
}

// SetLevel sets a new log level for the Handler.
func (handler *Handler) SetLevel(level loglevel.LogLevel) {
	handler.level = level
}

// Formatter returns formatter.Interface used by the Handler.
func (handler *Handler) Formatter() formatter.Interface {
	return handler.formatter
}

// Write writes log message to the defined by the Handler writer.
func (handler *Handler) Write(logName string, level loglevel.LogLevel, message string, parameters ...any) {
	formattedMessage := fmt.Sprintf(message, parameters...)

	writer := handler.writer

	if level >= loglevel.Error {
		writer = handler.errorWriter
	}

	var colored = false

	if consoleSupportsANSIColors() && (writer == osStdout || writer == osStderr) {
		colored = true
	}

	log := handler.formatter.Format(formattedMessage, logName, level, colored)

	if _, err := writer.Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}

// consoleSupportsANSIColors returns true, if current terminal supports ANSI
// colors, otherwise returns False.
func consoleSupportsANSIColors() bool {
	term := os.Getenv("TERM")
	return strings.Contains(term, "xterm") || strings.Contains(term, "color")
}
