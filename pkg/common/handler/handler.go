// Package handler provides common handlers for the logger, it contains logic for
// the logging messages.
package handler

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	"io"
	"os"
	"strings"
)

// Interface represents interface that shall be satisfied by Handler.
type Interface interface {
	Writer() io.Writer
	SetWriter(writer io.Writer)
	FromLevel() level.Level
	SetFromLevel(fromLevel level.Level)
	ToLevel() level.Level
	SetToLevel(toLevel level.Level)
}

// Handler struct contains information where it shall write log message, how to
// format them and their log fromLevel.
type Handler struct {
	fromLevel                 level.Level
	toLevel                   level.Level
	writer                    io.Writer
	ConsoleSupportsANSIColors func() bool
}

// New create a new instance of the Handler.
func New(fromLevel level.Level, toLevel level.Level, writer io.Writer) *Handler {
	return &Handler{
		fromLevel:                 fromLevel,
		toLevel:                   toLevel,
		writer:                    writer,
		ConsoleSupportsANSIColors: consoleSupportsANSIColors,
	}
}

// Writer returns writer of the Handler.
func (handler *Handler) Writer() io.Writer {
	return handler.writer
}

// SetWriter sets a new writer for the Handler.
func (handler *Handler) SetWriter(writer io.Writer) {
	handler.writer = writer
}

// FromLevel returns log fromLevel of the Handler.
func (handler *Handler) FromLevel() level.Level {
	return handler.fromLevel
}

// SetFromLevel sets a new log fromLevel for the Handler.
func (handler *Handler) SetFromLevel(fromLevel level.Level) {
	handler.fromLevel = fromLevel
}

// ToLevel returns log toLevel of the Handler.
func (handler *Handler) ToLevel() level.Level {
	return handler.toLevel
}

// SetToLevel sets a new log toLevel for the Handler.
func (handler *Handler) SetToLevel(toLevel level.Level) {
	handler.toLevel = toLevel
}

// consoleSupportsANSIColors returns true, if current terminal supports ANSI
// colors, otherwise returns False.
func consoleSupportsANSIColors() bool {
	term := os.Getenv("TERM")
	return strings.Contains(term, "xterm") || strings.Contains(term, "color")
}
