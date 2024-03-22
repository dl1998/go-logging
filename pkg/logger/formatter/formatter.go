// Package formatter contains formatter that interpolates template strings and
// formats them.
package formatter

import (
	commonformatter "github.com/dl1998/go-logging/pkg/common/formatter"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger/logrecord"
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
	Format(record logrecord.Interface, colored bool) string
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
func (formatter *Formatter) Format(record logrecord.Interface, colored bool) string {
	format := ParseTemplate(formatter.template, record)

	if colored {
		format = logLevelColors[record.Level()] + format + resetColor
	}

	return format + "\n"
}

// ParseTemplate parses template string and replaces keys with values from the log record.
func ParseTemplate(format string, record logrecord.Interface) string {
	format = ReplaceKey(format, "%(name)", record)
	format = ReplaceKey(format, "%(level)", record)
	format = ReplaceKey(format, "%(levelnr)", record)
	format = ReplaceKey(format, "%(datetime)", record)
	format = ReplaceKey(format, "%(timestamp)", record)
	format = ReplaceKey(format, "%(fname)", record)
	format = ReplaceKey(format, "%(fline)", record)
	format = ReplaceKey(format, "%(message)", record)
	return format
}

// ReplaceKey replaces key with value from the log record.
func ReplaceKey(format string, key string, record logrecord.Interface) string {
	if strings.Contains(format, key) {
		value := ParseKey(key, record)
		return strings.ReplaceAll(format, key, value)
	}
	return format
}

// ParseKey parses the key and returns the value.
func ParseKey(key string, record logrecord.Interface) string {
	switch key {
	case "%(message)":
		return record.Message()
	default:
		value := commonformatter.ParseKey(key, record)
		switch convertedValue := value.(type) {
		case int64:
			return strconv.FormatInt(convertedValue, 10)
		case int:
			return strconv.Itoa(convertedValue)
		case string:
			return convertedValue
		default:
			return value.(string)
		}
	}
}
