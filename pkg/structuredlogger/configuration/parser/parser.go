// Package parser contains the configuration parser for the structured logger.
package parser

import (
	"fmt"
	"github.com/dl1998/go-logging/pkg/common/configuration/parser"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/structuredlogger"
	"github.com/dl1998/go-logging/pkg/structuredlogger/formatter"
	"github.com/dl1998/go-logging/pkg/structuredlogger/handler"
	"strings"
)

// convertMap converts map[string]interface{} to map[string]string.
func convertMap(input interface{}) map[string]string {
	data := input.(map[string]interface{})
	result := make(map[string]string)
	for key, value := range data {
		result[key] = fmt.Sprintf("%v", value)
	}
	return result
}

// parseFormatter parses parser.FormatterConfiguration configuration and returns
// formatter.Interface.
func parseFormatter(configuration parser.FormatterConfiguration) formatter.Interface {
	switch configuration.Type {
	case "json":
		return formatter.NewJSON(convertMap(configuration.Template), configuration.PrettyPrint)
	case "key-value":
		return formatter.NewKeyValue(convertMap(configuration.Template), configuration.KeyValueDelimiter, configuration.PairSeparator)
	default:
		panic("unknown formatter type.")
	}
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
// structuredlogger.Logger.
func parseLogger(configuration parser.LoggerConfiguration) *structuredlogger.Logger {
	newLogger := structuredlogger.New(configuration.Name, configuration.TimeFormat)
	for _, handlerConfiguration := range configuration.Handlers {
		newLogger.AddHandler(parseHandler(handlerConfiguration))
	}
	return newLogger
}

// parseAsyncLogger parses parser.LoggerConfiguration configuration and returns
// structuredlogger.AsyncLogger.
func parseAsyncLogger(configuration parser.LoggerConfiguration) *structuredlogger.AsyncLogger {
	newLogger := structuredlogger.NewAsyncLogger(configuration.Name, configuration.TimeFormat, configuration.MessageQueueSize)
	for _, handlerConfiguration := range configuration.Handlers {
		newLogger.AddHandler(parseHandler(handlerConfiguration))
	}
	return newLogger
}

// GetLogger returns structuredlogger.Logger by name from the configuration.
func GetLogger(name string, configuration parser.Configuration) *structuredlogger.Logger {
	for _, loggerConfiguration := range configuration.Loggers {
		if loggerConfiguration.Name == name {
			return parseLogger(loggerConfiguration)
		}
	}
	return nil
}

// GetAsyncLogger returns structuredlogger.AsyncLogger by name from the
// configuration.
func GetAsyncLogger(name string, configuration parser.Configuration) *structuredlogger.AsyncLogger {
	for _, loggerConfiguration := range configuration.Loggers {
		if loggerConfiguration.Name == name {
			return parseAsyncLogger(loggerConfiguration)
		}
	}
	return nil
}
