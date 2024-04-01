package logger

import (
	"fmt"
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"github.com/dl1998/go-logging/pkg/logger/logrecord"
	"sync"
	"testing"
)

var (
	messageQueueSize = 1
)

// createBaseAsyncLogger creates a new baseAsyncLogger with the provided
// handlers, message queue size and channel open flag.
func createBaseAsyncLogger(handlers []handler.Interface, messageQueueSize int, isChannelOpen bool) *baseAsyncLogger {
	newBaseAsyncLogger := &baseAsyncLogger{
		baseLogger: &baseLogger{
			name:     loggerName,
			handlers: handlers,
		},
		messageQueue:  make(chan logrecord.Interface, messageQueueSize),
		isChannelOpen: isChannelOpen,
		waitGroup:     sync.WaitGroup{},
	}
	return newBaseAsyncLogger
}

// TestBaseAsyncLogger_startListeningMessages tests that
// baseAsyncLogger.startListeningMessages method processes logs from the message
// queue.
func TestBaseAsyncLogger_startListeningMessages(t *testing.T) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := createBaseAsyncLogger([]handler.Interface{mockHandler}, messageQueueSize, true)

	record := logrecord.New(loggerName, logLevel, timeFormat, message, parameters, skipCallers)
	newBaseAsyncLogger.messageQueue <- record

	go newBaseAsyncLogger.startListeningMessages()
	newBaseAsyncLogger.waitGroup.Add(1)
	newBaseAsyncLogger.waitGroup.Wait()

	testutils.AssertEquals(t, record, mockHandler.Parameters[0].(*logrecord.LogRecord))
}

// TestBaseAsyncLogger_WaitToFinishLogging tests that
// baseAsyncLogger.WaitToFinishLogging waits until async logger finish logging
// messages.
func TestBaseAsyncLogger_WaitToFinishLogging(t *testing.T) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := createBaseAsyncLogger([]handler.Interface{mockHandler}, messageQueueSize, true)

	waitChannel := make(chan bool)

	go func() {
		waitChannel <- true
		newBaseAsyncLogger.waitGroup.Add(1)
	}()

	newBaseAsyncLogger.WaitToFinishLogging()
	waited := <-waitChannel

	testutils.AssertEquals(t, true, waited)
}

// TestBaseAsyncLogger_Open tests that baseAsyncLogger.Open creates a new message
// queue and start listening messages.
func TestBaseAsyncLogger_Open(t *testing.T) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := createBaseAsyncLogger([]handler.Interface{mockHandler}, messageQueueSize, true)

	newBaseAsyncLogger.Close()

	err := newBaseAsyncLogger.Open(messageQueueSize)

	testutils.AssertNotNil(t, newBaseAsyncLogger.messageQueue)
	testutils.AssertNil(t, err)
	testutils.AssertEquals(t, true, newBaseAsyncLogger.isChannelOpen)
}

// TestBaseAsyncLogger_Open_Error tests that baseAsyncLogger.Open returns an
// error when trying to open a new channel with current channel being open.
func TestBaseAsyncLogger_Open_Error(t *testing.T) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := createBaseAsyncLogger([]handler.Interface{mockHandler}, messageQueueSize, true)

	err := newBaseAsyncLogger.Open(messageQueueSize)

	testutils.AssertNotNil(t, newBaseAsyncLogger.messageQueue)
	testutils.AssertNotNil(t, err)
	testutils.AssertEquals(t, true, newBaseAsyncLogger.isChannelOpen)
}

// isChannelClosed checks if the provided channel is closed.
func isChannelClosed(ch <-chan logrecord.Interface) bool {
	select {
	case _, ok := <-ch:
		return !ok
	default:
		return false
	}
}

// TestBaseAsyncLogger_Close tests that baseAsyncLogger.Close closes message
// queue channel.
func TestBaseAsyncLogger_Close(t *testing.T) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := createBaseAsyncLogger([]handler.Interface{mockHandler}, messageQueueSize, true)

	newBaseAsyncLogger.Close()

	testutils.AssertEquals(t, true, isChannelClosed(newBaseAsyncLogger.messageQueue))
	testutils.AssertEquals(t, false, newBaseAsyncLogger.isChannelOpen)
}

// TestBaseAsyncLogger_Log tests that baseAsyncLogger.Log sends a new record on
// the message queue.
func TestBaseAsyncLogger_Log(t *testing.T) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := createBaseAsyncLogger([]handler.Interface{mockHandler}, messageQueueSize, true)

	newBaseAsyncLogger.Log(logLevel, message, parameters...)
	record := <-newBaseAsyncLogger.messageQueue

	testutils.AssertEquals(t, logLevel, record.Level())
	testutils.AssertEquals(t, fmt.Sprintf(message, parameters...), record.Message())
}

// BenchmarkBaseAsyncLogger_Log benchmarks baseAsyncLogger.Log method of the
// baseAsyncLogger.
func BenchmarkBaseAsyncLogger_Log(b *testing.B) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := createBaseAsyncLogger([]handler.Interface{mockHandler}, b.N, true)

	b.ResetTimer()

	for index := 0; index < b.N; index++ {
		newBaseAsyncLogger.Log(logLevel, message, parameters...)
	}
}

// TestNewAsyncLogger tests that NewAsyncLogger creates and return a new
// AsyncLogger.
func TestNewAsyncLogger(t *testing.T) {
	newAsyncLogger := NewAsyncLogger(loggerName, timeFormat, messageQueueSize)

	testutils.AssertNotNil(t, newAsyncLogger)
	testutils.AssertEquals(t, loggerName, newAsyncLogger.Name())
	testutils.AssertEquals(t, messageQueueSize, cap(newAsyncLogger.baseLogger.(*baseAsyncLogger).messageQueue))
}

// BenchmarkNewAsyncLogger benchmarks NewAsyncLogger.
func BenchmarkNewAsyncLogger(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NewAsyncLogger(loggerName, timeFormat, messageQueueSize)
	}
}

// TestAsyncLogger_WaitToFinishLogging tests that AsyncLogger.WaitToFinishLogging
// waits until async logger finish logging messages.
func TestAsyncLogger_WaitToFinishLogging(t *testing.T) {
	mockHandler := &MockHandler{}
	newAsyncLogger := &AsyncLogger{
		Logger: &Logger{
			baseLogger: createBaseAsyncLogger([]handler.Interface{mockHandler}, messageQueueSize, true),
		},
	}

	waitChannel := make(chan bool)

	go func() {
		waitChannel <- true
		newAsyncLogger.WaitToFinishLogging()
	}()

	waited := <-waitChannel

	testutils.AssertEquals(t, true, waited)
}

// TestAsyncLogger_Open tests that AsyncLogger.Open creates a new message queue
// and start listening messages.
func TestAsyncLogger_Open(t *testing.T) {
	mockHandler := &MockHandler{}
	newAsyncLogger := &AsyncLogger{
		Logger: &Logger{
			baseLogger: createBaseAsyncLogger([]handler.Interface{mockHandler}, messageQueueSize, true),
		},
	}

	newAsyncLogger.Close()

	err := newAsyncLogger.Open(messageQueueSize)

	testutils.AssertNotNil(t, newAsyncLogger.baseLogger.(*baseAsyncLogger).messageQueue)
	testutils.AssertNil(t, err)
	testutils.AssertEquals(t, true, newAsyncLogger.baseLogger.(*baseAsyncLogger).isChannelOpen)
}

// TestAsyncLogger_Open_Error tests that AsyncLogger.Open returns an error when
// trying to open a new channel with current channel being open.
func TestAsyncLogger_Open_Error(t *testing.T) {
	mockHandler := &MockHandler{}
	newAsyncLogger := &AsyncLogger{
		Logger: &Logger{
			baseLogger: createBaseAsyncLogger([]handler.Interface{mockHandler}, messageQueueSize, true),
		},
	}

	err := newAsyncLogger.Open(messageQueueSize)

	testutils.AssertNotNil(t, newAsyncLogger.baseLogger.(*baseAsyncLogger).messageQueue)
	testutils.AssertNotNil(t, err)
	testutils.AssertEquals(t, true, newAsyncLogger.baseLogger.(*baseAsyncLogger).isChannelOpen)
}

// TestAsyncLogger_Close tests that AsyncLogger.Close closes message queue
// channel.
func TestAsyncLogger_Close(t *testing.T) {
	mockHandler := &MockHandler{}
	newAsyncLogger := &AsyncLogger{
		Logger: &Logger{
			baseLogger: createBaseAsyncLogger([]handler.Interface{mockHandler}, messageQueueSize, true),
		},
	}

	newAsyncLogger.Close()

	testutils.AssertEquals(t, true, isChannelClosed(newAsyncLogger.baseLogger.(*baseAsyncLogger).messageQueue))
	testutils.AssertEquals(t, false, newAsyncLogger.baseLogger.(*baseAsyncLogger).isChannelOpen)
}
