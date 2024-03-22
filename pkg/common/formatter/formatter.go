// Package formatter provides common formatting methods for the loggers.
package formatter

import (
	"github.com/dl1998/go-logging/pkg/common/logrecord"
)

// ParseKey parses the key and returns the value.
func ParseKey(key string, record logrecord.Interface) interface{} {
	var value interface{}

	switch key {
	case "%(name)":
		value = record.Name()
	case "%(level)":
		value = record.Level().String()
	case "%(levelnr)":
		value = record.Level().DigitRepresentation()
	case "%(datetime)":
		value = record.Time()
	case "%(timestamp)":
		value = record.Timestamp()
	case "%(fname)":
		value = record.FileName()
	case "%(fline)":
		value = record.FileLine()
	default:
		value = key
	}

	return value
}
