// Package logger contains logger implementation.
package logger

import (
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"github.com/dl1998/go-logging/pkg/logger/loglevel"
)

var rootLogger = GetDefaultLogger()

// baseLoggerInterface defines low level logging interface.
type baseLoggerInterface interface {
	Log(level loglevel.LogLevel, message string, parameters ...any)
	Name() string
	SetName(name string)
	Handlers() []handler.Interface
	AddHandler(handler handler.Interface)
}

// baseLogger struct contains basic fields for the logger.
type baseLogger struct {
	name     string
	handlers []handler.Interface
}

// Log logs interpolated message with the provided loglevel.LogLevel.
func (logger *baseLogger) Log(level loglevel.LogLevel, message string, parameters ...any) {
	for _, registeredHandler := range logger.handlers {
		if level >= registeredHandler.Level() {
			registeredHandler.Write(logger.name, level, message, parameters...)
		}
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
func (logger *baseLogger) AddHandler(handler handler.Interface) {
	logger.handlers = append(logger.handlers, handler)
}

// Interface represents interface that shall be satisfied by Logger.
type Interface interface {
	Name() string
	Handlers() []handler.Interface
	AddHandler(handler handler.Interface)
	Trace(message string, parameters ...any)
	Debug(message string, parameters ...any)
	Verbose(message string, parameters ...any)
	Info(message string, parameters ...any)
	Notice(message string, parameters ...any)
	Warning(message string, parameters ...any)
	Severe(message string, parameters ...any)
	Error(message string, parameters ...any)
	Alert(message string, parameters ...any)
	Critical(message string, parameters ...any)
	Emergency(message string, parameters ...any)
}

// Logger struct encapsulates baseLogger implementation.
type Logger struct {
	baseLogger baseLoggerInterface
}

// New creates a new instance of the Logger.
func New(name string) *Logger {
	return &Logger{
		baseLogger: &baseLogger{
			name:     name,
			handlers: make([]handler.Interface, 0),
		},
	}
}

// GetDefaultLogger creates a new default logger.
func GetDefaultLogger() *Logger {
	newLogger := New("root")

	newFormatter := formatter.New("%(level):%(name):%(message)")

	newHandler := handler.NewConsoleHandler(loglevel.Warning, newFormatter)

	newLogger.baseLogger.AddHandler(newHandler)

	return newLogger
}

// Name returns logger name for the Logger.
func (logger *Logger) Name() string {
	return logger.baseLogger.Name()
}

// Handlers returns a list of the registered handler.Interface objects for the
// Logger.
func (logger *Logger) Handlers() []handler.Interface {
	return logger.baseLogger.Handlers()
}

// AddHandler registers a new handler.Interface for the Logger.
func (logger *Logger) AddHandler(handler handler.Interface) {
	logger.baseLogger.AddHandler(handler)
}

// Trace logs a new message using Logger with loglevel.Trace level.
func (logger *Logger) Trace(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Trace, message, parameters...)
}

// Debug logs a new message using Logger with loglevel.Debug level.
func (logger *Logger) Debug(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Debug, message, parameters...)
}

// Verbose logs a new message using Logger with loglevel.Verbose level.
func (logger *Logger) Verbose(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Verbose, message, parameters...)
}

// Info logs a new message using Logger with loglevel.Info level.
func (logger *Logger) Info(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Info, message, parameters...)
}

// Notice logs a new message using Logger with loglevel.Notice level.
func (logger *Logger) Notice(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Notice, message, parameters...)
}

// Warning logs a new message using Logger with loglevel.Warning level.
func (logger *Logger) Warning(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Warning, message, parameters...)
}

// Severe logs a new message using Logger with loglevel.Severe level.
func (logger *Logger) Severe(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Severe, message, parameters...)
}

// Error logs a new message using Logger with loglevel.Error level.
func (logger *Logger) Error(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Error, message, parameters...)
}

// Alert logs a new message using Logger with loglevel.Alert level.
func (logger *Logger) Alert(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Alert, message, parameters...)
}

// Critical logs a new message using Logger with loglevel.Critical level.
func (logger *Logger) Critical(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Critical, message, parameters...)
}

// Emergency logs a new message using Logger with loglevel.Emergency level.
func (logger *Logger) Emergency(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Emergency, message, parameters...)
}

// SetLevel sets a new loglevel.LogLevel for the default logger.
func SetLevel(level loglevel.LogLevel) {
	handlerInterface := rootLogger.baseLogger.Handlers()[0]
	if handlerInterface != nil {
		handlerInterface.SetLevel(level)
	}
}

// Trace logs a new message using default logger with loglevel.Trace level.
func Trace(message string, parameters ...any) {
	rootLogger.Trace(message, parameters...)
}

// Debug logs a new message using default logger with loglevel.Debug level.
func Debug(message string, parameters ...any) {
	rootLogger.Debug(message, parameters...)
}

// Verbose logs a new message using default logger with loglevel.Verbose level.
func Verbose(message string, parameters ...any) {
	rootLogger.Verbose(message, parameters...)
}

// Info logs a new message using default logger with loglevel.Info level.
func Info(message string, parameters ...any) {
	rootLogger.Info(message, parameters...)
}

// Notice logs a new message using default logger with loglevel.Notice level.
func Notice(message string, parameters ...any) {
	rootLogger.Notice(message, parameters...)
}

// Warning logs a new message using default logger with loglevel.Warning level.
func Warning(message string, parameters ...any) {
	rootLogger.Warning(message, parameters...)
}

// Severe logs a new message using default logger with loglevel.Severe level.
func Severe(message string, parameters ...any) {
	rootLogger.Severe(message, parameters...)
}

// Error logs a new message using default logger with loglevel.Error level.
func Error(message string, parameters ...any) {
	rootLogger.Error(message, parameters...)
}

// Alert logs a new message using default logger with loglevel.Alert level.
func Alert(message string, parameters ...any) {
	rootLogger.Alert(message, parameters...)
}

// Critical logs a new message using default logger with loglevel.Critical level.
func Critical(message string, parameters ...any) {
	rootLogger.Critical(message, parameters...)
}

// Emergency logs a new message using default logger with loglevel.Emergency level.
func Emergency(message string, parameters ...any) {
	rootLogger.Emergency(message, parameters...)
}
