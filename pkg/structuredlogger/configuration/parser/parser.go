// Package parser contains the configuration parser for the structured logger.
package parser

import (
	"github.com/dl1998/go-logging/pkg/common/configuration/parser"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/structuredlogger"
	"github.com/dl1998/go-logging/pkg/structuredlogger/formatter"
	"github.com/dl1998/go-logging/pkg/structuredlogger/handler"
	"strings"
)

// Parser is the configuration parser for the structured logger.
type Parser struct {
	// configuration is the parser.Configuration with the structured logger
	// configuration.
	configuration *parser.Configuration
}

// NewParser returns a new instance of the Parser with the given
// parser.Configuration.
func NewParser(configuration parser.Configuration) *Parser {
	return &Parser{configuration: &configuration}
}

// parseFile parses the file with the given parser function and returns the
// Parser.
func parseFile(file string, parserFunction func(string) (*parser.Configuration, error)) (*Parser, error) {
	configuration, err := parserFunction(file)
	if err != nil {
		return nil, err
	}
	return NewParser(*configuration), nil
}

// ParseJSON parses the JSON file and returns the Parser.
func ParseJSON(file string) (*Parser, error) {
	return parseFile(file, parser.ReadFromJSON)
}

// ParseYAML parses the YAML file and returns the Parser.
func ParseYAML(file string) (*Parser, error) {
	return parseFile(file, parser.ReadFromYAML)
}

// ParseXML parses the XML file and returns the Parser.
func ParseXML(file string) (*Parser, error) {
	return parseFile(file, parser.ReadFromXML)
}

// parseFormatter parses parser.FormatterConfiguration configuration and returns
// formatter.Interface.
func (parser *Parser) parseFormatter(configuration parser.FormatterConfiguration) formatter.Interface {
	switch configuration.Type {
	case "json":
		return formatter.NewJSON(configuration.Template.MapValue, configuration.PrettyPrint)
	case "key-value":
		return formatter.NewKeyValue(configuration.Template.MapValue, configuration.KeyValueDelimiter, configuration.PairSeparator)
	default:
		panic("unknown formatter type.")
	}
}

// parseHandler parses parser.HandlerConfiguration configuration and returns
// handler.Interface.
func (parser *Parser) parseHandler(configuration parser.HandlerConfiguration) handler.Interface {
	fromLevel := level.ParseLevel(strings.ToLower(configuration.FromLevel))
	toLevel := level.ParseLevel(strings.ToLower(configuration.ToLevel))
	switch configuration.Type {
	case "stdout":
		return handler.NewConsoleHandler(fromLevel, toLevel, parser.parseFormatter(configuration.Formatter))
	case "stderr":
		return handler.NewConsoleErrorHandler(fromLevel, toLevel, parser.parseFormatter(configuration.Formatter))
	case "file":
		if configuration.File == "" {
			panic("file handler requires file option.")
		}
		return handler.NewFileHandler(fromLevel, toLevel, parser.parseFormatter(configuration.Formatter), configuration.File)
	default:
		return nil
	}
}

// parseLogger parses parser.LoggerConfiguration configuration and returns
// structuredlogger.Logger.
func (parser *Parser) parseLogger(configuration parser.LoggerConfiguration) *structuredlogger.Logger {
	newLogger := structuredlogger.New(configuration.Name, configuration.TimeFormat)
	newLogger.SetErrorLevel(level.ParseLevel(strings.ToLower(configuration.ErrorLevel)))
	newLogger.SetPanicLevel(level.ParseLevel(strings.ToLower(configuration.PanicLevel)))
	for _, handlerConfiguration := range configuration.Handlers {
		newLogger.AddHandler(parser.parseHandler(handlerConfiguration))
	}
	return newLogger
}

// parseAsyncLogger parses parser.LoggerConfiguration configuration and returns
// structuredlogger.AsyncLogger.
func (parser *Parser) parseAsyncLogger(configuration parser.LoggerConfiguration) *structuredlogger.AsyncLogger {
	newLogger := structuredlogger.NewAsyncLogger(configuration.Name, configuration.TimeFormat, configuration.MessageQueueSize)
	newLogger.SetErrorLevel(level.ParseLevel(strings.ToLower(configuration.ErrorLevel)))
	newLogger.SetPanicLevel(level.ParseLevel(strings.ToLower(configuration.PanicLevel)))
	for _, handlerConfiguration := range configuration.Handlers {
		newLogger.AddHandler(parser.parseHandler(handlerConfiguration))
	}
	return newLogger
}

// GetLogger returns structuredlogger.Logger by name from the configuration.
func (parser *Parser) GetLogger(name string) *structuredlogger.Logger {
	for _, loggerConfiguration := range parser.configuration.Loggers {
		if loggerConfiguration.Name == name {
			return parser.parseLogger(loggerConfiguration)
		}
	}
	return nil
}

// GetAsyncLogger returns structuredlogger.AsyncLogger by name from the
// configuration.
func (parser *Parser) GetAsyncLogger(name string) *structuredlogger.AsyncLogger {
	for _, loggerConfiguration := range parser.configuration.Loggers {
		if loggerConfiguration.Name == name {
			return parser.parseAsyncLogger(loggerConfiguration)
		}
	}
	return nil
}
