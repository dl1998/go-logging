package formatter

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	"runtime"
	"time"
)

// EvaluatePreset evaluates pre-defined set of formatting options and returns map
// with mapping of the option to interpolated value.
func EvaluatePreset(loggerName string, logLevel level.Level, skipCaller int) map[string]interface{} {
	_, functionName, functionLine, _ := runtime.Caller(skipCaller)
	var presets = map[string]interface{}{
		"%(name)":      loggerName,                       // Logger name
		"%(time)":      time.Now().Format(time.TimeOnly), // Current time (format: HH:MM:ss)
		"%(date)":      time.Now().Format(time.DateOnly), // Current date (format: yyyy-mm-dd)
		"%(isotime)":   time.Now().Format(time.RFC3339),  // Current date and time (format: yyyy-mm-ddTHH:MM:ssGMT)
		"%(timestamp)": time.Now().Unix(),                // Current timestamp
		"%(level)":     logLevel.String(),                // Logging log level name
		"%(levelnr)":   logLevel.DigitRepresentation(),   // Logging log level number
		"%(fname)":     functionName,                     // Name of the function from which logger has been called
		"%(fline)":     functionLine,                     // Line number from which logger has been called
	}
	return presets
}
