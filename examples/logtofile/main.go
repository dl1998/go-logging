// Example that shows how to create logger that logs to file.
package main

import (
	"fmt"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"os"
	"time"
)

func main() {
	directory := "./tmp"

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.Mkdir(directory, 0777)

		if err != nil {
			fmt.Println("Cannot create 'tmp' directory.")
		}
	}

	applicationLogger := logger.New("file-logger", time.RFC3339)

	applicationFormatter := formatter.New("%(datetime) [%(level)] %(message)")
	fileHandler := handler.NewFileHandler(level.Warning, level.Null, applicationFormatter, fmt.Sprintf("%s/file.log", directory))
	applicationLogger.AddHandler(fileHandler)

	applicationLogger.Warning("This file has only Warning level logs or higher.")

	applicationLogger.Debug("This message will not be logged.")

	applicationLogger.Error("This message will be logged, because it has higher level.")
}
