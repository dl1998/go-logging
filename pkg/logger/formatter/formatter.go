// Package formatter contains formatter that interpolates template strings and
// formats them.
package formatter

import (
	commonformatter "github.com/dl1998/go-logging/pkg/common/formatter"
	"github.com/dl1998/go-logging/pkg/common/level"
	"strconv"
	"strings"
)

// logLevelColors maps level.Level values to ANSI color codes.
var logLevelColors = map[level.Level]string{
	level.Trace:     "\033[90m",         // Dark Grey
	level.Debug:     "\033[36m",         // Cyan
	level.Verbose:   "\033[96m",         // Light Cyan
	level.Info:      "\033[97m",         // Default terminal text color (ANSI Bright White)
	level.Notice:    "\033[94m",         // Light Blue
	level.Warning:   "\033[93m",         // Bright Yellow
	level.Severe:    "\033[38;5;208m",   // Orange
	level.Error:     "\033[31m",         // Red
	level.Alert:     "\033[38;5;202m",   // Dark Orange
	level.Critical:  "\033[1;31m",       // Red or magenta (ANSI Bright Magenta)
	level.Emergency: "\033[97m\033[41m", // Bright White on Red Background
}

// Reset color
const resetColor = "\033[0m"

// Interface represents interface that shall be satisfied by Formatter.
type Interface interface {
	Template() string
	Format(message string, loggerName string, logLevel level.Level, colored bool) string
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

// Template returns template string used by formatter.
func (formatter *Formatter) Template() string {
	return formatter.template
}

// Format formats provided message template to the interpolated string.
func (formatter *Formatter) Format(message string, loggerName string, logLevel level.Level, colored bool) string {
	var presets = commonformatter.EvaluatePreset(loggerName, logLevel, 2)
	var convertedPresets = make(map[string]string, len(presets)+1)
	for key, value := range presets {
		switch convertedValue := value.(type) {
		case string:
			convertedPresets[key] = convertedValue
		case int:
			convertedPresets[key] = strconv.Itoa(convertedValue)
		case int64:
			convertedPresets[key] = strconv.FormatInt(convertedValue, 10)
		}
	}
	convertedPresets["%(message)"] = message

	format := formatter.template

	for key, value := range convertedPresets {
		format = strings.ReplaceAll(format, key, value)
	}

	if colored {
		format = logLevelColors[logLevel] + format + resetColor
	}

	return format + "\n"
}
