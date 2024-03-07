package logger

import (
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"github.com/dl1998/go-logging/pkg/logger/loglevel"
)

var rootLogger = GetDefaultLogger()

type baseLoggerInterface interface {
	Log(level loglevel.LogLevel, message string, parameters ...any)
	Name() string
	SetName(name string)
	Handlers() []handler.Interface
	AddHandler(handler handler.Interface)
}

type baseLogger struct {
	name     string
	handlers []handler.Interface
}

func (logger *baseLogger) Log(level loglevel.LogLevel, message string, parameters ...any) {
	for _, registeredHandler := range logger.handlers {
		if level >= registeredHandler.Level() {
			registeredHandler.Write(logger.name, level, message, parameters...)
		}
	}
}

func (logger *baseLogger) Name() string {
	return logger.name
}

func (logger *baseLogger) SetName(name string) {
	logger.name = name
}

func (logger *baseLogger) Handlers() []handler.Interface {
	return logger.handlers
}

func (logger *baseLogger) AddHandler(handler handler.Interface) {
	logger.handlers = append(logger.handlers, handler)
}

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

type Logger struct {
	baseLogger baseLoggerInterface
}

func New(name string) *Logger {
	return &Logger{
		baseLogger: &baseLogger{
			name:     name,
			handlers: make([]handler.Interface, 0),
		},
	}
}

func GetDefaultLogger() *Logger {
	newLogger := New("root")

	newFormatter := formatter.New("%(level):%(name):%(message)")

	newHandler := handler.NewConsoleHandler(loglevel.Warning, newFormatter)

	newLogger.baseLogger.AddHandler(newHandler)

	return newLogger
}

func (logger *Logger) Name() string {
	return logger.baseLogger.Name()
}

func (logger *Logger) Handlers() []handler.Interface {
	return logger.baseLogger.Handlers()
}

func (logger *Logger) AddHandler(handler handler.Interface) {
	logger.baseLogger.AddHandler(handler)
}

func (logger *Logger) Trace(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Trace, message, parameters...)
}

func (logger *Logger) Debug(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Debug, message, parameters...)
}

func (logger *Logger) Verbose(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Verbose, message, parameters...)
}

func (logger *Logger) Info(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Info, message, parameters...)
}

func (logger *Logger) Notice(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Notice, message, parameters...)
}

func (logger *Logger) Warning(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Warning, message, parameters...)
}

func (logger *Logger) Severe(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Severe, message, parameters...)
}

func (logger *Logger) Error(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Error, message, parameters...)
}

func (logger *Logger) Alert(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Alert, message, parameters...)
}

func (logger *Logger) Critical(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Critical, message, parameters...)
}

func (logger *Logger) Emergency(message string, parameters ...any) {
	logger.baseLogger.Log(loglevel.Emergency, message, parameters...)
}

func SetLevel(level loglevel.LogLevel) {
	handlerInterface := rootLogger.baseLogger.Handlers()[0]
	if handlerInterface != nil {
		handlerInterface.SetLevel(level)
	}
}

func Trace(message string, parameters ...any) {
	rootLogger.Trace(message, parameters...)
}

func Debug(message string, parameters ...any) {
	rootLogger.Debug(message, parameters...)
}

func Verbose(message string, parameters ...any) {
	rootLogger.Verbose(message, parameters...)
}

func Info(message string, parameters ...any) {
	rootLogger.Info(message, parameters...)
}

func Notice(message string, parameters ...any) {
	rootLogger.Notice(message, parameters...)
}

func Warning(message string, parameters ...any) {
	rootLogger.Warning(message, parameters...)
}

func Severe(message string, parameters ...any) {
	rootLogger.Severe(message, parameters...)
}

func Error(message string, parameters ...any) {
	rootLogger.Error(message, parameters...)
}

func Alert(message string, parameters ...any) {
	rootLogger.Alert(message, parameters...)
}

func Critical(message string, parameters ...any) {
	rootLogger.Critical(message, parameters...)
}

func Emergency(message string, parameters ...any) {
	rootLogger.Emergency(message, parameters...)
}
