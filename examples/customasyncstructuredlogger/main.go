// Example that shows how to create and use custom async structured logger.
package main

import (
	"fmt"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/structuredlogger"
	"github.com/dl1998/go-logging/pkg/structuredlogger/formatter"
	"github.com/dl1998/go-logging/pkg/structuredlogger/handler"
	"time"
)

func main() {
	asyncQueueSize := 10

	applicationLogger := structuredlogger.NewAsyncStructuredLogger("example", time.DateTime, asyncQueueSize)

	applicationFormatter := formatter.NewJSON(map[string]string{
		"time":  "%(datetime)",
		"level": "%(level)",
		"name":  "%(name)",
	}, false)
	consoleHandler := handler.NewConsoleHandler(level.Debug, level.Null, applicationFormatter)
	applicationLogger.AddHandler(consoleHandler)

	for index := 0; index < asyncQueueSize; index++ {
		applicationLogger.Warning("message", "This message will be displayed.")
	}

	fmt.Println("This will be printed before the last warning log message.")

	applicationLogger.WaitToFinishLogging()
}
