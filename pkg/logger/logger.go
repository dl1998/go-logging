// Package logger contains logger implementation.
package logger

import (
	"fmt"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"time"
)

var rootLogger Interface
var fromLevel level.Level
var toLevel level.Level
var template string

func init() {
	Configure(NewConfiguration())
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
	ErrorLevel() level.Level
	SetErrorLevel(newLevel level.Level)
	PanicLevel() level.Level
	SetPanicLevel(newLevel level.Level)
	RaiseError(message string, parameters ...any) error
	CaptureError(message error)
	Panic(message string, parameters ...any)
}

// Logger struct encapsulates baseLogger implementation.
type Logger struct {
	baseLogger baseLoggerInterface
	errorLevel level.Level
	panicLevel level.Level
}

// New creates a new instance of the Logger.
func New(name string, timeFormat string) *Logger {
	return &Logger{
		baseLogger: &baseLogger{
			name:       name,
			timeFormat: timeFormat,
			handlers:   make([]handler.Interface, 0),
		},
		errorLevel: level.Error,
		panicLevel: level.Critical,
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

// ErrorLevel returns errorLevel for the Logger.
func (logger *Logger) ErrorLevel() level.Level {
	return logger.errorLevel
}

// SetErrorLevel sets errorLevel in the Logger that is used in the RaiseError and
// CaptureError methods.
func (logger *Logger) SetErrorLevel(newLevel level.Level) {
	if newLevel > level.All && newLevel < level.Null {
		logger.errorLevel = newLevel
	}
}

// PanicLevel returns panicLevel for the Logger.
func (logger *Logger) PanicLevel() level.Level {
	return logger.panicLevel
}

// SetPanicLevel sets panicLevel in the Logger that is used in the Panic method.
func (logger *Logger) SetPanicLevel(newLevel level.Level) {
	if newLevel > level.All && newLevel < level.Null {
		logger.panicLevel = newLevel
	}
}

// RaiseError logs a new message using Logger and returns a new error with logged
// error message.
func (logger *Logger) RaiseError(message string, parameters ...any) error {
	logger.baseLogger.Log(logger.errorLevel, message, parameters...)
	return fmt.Errorf(message, parameters...)
}

// CaptureError logs a new message from the error using Logger.
func (logger *Logger) CaptureError(message error) {
	logger.baseLogger.Log(logger.errorLevel, message.Error())
}

// Panic logs a new message using Logger and panics with the message.
func (logger *Logger) Panic(message string, parameters ...any) {
	logger.baseLogger.Log(logger.panicLevel, message, parameters...)
	panic(fmt.Sprintf(message, parameters...))
}

// Configuration struct contains configuration for the logger.
type Configuration struct {
	errorLevel level.Level
	panicLevel level.Level
	fromLevel  level.Level
	toLevel    level.Level
	template   string
	file       string
	name       string
	timeFormat string
}

// Option represents option for the Configuration.
type Option func(*Configuration)

// WithErrorLevel sets errorLevel for the Configuration.
func WithErrorLevel(errorLevel level.Level) Option {
	return func(configuration *Configuration) {
		configuration.errorLevel = errorLevel
	}
}

// WithPanicLevel sets panicLevel for the Configuration.
func WithPanicLevel(panicLevel level.Level) Option {
	return func(configuration *Configuration) {
		configuration.panicLevel = panicLevel
	}
}

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

// WithTimeFormat sets timeFormat for the Configuration.
func WithTimeFormat(timeFormat string) Option {
	return func(configuration *Configuration) {
		configuration.timeFormat = timeFormat
	}
}

// NewConfiguration creates a new instance of the Configuration.
func NewConfiguration(options ...Option) *Configuration {
	newConfiguration := &Configuration{
		errorLevel: level.Error,
		panicLevel: level.Critical,
		fromLevel:  level.Warning,
		toLevel:    level.Null,
		template:   "%(level):%(name):%(message)",
		file:       "",
		name:       "root",
		timeFormat: time.RFC3339,
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

	newLogger := New(configuration.name, configuration.timeFormat)
	newLogger.SetErrorLevel(configuration.errorLevel)
	newLogger.SetPanicLevel(configuration.panicLevel)

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

// ErrorLevel returns errorLevel in the default logger that is used in the
// RaiseError and CaptureError methods.
func ErrorLevel() level.Level {
	return rootLogger.ErrorLevel()
}

// SetErrorLevel sets errorLevel in the default logger that is used in the
// RaiseError and CaptureError methods.
func SetErrorLevel(newLevel level.Level) {
	rootLogger.SetErrorLevel(newLevel)
}

// PanicLevel returns panicLevel in the default logger that is used in the Panic
// method.
func PanicLevel() level.Level {
	return rootLogger.PanicLevel()
}

// SetPanicLevel sets panicLevel in the default logger that is used in the Panic
// method.
func SetPanicLevel(newLevel level.Level) {
	rootLogger.SetPanicLevel(newLevel)
}

// RaiseError logs a new message using default logger and returns a new error
// with logged error message.
func RaiseError(message string, parameters ...any) error {
	return rootLogger.RaiseError(message, parameters...)
}

// CaptureError logs a new message from the error using default logger.
func CaptureError(message error) {
	rootLogger.CaptureError(message)
}

// Panic logs a new message using default logger and panics with the message.
func Panic(message string, parameters ...any) {
	rootLogger.Panic(message, parameters...)
}
