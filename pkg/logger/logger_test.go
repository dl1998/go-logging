// Package logger_test has tests for logger package.
package logger

import (
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"testing"
	"time"
)

// TestNew tests that New creates a new logger.
func TestNew(t *testing.T) {
	newLogger := New(loggerName, timeFormat)

	testutils.AssertEquals(t, loggerName, newLogger.Name())

	handlersSize := len(newLogger.Handlers())

	testutils.AssertEquals(t, 0, handlersSize)
	testutils.AssertEquals(t, level.Error, newLogger.errorLevel)
	testutils.AssertEquals(t, level.Critical, newLogger.panicLevel)
}

// BenchmarkNew perform benchmarking of the New().
func BenchmarkNew(b *testing.B) {
	for index := 0; index < b.N; index++ {
		New(loggerName, timeFormat)
	}
}

// TestLogger_Name tests that Logger.Name returns loggerName of the logger.
func TestLogger_Name(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	testutils.AssertEquals(t, loggerName, newLogger.Name())
}

// BenchmarkLogger_Name perform benchmarking of the Logger.Name().
func BenchmarkLogger_Name(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Name()
	}
}

// TestLogger_Handlers tests that Logger.Handlers returns the list of handlers
// for the logger.
func TestLogger_Handlers(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	testutils.AssertEquals(t, make([]handler.Interface, 0), newLogger.Handlers())
}

// BenchmarkLogger_Handlers perform benchmarking of the Logger.Handlers().
func BenchmarkLogger_Handlers(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Handlers()
	}
}

// TestLogger_AddHandler tests that Logger.AddHandler adds a new Handler on the
// list of handlers.
func TestLogger_AddHandler(t *testing.T) {
	mockHandler := &MockHandler{}

	newLogger := &Logger{baseLogger: &baseLogger{
		name:     loggerName,
		handlers: make([]handler.Interface, 0),
	}}

	newLogger.AddHandler(mockHandler)

	testutils.AssertEquals(t, []handler.Interface{mockHandler}, newLogger.baseLogger.Handlers())
}

// BenchmarkLogger_AddHandler perform benchmarking of the Logger.AddHandler().
func BenchmarkLogger_AddHandler(b *testing.B) {
	mockHandler := &MockHandler{}

	newLogger := &Logger{baseLogger: &baseLogger{
		name:     loggerName,
		handlers: make([]handler.Interface, 0),
	}}

	for index := 0; index < b.N; index++ {
		newLogger.AddHandler(mockHandler)
	}
}

// TestLogger_RemoveHandler tests that Logger.RemoveHandler removes a Handler from the list of handlers.
func TestLogger_RemoveHandler(t *testing.T) {
	mockHandler1 := &MockHandler{}
	mockHandler2 := &MockHandler{}

	newLogger := &Logger{baseLogger: &baseLogger{
		name:     loggerName,
		handlers: []handler.Interface{mockHandler1, mockHandler2},
	}}

	newLogger.RemoveHandler(mockHandler1)
	newLogger.RemoveHandler(mockHandler2)

	testutils.AssertEquals(t, make([]handler.Interface, 0), newLogger.baseLogger.Handlers())
}

// BenchmarkLogger_RemoveHandler perform benchmarking of the Logger.RemoveHandler().
func BenchmarkLogger_RemoveHandler(b *testing.B) {
	mockHandler := &MockHandler{}

	newLogger := &Logger{baseLogger: &baseLogger{
		name:     loggerName,
		handlers: []handler.Interface{mockHandler},
	}}

	for index := 0; index < b.N; index++ {
		newLogger.RemoveHandler(mockHandler)
	}
}

// TestLogger_Trace tests that Logger.Trace logs message with parameters on trace
// level.
func TestLogger_Trace(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Trace(message, parameters...)

	testutils.AssertEquals(t, level.Trace, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Trace perform benchmarking of the Logger.Trace().
func BenchmarkLogger_Trace(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Trace(message, parameters...)
	}
}

// TestLogger_Debug tests that Logger.Debug logs message with parameters on debug
// level.
func TestLogger_Debug(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Debug(message, parameters...)

	testutils.AssertEquals(t, level.Debug, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Debug perform benchmarking of the Logger.Debug().
func BenchmarkLogger_Debug(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Debug(message, parameters...)
	}
}

// TestLogger_Verbose tests that Logger.Verbose logs message with parameters on
// verbose level.
func TestLogger_Verbose(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Verbose(message, parameters...)

	testutils.AssertEquals(t, level.Verbose, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Verbose perform benchmarking of the Logger.Verbose().
func BenchmarkLogger_Verbose(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Verbose(message, parameters...)
	}
}

// TestLogger_Info tests that Logger.Info logs message with parameters on info
// level.
func TestLogger_Info(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Info(message, parameters...)

	testutils.AssertEquals(t, level.Info, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Info perform benchmarking of the Logger.Info().
func BenchmarkLogger_Info(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Info(message, parameters...)
	}
}

// TestLogger_Notice tests that Logger.Notice logs message with parameters on
// notice level.
func TestLogger_Notice(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Notice(message, parameters...)

	testutils.AssertEquals(t, level.Notice, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Notice perform benchmarking of the Logger.Notice().
func BenchmarkLogger_Notice(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Notice(message, parameters...)
	}
}

// TestLogger_Warning tests that Logger.Warning logs message with parameters on
// warning level.
func TestLogger_Warning(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Warning(message, parameters...)

	testutils.AssertEquals(t, level.Warning, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Warning perform benchmarking of the Logger.Warning().
func BenchmarkLogger_Warning(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Warning(message, parameters...)
	}
}

// TestLogger_Severe tests that Logger.Severe logs message with parameters on
// severe level.
func TestLogger_Severe(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Severe(message, parameters...)

	testutils.AssertEquals(t, level.Severe, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Severe perform benchmarking of the Logger.Severe().
func BenchmarkLogger_Severe(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Severe(message, parameters...)
	}
}

// TestLogger_Error tests that Logger.Error logs message with parameters on error
// level.
func TestLogger_Error(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Error(message, parameters...)

	testutils.AssertEquals(t, level.Error, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Error perform benchmarking of the Logger.Error().
func BenchmarkLogger_Error(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Error(message, parameters...)
	}
}

// TestLogger_Alert tests that Logger.Alert logs message with parameters on alert
// level.
func TestLogger_Alert(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Alert(message, parameters...)

	testutils.AssertEquals(t, level.Alert, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Alert perform benchmarking of the Logger.Alert().
func BenchmarkLogger_Alert(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Alert(message, parameters...)
	}
}

// TestLogger_Critical tests that Logger.Critical logs message with parameters on
// critical level.
func TestLogger_Critical(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Critical(message, parameters...)

	testutils.AssertEquals(t, level.Critical, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Critical perform benchmarking of the Logger.Critical().
func BenchmarkLogger_Critical(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Critical(message, parameters...)
	}
}

// TestLogger_Emergency tests that Logger.Emergency logs message with parameters
// on emergency level.
func TestLogger_Emergency(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Emergency(message, parameters...)

	testutils.AssertEquals(t, level.Emergency, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Emergency perform benchmarking of the Logger.Emergency().
func BenchmarkLogger_Emergency(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Emergency(message, parameters...)
	}
}

// TestLogger_ErrorLevel tests that Logger.ErrorLevel returns the error level of
// the logger.
func TestLogger_ErrorLevel(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	testutils.AssertEquals(t, level.Error, newLogger.ErrorLevel())
}

// BenchmarkLogger_ErrorLevel perform benchmarking of the Logger.ErrorLevel().
func BenchmarkLogger_ErrorLevel(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	for index := 0; index < b.N; index++ {
		newLogger.ErrorLevel()
	}
}

// TestLogger_SetErrorLevel tests that Logger.SetErrorLevel sets the error level
// of the logger.
func TestLogger_SetErrorLevel(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	newLogger.SetErrorLevel(level.Warning)

	testutils.AssertEquals(t, level.Warning, newLogger.errorLevel)
}

// BenchmarkLogger_SetErrorLevel perform benchmarking of the Logger.SetErrorLevel().
func BenchmarkLogger_SetErrorLevel(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	for index := 0; index < b.N; index++ {
		newLogger.SetErrorLevel(level.Warning)
	}
}

// TestLogger_PanicLevel tests that Logger.PanicLevel returns the panic level of
// the logger.
func TestLogger_PanicLevel(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	testutils.AssertEquals(t, level.Critical, newLogger.PanicLevel())
}

// BenchmarkLogger_PanicLevel perform benchmarking of the Logger.PanicLevel().
func BenchmarkLogger_PanicLevel(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	for index := 0; index < b.N; index++ {
		newLogger.PanicLevel()
	}
}

// TestLogger_SetPanicLevel tests that Logger.SetPanicLevel sets the panic level
// of the logger.
func TestLogger_SetPanicLevel(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	newLogger.SetPanicLevel(level.Warning)

	testutils.AssertEquals(t, level.Warning, newLogger.panicLevel)
}

// BenchmarkLogger_SetPanicLevel perform benchmarking of the Logger.SetPanicLevel().
func BenchmarkLogger_SetPanicLevel(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	for index := 0; index < b.N; index++ {
		newLogger.SetPanicLevel(level.Warning)
	}
}

// TestLogger_RaiseError tests that Logger.RaiseError logs message with
// parameters on error level and returns a new error.
func TestLogger_RaiseError(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	expectedError := fmt.Errorf(message, parameters...)

	err := newLogger.RaiseError(message, parameters...)

	testutils.AssertEquals(t, level.Error, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
	testutils.AssertNotNil(t, err)
	testutils.AssertEquals(t, expectedError, err)
}

// BenchmarkLogger_RaiseError perform benchmarking of the Logger.RaiseError().
func BenchmarkLogger_RaiseError(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	for index := 0; index < b.N; index++ {
		_ = newLogger.RaiseError(message, parameters...)
	}
}

// TestLogger_CaptureError tests that Logger.CaptureError logs message from the
// error.
func TestLogger_CaptureError(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	err := fmt.Errorf(message, parameters...)

	newLogger.CaptureError(err)

	testutils.AssertEquals(t, level.Error, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, err.Error(), mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, make([]any, 0), mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_CaptureError perform benchmarking of the Logger.CaptureError().
func BenchmarkLogger_CaptureError(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	err := fmt.Errorf(message, parameters...)

	for index := 0; index < b.N; index++ {
		newLogger.CaptureError(err)
	}
}

// TestLogger_Panic tests that Logger.Panic logs message with parameters on panic
// level and panics.
func TestLogger_Panic(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic")
		} else {
			testutils.AssertEquals(t, fmt.Sprintf(message, parameters...), r.(string))
		}
	}()

	newLogger.Panic(message, parameters...)

	testutils.AssertNotNil(t, mockLogger.Parameters)
	testutils.AssertEquals(t, level.Critical, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Panic perform benchmarking of the Logger.Panic().
func BenchmarkLogger_Panic(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	for index := 0; index < b.N; index++ {
		func() {
			defer func() {
				if r := recover(); r == nil {
					return
				}
			}()

			newLogger.Panic(message, parameters...)
		}()
	}
}

// TestWithErrorLevel tests that WithErrorLevel sets the error level in the
// Configuration.
func TestWithErrorLevel(t *testing.T) {
	configuration := NewConfiguration()

	option := WithErrorLevel(logLevel)

	option(configuration)

	testutils.AssertEquals(t, logLevel, configuration.errorLevel)
}

// BenchmarkWithErrorLevel perform benchmarking of the WithErrorLevel().
func BenchmarkWithErrorLevel(b *testing.B) {
	configuration := NewConfiguration()

	option := WithErrorLevel(logLevel)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithPanicLevel tests that WithPanicLevel sets the panic level in the
// Configuration.
func TestWithPanicLevel(t *testing.T) {
	configuration := NewConfiguration()

	option := WithPanicLevel(logLevel)

	option(configuration)

	testutils.AssertEquals(t, logLevel, configuration.panicLevel)
}

// BenchmarkWithPanicLevel perform benchmarking of the WithPanicLevel().
func BenchmarkWithPanicLevel(b *testing.B) {
	configuration := NewConfiguration()

	option := WithPanicLevel(logLevel)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithFromLevel tests that WithFromLevel sets the from level in the
// Configuration.
func TestWithFromLevel(t *testing.T) {
	configuration := NewConfiguration()

	option := WithFromLevel(logLevel)

	option(configuration)

	testutils.AssertEquals(t, logLevel, configuration.fromLevel)
}

// BenchmarkWithFromLevel perform benchmarking of the WithFromLevel().
func BenchmarkWithFromLevel(b *testing.B) {
	configuration := NewConfiguration()

	option := WithFromLevel(logLevel)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithToLevel tests that WithToLevel sets the to level in the Configuration.
func TestWithToLevel(t *testing.T) {
	configuration := NewConfiguration()

	option := WithToLevel(logLevel)

	option(configuration)

	testutils.AssertEquals(t, logLevel, configuration.toLevel)
}

// BenchmarkWithToLevel perform benchmarking of the WithToLevel().
func BenchmarkWithToLevel(b *testing.B) {
	configuration := NewConfiguration()

	option := WithToLevel(logLevel)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithTemplate tests that WithTemplate sets the template in the
// Configuration.
func TestWithTemplate(t *testing.T) {
	configuration := NewConfiguration()

	option := WithTemplate(loggerTemplate)

	option(configuration)

	testutils.AssertEquals(t, loggerTemplate, configuration.template)
}

// BenchmarkWithTemplate perform benchmarking of the WithTemplate().
func BenchmarkWithTemplate(b *testing.B) {
	configuration := NewConfiguration()

	option := WithTemplate(loggerTemplate)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithFile tests that WithFile sets the file in the Configuration.
func TestWithFile(t *testing.T) {
	configuration := NewConfiguration()

	file := "file.log"

	option := WithFile(file)

	option(configuration)

	testutils.AssertEquals(t, file, configuration.file)
}

// BenchmarkWithFile perform benchmarking of the WithFile().
func BenchmarkWithFile(b *testing.B) {
	configuration := NewConfiguration()

	file := "file.log"

	option := WithFile(file)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithName tests that WithName sets the loggerName in the Configuration.
func TestWithName(t *testing.T) {
	configuration := NewConfiguration()

	name := "test"

	option := WithName(name)

	option(configuration)

	testutils.AssertEquals(t, name, configuration.name)
}

// BenchmarkWithName perform benchmarking of the WithName().
func BenchmarkWithName(b *testing.B) {
	configuration := NewConfiguration()

	name := "test"

	option := WithName(name)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithTimeFormat tests that WithTimeFormat sets the time format in the
// Configuration.
func TestWithTimeFormat(t *testing.T) {
	configuration := NewConfiguration()

	timeFormat := time.RFC3339

	option := WithTimeFormat(timeFormat)

	option(configuration)

	testutils.AssertEquals(t, timeFormat, configuration.timeFormat)
}

// BenchmarkWithTimeFormat perform benchmarking of the WithTimeFormat().
func BenchmarkWithTimeFormat(b *testing.B) {
	configuration := NewConfiguration()

	timeFormat := time.RFC3339

	option := WithTimeFormat(timeFormat)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestNewConfiguration tests that NewConfiguration creates a new Configuration.
func TestNewConfiguration(t *testing.T) {
	tests := map[string]struct {
		options            []Option
		expectedErrorLevel level.Level
		expectedPanicLevel level.Level
		expectedFromLevel  level.Level
		expectedToLevel    level.Level
		expectedTemplate   string
		expectedFile       string
		expectedName       string
		expectedTimeFormat string
	}{
		"Empty": {
			options:            []Option{},
			expectedErrorLevel: level.Error,
			expectedPanicLevel: level.Critical,
			expectedFromLevel:  level.Warning,
			expectedToLevel:    level.Null,
			expectedTemplate:   loggerTemplate,
			expectedFile:       "",
			expectedName:       "root",
			expectedTimeFormat: time.RFC3339,
		},
		"Non Standard": {
			options: []Option{
				WithErrorLevel(level.Warning),
				WithPanicLevel(level.Alert),
				WithFromLevel(level.All),
				WithToLevel(level.Emergency),
				WithTemplate("%(message):%(name):%(level)"),
				WithFile("file.log"),
				WithName("test"),
				WithTimeFormat(time.DateTime),
			},
			expectedErrorLevel: level.Warning,
			expectedPanicLevel: level.Alert,
			expectedFromLevel:  level.All,
			expectedToLevel:    level.Emergency,
			expectedTemplate:   "%(message):%(name):%(level)",
			expectedFile:       "file.log",
			expectedName:       "test",
			expectedTimeFormat: time.DateTime,
		},
	}
	for name, configuration := range tests {
		t.Run(name, func(t *testing.T) {
			newConfiguration := NewConfiguration(configuration.options...)

			testutils.AssertEquals(t, configuration.expectedErrorLevel, newConfiguration.errorLevel)
			testutils.AssertEquals(t, configuration.expectedPanicLevel, newConfiguration.panicLevel)
			testutils.AssertEquals(t, configuration.expectedFromLevel, newConfiguration.fromLevel)
			testutils.AssertEquals(t, configuration.expectedToLevel, newConfiguration.toLevel)
			testutils.AssertEquals(t, configuration.expectedTemplate, newConfiguration.template)
			testutils.AssertEquals(t, configuration.expectedFile, newConfiguration.file)
			testutils.AssertEquals(t, configuration.expectedName, newConfiguration.name)
			testutils.AssertEquals(t, configuration.expectedTimeFormat, newConfiguration.timeFormat)
		})
	}
}

// BenchmarkNewConfiguration perform benchmarking of the NewConfiguration().
func BenchmarkNewConfiguration(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NewConfiguration()
	}
}

// TestConfigure tests that Configure sets the configuration for the default
// logger.
func TestConfigure(t *testing.T) {
	configuration := NewConfiguration(
		WithFromLevel(level.All),
		WithToLevel(level.Emergency),
		WithTemplate("%(message):%(name):%(level)"),
		WithFile(""),
		WithName("test"),
	)

	Configure(configuration)

	testutils.AssertEquals(t, "test", rootLogger.Name())
	testutils.AssertEquals(t, 2, len(rootLogger.Handlers()))
}

// TestConfigure_IncorrectLevels tests that Configure returns an error when
// 'from' level is greater than 'to' level.
func TestConfigure_IncorrectLevels(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic when 'from' level was greater than 'to' level")
		}
	}()

	configuration := NewConfiguration(
		WithFromLevel(level.Warning),
		WithToLevel(level.Debug),
	)

	Configure(configuration)
}

// BenchmarkConfigure perform benchmarking of the Configure().
func BenchmarkConfigure(b *testing.B) {
	configuration := NewConfiguration(
		WithFromLevel(level.All),
		WithToLevel(level.Emergency),
		WithTemplate("%(message):%(name):%(level)"),
		WithFile(""),
		WithName("test"),
	)

	for index := 0; index < b.N; index++ {
		Configure(configuration)
	}
}

// TestName tests that Name returns loggerName of the default logger.
func TestName(t *testing.T) {
	Configure(NewConfiguration())

	testutils.AssertEquals(t, "root", Name())
}

// BenchmarkName perform benchmarking of the Name().
func BenchmarkName(b *testing.B) {
	Configure(NewConfiguration())

	for index := 0; index < b.N; index++ {
		Name()
	}
}

// TestTemplate tests that Template returns template of the default logger.
func TestTemplate(t *testing.T) {
	Configure(NewConfiguration())

	testutils.AssertEquals(t, loggerTemplate, Template())
}

// BenchmarkTemplate perform benchmarking of the Template().
func BenchmarkTemplate(b *testing.B) {
	Configure(NewConfiguration())

	for index := 0; index < b.N; index++ {
		Template()
	}
}

// TestFromLevel tests that FromLevel returns from level of the default logger.
func TestFromLevel(t *testing.T) {
	Configure(NewConfiguration())

	testutils.AssertEquals(t, level.Warning, FromLevel())
}

// BenchmarkFromLevel perform benchmarking of the FromLevel().
func BenchmarkFromLevel(b *testing.B) {
	Configure(NewConfiguration())

	for index := 0; index < b.N; index++ {
		FromLevel()
	}
}

// TestToLevel tests that ToLevel returns to level of the default logger.
func TestToLevel(t *testing.T) {
	Configure(NewConfiguration())

	testutils.AssertEquals(t, level.Null, ToLevel())
}

// BenchmarkToLevel perform benchmarking of the ToLevel().
func BenchmarkToLevel(b *testing.B) {
	Configure(NewConfiguration())

	for index := 0; index < b.N; index++ {
		ToLevel()
	}
}

// TestTrace tests that Trace logs message with parameters on trace level using
// default logger.
func TestTrace(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Trace(message, parameters...)

	testutils.AssertEquals(t, level.Trace, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkTrace perform benchmarking of the Trace().
func BenchmarkTrace(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Trace(message, parameters...)
	}
}

// TestDebug tests that Debug logs message with parameters on debug level using
// default logger.
func TestDebug(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Debug(message, parameters...)

	testutils.AssertEquals(t, level.Debug, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkDebug perform benchmarking of the Debug().
func BenchmarkDebug(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Debug(message, parameters...)
	}
}

// TestVerbose tests that Verbose logs message with parameters on verbose level
// using default logger.
func TestVerbose(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Verbose(message, parameters...)

	testutils.AssertEquals(t, level.Verbose, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkVerbose perform benchmarking of the Verbose().
func BenchmarkVerbose(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Verbose(message, parameters...)
	}
}

// TestInfo tests that Info logs message with parameters on info level using
// default logger.
func TestInfo(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Info(message, parameters...)

	testutils.AssertEquals(t, level.Info, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkInfo perform benchmarking of the Info().
func BenchmarkInfo(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Info(message, parameters...)
	}
}

// TestNotice tests that Notice logs message with parameters on notice level
// using default logger.
func TestNotice(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Notice(message, parameters...)

	testutils.AssertEquals(t, level.Notice, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkNotice perform benchmarking of the Notice().
func BenchmarkNotice(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Notice(message, parameters...)
	}
}

// TestWarning tests that Warning logs message with parameters on warning level
// using default logger.
func TestWarning(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Warning(message, parameters...)

	testutils.AssertEquals(t, level.Warning, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkWarning perform benchmarking of the Warning().
func BenchmarkWarning(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Warning(message, parameters...)
	}
}

// TestSevere tests that Severe logs message with parameters on severe level
// using default logger.
func TestSevere(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Severe(message, parameters...)

	testutils.AssertEquals(t, level.Severe, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkSevere perform benchmarking of the Severe().
func BenchmarkSevere(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Severe(message, parameters...)
	}
}

// TestError tests that Error logs message with parameters on error level using
// default logger.
func TestError(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Error(message, parameters...)

	testutils.AssertEquals(t, level.Error, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkError perform benchmarking of the Error().
func BenchmarkError(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Error(message, parameters...)
	}
}

// TestAlert tests that Alert logs message with parameters on alert level using
// default logger.
func TestAlert(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Alert(message, parameters...)

	testutils.AssertEquals(t, level.Alert, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkAlert perform benchmarking of the Alert().
func BenchmarkAlert(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Alert(message, parameters...)
	}
}

// TestCritical tests that Critical logs message with parameters on critical
// level using default logger.
func TestCritical(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Critical(message, parameters...)

	testutils.AssertEquals(t, level.Critical, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkCritical perform benchmarking of the Critical().
func BenchmarkCritical(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Critical(message, parameters...)
	}
}

// TestEmergency tests that Emergency logs message with parameters on emergency
// level using default logger.
func TestEmergency(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Emergency(message, parameters...)

	testutils.AssertEquals(t, level.Emergency, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkEmergency perform benchmarking of the Emergency().
func BenchmarkEmergency(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Emergency(message, parameters...)
	}
}

// TestErrorLevel tests that ErrorLevel returns the error level of the default
// logger.
func TestErrorLevel(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	testutils.AssertEquals(t, level.Error, ErrorLevel())
}

// BenchmarkErrorLevel perform benchmarking of the ErrorLevel().
func BenchmarkErrorLevel(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	for index := 0; index < b.N; index++ {
		ErrorLevel()
	}
}

// TestSetErrorLevel tests that SetErrorLevel sets the error level of the default
// logger.
func TestSetErrorLevel(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	SetErrorLevel(level.Warning)

	testutils.AssertEquals(t, level.Warning, rootLogger.ErrorLevel())
}

// BenchmarkSetErrorLevel perform benchmarking of the SetErrorLevel().
func BenchmarkSetErrorLevel(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	for index := 0; index < b.N; index++ {
		SetErrorLevel(level.Warning)
	}
}

// TestPanicLevel tests that PanicLevel returns the panic level of the default
// logger.
func TestPanicLevel(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	testutils.AssertEquals(t, level.Critical, PanicLevel())
}

// BenchmarkPanicLevel perform benchmarking of the PanicLevel().
func BenchmarkPanicLevel(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	for index := 0; index < b.N; index++ {
		PanicLevel()
	}
}

// TestSetPanicLevel tests that SetPanicLevel sets the panic level of the default
// logger.
func TestSetPanicLevel(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	SetPanicLevel(level.Warning)

	testutils.AssertEquals(t, level.Warning, rootLogger.PanicLevel())
}

// BenchmarkSetPanicLevel perform benchmarking of the SetPanicLevel().
func BenchmarkSetPanicLevel(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	for index := 0; index < b.N; index++ {
		SetPanicLevel(level.Warning)
	}
}

// TestRaiseError tests that RaiseError logs message with parameters on error
// level and returns a new error using default logger.
func TestRaiseError(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	expectedError := fmt.Errorf(message, parameters...)

	err := RaiseError(message, parameters...)

	testutils.AssertEquals(t, level.Error, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
	testutils.AssertNotNil(t, err)
	testutils.AssertEquals(t, expectedError, err)
}

// BenchmarkRaiseError perform benchmarking of the RaiseError().
func BenchmarkRaiseError(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	for index := 0; index < b.N; index++ {
		_ = RaiseError(message, parameters...)
	}
}

// TestCaptureError tests that CaptureError logs message from the error using the
// default logger.
func TestCaptureError(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	err := fmt.Errorf(message, parameters...)

	CaptureError(err)

	testutils.AssertEquals(t, level.Error, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, err.Error(), mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, make([]any, 0), mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkCaptureError perform benchmarking of the CaptureError().
func BenchmarkCaptureError(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	err := fmt.Errorf(message, parameters...)

	for index := 0; index < b.N; index++ {
		CaptureError(err)
	}
}

// TestPanic tests that Panic logs message with parameters on panic level and
// panics using the default logger.
func TestPanic(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic")
		} else {
			testutils.AssertEquals(t, fmt.Sprintf(message, parameters...), r.(string))
		}
	}()

	Panic(message, parameters...)

	testutils.AssertNotNil(t, mockLogger.Parameters)
	testutils.AssertEquals(t, level.Critical, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, message, mockLogger.Parameters[1].(string))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[2:len(mockLogger.Parameters)])
}

// BenchmarkPanic perform benchmarking of the Panic().
func BenchmarkPanic(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger, errorLevel: level.Error, panicLevel: level.Critical}

	for index := 0; index < b.N; index++ {
		func() {
			defer func() {
				if r := recover(); r == nil {
					return
				}
			}()
			Panic(message, parameters...)
		}()
	}
}
