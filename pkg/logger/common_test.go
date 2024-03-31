package logger

import (
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger/formatter"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"github.com/dl1998/go-logging/pkg/logger/logrecord"
	"io"
	"testing"
	"time"
)

var (
	loggerTemplate = "%(level):%(name):%(message)"
	loggerName     = "test"
	message        = "Test Message: %s."
	parameters     = []any{
		"test",
	}
	timeFormat = time.RFC3339
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
func (mock *MockLogger) Log(level level.Level, message string, parameters ...any) {
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
func (mock *MockLogger) AddHandler(handlerInterface handler.Interface) {
	mock.CalledName = "AddHandler"
	mock.Called = true
	mock.Parameters = append(make([]any, 0), handlerInterface)
	mock.Return = nil
}

// RemoveHandler mocks RemoveHandler from baseLogger.
func (mock *MockLogger) RemoveHandler(handlerInterface handler.Interface) {
	mock.CalledName = "RemoveHandler"
	mock.Called = true
	mock.Parameters = append(make([]any, 0), handlerInterface)
	mock.Return = nil
}

// MockHandler is used to mock Handler.
type MockHandler struct {
	writer     io.Writer
	CalledName string
	Called     bool
	Parameters []any
	Return     any
}

// Writer mocks Writer from Handler.
func (mock *MockHandler) Writer() io.Writer {
	mock.CalledName = "Writer"
	mock.Called = true
	mock.Parameters = make([]any, 0)
	mock.Return = mock.writer
	return mock.writer
}

// FromLevel mocks FromLevel from Handler.
func (mock *MockHandler) FromLevel() level.Level {
	mock.CalledName = "FromLevel"
	mock.Called = true
	mock.Parameters = make([]any, 0)
	returnValue := level.Debug
	mock.Return = returnValue
	return returnValue
}

// SetFromLevel mocks SetFromLevel from Handler.
func (mock *MockHandler) SetFromLevel(level level.Level) {
	mock.CalledName = "SetFromLevel"
	mock.Called = true
	mock.Parameters = append(make([]any, 0), level)
	mock.Return = nil
}

// ToLevel mocks ToLevel from Handler.
func (mock *MockHandler) ToLevel() level.Level {
	mock.CalledName = "ToLevel"
	mock.Called = true
	mock.Parameters = make([]any, 0)
	returnValue := level.Debug
	mock.Return = returnValue
	return returnValue
}

// SetToLevel mocks SetToLevel from Handler.
func (mock *MockHandler) SetToLevel(level level.Level) {
	mock.CalledName = "SetToLevel"
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
func (mock *MockHandler) Write(record logrecord.Interface) {
	mock.CalledName = "Write"
	mock.Called = true
	mock.Parameters = append(make([]any, 0), record)
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

	logLevel := level.Debug

	newBaseLogger.Log(logLevel, message, parameters...)

	handlerRecord := newHandler.Parameters[0].(*logrecord.LogRecord)

	testutils.AssertEquals(t, loggerName, handlerRecord.Name())
	testutils.AssertEquals(t, logLevel, handlerRecord.Level())
	testutils.AssertEquals(t, fmt.Sprintf(message, parameters...), handlerRecord.Message())
}

// BenchmarkBaseLogger_Log perform benchmarking of the baseLogger.Log().
func BenchmarkBaseLogger_Log(b *testing.B) {
	newBaseLogger := &baseLogger{
		name: loggerName,
		handlers: []handler.Interface{
			&MockHandler{},
		},
	}

	logLevel := level.Debug

	for index := 0; index < b.N; index++ {
		newBaseLogger.Log(logLevel, message, parameters...)
	}
}

// TestBaseLogger_Name tests that baseLogger.Name returns loggerName of the logger.
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

// TestBaseLogger_SetName tests that baseLogger.SetName set a new loggerName for the
// logger.
func TestBaseLogger_SetName(t *testing.T) {
	newBaseLogger := &baseLogger{
		name: loggerName,
		handlers: []handler.Interface{
			&MockHandler{},
		},
	}

	newName := "new-loggerName"

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

	newName := "new-loggerName"

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

// TestBaseLogger_RemoveHandler tests that baseLogger.RemoveHandler removes a
// Handler from the list of handlers.
func TestBaseLogger_RemoveHandler(t *testing.T) {
	newHandler := &MockHandler{}

	newBaseLogger := &baseLogger{
		name:     loggerName,
		handlers: []handler.Interface{newHandler},
	}

	newBaseLogger.RemoveHandler(newHandler)

	testutils.AssertEquals(t, make([]handler.Interface, 0), newBaseLogger.handlers)
}

// BenchmarkBaseLogger_RemoveHandler perform benchmarking of the baseLogger.RemoveHandler().
func BenchmarkBaseLogger_RemoveHandler(b *testing.B) {
	newHandler := &MockHandler{}

	newBaseLogger := &baseLogger{
		name:     loggerName,
		handlers: []handler.Interface{newHandler},
	}

	for index := 0; index < b.N; index++ {
		newBaseLogger.RemoveHandler(newHandler)
	}
}
