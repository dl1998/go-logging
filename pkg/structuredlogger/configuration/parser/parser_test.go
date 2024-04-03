package parser

import (
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/configuration/parser"
	"github.com/dl1998/go-logging/pkg/common/level"
	"io"
	"os"
	"testing"
)

var (
	name     = "test"
	template = map[string]interface{}{
		"level": "%(level)",
		"name":  "%(name)",
	}
	fromLevel         = level.Info
	toLevel           = level.Error
	testConfiguration = &parser.Configuration{
		Loggers: []parser.LoggerConfiguration{
			{
				Name: name,
				Handlers: []parser.HandlerConfiguration{
					createHandlerConfiguration("stdout", ""),
				},
			},
		},
	}
)

func createHandlerConfiguration(handlerType string, file string) parser.HandlerConfiguration {
	return parser.HandlerConfiguration{
		Type:      handlerType,
		FromLevel: fromLevel.String(),
		ToLevel:   toLevel.String(),
		File:      file,
		Formatter: parser.FormatterConfiguration{
			Type:        "json",
			Template:    template,
			PrettyPrint: false,
		},
	}
}

// TestParseFormatter_JSON tests that parseFormatter returns formatter.Interface.
func TestParseFormatter_JSON(t *testing.T) {
	formatterConfiguration := parser.FormatterConfiguration{
		Type:        "json",
		PrettyPrint: false,
		Template:    template,
	}
	formatter := parseFormatter(formatterConfiguration)

	testutils.AssertNotNil(t, formatter)
	testutils.AssertEquals(t, convertMap(template), formatter.Template())
}

// TestParseFormatter_KeyValue tests that parseFormatter returns formatter.Interface.
func TestParseFormatter_KeyValue(t *testing.T) {
	formatterConfiguration := parser.FormatterConfiguration{
		Type:              "key-value",
		KeyValueDelimiter: "=",
		PairSeparator:     ",",
		Template:          template,
	}
	formatter := parseFormatter(formatterConfiguration)

	testutils.AssertNotNil(t, formatter)
	testutils.AssertEquals(t, convertMap(template), formatter.Template())
}

// TestParseFormatter_Default tests that parseFormatter panics if unknown
// formatter type was provided.
func TestParseFormatter_Default(t *testing.T) {
	defer func() {
		if recovery := recover(); recovery != nil {
			testutils.AssertNotNil(t, recovery)
		}
	}()

	parseFormatter(parser.FormatterConfiguration{})
}

// BenchmarkParseFormatter benchmarks the parseFormatter function.
func BenchmarkParseFormatter(b *testing.B) {
	formatter := testConfiguration.Loggers[0].Handlers[0].Formatter
	for index := 0; index < b.N; index++ {
		parseFormatter(formatter)
	}
}

// TestParseHandler tests that parseHandler returns handler.Interface.
func TestParseHandler(t *testing.T) {
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
			handler := parseHandler(createHandlerConfiguration(test.handlerType, ""))

			testutils.AssertNotNil(t, handler)
			testutils.AssertNotNil(t, handler.Formatter())
			testutils.AssertEquals(t, test.expectedWriter, handler.Writer())
			testutils.AssertEquals(t, fromLevel, handler.FromLevel())
			testutils.AssertEquals(t, toLevel, handler.ToLevel())
		})
	}
}

// TestParseHandler_File tests that parseHandler returns handler.Interface with
// file writer.
func TestParseHandler_File(t *testing.T) {
	handler := parseHandler(createHandlerConfiguration("file", "/dev/null"))

	testutils.AssertNotNil(t, handler)
	testutils.AssertNotNil(t, handler.Formatter())
	testutils.AssertNotNil(t, handler.Writer())
	testutils.AssertEquals(t, fromLevel, handler.FromLevel())
	testutils.AssertEquals(t, toLevel, handler.ToLevel())
}

// TestParseHandler_File_Error tests that parseHandler panics if empty string was
// provided for file handler.
func TestParseHandler_File_Error(t *testing.T) {
	defer func() {
		if recovery := recover(); recovery != nil {
			testutils.AssertNotNil(t, recovery)
		}
	}()

	parseHandler(createHandlerConfiguration("file", ""))
}

// TestParseHandler_Default tests that parseHandler returns nil if unknown
// handler type was provided.
func TestParseHandler_Default(t *testing.T) {
	handler := parseHandler(createHandlerConfiguration("", ""))

	testutils.AssertNil(t, handler)
}

// BenchmarkParseHandler benchmarks the parseHandler function.
func BenchmarkParseHandler(b *testing.B) {
	handler := testConfiguration.Loggers[0].Handlers[0]
	for index := 0; index < b.N; index++ {
		parseHandler(handler)
	}
}

// TestParseLogger tests that parseLogger returns logger.Logger.
func TestParseLogger(t *testing.T) {
	logger := parseLogger(testConfiguration.Loggers[0])

	testutils.AssertNotNil(t, logger)
	testutils.AssertEquals(t, name, logger.Name())
	testutils.AssertEquals(t, len(testConfiguration.Loggers), len(logger.Handlers()))
}

// BenchmarkParseLogger benchmarks the parseLogger function.
func BenchmarkParseLogger(b *testing.B) {
	logger := testConfiguration.Loggers[0]
	for index := 0; index < b.N; index++ {
		parseLogger(logger)
	}
}

// TestParseAsyncLogger tests that parseAsyncLogger returns logger.AsyncLogger.
func TestParseAsyncLogger(t *testing.T) {
	logger := parseAsyncLogger(testConfiguration.Loggers[0])

	testutils.AssertNotNil(t, logger)
	testutils.AssertEquals(t, name, logger.Name())
	testutils.AssertEquals(t, len(testConfiguration.Loggers), len(logger.Handlers()))
}

// BenchmarkParseAsyncLogger benchmarks the parseAsyncLogger function.
func BenchmarkParseAsyncLogger(b *testing.B) {
	logger := testConfiguration.Loggers[0]
	for index := 0; index < b.N; index++ {
		parseAsyncLogger(logger)
	}
}

// TestGetLogger tests that GetLogger returns logger.Logger by name from the
// configuration.
func TestGetLogger(t *testing.T) {
	logger := GetLogger(name, *testConfiguration)

	testutils.AssertNotNil(t, logger)
	testutils.AssertEquals(t, name, logger.Name())
}

// TestGetLogger_Default tests that GetLogger returns nil if logger with the
// name was not found.
func TestGetLogger_Default(t *testing.T) {
	logger := GetLogger("", *testConfiguration)

	testutils.AssertNil(t, logger)
}

// BenchmarkGetLogger benchmarks the GetLogger function.
func BenchmarkGetLogger(b *testing.B) {
	for index := 0; index < b.N; index++ {
		_ = GetLogger(name, *testConfiguration)
	}
}

// TestGetAsyncLogger tests that GetAsyncLogger returns logger.AsyncLogger by
// name from the configuration.
func TestGetAsyncLogger(t *testing.T) {
	logger := GetAsyncLogger(name, *testConfiguration)

	testutils.AssertNotNil(t, logger)
	testutils.AssertEquals(t, name, logger.Name())
}

// TestGetAsyncLogger_Default tests that GetAsyncLogger returns nil if logger
// with the name was not found.
func TestGetAsyncLogger_Default(t *testing.T) {
	logger := GetAsyncLogger("", *testConfiguration)

	testutils.AssertNil(t, logger)
}

// BenchmarkGetAsyncLogger benchmarks the GetAsyncLogger function.
func BenchmarkGetAsyncLogger(b *testing.B) {
	for index := 0; index < b.N; index++ {
		_ = GetAsyncLogger(name, *testConfiguration)
	}
}
