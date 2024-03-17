// Example that shows how to use default structured key-value logger.
package main

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	logger "github.com/dl1998/go-logging/pkg/structuredlogger"
)

func main() {
	format := logger.WithFormat("key-value")

	logger.Configure(logger.NewConfiguration(format))

	logger.Debug("message", "This message will not be displayed, as default Level is Warning.")

	logger.Warning("message", "This message will be displayed in standard key-value format.")

	logger.Configure(logger.NewConfiguration(format, logger.WithFromLevel(level.All), logger.WithKeyValueDelimiter(":"), logger.WithPairSeparator(", ")))

	logger.Debug("message", "This message will be displayed with ':' as delimiter and ', ' as pair separator.")

	logger.Configure(logger.NewConfiguration(format, logger.WithPairSeparator("\n")))

	logger.Warning("message", "This message will be displayed with each key on the new line.")
}
