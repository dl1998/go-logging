// Package structuredlogger contains structured logger implementation.
package structuredlogger

import (
	"fmt"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/structuredlogger/formatter"
	"github.com/dl1998/go-logging/pkg/structuredlogger/handler"
	"time"
)

const (
	JSONFormatterType     = "json"
	KeyValueFormatterType = "key-value"
)

var rootLogger *Logger
var fromLevel level.Level
var toLevel level.Level
var template map[string]string

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
	CaptureError(message error, parameters ...any)
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
func (logger *Logger) Trace(parameters ...any) {
	logger.baseLogger.Log(level.Trace, parameters...)
}

// Debug logs a new message using Logger with level.Debug level.
func (logger *Logger) Debug(parameters ...any) {
	logger.baseLogger.Log(level.Debug, parameters...)
}

// Verbose logs a new message using Logger with level.Verbose level.
func (logger *Logger) Verbose(parameters ...any) {
	logger.baseLogger.Log(level.Verbose, parameters...)
}

// Info logs a new message using Logger with level.Info level.
func (logger *Logger) Info(parameters ...any) {
	logger.baseLogger.Log(level.Info, parameters...)
}

// Notice logs a new message using Logger with level.Notice level.
func (logger *Logger) Notice(parameters ...any) {
	logger.baseLogger.Log(level.Notice, parameters...)
}

// Warning logs a new message using Logger with level.Warning level.
func (logger *Logger) Warning(parameters ...any) {
	logger.baseLogger.Log(level.Warning, parameters...)
}

// Severe logs a new message using Logger with level.Severe level.
func (logger *Logger) Severe(parameters ...any) {
	logger.baseLogger.Log(level.Severe, parameters...)
}

// Error logs a new message using Logger with level.Error level.
func (logger *Logger) Error(parameters ...any) {
	logger.baseLogger.Log(level.Error, parameters...)
}

// Alert logs a new message using Logger with level.Alert level.
func (logger *Logger) Alert(parameters ...any) {
	logger.baseLogger.Log(level.Alert, parameters...)
}

// Critical logs a new message using Logger with level.Critical level.
func (logger *Logger) Critical(parameters ...any) {
	logger.baseLogger.Log(level.Critical, parameters...)
}

// Emergency logs a new message using Logger with level.Emergency level.
func (logger *Logger) Emergency(parameters ...any) {
	logger.baseLogger.Log(level.Emergency, parameters...)
}

// ErrorLevel returns errorLevel for the Logger.
func (logger *Logger) ErrorLevel() level.Level {
	return logger.errorLevel
}

// SetErrorLevel sets errorLevel in the Logger that is used in the RaiseError and
// CaptureError methods.
func (logger *Logger) SetErrorLevel(newLevel level.Level) {
	logger.errorLevel = newLevel
}

// PanicLevel returns panicLevel for the Logger.
func (logger *Logger) PanicLevel() level.Level {
	return logger.panicLevel
}

// SetPanicLevel sets panicLevel in the Logger that is used in the Panic method.
func (logger *Logger) SetPanicLevel(newLevel level.Level) {
	logger.panicLevel = newLevel
}

// RaiseError logs a new message using Logger and returns a new error with logged
// error message.
func (logger *Logger) RaiseError(message string, parameters ...any) error {
	parametersArray := []any{"error", message}
	if parameters != nil {
		parametersArray = append(parametersArray, parameters...)
	}
	logger.baseLogger.Log(logger.errorLevel, parametersArray...)
	return fmt.Errorf(message)
}

// CaptureError logs a new message from the error using Logger.
func (logger *Logger) CaptureError(message error, parameters ...any) {
	parametersArray := []any{"error", message.Error()}
	if parameters != nil {
		parametersArray = append(parametersArray, parameters...)
	}
	logger.baseLogger.Log(logger.errorLevel, parametersArray...)
}

// Panic logs a new message using Logger and panics with the message.
func (logger *Logger) Panic(message string, parameters ...any) {
	parametersArray := []any{"panic", message}
	if parameters != nil {
		parametersArray = append(parametersArray, parameters...)
	}
	logger.baseLogger.Log(logger.panicLevel, parametersArray...)
	panic(message)
}

// Configuration struct contains configuration for the logger.
type Configuration struct {
	errorLevel        level.Level
	panicLevel        level.Level
	fromLevel         level.Level
	toLevel           level.Level
	template          map[string]string
	format            string
	pretty            bool
	keyValueDelimiter string
	pairSeparator     string
	file              string
	name              string
	timeFormat        string
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
func WithTemplate(template map[string]string) Option {
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

// WithFormat sets format for the Configuration.
func WithFormat(format string) Option {
	return func(configuration *Configuration) {
		configuration.format = format
	}
}

// WithPretty sets pretty for the Configuration.
func WithPretty(pretty bool) Option {
	return func(configuration *Configuration) {
		configuration.pretty = pretty
	}
}

// WithKeyValueDelimiter sets keyValueDelimiter for the Configuration.
func WithKeyValueDelimiter(keyValueDelimiter string) Option {
	return func(configuration *Configuration) {
		configuration.keyValueDelimiter = keyValueDelimiter
	}
}

// WithPairSeparator sets pairSeparator for the Configuration.
func WithPairSeparator(pairSeparator string) Option {
	return func(configuration *Configuration) {
		configuration.pairSeparator = pairSeparator
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
		template: map[string]string{
			"timestamp": "%(timestamp)",
			"level":     "%(level)",
			"name":      "%(name)",
		},
		format:            JSONFormatterType,
		pretty:            false,
		keyValueDelimiter: "=",
		pairSeparator:     " ",
		file:              "",
		name:              "root",
		timeFormat:        time.RFC3339,
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

	var defaultFormatter formatter.Interface

	if configuration.format == JSONFormatterType {
		defaultFormatter = formatter.NewJSON(configuration.template, configuration.pretty)
	} else if configuration.format == KeyValueFormatterType {
		defaultFormatter = formatter.NewKeyValue(configuration.template, configuration.keyValueDelimiter, configuration.pairSeparator)
	} else {
		panic("unsupported format")
	}

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
func Template() map[string]string {
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
func Trace(parameters ...any) {
	rootLogger.Trace(parameters...)
}

// Debug logs a new message using default logger with level.Debug level.
func Debug(parameters ...any) {
	rootLogger.Debug(parameters...)
}

// Verbose logs a new message using default logger with level.Verbose level.
func Verbose(parameters ...any) {
	rootLogger.Verbose(parameters...)
}

// Info logs a new message using default logger with level.Info level.
func Info(parameters ...any) {
	rootLogger.Info(parameters...)
}

// Notice logs a new message using default logger with level.Notice level.
func Notice(parameters ...any) {
	rootLogger.Notice(parameters...)
}

// Warning logs a new message using default logger with level.Warning level.
func Warning(parameters ...any) {
	rootLogger.Warning(parameters...)
}

// Severe logs a new message using default logger with level.Severe level.
func Severe(parameters ...any) {
	rootLogger.Severe(parameters...)
}

// Error logs a new message using default logger with level.Error level.
func Error(parameters ...any) {
	rootLogger.Error(parameters...)
}

// Alert logs a new message using default logger with level.Alert level.
func Alert(parameters ...any) {
	rootLogger.Alert(parameters...)
}

// Critical logs a new message using default logger with level.Critical level.
func Critical(parameters ...any) {
	rootLogger.Critical(parameters...)
}

// Emergency logs a new message using default logger with level.Emergency level.
func Emergency(parameters ...any) {
	rootLogger.Emergency(parameters...)
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
func CaptureError(message error, parameters ...any) {
	rootLogger.CaptureError(message, parameters...)
}

// Panic logs a new message using default logger and panics with the message.
func Panic(message string, parameters ...any) {
	rootLogger.Panic(message, parameters...)
}
