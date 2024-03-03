package logger

import (
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"github.com/dl1998/go-logging/pkg/logger/loglevel"
)

var rootLogger = GetDefaultLogger()

type Logger struct {
	Name     string
	Handlers []*handler.Handler
}

func New(name string) *Logger {
	return &Logger{
		Name:     name,
		Handlers: make([]*handler.Handler, 0),
	}
}

func GetDefaultLogger() *Logger {
	newLogger := New("root")

	newFormatter := formatter.New("%(level):%(name):%(message)")

	newHandler := handler.NewConsoleHandler(loglevel.Warning, *newFormatter)

	newLogger.AddHandler(newHandler)

	return newLogger
}

func (logger *Logger) log(level loglevel.LogLevel, message string, parameters ...any) {
	for _, registeredHandler := range logger.Handlers {
		if level >= registeredHandler.Level() {
			registeredHandler.Write(logger.Name, level, message, parameters...)
		}
	}
}

func (logger *Logger) Trace(message string, parameters ...any) {
	logger.log(loglevel.Trace, message, parameters...)
}

func (logger *Logger) Debug(message string, parameters ...any) {
	logger.log(loglevel.Debug, message, parameters...)
}

func (logger *Logger) Verbose(message string, parameters ...any) {
	logger.log(loglevel.Verbose, message, parameters...)
}

func (logger *Logger) Info(message string, parameters ...any) {
	logger.log(loglevel.Info, message, parameters...)
}

func (logger *Logger) Notice(message string, parameters ...any) {
	logger.log(loglevel.Notice, message, parameters...)
}

func (logger *Logger) Warning(message string, parameters ...any) {
	logger.log(loglevel.Warning, message, parameters...)
}

func (logger *Logger) Severe(message string, parameters ...any) {
	logger.log(loglevel.Severe, message, parameters...)
}

func (logger *Logger) Error(message string, parameters ...any) {
	logger.log(loglevel.Error, message, parameters...)
}

func (logger *Logger) Alert(message string, parameters ...any) {
	logger.log(loglevel.Alert, message, parameters...)
}

func (logger *Logger) Critical(message string, parameters ...any) {
	logger.log(loglevel.Critical, message, parameters...)
}

func (logger *Logger) Emergency(message string, parameters ...any) {
	logger.log(loglevel.Emergency, message, parameters...)
}

func (logger *Logger) AddHandler(handler *handler.Handler) {
	logger.Handlers = append(logger.Handlers, handler)
}

func SetLevel(level loglevel.LogLevel) {
	if rootLogger.Handlers[0] != nil {
		rootLogger.Handlers[0].SetLevel(level)
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
