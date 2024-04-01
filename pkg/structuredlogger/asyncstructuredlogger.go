package structuredlogger

import (
	"fmt"
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
	// isChannelOpen is a flag that indicates if the messageQueue is open.
	isChannelOpen bool
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
func (logger *baseAsyncLogger) Open(queueSize int) error {
	if logger.isChannelOpen {
		return fmt.Errorf("cannot open a new message queue, current queue is already open")
	}
	logger.messageQueue = make(chan logrecord.Interface, queueSize)
	logger.isChannelOpen = true
	logger.waitGroup = sync.WaitGroup{}
	go logger.startListeningMessages()
	return nil
}

// Close closes the messageQueue.
func (logger *baseAsyncLogger) Close() {
	close(logger.messageQueue)
	logger.isChannelOpen = false
}

// Log logs interpolated message with the provided level.Level.
func (logger *baseAsyncLogger) Log(logLevel level.Level, parameters ...any) {
	logger.waitGroup.Add(1)
	var parametersMap = convertParametersToMap(parameters...)
	logRecord := logrecord.New(logger.name, logLevel, logger.timeFormat, parametersMap, 3)
	logger.messageQueue <- logRecord
}

// AsyncLoggerInterface defines async structured logger interface.
type AsyncLoggerInterface interface {
	Interface
	WaitToFinishLogging()
	Open(queueSize int)
	Close()
}

// AsyncLogger is a structured logger that logs messages asynchronously.
type AsyncLogger struct {
	// Logger is a standard structured logger.
	*Logger
}

// NewAsyncLogger creates a new async structured logger.
func NewAsyncLogger(name string, timeFormat string, queueSize int) *AsyncLogger {
	newBaseLogger := &baseAsyncLogger{
		baseLogger: &baseLogger{
			name:       name,
			timeFormat: timeFormat,
			handlers:   make([]handler.Interface, 0),
		},
		messageQueue:  make(chan logrecord.Interface, queueSize),
		isChannelOpen: true,
		waitGroup:     sync.WaitGroup{},
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
func (logger *AsyncLogger) Open(queueSize int) error {
	return logger.baseLogger.(*baseAsyncLogger).Open(queueSize)
}

// Close closes the messageQueue.
func (logger *AsyncLogger) Close() {
	logger.baseLogger.(*baseAsyncLogger).Close()
}
