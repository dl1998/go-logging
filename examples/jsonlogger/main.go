// Example that shows how to use default structured json logger.
package main

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	logger "github.com/dl1998/go-logging/pkg/structuredlogger"
)

func main() {
	logger.Debug("message", "This message will not be displayed, as default Level is Warning.")

	logger.Warning("message", "This message will be displayed.")

	logger.Configure(logger.NewConfiguration(logger.WithFromLevel(level.All)))

	logger.Debug("message", "This message will be displayed, because Level has been changed.")

	logger.Configure(logger.NewConfiguration(logger.WithPretty(true)))

	logger.Warning("message", "This message will be displayed in pretty format.")
}
