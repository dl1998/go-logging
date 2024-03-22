// Package logger contains logger implementation.
package logger

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"github.com/dl1998/go-logging/pkg/logger/logrecord"
)

var rootLogger *Logger
var fromLevel level.Level
var toLevel level.Level
var template string

func init() {
	Configure(NewConfiguration())
}

// baseLoggerInterface defines low level logging interface.
type baseLoggerInterface interface {
	Log(level level.Level, message string, parameters ...any)
	Name() string
	SetName(name string)
	Handlers() []handler.Interface
	AddHandler(handlerInterface handler.Interface)
	RemoveHandler(handlerInterface handler.Interface)
}

// baseLogger struct contains basic fields for the logger.
type baseLogger struct {
	name     string
	handlers []handler.Interface
}

// Log logs interpolated message with the provided level.Level.
func (logger *baseLogger) Log(level level.Level, message string, parameters ...any) {
	record := logrecord.New(logger.name, level, template, message, parameters, 3)
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

// Interface represents interface that shall be satisfied by Logger.
type Interface interface {
	Name() string
	Handlers() []handler.Interface
	AddHandler(handlerInterface handler.Interface)
	RemoveHandler(handlerInterface handler.Interface)
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
func (logger *Logger) AddHandler(handlerInterface handler.Interface) {
	logger.baseLogger.AddHandler(handlerInterface)
}

// RemoveHandler removes a handler.Interface from the Logger handlers.
func (logger *Logger) RemoveHandler(handlerInterface handler.Interface) {
	logger.baseLogger.RemoveHandler(handlerInterface)
}

// Trace logs a new message using Logger with level.Trace level.
func (logger *Logger) Trace(message string, parameters ...any) {
	logger.baseLogger.Log(level.Trace, message, parameters...)
}

// Debug logs a new message using Logger with level.Debug level.
func (logger *Logger) Debug(message string, parameters ...any) {
	logger.baseLogger.Log(level.Debug, message, parameters...)
}

// Verbose logs a new message using Logger with level.Verbose level.
func (logger *Logger) Verbose(message string, parameters ...any) {
	logger.baseLogger.Log(level.Verbose, message, parameters...)
}

// Info logs a new message using Logger with level.Info level.
func (logger *Logger) Info(message string, parameters ...any) {
	logger.baseLogger.Log(level.Info, message, parameters...)
}

// Notice logs a new message using Logger with level.Notice level.
func (logger *Logger) Notice(message string, parameters ...any) {
	logger.baseLogger.Log(level.Notice, message, parameters...)
}

// Warning logs a new message using Logger with level.Warning level.
func (logger *Logger) Warning(message string, parameters ...any) {
	logger.baseLogger.Log(level.Warning, message, parameters...)
}

// Severe logs a new message using Logger with level.Severe level.
func (logger *Logger) Severe(message string, parameters ...any) {
	logger.baseLogger.Log(level.Severe, message, parameters...)
}

// Error logs a new message using Logger with level.Error level.
func (logger *Logger) Error(message string, parameters ...any) {
	logger.baseLogger.Log(level.Error, message, parameters...)
}

// Alert logs a new message using Logger with level.Alert level.
func (logger *Logger) Alert(message string, parameters ...any) {
	logger.baseLogger.Log(level.Alert, message, parameters...)
}

// Critical logs a new message using Logger with level.Critical level.
func (logger *Logger) Critical(message string, parameters ...any) {
	logger.baseLogger.Log(level.Critical, message, parameters...)
}

// Emergency logs a new message using Logger with level.Emergency level.
func (logger *Logger) Emergency(message string, parameters ...any) {
	logger.baseLogger.Log(level.Emergency, message, parameters...)
}

// Configuration struct contains configuration for the logger.
type Configuration struct {
	fromLevel level.Level
	toLevel   level.Level
	template  string
	file      string
	name      string
}

// Option represents option for the Configuration.
type Option func(*Configuration)

// WithFromLevel sets fromLevel for the Configuration.
func WithFromLevel(fromLevel level.Level) Option {
	return func(configuration *Configuration) {
		configuration.fromLevel = fromLevel
	}
}

// WithToLevel sets toLevel for the Configuration.
func WithToLevel(toLevel level.Level) Option {
	return func(configuration *Configuration) {
		configuration.toLevel = toLevel
	}
}

// WithTemplate sets template for the Configuration.
func WithTemplate(template string) Option {
	return func(configuration *Configuration) {
		configuration.template = template
	}
}

// WithFile sets file for the Configuration.
func WithFile(file string) Option {
	return func(configuration *Configuration) {
		configuration.file = file
	}
}

// WithName sets name for the Configuration.
func WithName(name string) Option {
	return func(configuration *Configuration) {
		configuration.name = name
	}
}

// NewConfiguration creates a new instance of the Configuration.
func NewConfiguration(options ...Option) *Configuration {
	newConfiguration := &Configuration{
		fromLevel: level.Warning,
		toLevel:   level.Null,
		template:  "%(level):%(name):%(message)",
		file:      "",
		name:      "root",
	}

	for _, option := range options {
		option(newConfiguration)
	}

	return newConfiguration
}

// Configure configures the logger with the provided configuration.
func Configure(configuration *Configuration) {
	if configuration.fromLevel.DigitRepresentation() > configuration.toLevel.DigitRepresentation() {
		panic("fromLevel cannot be higher than toLevel")
	}

	fromLevel = configuration.fromLevel
	toLevel = configuration.toLevel
	template = configuration.template

	newLogger := New(configuration.name)

	defaultFormatter := formatter.New(configuration.template)

	var createStdoutHandler = configuration.fromLevel.DigitRepresentation() <= level.Severe.DigitRepresentation()
	var createStderrHandler = configuration.toLevel.DigitRepresentation() >= level.Error.DigitRepresentation()
	var createFileHandler = configuration.file != ""

	if createStdoutHandler {
		stdoutToLevel := toLevel
		if stdoutToLevel > level.Severe {
			stdoutToLevel = level.Severe
		}
		newHandler := handler.NewConsoleHandler(configuration.fromLevel, stdoutToLevel, defaultFormatter)
		newLogger.baseLogger.AddHandler(newHandler)
	}

	if createStderrHandler {
		stderrFromLevel := fromLevel
		if stderrFromLevel < level.Error {
			stderrFromLevel = level.Error
		}
		newHandler := handler.NewConsoleErrorHandler(stderrFromLevel, configuration.toLevel, defaultFormatter)
		newLogger.baseLogger.AddHandler(newHandler)
	}

	if createFileHandler {
		newHandler := handler.NewFileHandler(configuration.fromLevel, configuration.toLevel, defaultFormatter, configuration.file)
		newLogger.baseLogger.AddHandler(newHandler)
	}

	rootLogger = newLogger
}

// Name returns name of the rootLogger.
func Name() string {
	return rootLogger.Name()
}

// Template returns template of the rootLogger.
func Template() string {
	return template
}

// FromLevel returns fromLevel of the rootLogger.
func FromLevel() level.Level {
	return fromLevel
}

// ToLevel returns toLevel of the rootLogger.
func ToLevel() level.Level {
	return toLevel
}

// Trace logs a new message using default logger with level.Trace level.
func Trace(message string, parameters ...any) {
	rootLogger.Trace(message, parameters...)
}

// Debug logs a new message using default logger with level.Debug level.
func Debug(message string, parameters ...any) {
	rootLogger.Debug(message, parameters...)
}

// Verbose logs a new message using default logger with level.Verbose level.
func Verbose(message string, parameters ...any) {
	rootLogger.Verbose(message, parameters...)
}

// Info logs a new message using default logger with level.Info level.
func Info(message string, parameters ...any) {
	rootLogger.Info(message, parameters...)
}

// Notice logs a new message using default logger with level.Notice level.
func Notice(message string, parameters ...any) {
	rootLogger.Notice(message, parameters...)
}

// Warning logs a new message using default logger with level.Warning level.
func Warning(message string, parameters ...any) {
	rootLogger.Warning(message, parameters...)
}

// Severe logs a new message using default logger with level.Severe level.
func Severe(message string, parameters ...any) {
	rootLogger.Severe(message, parameters...)
}

// Error logs a new message using default logger with level.Error level.
func Error(message string, parameters ...any) {
	rootLogger.Error(message, parameters...)
}

// Alert logs a new message using default logger with level.Alert level.
func Alert(message string, parameters ...any) {
	rootLogger.Alert(message, parameters...)
}

// Critical logs a new message using default logger with level.Critical level.
func Critical(message string, parameters ...any) {
	rootLogger.Critical(message, parameters...)
}

// Emergency logs a new message using default logger with level.Emergency level.
func Emergency(message string, parameters ...any) {
	rootLogger.Emergency(message, parameters...)
}
