package loglevel

type LogLevel int

const (
	None = iota * 5
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
)

func (level *LogLevel) String() string {
	mapping := map[LogLevel]string{
		None:      "none",
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
	}
	return mapping[*level]
}

func (level *LogLevel) DigitRepresentation() int {
	return int(*level)
}