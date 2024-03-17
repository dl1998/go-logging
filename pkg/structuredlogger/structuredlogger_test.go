// Package structuredlogger provides tests for the structuredlogger package.
package structuredlogger

import (
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/structuredlogger/formatter"
	"github.com/dl1998/go-logging/pkg/structuredlogger/handler"
	"io"
	"testing"
)

var (
	loggerName = "test"
	parameters = []any{
		"message", "test",
	}
	parametersWithMap = []any{
		map[string]interface{}{
			"message": "test",
		},
	}
	pretty = false
)

var loggerTemplate = map[string]string{
	"level": "%(level)",
	"name":  "%(name)",
}

// MockLogger is used to mock baseLogger.
type MockLogger struct {
	handlers   []handler.Interface
	CalledName string
	Called     bool
	Parameters []any
	Return     any
}

// Log mocks Log from baseLogger.
func (mock *MockLogger) Log(level level.Level, parameters ...any) {
	mock.CalledName = "Log"
	mock.Called = true
	mock.Parameters = append(make([]any, 0), level)
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
	returnValue := formatter.NewJSON(loggerTemplate, pretty)
	mock.Return = returnValue
	return returnValue
}

// Write mocks Write from Handler.
func (mock *MockHandler) Write(logName string, logLevel level.Level, parameters ...any) {
	mock.CalledName = "Write"
	mock.Called = true
	mock.Parameters = append(make([]any, 0), logName, logLevel)
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

	logLevel := level.Debug

	tests := map[string]struct {
		parameters         []any
		expectedParameters []any
	}{
		"Varargs": {
			parameters:         parameters,
			expectedParameters: parameters,
		},
		"Map": {
			parameters:         parametersWithMap,
			expectedParameters: parameters,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			newBaseLogger.Log(logLevel, test.parameters...)

			testutils.AssertEquals(t, loggerName, newHandler.Parameters[0].(string))
			testutils.AssertEquals(t, logLevel, newHandler.Parameters[1].(level.Level))
			testutils.AssertEquals(t, test.expectedParameters, newHandler.Parameters[2:len(newHandler.Parameters)])
		})
	}
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

	benchmarks := map[string]struct {
		parameters []any
	}{
		"Varargs": {
			parameters: parameters,
		},
		"Map": {
			parameters: parametersWithMap,
		},
	}

	for name, benchmark := range benchmarks {
		b.Run(name, func(b *testing.B) {
			for index := 0; index < b.N; index++ {
				newBaseLogger.Log(logLevel, benchmark.parameters...)
			}
		})
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

	newLogger.Trace(parameters...)

	testutils.AssertEquals(t, level.Trace, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Trace perform benchmarking of the Logger.Trace().
func BenchmarkLogger_Trace(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Trace(parameters...)
	}
}

// TestLogger_Debug tests that Logger.Debug logs message with parameters on debug
// level.
func TestLogger_Debug(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Debug(parameters...)

	testutils.AssertEquals(t, level.Debug, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Debug perform benchmarking of the Logger.Debug().
func BenchmarkLogger_Debug(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Debug(parameters...)
	}
}

// TestLogger_Verbose tests that Logger.Verbose logs message with parameters on
// verbose level.
func TestLogger_Verbose(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Verbose(parameters...)

	testutils.AssertEquals(t, level.Verbose, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Verbose perform benchmarking of the Logger.Verbose().
func BenchmarkLogger_Verbose(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Verbose(parameters...)
	}
}

// TestLogger_Info tests that Logger.Info logs message with parameters on info
// level.
func TestLogger_Info(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Info(parameters...)

	testutils.AssertEquals(t, level.Info, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Info perform benchmarking of the Logger.Info().
func BenchmarkLogger_Info(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Info(parameters...)
	}
}

// TestLogger_Notice tests that Logger.Notice logs message with parameters on
// notice level.
func TestLogger_Notice(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Notice(parameters...)

	testutils.AssertEquals(t, level.Notice, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Notice perform benchmarking of the Logger.Notice().
func BenchmarkLogger_Notice(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Notice(parameters...)
	}
}

// TestLogger_Warning tests that Logger.Warning logs message with parameters on
// warning level.
func TestLogger_Warning(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Warning(parameters...)

	testutils.AssertEquals(t, level.Warning, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Warning perform benchmarking of the Logger.Warning().
func BenchmarkLogger_Warning(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Warning(parameters...)
	}
}

// TestLogger_Severe tests that Logger.Severe logs message with parameters on
// severe level.
func TestLogger_Severe(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Severe(parameters...)

	testutils.AssertEquals(t, level.Severe, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Severe perform benchmarking of the Logger.Severe().
func BenchmarkLogger_Severe(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Severe(parameters...)
	}
}

// TestLogger_Error tests that Logger.Error logs message with parameters on error
// level.
func TestLogger_Error(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Error(parameters...)

	testutils.AssertEquals(t, level.Error, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Error perform benchmarking of the Logger.Error().
func BenchmarkLogger_Error(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Error(parameters...)
	}
}

// TestLogger_Alert tests that Logger.Alert logs message with parameters on alert
// level.
func TestLogger_Alert(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Alert(parameters...)

	testutils.AssertEquals(t, level.Alert, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Alert perform benchmarking of the Logger.Alert().
func BenchmarkLogger_Alert(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Alert(parameters...)
	}
}

// TestLogger_Critical tests that Logger.Critical logs message with parameters on
// critical level.
func TestLogger_Critical(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Critical(parameters...)

	testutils.AssertEquals(t, level.Critical, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Critical perform benchmarking of the Logger.Critical().
func BenchmarkLogger_Critical(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Critical(parameters...)
	}
}

// TestLogger_Emergency tests that Logger.Emergency logs message with parameters
// on emergency level.
func TestLogger_Emergency(t *testing.T) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	newLogger.Emergency(parameters...)

	testutils.AssertEquals(t, level.Emergency, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkLogger_Emergency perform benchmarking of the Logger.Emergency().
func BenchmarkLogger_Emergency(b *testing.B) {
	mockLogger := &MockLogger{}

	newLogger := &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		newLogger.Emergency(parameters...)
	}
}

// TestWithFromLevel tests that WithFromLevel sets the from level in the Configuration.
func TestWithFromLevel(t *testing.T) {
	configuration := NewConfiguration()

	option := WithFromLevel(level.Trace)

	option(configuration)

	testutils.AssertEquals(t, configuration.fromLevel, level.Trace)
}

// BenchmarkWithFromLevel perform benchmarking of the WithFromLevel().
func BenchmarkWithFromLevel(b *testing.B) {
	configuration := NewConfiguration()

	option := WithFromLevel(level.Trace)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithToLevel tests that WithToLevel sets the to level in the Configuration.
func TestWithToLevel(t *testing.T) {
	configuration := NewConfiguration()

	option := WithToLevel(level.Trace)

	option(configuration)

	testutils.AssertEquals(t, configuration.toLevel, level.Trace)
}

// BenchmarkWithToLevel perform benchmarking of the WithToLevel().
func BenchmarkWithToLevel(b *testing.B) {
	configuration := NewConfiguration()

	option := WithToLevel(level.Trace)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithTemplate tests that WithTemplate sets the template in the
// Configuration.
func TestWithTemplate(t *testing.T) {
	configuration := NewConfiguration()

	template := map[string]string{
		"level-number": "%(levelnr)",
		"name":         "%(name)",
	}

	option := WithTemplate(template)

	option(configuration)

	testutils.AssertEquals(t, configuration.template, template)
}

// BenchmarkWithTemplate perform benchmarking of the WithTemplate().
func BenchmarkWithTemplate(b *testing.B) {
	configuration := NewConfiguration()

	template := map[string]string{
		"level-number": "%(levelnr)",
		"name":         "%(name)",
	}

	option := WithTemplate(template)

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

	testutils.AssertEquals(t, configuration.file, file)
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

// TestWithFormat tests that WithFormat sets the format in the Configuration.
func TestWithFormat(t *testing.T) {
	configuration := NewConfiguration()

	format := "json"

	option := WithFormat(format)

	option(configuration)

	testutils.AssertEquals(t, configuration.format, format)
}

// BenchmarkWithFormat perform benchmarking of the WithFormat().
func BenchmarkWithFormat(b *testing.B) {
	configuration := NewConfiguration()

	format := "json"

	option := WithFormat(format)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithPretty tests that WithPretty sets the pretty in the Configuration.
func TestWithPretty(t *testing.T) {
	configuration := NewConfiguration()

	option := WithPretty(true)

	option(configuration)

	testutils.AssertEquals(t, configuration.pretty, true)
}

// BenchmarkWithPretty perform benchmarking of the WithPretty().
func BenchmarkWithPretty(b *testing.B) {
	configuration := NewConfiguration()

	option := WithPretty(true)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithKeyValueDelimiter tests that WithKeyValueDelimiter sets the key value
// delimiter in the Configuration.
func TestWithKeyValueDelimiter(t *testing.T) {
	configuration := NewConfiguration()

	delimiter := "="

	option := WithKeyValueDelimiter(delimiter)

	option(configuration)

	testutils.AssertEquals(t, configuration.keyValueDelimiter, delimiter)
}

// BenchmarkWithKeyValueDelimiter perform benchmarking of the WithKeyValueDelimiter().
func BenchmarkWithKeyValueDelimiter(b *testing.B) {
	configuration := NewConfiguration()

	delimiter := "="

	option := WithKeyValueDelimiter(delimiter)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithPairSeparator tests that WithPairSeparator sets the pair separator in the Configuration.
func TestWithPairSeparator(t *testing.T) {
	configuration := NewConfiguration()

	separator := ","

	option := WithPairSeparator(separator)

	option(configuration)

	testutils.AssertEquals(t, configuration.pairSeparator, separator)
}

// BenchmarkWithPairSeparator perform benchmarking of the WithPairSeparator().
func BenchmarkWithPairSeparator(b *testing.B) {
	configuration := NewConfiguration()

	separator := ","

	option := WithPairSeparator(separator)

	for index := 0; index < b.N; index++ {
		option(configuration)
	}
}

// TestWithName tests that WithName sets the name in the Configuration.
func TestWithName(t *testing.T) {
	configuration := NewConfiguration()

	name := "test"

	option := WithName(name)

	option(configuration)

	testutils.AssertEquals(t, configuration.name, name)
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

// TestNewConfiguration tests that NewConfiguration creates a new Configuration.
func TestNewConfiguration(t *testing.T) {
	tests := map[string]struct {
		options           []Option
		expectedFromLevel level.Level
		expectedToLevel   level.Level
		expectedTemplate  map[string]string
		expectedFile      string
		expectedName      string
	}{
		"Empty": {
			options:           []Option{},
			expectedFromLevel: level.Warning,
			expectedToLevel:   level.Null,
			expectedTemplate:  template,
			expectedFile:      "",
			expectedName:      "root",
		},
		"Non Standard": {
			options: []Option{
				WithFromLevel(level.All),
				WithToLevel(level.Emergency),
				WithTemplate(map[string]string{
					"level-number": "%(levelnr)",
					"name":         "%(name)",
				}),
				WithFile("file.log"),
				WithName("test"),
			},
			expectedFromLevel: level.All,
			expectedToLevel:   level.Emergency,
			expectedTemplate: map[string]string{
				"level-number": "%(levelnr)",
				"name":         "%(name)",
			},
			expectedFile: "file.log",
			expectedName: "test",
		},
	}
	for name, configuration := range tests {
		t.Run(name, func(t *testing.T) {
			newConfiguration := NewConfiguration(configuration.options...)

			testutils.AssertEquals(t, configuration.expectedFromLevel, newConfiguration.fromLevel)
			testutils.AssertEquals(t, configuration.expectedToLevel, newConfiguration.toLevel)
			testutils.AssertEquals(t, configuration.expectedTemplate, newConfiguration.template)
			testutils.AssertEquals(t, configuration.expectedFile, newConfiguration.file)
			testutils.AssertEquals(t, configuration.expectedName, newConfiguration.name)
		})
	}
}

// BenchmarkNewConfiguration perform benchmarking of the NewConfiguration().
func BenchmarkNewConfiguration(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NewConfiguration()
	}
}

// TestConfigure_JSON tests that Configure sets the configuration for the default
// JSON logger.
func TestConfigure_JSON(t *testing.T) {
	configuration := NewConfiguration(
		WithFromLevel(level.All),
		WithToLevel(level.Emergency),
		WithTemplate(template),
		WithFile(""),
		WithName("test"),
	)

	Configure(configuration)

	testutils.AssertEquals(t, "test", rootLogger.baseLogger.Name())
	testutils.AssertEquals(t, 2, len(rootLogger.baseLogger.Handlers()))
}

// TestConfigure_KeyValue tests that Configure sets the configuration for the default
// KeyValue logger.
func TestConfigure_KeyValue(t *testing.T) {
	configuration := NewConfiguration(
		WithFromLevel(level.All),
		WithToLevel(level.Emergency),
		WithTemplate(template),
		WithFile(""),
		WithName("test"),
		WithFormat("key-value"),
	)

	Configure(configuration)

	testutils.AssertEquals(t, "test", rootLogger.baseLogger.Name())
	testutils.AssertEquals(t, 2, len(rootLogger.baseLogger.Handlers()))
}

// TestConfigure_IncorrectFormat tests that Configure panics when receive an incorrect format.
func TestConfigure_IncorrectFormat(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic when received an incorrect format")
		}
	}()

	configuration := NewConfiguration(
		WithFormat("incorrect"),
	)

	Configure(configuration)
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
		WithTemplate(map[string]string{
			"level-number": "%(levelnr)",
			"name":         "%(name)",
		}),
		WithFile(""),
		WithName("test"),
	)

	for index := 0; index < b.N; index++ {
		Configure(configuration)
	}
}

// TestName tests that Name returns name of the default logger.
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

	testutils.AssertEquals(t, template, Template())
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

	Trace(parameters...)

	testutils.AssertEquals(t, level.Trace, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkTrace perform benchmarking of the Trace().
func BenchmarkTrace(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Trace(parameters...)
	}
}

// TestDebug tests that Debug logs message with parameters on debug level using
// default logger.
func TestDebug(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Debug(parameters...)

	testutils.AssertEquals(t, level.Debug, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkDebug perform benchmarking of the Debug().
func BenchmarkDebug(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Debug(parameters...)
	}
}

// TestVerbose tests that Verbose logs message with parameters on verbose level
// using default logger.
func TestVerbose(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Verbose(parameters...)

	testutils.AssertEquals(t, level.Verbose, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkVerbose perform benchmarking of the Verbose().
func BenchmarkVerbose(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Verbose(parameters...)
	}
}

// TestInfo tests that Info logs message with parameters on info level using
// default logger.
func TestInfo(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Info(parameters...)

	testutils.AssertEquals(t, level.Info, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkInfo perform benchmarking of the Info().
func BenchmarkInfo(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Info(parameters...)
	}
}

// TestNotice tests that Notice logs message with parameters on notice level
// using default logger.
func TestNotice(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Notice(parameters...)

	testutils.AssertEquals(t, level.Notice, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkNotice perform benchmarking of the Notice().
func BenchmarkNotice(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Notice(parameters...)
	}
}

// TestWarning tests that Warning logs message with parameters on warning level
// using default logger.
func TestWarning(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Warning(parameters...)

	testutils.AssertEquals(t, level.Warning, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkWarning perform benchmarking of the Warning().
func BenchmarkWarning(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Warning(parameters...)
	}
}

// TestSevere tests that Severe logs message with parameters on severe level
// using default logger.
func TestSevere(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Severe(parameters...)

	testutils.AssertEquals(t, level.Severe, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkSevere perform benchmarking of the Severe().
func BenchmarkSevere(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Severe(parameters...)
	}
}

// TestError tests that Error logs message with parameters on error level using
// default logger.
func TestError(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Error(parameters...)

	testutils.AssertEquals(t, level.Error, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkError perform benchmarking of the Error().
func BenchmarkError(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Error(parameters...)
	}
}

// TestAlert tests that Alert logs message with parameters on alert level using
// default logger.
func TestAlert(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Alert(parameters...)

	testutils.AssertEquals(t, level.Alert, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkAlert perform benchmarking of the Alert().
func BenchmarkAlert(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Alert(parameters...)
	}
}

// TestCritical tests that Critical logs message with parameters on critical
// level using default logger.
func TestCritical(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Critical(parameters...)

	testutils.AssertEquals(t, level.Critical, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkCritical perform benchmarking of the Critical().
func BenchmarkCritical(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Critical(parameters...)
	}
}

// TestEmergency tests that Emergency logs message with parameters on emergency
// level using default logger.
func TestEmergency(t *testing.T) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	Emergency(parameters...)

	testutils.AssertEquals(t, level.Emergency, mockLogger.Parameters[0].(level.Level))
	testutils.AssertEquals(t, parameters, mockLogger.Parameters[1:len(mockLogger.Parameters)])
}

// BenchmarkEmergency perform benchmarking of the Emergency().
func BenchmarkEmergency(b *testing.B) {
	mockLogger := &MockLogger{}

	rootLogger = &Logger{baseLogger: mockLogger}

	for index := 0; index < b.N; index++ {
		Emergency(parameters...)
	}
}
