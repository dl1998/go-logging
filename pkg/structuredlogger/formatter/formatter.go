// Package formatter provides formatters for the structured logger, it contains logic for the formatting messages.
package formatter

import (
	"encoding/json"
	commonFormatter "github.com/dl1998/go-logging/pkg/common/formatter"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/structuredlogger/logrecord"
	"sort"
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

// baseFormatter struct that contains necessary for the formatting fields.
type baseFormatter struct {
	// template contains key-value pairs with template for the formatter.
	template map[string]string
}

// Template returns template string used by formatter.
func (formatter *baseFormatter) Template() map[string]string {
	return formatter.template
}

// Format formats provided message template to the interpolated string.
func (formatter *baseFormatter) Format(record logrecord.Interface) map[string]interface{} {
	format := make(map[string]interface{})

	for key, value := range formatter.template {
		format[key] = commonFormatter.ParseKey(value, record)
	}

	for key, value := range record.Parameters() {
		if stringValue, ok := value.(string); ok {
			format[key] = commonFormatter.ParseKey(stringValue, record)
		} else {
			format[key] = value
		}
	}

	return format
}

// Interface represents interface that shall be satisfied by Formatter.
type Interface interface {
	Template() map[string]string
	Format(record logrecord.Interface, colored bool) string
}

// JSONFormatter struct that contains necessary for the formatting fields.
type JSONFormatter struct {
	*baseFormatter
	pretty bool
}

// NewJSON create a new instance of the JSONFormatter.
func NewJSON(template map[string]string, pretty bool) *JSONFormatter {
	return &JSONFormatter{
		baseFormatter: &baseFormatter{
			template: template,
		},
		pretty: pretty,
	}
}

// Format formats provided message template to the interpolated string.
func (formatter *JSONFormatter) Format(record logrecord.Interface, colored bool) string {
	var format = formatter.baseFormatter.Format(record)

	var data []byte
	var err error

	if formatter.pretty {
		data, err = json.MarshalIndent(format, "", "  ")
	} else {
		data, err = json.Marshal(format)
	}

	if err != nil {
		return ""
	}

	formattedString := string(data)

	if colored && !formatter.pretty {
		formattedString = logLevelColors[record.Level()] + formattedString + resetColor
	} else if colored && formatter.pretty {
		formattedString = logLevelColors[record.Level()] + strings.ReplaceAll(formattedString, "\n", resetColor+"\n"+logLevelColors[record.Level()]) + resetColor
	}

	return formattedString + "\n"
}

// KeyValueFormatter struct that contains necessary for the formatting fields.
type KeyValueFormatter struct {
	// baseFormatter is a base formatter.
	*baseFormatter
	// keyValueDelimiter is a delimiter between key and value.
	keyValueDelimiter string
	// pairSeparator is a separator between key-value pairs.
	pairSeparator string
}

// NewKeyValue create a new instance of the KeyValueFormatter.
func NewKeyValue(template map[string]string, keyValueDelimiter string, pairSeparator string) *KeyValueFormatter {
	return &KeyValueFormatter{
		baseFormatter: &baseFormatter{
			template: template,
		},
		keyValueDelimiter: keyValueDelimiter,
		pairSeparator:     pairSeparator,
	}
}

// Format formats provided message template to the interpolated string.
func (formatter *KeyValueFormatter) Format(record logrecord.Interface, colored bool) string {
	var format = formatter.baseFormatter.Format(record)

	var result strings.Builder

	var keys = make([]string, 0, len(format))

	for key := range format {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		value := format[key]
		result.WriteString(key)
		result.WriteString(formatter.keyValueDelimiter)
		switch convertedValue := value.(type) {
		case string:
			result.WriteString("\"")
			result.WriteString(convertedValue)
			result.WriteString("\"")
		case bool:
			result.WriteString(strconv.FormatBool(convertedValue))
		case int:
			result.WriteString(strconv.Itoa(convertedValue))
		case int64:
			result.WriteString(strconv.FormatInt(convertedValue, 10))
		case float64:
			result.WriteString(strconv.FormatFloat(convertedValue, 'f', -1, 64))
		case float32:
			result.WriteString(strconv.FormatFloat(float64(convertedValue), 'f', -1, 32))
		}
		result.WriteString(formatter.pairSeparator)
	}

	formattedString := strings.TrimSuffix(result.String(), formatter.pairSeparator)

	if colored {
		formattedString = logLevelColors[record.Level()] + formattedString + resetColor
		if formatter.pairSeparator == "\n" {
			formattedString = strings.ReplaceAll(formattedString, "\n", resetColor+"\n"+logLevelColors[record.Level()])
		}
	}

	return formattedString + "\n"
}
