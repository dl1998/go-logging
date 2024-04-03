// Package parser contains the configuration parser for the logger.
package parser

import (
	"github.com/dl1998/go-logging/pkg/common/configuration/parser"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"strings"
)

// parseFormatter parses parser.FormatterConfiguration configuration and returns
// formatter.Interface.
func parseFormatter(configuration parser.FormatterConfiguration) formatter.Interface {
	return formatter.New(configuration.Template.(string))
}

// parseHandler parses parser.HandlerConfiguration configuration and returns
// handler.Interface.
func parseHandler(configuration parser.HandlerConfiguration) handler.Interface {
	fromLevel := level.ParseLevel(strings.ToLower(configuration.FromLevel))
	toLevel := level.ParseLevel(strings.ToLower(configuration.ToLevel))
	switch configuration.Type {
	case "stdout":
		return handler.NewConsoleHandler(fromLevel, toLevel, parseFormatter(configuration.Formatter))
	case "stderr":
		return handler.NewConsoleErrorHandler(fromLevel, toLevel, parseFormatter(configuration.Formatter))
	case "file":
		if configuration.File == "" {
			panic("file handler requires file option.")
		}
		return handler.NewFileHandler(fromLevel, toLevel, parseFormatter(configuration.Formatter), configuration.File)
	default:
		return nil
	}
}

// parseLogger parses parser.LoggerConfiguration configuration and returns
// logger.Logger.
func parseLogger(configuration parser.LoggerConfiguration) *logger.Logger {
	newLogger := logger.New(configuration.Name, configuration.TimeFormat)
	for _, handlerConfiguration := range configuration.Handlers {
		newLogger.AddHandler(parseHandler(handlerConfiguration))
	}
	return newLogger
}

// parseAsyncLogger parses parser.LoggerConfiguration configuration and returns
// logger.AsyncLogger.
func parseAsyncLogger(configuration parser.LoggerConfiguration) *logger.AsyncLogger {
	newLogger := logger.NewAsyncLogger(configuration.Name, configuration.TimeFormat, configuration.MessageQueueSize)
	for _, handlerConfiguration := range configuration.Handlers {
		newLogger.AddHandler(parseHandler(handlerConfiguration))
	}
	return newLogger
}

// GetLogger returns logger.Logger by name from the configuration.
func GetLogger(name string, configuration parser.Configuration) *logger.Logger {
	for _, loggerConfiguration := range configuration.Loggers {
		if loggerConfiguration.Name == name {
			return parseLogger(loggerConfiguration)
		}
	}
	return nil
}

// GetAsyncLogger returns logger.AsyncLogger by name from the configuration.
func GetAsyncLogger(name string, configuration parser.Configuration) *logger.AsyncLogger {
	for _, loggerConfiguration := range configuration.Loggers {
		if loggerConfiguration.Name == name {
			return parseAsyncLogger(loggerConfiguration)
		}
	}
	return nil
}
