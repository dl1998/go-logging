// Package logger_test has tests for logger package.
package logger

import (
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"github.com/dl1998/go-logging/pkg/logger/loglevel"
	"testing"
)

var (
	loggerTemplate = "%(level):%(name):%(message)"
	loggerName     = "test"
	message        = "Test Message: %s."
	parameters     = []any{
		"test",
	}
)

// MockLogger is used to mock baseLogger.
type MockLogger struct {
	handlers   []handler.Interface
	CalledName string
	Called     bool
	Parameters []any
	Return     any
}

// Log mocks Log from baseLogger.
func (mock *MockLogger) Log(level loglevel.LogLevel, message string, parameters ...any) {
	mock.CalledName = "Log"
	mock.Called = true
	mock.Parameters = append(make([]any, 0), level, message)
	mock.Parameters = append(mock.Parameters, parameters...)
	mock.Return = nil
}

// Name mocks Name from baseLogger.
func (mock *MockLogger) Name() string {
	mock.CalledName = "SetName"
	mock.Called = true
	mock.Parameters = make([]any, 0)
	returnValue := "test"
	mock.Return = returnValue
	return returnValue
}

// SetName mocks SetName from baseLogger.
func (mock *MockLogger) SetName(name string) {
	mock.CalledName = "SetName"
	mock.Called = true
	mock.Parameters = append(make([]any, 0), name)
	mock.Return = nil
}

// Handlers mocks Handlers from baseLogger.
func (mock *MockLogger) Handlers() []handler.Interface {
	mock.CalledName = "Handlers"
	mock.Called = true
	mock.Parameters = make([]any, 0)
	returnValue := mock.handlers
	if mock.handlers == nil {
		returnValue = make([]handler.Interface, 0)
	}
	mock.Return = returnValue
	return returnValue
}

// AddHandler mocks AddHandler from baseLogger.
func (mock *MockLogger) AddHandler(handler handler.Interface) {
	mock.CalledName = "AddHandler"
	mock.Called = true
	mock.Parameters = append(make([]any, 0), handler)
	mock.Return = nil
}

// MockHandler is used to mock Handler.
type MockHandler struct {
	CalledName string
	Called     bool
	Parameters []any
	Return     any
}

// Level mocks Level from Handler.
func (mock *MockHandler) Level() loglevel.LogLevel {
	mock.CalledName = "Level"
	mock.Called = true
	mock.Parameters = make([]any, 0)
	returnValue := loglevel.Debug
	mock.Return = returnValue
	return returnValue
}

// SetLevel mocks SetLevel from Handler.
func (mock *MockHandler) SetLevel(level loglevel.LogLevel) {
	mock.CalledName = "SetLevel"
	mock.Called = true
	mock.Parameters = append(make([]any, 0), level)
	mock.Return = nil
}

// Formatter mocks Formatter from Handler.
func (mock *MockHandler) Formatter() formatter.Interface {
	mock.CalledName = "Formatter"
	mock.Called = true
	mock.Parameters = make([]any, 0)
	returnValue := formatter.New(loggerTemplate)
	mock.Return = returnValue
	return returnValue
}

// Write mocks Write from Handler.
func (mock *MockHandler) Write(logName string, level loglevel.LogLevel, message string, parameters ...any) {
	mock.CalledName = "Write"
	mock.Called = true
	mock.Parameters = append(make([]any, 0), logName, level, message)
	mock.Parameters = append(mock.Parameters, parameters...)
	mock.Return = nil
}

// TestBaseLogger_Log tests that baseLogger.Log method works correctly.
func TestBaseLogger_Log(t *testing.T) {
	newHandler := &MockHandler{}

	newBaseLogger := &baseLogger{
		name: loggerName,
		handlers: []handler.Interface{
			newHandler,
		},
	}

	logLevel := loglevel.Debug

	newBaseLogger.Log(logLevel, message, parameters...)

	testutils.AssertEquals(t, loggerName, newHandler.Parameters[0].(string))
	testutils.AssertEquals(t, logLevel, newHandler.Parameters[1].(loglevel.LogLevel))
	testutils.AssertEquals(t, message, newHandler.Parameters[2].(string))
	testutils.AssertEquals(t, parameters, newHandler.Parameters[3:len(newHandler.Parameters)])
}

// BenchmarkBaseLogger_Log perform benchmarking of the baseLogger.Log().
func BenchmarkBaseLogger_Log(b *testing.B) {
	newBaseLogger := &baseLogger{
		name: loggerName,
		handlers: []handler.Interface{
			&MockHandler{},
		},
	}

	logLevel := loglevel.Debug

	for index := 0; index < b.N; index++ {
		newBaseLogger.Log(logLevel, message, parameters...)
	}
}

// TestBaseLogger_Name tests that baseLogger.Name returns name of the logger.
func TestBaseLogger_Name(t *testing.T) {
	newBaseLogger := &baseLogger{
		name: loggerName,
		handlers: []handler.Interface{
			&MockHandler{},
		},
	}

	testutils.AssertEquals(t, loggerName, newBaseLogger.Name())
}

// BenchmarkBaseLogger_Name perform benchmarking of the baseLogger.Name().
func BenchmarkBaseLogger_Name(b *testing.B) {
	newBaseLogger := &baseLogger{
		name: loggerName,
		handlers: []handler.Interface{
			&MockHandler{},
		},
	}

	for index := 0; index < b.N; index++ {
		newBaseLogger.Name()
	}
}

// TestBaseLogger_SetName tests that baseLogger.SetName set a new name for the
// logger.
func TestBaseLogger_SetName(t *testing.T) {
	newBaseLogger := &baseLogger{
		name: loggerName,
		handlers: []handler.Interface{
			&MockHandler{},
		},
	}

	newName := "new-name"

	newBaseLogger.SetName(newName)

	testutils.AssertEquals(t, newName, newBaseLogger.name)
}

// BenchmarkBaseLogger_SetName perform benchmarking of the baseLogger.SetName().
func BenchmarkBaseLogger_SetName(b *testing.B) {
	newBaseLogger := &baseLogger{
		name: loggerName,
		handlers: []handler.Interface{
			&MockHandler{},
		},
	}

	newName := "new-name"

	for index := 0; index < b.N; index++ {
		newBaseLogger.SetName(newName)
	}
}

// TestBaseLogger_Handlers tests that baseLogger.Handlers returns a list of
// handlers for the logger.
func TestBaseLogger_Handlers(t *testing.T) {
	handlers := []handler.Interface{
		&MockHandler{},
	}

	newBaseLogger := &baseLogger{
		name:     loggerName,
		handlers: handlers,
	}

	testutils.AssertEquals(t, handlers, newBaseLogger.Handlers())
}

// BenchmarkBaseLogger_Handlers perform benchmarking of the baseLogger.Handlers().
func BenchmarkBaseLogger_Handlers(b *testing.B) {
	handlers := []handler.Interface{
		&MockHandler{},
	}

	newBaseLogger := &baseLogger{
		name:     loggerName,
		handlers: handlers,
	}

	for index := 0; index < b.N; index++ {
		newBaseLogger.Handlers()
	}
}

// TestBaseLogger_AddHandler tests that baseLogger.AddHandler adds a new Handler
// on the list of handlers.
func TestBaseLogger_AddHandler(t *testing.T) {
	newHandler := &MockHandler{}

	newBaseLogger := &baseLogger{
		name:     loggerName,
		handlers: make([]handler.Interface, 0),
	}

	newBaseLogger.AddHandler(newHandler)

	testutils.AssertEquals(t, []handler.Interface{newHandler}, newBaseLogger.handlers)
}

// BenchmarkBaseLogger_AddHandler perform benchmarking of the baseLogger.AddHandler().
func BenchmarkBaseLogger_AddHandler(b *testing.B) {
	newHandler := &MockHandler{}

	newBaseLogger := &baseLogger{
		name:     loggerName,
		handlers: make([]handler.Interface, 0),
	}

	for index := 0; index < b.N; index++ {
		newBaseLogger.AddHandler(newHandler)
	}
}

// TestNew tests that New creates a new logger.
func TestNew(t *testing.T) {
	newLogger := New(loggerName)

	testutils.AssertEquals(t, loggerName, newLogger.Name())

	handlersSize := len(newLogger.Handlers())

	testutils.AssertEquals(t, 0, handlersSize)
}

// BenchmarkNew perform benchmarking of the New().
func BenchmarkNew(b *testing.B) {
	for index := 0; index < b.N; index++ {
		New(loggerName)
	}
}

// TestGetDefaultLogger tests that GetDefaultLogger returns default logger
// instance.
func TestGetDefaultLogger(t *testing.T) {
	rootLoggerName := "root"

	rootLogger := GetDefaultLogger()

	testutils.AssertEquals(t, rootLoggerName, rootLogger.Name())

	handlersSize := len(rootLogger.Handlers())

	handlerInterface := rootLogger.Handlers()[0]

	testutils.AssertEquals(t, 1, handlersSize)
	testutils.AssertEquals(t, loglevel.Warning, handlerInterface.Level())
	testutils.AssertEquals(t, "%(level):%(name):%(message)", handlerInterface.Formatter().Template())
}

// BenchmarkGetDefaultLogger perform benchmarking of the GetDefaultLogger().
func BenchmarkGetDefaultLogger(b *testing.B) {
	for index := 0; index < b.N; index++ {
		GetDefaultLogger()
	}
}

// TestLogger_Name tests that Logger.Name returns name of the logger.
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

// TestLogger_Trace tests that Logger.Trace logs message with parameters on trace
// level.
func TestLogger_Trace(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Trace(message, parameters...)

	testutils.AssertEquals(t, loglevel.Trace, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Debug, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Verbose, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Info, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Notice, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Warning, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Severe, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Error, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Alert, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Critical, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Emergency, mockLogger.Parameters[0].(loglevel.LogLevel))
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

// TestSetLevel tests that SetLevel set a new log level for the default logger.
func TestSetLevel(t *testing.T) {
	mockHandler := &MockHandler{}
	mockLogger := &MockLogger{}
	mockLogger.handlers = []handler.Interface{mockHandler}

	rootLogger = &Logger{baseLogger: mockLogger}

	newLevel := loglevel.Error

	SetLevel(newLevel)

	testutils.AssertEquals(t, newLevel, mockHandler.Parameters[0].(loglevel.LogLevel))
}

// BenchmarkSetLevel perform benchmarking of the SetLevel().
func BenchmarkSetLevel(b *testing.B) {
	mockHandler := &MockHandler{}
	mockLogger := &MockLogger{}
	mockLogger.handlers = []handler.Interface{mockHandler}

	rootLogger = &Logger{baseLogger: mockLogger}

	newLevel := loglevel.Error

	for index := 0; index < b.N; index++ {
		SetLevel(newLevel)
	}
}

// TestTrace tests that Trace logs message with parameters on trace level using
// default logger.
func TestTrace(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Trace(message, parameters...)

	testutils.AssertEquals(t, loglevel.Trace, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Debug, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Verbose, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Info, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Notice, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Warning, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Severe, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Error, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Alert, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Critical, mockLogger.Parameters[0].(loglevel.LogLevel))
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

	testutils.AssertEquals(t, loglevel.Emergency, mockLogger.Parameters[0].(loglevel.LogLevel))
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
