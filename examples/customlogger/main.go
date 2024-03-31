// Example that shows how to create and use custom logger.
package main

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"time"
)

func main() {
	applicationLogger := logger.New("example", time.RFC3339)

	applicationLogger.Debug("This message will not be displayed, because there are no handlers registered.")

	applicationFormatter := formatter.New("%(datetime) [%(level)] %(message)")
	consoleHandler := handler.NewConsoleHandler(level.Debug, level.Null, applicationFormatter)
	applicationLogger.AddHandler(consoleHandler)

	applicationLogger.Warning("This message will be displayed.")

	applicationLogger.Debug("This message will be displayed, because Level of custom logger set to Debug.")

	applicationLogger.Trace("This message will not be displayed, because Trace has lower level than Debug.")

	consoleHandler.SetFromLevel(level.Trace)

	applicationLogger.Trace("This message will be displayed, because Level has been changed to Trace.")
}
