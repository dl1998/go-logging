package logger

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"github.com/dl1998/go-logging/pkg/logger/logrecord"
	"sync"
)

// baseAsyncLogger struct contains basic fields for the async logger.
type baseAsyncLogger struct {
	// baseLogger is a base logger.
	*baseLogger
	// messageQueue is a channel for the log messages.
	messageQueue chan logrecord.Interface
	// waitGroup is a wait group for the async logger messages.
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
func (logger *baseAsyncLogger) Log(level level.Level, message string, parameters ...any) {
	logger.waitGroup.Add(1)
	record := logrecord.New(logger.name, level, logger.timeFormat, message, parameters, 3)
	logger.messageQueue <- record
}

// AsyncLoggerInterface defines async logger interface.
type AsyncLoggerInterface interface {
	Interface
	WaitToFinishLogging()
	Open(queueSize int)
	Close()
}

// AsyncLogger represents an asynchronous logger.
type AsyncLogger struct {
	// Logger is a standard logger.
	*Logger
}

// NewAsyncLogger creates a new AsyncLogger with the provided name, timeFormat, and queueSize.
func NewAsyncLogger(name string, timeFormat string, queueSize int) *AsyncLogger {
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
	return &AsyncLogger{
		Logger: &Logger{
			baseLogger: newBaseLogger,
		},
	}
}

// WaitToFinishLogging waits for all messages to be logged.
func (logger *AsyncLogger) WaitToFinishLogging() {
	logger.baseLogger.(*baseAsyncLogger).WaitToFinishLogging()
}

// Open opens the messageQueue with the provided queueSize and starts listening
// for messages.
func (logger *AsyncLogger) Open(queueSize int) {
	logger.baseLogger.(*baseAsyncLogger).Open(queueSize)
}

// Close closes the messageQueue.
func (logger *AsyncLogger) Close() {
	logger.baseLogger.(*baseAsyncLogger).Close()
}
