package logger

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"github.com/dl1998/go-logging/pkg/logger/logrecord"
)

// baseLoggerInterface defines low level logging interface.
type baseLoggerInterface interface {
	Log(level level.Level, skipCaller int, message string, parameters ...any)
	Name() string
	SetName(name string)
	Handlers() []handler.Interface
	AddHandler(handlerInterface handler.Interface)
	RemoveHandler(handlerInterface handler.Interface)
}

// baseLogger struct contains basic fields for the logger.
type baseLogger struct {
	name       string
	timeFormat string
	handlers   []handler.Interface
}

// Log logs interpolated message with the provided level.Level.
func (logger *baseLogger) Log(level level.Level, skipCaller int, message string, parameters ...any) {
	record := logrecord.New(logger.name, level, logger.timeFormat, message, parameters, skipCaller)
	for _, registeredHandler := range logger.handlers {
		registeredHandler.Write(record)
	}
}

// Name return baseLogger name.
func (logger *baseLogger) Name() string {
	return logger.name
}

// SetName sets a new name for the baseLogger.
func (logger *baseLogger) SetName(name string) {
	logger.name = name
}

// Handlers returns a list of the registered handler.Interface objects for the
// baseLogger.
func (logger *baseLogger) Handlers() []handler.Interface {
	return logger.handlers
}

// AddHandler register a new handler.Interface for the baseLogger.
func (logger *baseLogger) AddHandler(handlerInterface handler.Interface) {
	logger.handlers = append(logger.handlers, handlerInterface)
}

// RemoveHandler removes a handler.Interface from the baseLogger handlers.
func (logger *baseLogger) RemoveHandler(handlerInterface handler.Interface) {
	newSlice := make([]handler.Interface, 0)
	for _, element := range logger.handlers {
		if element != handlerInterface {
			newSlice = append(newSlice, element)
		}
	}
	logger.handlers = newSlice
}
