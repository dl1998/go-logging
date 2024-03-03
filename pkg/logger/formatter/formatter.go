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

type Formatter struct {
	format string
}

func New(format string) *Formatter {
	return &Formatter{format: format}
}

func (formatter *Formatter) IsEqual(anotherFormatter *Formatter) bool {
	return formatter.format == anotherFormatter.format
}

func (formatter *Formatter) EvaluatePreset(message string, loggerName string, level loglevel.LogLevel) map[string]string {
	_, functionName, functionLine, _ := runtime.Caller(2)
	var presets = map[string]string{
		"%(name)":      loggerName,
		"%(message)":   message,
		"%(time)":      time.Now().Format(time.TimeOnly),
		"%(date)":      time.Now().Format(time.DateOnly),
		"%(isotime)":   time.Now().Format(time.RFC3339),
		"%(timestamp)": strconv.FormatInt(time.Now().Unix(), 10),
		"%(level)":     level.String(),
		"%(levelnr)":   strconv.Itoa(level.DigitRepresentation()),
		"%(fname)":     functionName,
		"%(fline)":     strconv.Itoa(functionLine),
	}
	return presets
}

func (formatter *Formatter) Format(message string, loggerName string, level loglevel.LogLevel, colored bool) string {
	var presets = formatter.EvaluatePreset(message, loggerName, level)

	format := formatter.format

	for key, value := range presets {
		format = strings.ReplaceAll(format, key, value)
	}

	if colored {
		format = logLevelColors[level] + format + resetColor
	}

	return format + "\n"
}
