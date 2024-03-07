// Package formatter contains formatter that interpolates template strings and
// formats them.
package formatter

import (
	"github.com/dl1998/go-logging/pkg/logger/loglevel"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// logLevelColors maps LogLevel values to ANSI color codes.
var logLevelColors = map[loglevel.LogLevel]string{
	loglevel.Trace:     "\033[90m",         // Dark Grey
	loglevel.Debug:     "\033[36m",         // Cyan
	loglevel.Verbose:   "\033[96m",         // Light Cyan
	loglevel.Info:      "\033[97m",         // Default terminal text color (ANSI Bright White)
	loglevel.Notice:    "\033[94m",         // Light Blue
	loglevel.Warning:   "\033[93m",         // Bright Yellow
	loglevel.Severe:    "\033[38;5;208m",   // Orange
	loglevel.Error:     "\033[31m",         // Red
	loglevel.Alert:     "\033[38;5;202m",   // Dark Orange
	loglevel.Critical:  "\033[1;31m",       // Red or magenta (ANSI Bright Magenta)
	loglevel.Emergency: "\033[97m\033[41m", // Bright White on Red Background
}

// Reset color
const resetColor = "\033[0m"

// Interface represents interface that shall be satisfied by Formatter.
type Interface interface {
	Template() string
	Format(message string, loggerName string, level loglevel.LogLevel, colored bool) string
}

// Formatter struct that contains necessary for the formatting fields.
type Formatter struct {
	template string
}

// New create a new instance of the Formatter.
func New(template string) *Formatter {
	return &Formatter{template: template}
}

// IsEqual checks that two formatters are the same and returns result of the
// comparison.
func (formatter *Formatter) IsEqual(anotherFormatter *Formatter) bool {
	return formatter.template == anotherFormatter.template
}

// EvaluatePreset evaluates pre-defined set of formatting options and returns map
// with mapping of the option to interpolated value.
func (formatter *Formatter) EvaluatePreset(message string, loggerName string, level loglevel.LogLevel) map[string]string {
	_, functionName, functionLine, _ := runtime.Caller(2)
	var presets = map[string]string{
		"%(name)":      loggerName,                                // Logger name
		"%(message)":   message,                                   // Logged message
		"%(time)":      time.Now().Format(time.TimeOnly),          // Current time (format: HH:MM:ss)
		"%(date)":      time.Now().Format(time.DateOnly),          // Current date (format: yyyy-mm-dd)
		"%(isotime)":   time.Now().Format(time.RFC3339),           // Current date and time (format: yyyy-mm-ddTHH:MM:ssGMT)
		"%(timestamp)": strconv.FormatInt(time.Now().Unix(), 10),  // Current timestamp
		"%(level)":     level.String(),                            // Logging level name
		"%(levelnr)":   strconv.Itoa(level.DigitRepresentation()), // Logging level number
		"%(fname)":     functionName,                              // Name of the function from which logger has been called
		"%(fline)":     strconv.Itoa(functionLine),                // Line number from which logger has been called
	}
	return presets
}

// Template returns template string used by formatter.
func (formatter *Formatter) Template() string {
	return formatter.template
}

// Format formats provided message template to the interpolated string.
func (formatter *Formatter) Format(message string, loggerName string, level loglevel.LogLevel, colored bool) string {
	var presets = formatter.EvaluatePreset(message, loggerName, level)

	format := formatter.template

	for key, value := range presets {
		format = strings.ReplaceAll(format, key, value)
	}

	if colored {
		format = logLevelColors[level] + format + resetColor
	}

	return format + "\n"
}
