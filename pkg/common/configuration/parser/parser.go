package parser

import (
	"encoding/json"
	"encoding/xml"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	readFile = os.ReadFile
)

// FormatterConfiguration is a struct that represents the configuration of a formatter.
type FormatterConfiguration struct {
	// Type is the type of the formatter.
	Type string `json:"type" yaml:"type" xml:"type"`
	// PrettyPrint is a flag used by json formatter that indicates whether the
	// formatter should pretty print the output.
	PrettyPrint bool `json:"pretty-print" yaml:"pretty-print" xml:"pretty-print"`
	// KeyValueDelimiter is a delimiter used by key-value formatter to separate key
	// and value.
	KeyValueDelimiter string `json:"key-value-delimiter" yaml:"key-value-delimiter" xml:"key-value-delimiter"`
	// PairSeparator is a separator used by key-value formatter to separate key-value
	// pairs.
	PairSeparator string `json:"pair-separator" yaml:"pair-separator" xml:"pair-separator"`
	// Template is a template used by the formatter.
	Template interface{} `json:"template" yaml:"template" xml:"template"`
}

// HandlerConfiguration is a struct that represents the configuration of a handler.
type HandlerConfiguration struct {
	// Type is the type of the handler.
	Type string `json:"type" yaml:"type" xml:"type"`
	// FromLevel is the level from which the handler should log messages.
	FromLevel string `json:"from-level" yaml:"from-level" xml:"from-level"`
	// ToLevel is the level to which the handler should log messages.
	ToLevel string `json:"to-level" yaml:"to-level" xml:"to-level"`
	// File is the file used by the handler, it is needed for file handler to specify
	// where to write logs.
	File string `json:"file" yaml:"file" xml:"file"`
	// Formatter is the formatter used by the handler to format log messages.
	Formatter FormatterConfiguration `json:"formatter" yaml:"formatter" xml:"formatter"`
}

// LoggerConfiguration is a struct that represents the configuration of a logger.
type LoggerConfiguration struct {
	// Name is the name of the logger.
	Name string `json:"name" yaml:"name" xml:"name"`
	// TimeFormat is the time format used by the logger.
	TimeFormat string `json:"time-format" yaml:"time-format" xml:"time-format"`
	// MessageQueueSize is the size of the message queue used by async logger.
	MessageQueueSize int `json:"message-queue-size" yaml:"message-queue-size" xml:"message-queue-size"`
	// Handlers is the list of handlers used by the logger.
	Handlers []HandlerConfiguration `json:"handlers" yaml:"handlers" xml:"handlers>handler"`
}

// Configuration is a struct that represents the configuration.
type Configuration struct {
	// Loggers is the list of loggers present in the configuration.
	Loggers []LoggerConfiguration `json:"loggers" yaml:"loggers" xml:"loggers>logger"`
}

// readFromFile reads the configuration from the file and unmarshal it into the
// configuration struct.
func readFromFile(path string, unmarshal func([]byte, any) error) (*Configuration, error) {
	data, err := readFile(path)
	if err != nil {
		return nil, err
	}
	var configuration *Configuration
	err = unmarshal(data, &configuration)
	if err != nil {
		return nil, err
	}
	return configuration, nil
}

// ReadFromJSON reads the configuration from the JSON file and returns the
// configuration.
func ReadFromJSON(path string) (*Configuration, error) {
	return readFromFile(path, json.Unmarshal)
}

// ReadFromYAML reads the configuration from the YAML file and returns the
// configuration.
func ReadFromYAML(path string) (*Configuration, error) {
	return readFromFile(path, yaml.Unmarshal)
}

// ReadFromXML reads the configuration from the XML file and returns the
// configuration.
func ReadFromXML(path string) (*Configuration, error) {
	return readFromFile(path, xml.Unmarshal)
}
