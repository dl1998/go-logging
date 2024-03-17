// Package formatter provides formatters for the structured logger, it contains logic for the formatting messages.
package formatter

import (
	"encoding/json"
	commonformatter "github.com/dl1998/go-logging/pkg/common/formatter"
	"github.com/dl1998/go-logging/pkg/common/level"
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

type baseInterface interface {
	Template() map[string]string
	Format(loggerName string, logLevel level.Level, parameters ...any) map[string]interface{}
}

type baseFormatter struct {
	template map[string]string
}

func (formatter *baseFormatter) Template() map[string]string {
	return formatter.template
}

func (formatter *baseFormatter) Format(loggerName string, logLevel level.Level, parameters ...any) map[string]interface{} {
	var presets = commonformatter.EvaluatePreset(loggerName, logLevel, 3)

	format := make(map[string]interface{})

	for key, value := range formatter.template {
		if presetValue, ok := presets[value]; ok {
			format[key] = presetValue
		} else {
			format[key] = value
		}
	}

	var key string

	for index := 0; index < len(parameters); index++ {
		if index%2 == 0 {
			key = parameters[index].(string)
		} else {
			format[key] = parameters[index]
		}
	}

	return format
}

type Interface interface {
	Template() map[string]string
	Format(loggerName string, logLevel level.Level, colored bool, parameters ...any) string
}

type JSONFormatter struct {
	*baseFormatter
	pretty bool
}

func NewJSON(template map[string]string, pretty bool) *JSONFormatter {
	return &JSONFormatter{
		baseFormatter: &baseFormatter{
			template: template,
		},
		pretty: pretty,
	}
}

func (formatter *JSONFormatter) Format(loggerName string, logLevel level.Level, colored bool, parameters ...any) string {
	var format = formatter.baseFormatter.Format(loggerName, logLevel, parameters...)

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
		formattedString = logLevelColors[logLevel] + formattedString + resetColor
	} else if colored && formatter.pretty {
		formattedString = logLevelColors[logLevel] + strings.ReplaceAll(formattedString, "\n", resetColor+"\n"+logLevelColors[logLevel]) + resetColor
	}

	return formattedString + "\n"
}

type KeyValueFormatter struct {
	*baseFormatter
	keyValueDelimiter string
	pairSeparator     string
}

func NewKeyValue(template map[string]string, keyValueDelimiter string, pairSeparator string) *KeyValueFormatter {
	return &KeyValueFormatter{
		baseFormatter: &baseFormatter{
			template: template,
		},
		keyValueDelimiter: keyValueDelimiter,
		pairSeparator:     pairSeparator,
	}
}

func (formatter *KeyValueFormatter) Format(loggerName string, logLevel level.Level, colored bool, parameters ...any) string {
	var format = formatter.baseFormatter.Format(loggerName, logLevel, parameters...)

	var result strings.Builder

	var keys = make([]string, 0, len(format))

	for key, _ := range format {
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
		formattedString = logLevelColors[logLevel] + formattedString + resetColor
	}

	return formattedString + "\n"
}
