// Example that shows how to use error/panic wrappers with the
// standard/structured logger.
package main

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	standard "github.com/dl1998/go-logging/pkg/logger"
	standardFormatter "github.com/dl1998/go-logging/pkg/logger/formatter"
	standardHandler "github.com/dl1998/go-logging/pkg/logger/handler"
	structured "github.com/dl1998/go-logging/pkg/structuredlogger"
	structuredFormatter "github.com/dl1998/go-logging/pkg/structuredlogger/formatter"
	structuredHandler "github.com/dl1998/go-logging/pkg/structuredlogger/handler"
	"time"
)

func main() {
	exampleStandardLogger()
	exampleStructuredLogger()
}

// exampleStandardLogger is a sample function to show how to use error/panic
// wrappers with the standard logger.
func exampleStandardLogger() {
	applicationLogger := standard.New("main", time.DateTime)

	applicationFormatter := standardFormatter.New("%(datetime)\t[%(level)]\t%(message)")

	applicationHandler := standardHandler.NewConsoleErrorHandler(level.All, level.Null, applicationFormatter)

	applicationLogger.AddHandler(applicationHandler)

	applicationLogger.SetErrorLevel(level.Critical)
	applicationLogger.SetPanicLevel(level.Emergency)

	err := applicationLogger.RaiseError("exit code: %d", 1)

	applicationLogger.CaptureError(err)

	// This will log message after the panic message.
	defer func() {
		if r := recover(); r != nil {
			applicationLogger.Info("panic recovered: %v", r)
		}
	}()

	applicationLogger.Panic("panic with error code: %d", 123)
}

// exampleStructuredLogger is a sample function to show how to use error/panic
// wrappers with the structured logger.
func exampleStructuredLogger() {
	applicationLogger := structured.New("main", time.DateTime)

	applicationFormatter := structuredFormatter.NewJSON(
		map[string]string{
			"timestamp": "%(timestamp)",
			"level":     "%(level)",
			"name":      "%(name)",
		},
		false,
	)

	applicationHandler := structuredHandler.NewConsoleErrorHandler(level.All, level.Null, applicationFormatter)

	applicationLogger.AddHandler(applicationHandler)

	applicationLogger.SetErrorLevel(level.Critical)
	applicationLogger.SetPanicLevel(level.Emergency)

	err := applicationLogger.RaiseError("error message", "exit code", 1)

	applicationLogger.CaptureError(err, "hostname", "localhost")

	// This will log message after the panic message.
	defer func() {
		if r := recover(); r != nil {
			applicationLogger.Info("panic", r)
		}
	}()

	applicationLogger.Panic("panic message", "hostname", "localhost")
}
