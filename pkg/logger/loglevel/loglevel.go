// Package loglevel contains LogLevel definition with supported levels.
package loglevel

type LogLevel int

const step = 5

const (
	All = LogLevel(iota * step)
	Trace
	Debug
	Verbose
	Info
	Notice
	Warning
	Severe
	Error
	Alert
	Critical
	Emergency
	Null
)

// String returns string representation of the LogLevel.
func (level LogLevel) String() string {
	mapping := map[LogLevel]string{
		All:       "all",
		Trace:     "trace",
		Debug:     "debug",
		Verbose:   "verbose",
		Info:      "info",
		Notice:    "notice",
		Warning:   "warning",
		Severe:    "severe",
		Error:     "error",
		Alert:     "alert",
		Critical:  "critical",
		Emergency: "emergency",
		Null:      "null",
	}
	return mapping[level]
}

// DigitRepresentation returns digit representations of the LogLevel.
func (level LogLevel) DigitRepresentation() int {
	return int(level)
}

// Next returns next LogLevel.
func (level LogLevel) Next() LogLevel {
	if level == Null {
		return level
	}
	return LogLevel(level.DigitRepresentation() + step)
}

// Previous returns previous LogLevel.
func (level LogLevel) Previous() LogLevel {
	if level == All {
		return level
	}
	return LogLevel(level.DigitRepresentation() - step)
}
