package handler

import (
	"fmt"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/loglevel"
	"io"
	"os"
	"strings"
)

type Handler struct {
	level       loglevel.LogLevel
	formatter   formatter.Formatter
	writer      io.Writer
	errorWriter io.Writer
}

func New(level loglevel.LogLevel, newFormatter formatter.Formatter, writer io.Writer, errorWriter io.Writer) *Handler {
	return &Handler{
		level:       level,
		formatter:   newFormatter,
		writer:      writer,
		errorWriter: errorWriter,
	}
}

func NewConsoleHandler(level loglevel.LogLevel, newFormatter formatter.Formatter) *Handler {
	writer := os.Stdout
	errorWriter := os.Stderr

	return New(level, newFormatter, writer, errorWriter)
}

func NewFileHandler(level loglevel.LogLevel, newFormatter formatter.Formatter, file string) *Handler {
	writer, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	return New(level, newFormatter, writer, writer)
}

func (handler *Handler) Level() loglevel.LogLevel {
	return handler.level
}

func (handler *Handler) SetLevel(level loglevel.LogLevel) {
	handler.level = level
}

func (handler *Handler) Write(logName string, level loglevel.LogLevel, message string, parameters ...any) {
	formattedMessage := fmt.Sprintf(message, parameters...)

	var colored = false

	if consoleSupportsANSIColors() && handler.writer == os.Stdout {
		colored = true
	}

	log := handler.formatter.Format(formattedMessage, logName, level, colored)

	writer := handler.writer

	if level >= loglevel.Error {
		writer = handler.errorWriter
	}

	if _, err := writer.Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}

func consoleSupportsANSIColors() bool {
	term := os.Getenv("TERM")
	return strings.Contains(term, "xterm") || strings.Contains(term, "color")
}
