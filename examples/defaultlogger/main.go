// Example that shows how to use default logger.
package main

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger"
)

func main() {
	logger.Debug("This message will not be displayed, as default Level is Warning.")

	logger.Warning("This message will be displayed.")

	logger.Configure(logger.NewConfiguration(logger.WithFromLevel(level.All)))

	logger.Debug("This message will be displayed, because Level has been changed.")
}
