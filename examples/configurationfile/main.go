// Example that shows how to create and use loggers from the configuration file.
package main

import (
	standard "github.com/dl1998/go-logging/pkg/logger/configuration/parser"
	structured "github.com/dl1998/go-logging/pkg/structuredlogger/configuration/parser"
	"path"
	"runtime"
)

var (
	ConfigurationsPath string
	LoggerName         = "example-logger"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	ConfigurationsPath = path.Join(path.Dir(fileName), "inputs")
}

func main() {
	exampleStandardLogsFromJSONConfiguration()
	exampleStandardLogsFromYAMLConfiguration()
	exampleStandardLogsFromXMLConfiguration()
	exampleStructuredLogsFromJSONConfiguration()
	exampleStructuredLogsFromYAMLConfiguration()
	exampleStructuredLogsFromXMLConfiguration()
}

func exampleStandardLogsFromJSONConfiguration() {
	newParser, err := standard.ParseJSON(path.Join(ConfigurationsPath, "example.json"))
	if err != nil {
		panic(err)
	}

	newLogger := newParser.GetLogger(LoggerName)

	newLogger.Info("JSON: Message")

	newAsyncLogger := newParser.GetAsyncLogger(LoggerName)

	newAsyncLogger.Info("JSON: Async Message")

	newAsyncLogger.WaitToFinishLogging()
}

func exampleStandardLogsFromYAMLConfiguration() {
	newParser, err := standard.ParseYAML(path.Join(ConfigurationsPath, "example.yaml"))
	if err != nil {
		panic(err)
	}

	newLogger := newParser.GetLogger(LoggerName)

	newLogger.Info("YAML: Message")

	newAsyncLogger := newParser.GetAsyncLogger(LoggerName)

	newAsyncLogger.Info("YAML: Async Message")

	newAsyncLogger.WaitToFinishLogging()
}

func exampleStandardLogsFromXMLConfiguration() {
	newParser, err := standard.ParseXML(path.Join(ConfigurationsPath, "example.xml"))
	if err != nil {
		panic(err)
	}

	newLogger := newParser.GetLogger(LoggerName)

	newLogger.Info("XML: Message")

	newAsyncLogger := newParser.GetAsyncLogger(LoggerName)

	newAsyncLogger.Info("XML: Async Message")

	newAsyncLogger.WaitToFinishLogging()
}

func exampleStructuredLogsFromJSONConfiguration() {
	newParser, err := structured.ParseJSON(path.Join(ConfigurationsPath, "example.json"))
	if err != nil {
		panic(err)
	}

	newLogger := newParser.GetLogger(LoggerName)

	newLogger.Info("message", "JSON: Message")

	newAsyncLogger := newParser.GetAsyncLogger(LoggerName)

	newAsyncLogger.Info("message", "JSON: Async Message")

	newAsyncLogger.WaitToFinishLogging()
}

func exampleStructuredLogsFromYAMLConfiguration() {
	newParser, err := structured.ParseYAML(path.Join(ConfigurationsPath, "example.yaml"))
	if err != nil {
		panic(err)
	}

	newLogger := newParser.GetLogger(LoggerName)

	newLogger.Info("message", "YAML: Message")

	newAsyncLogger := newParser.GetAsyncLogger(LoggerName)

	newAsyncLogger.Info("message", "YAML: Async Message")

	newAsyncLogger.WaitToFinishLogging()
}

func exampleStructuredLogsFromXMLConfiguration() {
	newParser, err := structured.ParseXML(path.Join(ConfigurationsPath, "example.xml"))
	if err != nil {
		panic(err)
	}

	newLogger := newParser.GetLogger(LoggerName)

	newLogger.Info("message", "XNL: Message")

	newAsyncLogger := newParser.GetAsyncLogger(LoggerName)

	newAsyncLogger.Info("message", "XML: Async Message")

	newAsyncLogger.WaitToFinishLogging()
}
