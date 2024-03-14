// Example that shows how to create and use custom logger.
package main

import (
	"github.com/dl1998/go-logging/pkg/logger"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"github.com/dl1998/go-logging/pkg/logger/loglevel"
)

func main() {
	applicationLogger := logger.New("example")

	applicationLogger.Debug("This message will not be displayed, because there are no handlers registered.")

	applicationFormatter := formatter.New("%(isotime) [%(level)] %(message)")
	consoleHandler := handler.NewConsoleHandler(loglevel.Debug, loglevel.Null, applicationFormatter)
	applicationLogger.AddHandler(consoleHandler)

	applicationLogger.Warning("This message will be displayed.")

	applicationLogger.Debug("This message will be displayed, because LogLevel of custom logger set to Debug.")

	applicationLogger.Trace("This message will not be displayed, because Trace has lower level than Debug.")

	consoleHandler.SetFromLevel(loglevel.Trace)

	applicationLogger.Trace("This message will be displayed, because LogLevel has been changed to Trace.")
}
