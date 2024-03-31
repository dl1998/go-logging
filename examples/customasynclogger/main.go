// Example that shows how to create and use custom async logger.
package main

import (
	"fmt"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"time"
)

func main() {
	asyncQueueSize := 10

	applicationLogger := logger.NewAsyncLogger("example", time.DateTime, asyncQueueSize)

	applicationFormatter := formatter.New("%(datetime) [%(level)] %(message)")
	consoleHandler := handler.NewConsoleHandler(level.Debug, level.Null, applicationFormatter)
	applicationLogger.AddHandler(consoleHandler)

	for index := 0; index < asyncQueueSize; index++ {
		applicationLogger.Warning("This message will be displayed.")
	}

	fmt.Println("This will be printed before the last warning log message.")

	applicationLogger.WaitToFinishLogging()
}
