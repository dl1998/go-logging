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
	Writer() io.Writer
	FromLevel() loglevel.LogLevel
	SetFromLevel(fromLevel loglevel.LogLevel)
	ToLevel() loglevel.LogLevel
	SetToLevel(toLevel loglevel.LogLevel)
	Formatter() formatter.Interface
	Write(logName string, level loglevel.LogLevel, message string, parameters ...any)
}

// Handler struct contains information where it shall write log message, how to
// format them and their log fromLevel.
type Handler struct {
	fromLevel                 loglevel.LogLevel
	toLevel                   loglevel.LogLevel
	formatter                 formatter.Interface
	writer                    io.Writer
	consoleSupportsANSIColors func() bool
}

// New create a new instance of the Handler.
func New(fromLevel loglevel.LogLevel, toLevel loglevel.LogLevel, newFormatter formatter.Interface, writer io.Writer) *Handler {
	return &Handler{
		fromLevel:                 fromLevel,
		toLevel:                   toLevel,
		formatter:                 newFormatter,
		writer:                    writer,
		consoleSupportsANSIColors: consoleSupportsANSIColors,
	}
}

// NewConsoleHandler create a new instance of the Handler that writes log
// messages to the os.Stdout.
func NewConsoleHandler(fromLevel loglevel.LogLevel, toLevel loglevel.LogLevel, newFormatter formatter.Interface) *Handler {
	return New(fromLevel, toLevel, newFormatter, osStdout)
}

// NewConsoleErrorHandler create a new instance of the Handler that writes log
// messages to the os.Stderr.
func NewConsoleErrorHandler(fromLevel loglevel.LogLevel, toLevel loglevel.LogLevel, newFormatter formatter.Interface) *Handler {
	return New(fromLevel, toLevel, newFormatter, osStderr)
}

// NewFileHandler creates a new instance of the Handler that writes log message
// to the log file.
func NewFileHandler(fromLevel loglevel.LogLevel, toLevel loglevel.LogLevel, newFormatter formatter.Interface, file string) *Handler {
	writer, err := osOpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return New(fromLevel, toLevel, newFormatter, writer)
}

// Writer returns writer of the Handler.
func (handler *Handler) Writer() io.Writer {
	return handler.writer
}

// FromLevel returns log fromLevel of the Handler.
func (handler *Handler) FromLevel() loglevel.LogLevel {
	return handler.fromLevel
}

// SetFromLevel sets a new log fromLevel for the Handler.
func (handler *Handler) SetFromLevel(fromLevel loglevel.LogLevel) {
	handler.fromLevel = fromLevel
}

// ToLevel returns log toLevel of the Handler.
func (handler *Handler) ToLevel() loglevel.LogLevel {
	return handler.toLevel
}

// SetToLevel sets a new log toLevel for the Handler.
func (handler *Handler) SetToLevel(toLevel loglevel.LogLevel) {
	handler.toLevel = toLevel
}

// Formatter returns formatter.Interface used by the Handler.
func (handler *Handler) Formatter() formatter.Interface {
	return handler.formatter
}

// Write writes log message to the defined by the Handler writer.
func (handler *Handler) Write(logName string, level loglevel.LogLevel, message string, parameters ...any) {
	if level.DigitRepresentation() < handler.fromLevel.DigitRepresentation() || level.DigitRepresentation() > handler.toLevel.DigitRepresentation() {
		return
	}

	formattedMessage := fmt.Sprintf(message, parameters...)

	var colored = false

	if handler.consoleSupportsANSIColors() && (handler.writer == osStdout || handler.writer == osStderr) {
		colored = true
	}

	log := handler.formatter.Format(formattedMessage, logName, level, colored)

	if _, err := handler.writer.Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}

// consoleSupportsANSIColors returns true, if current terminal supports ANSI
// colors, otherwise returns False.
func consoleSupportsANSIColors() bool {
	term := os.Getenv("TERM")
	return strings.Contains(term, "xterm") || strings.Contains(term, "color")
}
