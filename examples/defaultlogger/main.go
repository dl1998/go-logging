// Example that shows how to use default logger.
package main

import (
	"github.com/dl1998/go-logging/pkg/logger"
	"github.com/dl1998/go-logging/pkg/logger/loglevel"
)

func main() {
	logger.Debug("This message will not be displayed, as default LogLevel is Warning.")

	logger.Warning("This message will be displayed.")

	logger.SetLevel(loglevel.None)

	logger.Debug("This message will be displayed, because LogLevel has been changed.")
}
