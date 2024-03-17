// Example that shows how to create and use custom structured logger.
package main

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	logger "github.com/dl1998/go-logging/pkg/structuredlogger"
	"github.com/dl1998/go-logging/pkg/structuredlogger/formatter"
	"github.com/dl1998/go-logging/pkg/structuredlogger/handler"
)

func main() {
	template := map[string]string{
		"timestamp": "%(timestamp)",
		"level":     "%(level)",
	}

	applicationLogger := logger.New("example")

	applicationLogger.Debug("message", "This message will not be displayed, because there are no handlers registered.")

	stdoutFormatter := formatter.NewKeyValue(template, "=", " ")
	consoleHandler := handler.NewConsoleHandler(level.Debug, level.Notice, stdoutFormatter)
	applicationLogger.AddHandler(consoleHandler)

	stderrFormatter := formatter.NewJSON(template, false)
	consoleErrorHandler := handler.NewConsoleErrorHandler(level.Warning, level.Null, stderrFormatter)
	applicationLogger.AddHandler(consoleErrorHandler)

	applicationLogger.Warning("message", "This message will be displayed.")

	applicationLogger.Debug("message", "This message will be displayed, because Level of custom logger set to Debug.")

	applicationLogger.Trace("message", "This message will not be displayed, because Trace has lower level than Debug.")

	consoleHandler.SetFromLevel(level.Trace)

	applicationLogger.Trace(map[string]interface{}{
		"message": "This message will be displayed, because Level has been changed to Trace.",
	})
}
