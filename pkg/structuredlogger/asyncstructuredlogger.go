package structuredlogger

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/structuredlogger/handler"
	"github.com/dl1998/go-logging/pkg/structuredlogger/logrecord"
	"sync"
)

// baseAsyncLogger struct contains basic fields for the async structured logger.
type baseAsyncLogger struct {
	// baseLogger is a base logger.
	*baseLogger
	// messageQueue is a channel for the log messages.
	messageQueue chan logrecord.Interface
	// waitGroup is a wait group for the async structured logger messages.
	waitGroup sync.WaitGroup
}

// startListeningMessages starts listening for messages in the messageQueue.
func (logger *baseAsyncLogger) startListeningMessages() {
	for record := range logger.messageQueue {
		for _, registeredHandler := range logger.handlers {
			registeredHandler.Write(record)
		}
		logger.waitGroup.Done()
	}
}

// WaitToFinishLogging waits for all messages to be logged.
func (logger *baseAsyncLogger) WaitToFinishLogging() {
	logger.waitGroup.Wait()
}

// Open opens the messageQueue with the provided queueSize and starts listening
// for messages.
func (logger *baseAsyncLogger) Open(queueSize int) {
	logger.messageQueue = make(chan logrecord.Interface, queueSize)
	go logger.startListeningMessages()
}

// Close closes the messageQueue.
func (logger *baseAsyncLogger) Close() {
	close(logger.messageQueue)
}

// Log logs interpolated message with the provided level.Level.
func (logger *baseAsyncLogger) Log(logLevel level.Level, parameters ...any) {
	logger.waitGroup.Add(1)
	var parametersMap = convertParametersToMap(parameters...)
	logRecord := logrecord.New(logger.name, logLevel, logger.timeFormat, parametersMap, 3)
	logger.messageQueue <- logRecord
}

// AsyncStructuredLoggerInterface defines async structured logger interface.
type AsyncStructuredLoggerInterface interface {
	Interface
	WaitToFinishLogging()
	Open(queueSize int)
	Close()
}

// AsyncStructuredLogger is a structured logger that logs messages
// asynchronously.
type AsyncStructuredLogger struct {
	// Logger is a standard structured logger.
	*Logger
}

// NewAsyncStructuredLogger creates a new async structured logger.
func NewAsyncStructuredLogger(name string, timeFormat string, queueSize int) *AsyncStructuredLogger {
	newBaseLogger := &baseAsyncLogger{
		baseLogger: &baseLogger{
			name:       name,
			timeFormat: timeFormat,
			handlers:   make([]handler.Interface, 0),
		},
		messageQueue: make(chan logrecord.Interface, queueSize),
		waitGroup:    sync.WaitGroup{},
	}
	go newBaseLogger.startListeningMessages()
	return &AsyncStructuredLogger{
		Logger: &Logger{
			baseLogger: newBaseLogger,
		},
	}
}

// WaitToFinishLogging waits for all messages to be logged.
func (logger *AsyncStructuredLogger) WaitToFinishLogging() {
	logger.baseLogger.(*baseAsyncLogger).WaitToFinishLogging()
}

// Open opens the messageQueue with the provided queueSize and starts listening
// for messages.
func (logger *AsyncStructuredLogger) Open(queueSize int) {
	logger.baseLogger.(*baseAsyncLogger).Open(queueSize)
}

// Close closes the messageQueue.
func (logger *AsyncStructuredLogger) Close() {
	logger.baseLogger.(*baseAsyncLogger).Close()
}
