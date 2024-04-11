// Package level contains Level definition with supported levels.
package level

type Level int

const step = 5

const (
	All = Level(iota * step)
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

var mapping = map[Level]string{
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

// ParseLevel returns Level from string.
func ParseLevel(level string) Level {
	for levelType, levelName := range mapping {
		if levelName == level {
			return levelType
		}
	}
	return Null
}

// String returns string representation of the Level.
func (level Level) String() string {
	return mapping[level]
}

// DigitRepresentation returns digit representations of the Level.
func (level Level) DigitRepresentation() int {
	return int(level)
}

// Next returns next Level.
func (level Level) Next() Level {
	if level == Null {
		return level
	}
	return Level(level.DigitRepresentation() + step)
}

// Previous returns previous Level.
func (level Level) Previous() Level {
	if level == All {
		return level
	}
	return Level(level.DigitRepresentation() - step)
}
