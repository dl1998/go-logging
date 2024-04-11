package parser

import (
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/configuration/parser"
	"github.com/dl1998/go-logging/pkg/common/level"
	"io"
	"os"
	"path"
	"testing"
	"time"
)

var (
	name       = "test"
	template   = "testTemplate"
	fromLevel  = level.Info
	toLevel    = level.Error
	testParser = &Parser{
		configuration: &parser.Configuration{
			Loggers: []parser.LoggerConfiguration{
				{
					Name: name,
					Handlers: []parser.HandlerConfiguration{
						createHandlerConfiguration("stdout", ""),
					},
				},
			},
		},
	}
	testDataParser = &Parser{
		configuration: &parser.Configuration{
			Loggers: []parser.LoggerConfiguration{
				{
					Name:             "test-logger",
					TimeFormat:       time.DateTime,
					MessageQueueSize: 100,
					Handlers: []parser.HandlerConfiguration{
						{
							Type:      "stdout",
							FromLevel: level.All.String(),
							ToLevel:   level.Severe.String(),
							Formatter: parser.FormatterConfiguration{
								Type:        "json",
								PrettyPrint: false,
								Template: parser.TemplateConfiguration{
									StringValue: "%(datetime) - %(level) - %(message)",
									MapValue: map[string]string{
										"timestamp": "%(datetime)",
										"level":     "%(level)",
										"name":      "%(name)",
									},
								},
							},
						},
						{
							Type:      "stderr",
							FromLevel: level.Error.String(),
							ToLevel:   level.Null.String(),
							Formatter: parser.FormatterConfiguration{
								Type:              "key-value",
								PairSeparator:     " ",
								KeyValueDelimiter: ":",
								Template: parser.TemplateConfiguration{
									StringValue: "%(datetime) - %(level) - %(message)",
									MapValue: map[string]string{
										"timestamp": "%(datetime)",
										"level":     "%(level)",
										"name":      "%(name)",
									},
								},
							},
						},
					},
				},
			},
		},
	}
)

// createHandlerConfiguration creates a new instance of the
// parser.HandlerConfiguration.
func createHandlerConfiguration(handlerType string, file string) parser.HandlerConfiguration {
	return parser.HandlerConfiguration{
		Type:      handlerType,
		FromLevel: fromLevel.String(),
		ToLevel:   toLevel.String(),
		File:      file,
		Formatter: parser.FormatterConfiguration{
			Type: "formatterType",
			Template: parser.TemplateConfiguration{
				StringValue: parser.EscapedString(template),
			},
		},
	}
}

// TestNewParser tests that NewParser returns a new instance of the Parser.
func TestNewParser(t *testing.T) {
	expected := parser.Configuration{}

	newParser := NewParser(expected)

	testutils.AssertNotNil(t, newParser)
	testutils.AssertEquals(t, &expected, newParser.configuration)
}

// BenchmarkNewParser benchmarks the NewParser function.
func BenchmarkNewParser(b *testing.B) {
	configuration := parser.Configuration{}
	for index := 0; index < b.N; index++ {
		_ = NewParser(configuration)
	}
}

// TestParseFile tests that parseFile parses the file with the given parser
// function and returns the Parser.
func TestParseFile(t *testing.T) {
	configuration := parser.Configuration{}

	newParser, err := parseFile(
		"example.json",
		func(string) (*parser.Configuration, error) {
			return &configuration, nil
		},
	)

	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, newParser)
	testutils.AssertEquals(t, &configuration, newParser.configuration)
}

// TestParseFile_Error tests that parseFile returns an error if the parser
// function fails.
func TestParseFile_Error(t *testing.T) {
	newParser, err := parseFile(
		"example.json",
		func(string) (*parser.Configuration, error) {
			return nil, fmt.Errorf("error")
		},
	)

	testutils.AssertNotNil(t, err)
	testutils.AssertNil(t, newParser)
}

// BenchmarkParseFile benchmarks the parseFile function.
func BenchmarkParseFile(b *testing.B) {
	configuration := parser.Configuration{}
	for index := 0; index < b.N; index++ {
		_, _ = parseFile("example.json", func(string) (*parser.Configuration, error) {
			return &configuration, nil
		})
	}
}

// TestParseJSON tests that ParseJSON parses JSON file and return an instance of
// the Parser.
func TestParseJSON(t *testing.T) {
	testJSON := path.Join(testutils.TestDataPath, "example.json")

	newParser, err := ParseJSON(testJSON)

	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, newParser)
	testutils.AssertEquals(t, testDataParser.configuration, newParser.configuration)
}

// BenchmarkParseJSON benchmarks the ParseJSON function.
func BenchmarkParseJSON(b *testing.B) {
	for index := 0; index < b.N; index++ {
		_, _ = ParseJSON(path.Join(testutils.TestDataPath, "example.json"))
	}
}

// TestParseYAML tests that ParseYAML parses YAML file and return an instance of
// the Parser.
func TestParseYAML(t *testing.T) {
	testYAML := path.Join(testutils.TestDataPath, "example.yaml")
	newParser, err := ParseYAML(testYAML)

	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, newParser)
	testutils.AssertEquals(t, testDataParser.configuration, newParser.configuration)
}

// BenchmarkParseYAML benchmarks the ParseYAML function.
func BenchmarkParseYAML(b *testing.B) {
	for index := 0; index < b.N; index++ {
		_, _ = ParseYAML(path.Join(testutils.TestDataPath, "example.yaml"))
	}
}

// TestParseXML tests that ParseXML parses XML file and return an instance of the
// Parser.
func TestParseXML(t *testing.T) {
	testXML := path.Join(testutils.TestDataPath, "example.xml")
	newParser, err := ParseXML(testXML)

	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, newParser)
	testutils.AssertEquals(t, testDataParser.configuration, newParser.configuration)
}

// BenchmarkParseXML benchmarks the ParseXML function.
func BenchmarkParseXML(b *testing.B) {
	for index := 0; index < b.N; index++ {
		_, _ = ParseXML(path.Join(testutils.TestDataPath, "example.xml"))
	}
}

// TestParser_ParseFormatter tests that Parser.parseFormatter returns
// formatter.Interface.
func TestParser_ParseFormatter(t *testing.T) {
	formatter := testParser.parseFormatter(testParser.configuration.Loggers[0].Handlers[0].Formatter)

	testutils.AssertNotNil(t, formatter)
	testutils.AssertEquals(t, template, formatter.Template())
}

// BenchmarkParser_ParseFormatter benchmarks the Parser.parseFormatter function.
func BenchmarkParser_ParseFormatter(b *testing.B) {
	formatter := testParser.configuration.Loggers[0].Handlers[0].Formatter
	for index := 0; index < b.N; index++ {
		testParser.parseFormatter(formatter)
	}
}

// TestParser_ParseHandler tests that Parser.parseHandler returns
// handler.Interface.
func TestParser_ParseHandler(t *testing.T) {
	tests := map[string]struct {
		handlerType    string
		expectedWriter io.Writer
	}{
		"Stdout": {
			handlerType:    "stdout",
			expectedWriter: os.Stdout,
		},
		"Stderr": {
			handlerType:    "stderr",
			expectedWriter: os.Stderr,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			handler := testParser.parseHandler(createHandlerConfiguration(test.handlerType, ""))

			testutils.AssertNotNil(t, handler)
			testutils.AssertNotNil(t, handler.Formatter())
			testutils.AssertEquals(t, test.expectedWriter, handler.Writer())
			testutils.AssertEquals(t, fromLevel, handler.FromLevel())
			testutils.AssertEquals(t, toLevel, handler.ToLevel())
		})
	}
}

// TestParser_ParseHandler_File tests that Parser.parseHandler returns
// handler.Interface with file writer.
func TestParser_ParseHandler_File(t *testing.T) {
	handler := testParser.parseHandler(createHandlerConfiguration("file", "/dev/null"))

	testutils.AssertNotNil(t, handler)
	testutils.AssertNotNil(t, handler.Formatter())
	testutils.AssertNotNil(t, handler.Writer())
	testutils.AssertEquals(t, fromLevel, handler.FromLevel())
	testutils.AssertEquals(t, toLevel, handler.ToLevel())
}

// TestParser_ParseHandler_File_Error tests that Parser.parseHandler panics if
// empty string was provided for file handler.
func TestParser_ParseHandler_File_Error(t *testing.T) {
	defer func() {
		if recovery := recover(); recovery != nil {
			testutils.AssertNotNil(t, recovery)
		}
	}()

	testParser.parseHandler(createHandlerConfiguration("file", ""))
}

// TestParser_ParseHandler_Default tests that Parser.parseHandler returns nil if
// unknown handler type was provided.
func TestParser_ParseHandler_Default(t *testing.T) {
	handler := testParser.parseHandler(createHandlerConfiguration("", ""))

	testutils.AssertNil(t, handler)
}

// BenchmarkParser_ParseHandler benchmarks the Parser.parseHandler function.
func BenchmarkParser_ParseHandler(b *testing.B) {
	handler := testParser.configuration.Loggers[0].Handlers[0]
	for index := 0; index < b.N; index++ {
		testParser.parseHandler(handler)
	}
}

// TestParser_ParseLogger tests that Parser.parseLogger returns logger.Logger.
func TestParser_ParseLogger(t *testing.T) {
	logger := testParser.parseLogger(testParser.configuration.Loggers[0])

	testutils.AssertNotNil(t, logger)
	testutils.AssertEquals(t, name, logger.Name())
	testutils.AssertEquals(t, len(testParser.configuration.Loggers), len(logger.Handlers()))
}

// BenchmarkParser_ParseLogger benchmarks the Parser.parseLogger function.
func BenchmarkParser_ParseLogger(b *testing.B) {
	logger := testParser.configuration.Loggers[0]
	for index := 0; index < b.N; index++ {
		testParser.parseLogger(logger)
	}
}

// TestParser_ParseAsyncLogger tests that Parser.parseAsyncLogger returns
// logger.AsyncLogger.
func TestParser_ParseAsyncLogger(t *testing.T) {
	logger := testParser.parseAsyncLogger(testParser.configuration.Loggers[0])

	testutils.AssertNotNil(t, logger)
	testutils.AssertEquals(t, name, logger.Name())
	testutils.AssertEquals(t, len(testParser.configuration.Loggers), len(logger.Handlers()))
}

// BenchmarkParser_ParseAsyncLogger benchmarks the parseAsyncLogger function.
func BenchmarkParser_ParseAsyncLogger(b *testing.B) {
	logger := testParser.configuration.Loggers[0]
	for index := 0; index < b.N; index++ {
		testParser.parseAsyncLogger(logger)
	}
}

// TestParser_GetLogger tests that Parser.GetLogger returns logger.Logger by name
// from the configuration.
func TestParser_GetLogger(t *testing.T) {
	logger := testParser.GetLogger(name)

	testutils.AssertNotNil(t, logger)
	testutils.AssertEquals(t, name, logger.Name())
}

// TestParser_GetLogger_Default tests that Parser.GetLogger returns nil if logger
// with the name was not found.
func TestParser_GetLogger_Default(t *testing.T) {
	logger := testParser.GetLogger("")

	testutils.AssertNil(t, logger)
}

// BenchmarkParser_GetLogger benchmarks the Parser.GetLogger function.
func BenchmarkParser_GetLogger(b *testing.B) {
	for index := 0; index < b.N; index++ {
		_ = testParser.GetLogger(name)
	}
}

// TestParser_GetAsyncLogger tests that Parser.GetAsyncLogger returns
// logger.AsyncLogger by name from the configuration.
func TestParser_GetAsyncLogger(t *testing.T) {
	logger := testParser.GetAsyncLogger(name)

	testutils.AssertNotNil(t, logger)
	testutils.AssertEquals(t, name, logger.Name())
}

// TestParser_GetAsyncLogger_Default tests that Parser.GetAsyncLogger returns nil
// if logger with the name was not found.
func TestParser_GetAsyncLogger_Default(t *testing.T) {
	logger := testParser.GetAsyncLogger("")

	testutils.AssertNil(t, logger)
}

// BenchmarkParser_GetAsyncLogger benchmarks the Parser.GetAsyncLogger function.
func BenchmarkParser_GetAsyncLogger(b *testing.B) {
	for index := 0; index < b.N; index++ {
		_ = testParser.GetAsyncLogger(name)
	}
}
