package structuredlogger

import (
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/structuredlogger/handler"
	"github.com/dl1998/go-logging/pkg/structuredlogger/logrecord"
	"sync"
	"testing"
)

var (
	messageQueueSize = 1
)

// TestBaseAsyncLogger_startListeningMessages tests that
// baseAsyncLogger.startListeningMessages method processes logs from the message
// queue.
func TestBaseAsyncLogger_startListeningMessages(t *testing.T) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := &baseAsyncLogger{
		baseLogger: &baseLogger{
			name:     loggerName,
			handlers: []handler.Interface{mockHandler},
		},
		messageQueue: make(chan logrecord.Interface, messageQueueSize),
		waitGroup:    sync.WaitGroup{},
	}

	record := logrecord.New(loggerName, 0, timeFormat, parametersWithMap, 3)
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
	newBaseAsyncLogger := &baseAsyncLogger{
		baseLogger: &baseLogger{
			name:     loggerName,
			handlers: []handler.Interface{mockHandler},
		},
		messageQueue: make(chan logrecord.Interface, messageQueueSize),
		waitGroup:    sync.WaitGroup{},
	}

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
	newBaseAsyncLogger := &baseAsyncLogger{
		baseLogger: &baseLogger{
			name:     loggerName,
			handlers: []handler.Interface{mockHandler},
		},
		messageQueue: make(chan logrecord.Interface, messageQueueSize),
		waitGroup:    sync.WaitGroup{},
	}

	newBaseAsyncLogger.Open(messageQueueSize)

	testutils.AssertNotNil(t, newBaseAsyncLogger.messageQueue)
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
	newBaseAsyncLogger := &baseAsyncLogger{
		baseLogger: &baseLogger{
			name:     loggerName,
			handlers: []handler.Interface{mockHandler},
		},
		messageQueue: make(chan logrecord.Interface, messageQueueSize),
		waitGroup:    sync.WaitGroup{},
	}

	newBaseAsyncLogger.Close()

	testutils.AssertEquals(t, true, isChannelClosed(newBaseAsyncLogger.messageQueue))
}

// TestBaseAsyncLogger_Log tests that baseAsyncLogger.Log sends a new record on
// the message queue.
func TestBaseAsyncLogger_Log(t *testing.T) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := &baseAsyncLogger{
		baseLogger: &baseLogger{
			name:     loggerName,
			handlers: []handler.Interface{mockHandler},
		},
		messageQueue: make(chan logrecord.Interface, messageQueueSize),
		waitGroup:    sync.WaitGroup{},
	}

	newBaseAsyncLogger.Log(level.Trace, parametersWithMap)

	record := <-newBaseAsyncLogger.messageQueue

	testutils.AssertEquals(t, loggerName, record.Name())
	testutils.AssertEquals(t, parametersWithMap, record.Parameters())
}

// BenchmarkBaseAsyncLogger_Log benchmarks baseAsyncLogger.Log method of the
// baseAsyncLogger.
func BenchmarkBaseAsyncLogger_Log(b *testing.B) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := &baseAsyncLogger{
		baseLogger: &baseLogger{
			name:     loggerName,
			handlers: []handler.Interface{mockHandler},
		},
		messageQueue: make(chan logrecord.Interface, b.N),
		waitGroup:    sync.WaitGroup{},
	}

	b.ResetTimer()

	for index := 0; index < b.N; index++ {
		newBaseAsyncLogger.Log(level.Trace, parametersWithMap)
	}
}

// TestNewAsyncStructuredLogger tests that NewAsyncStructuredLogger creates a new
// AsyncStructuredLogger.
func TestNewAsyncStructuredLogger(t *testing.T) {
	newAsyncStructuredLogger := NewAsyncStructuredLogger(loggerName, timeFormat, messageQueueSize)

	testutils.AssertNotNil(t, newAsyncStructuredLogger)
	testutils.AssertEquals(t, loggerName, newAsyncStructuredLogger.Name())
	testutils.AssertEquals(t, messageQueueSize, cap(newAsyncStructuredLogger.baseLogger.(*baseAsyncLogger).messageQueue))
}

// BenchmarkNewAsyncStructuredLogger benchmarks NewAsyncStructuredLogger.
func BenchmarkNewAsyncStructuredLogger(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NewAsyncStructuredLogger(loggerName, timeFormat, messageQueueSize)
	}
}

// TestAsyncStructuredLogger_WaitToFinishLogging tests that
// AsyncStructuredLogger.WaitToFinishLogging waits until async logger finish
// logging messages.
func TestAsyncStructuredLogger_WaitToFinishLogging(t *testing.T) {
	mockHandler := &MockHandler{}
	newAsyncStructuredLogger := &AsyncStructuredLogger{
		Logger: &Logger{
			baseLogger: &baseAsyncLogger{
				baseLogger: &baseLogger{
					name:     loggerName,
					handlers: []handler.Interface{mockHandler},
				},
				messageQueue: make(chan logrecord.Interface, messageQueueSize),
				waitGroup:    sync.WaitGroup{},
			},
		},
	}

	waitChannel := make(chan bool)

	go func() {
		waitChannel <- true
		newAsyncStructuredLogger.baseLogger.(*baseAsyncLogger).waitGroup.Add(1)
	}()

	newAsyncStructuredLogger.WaitToFinishLogging()
	waited := <-waitChannel

	testutils.AssertEquals(t, true, waited)
}

// TestAsyncStructuredLogger_Open tests that AsyncStructuredLogger.Open creates a
// new message queue and start listening messages.
func TestAsyncStructuredLogger_Open(t *testing.T) {
	mockHandler := &MockHandler{}
	newAsyncStructuredLogger := &AsyncStructuredLogger{
		Logger: &Logger{
			baseLogger: &baseAsyncLogger{
				baseLogger: &baseLogger{
					name:     loggerName,
					handlers: []handler.Interface{mockHandler},
				},
				messageQueue: make(chan logrecord.Interface, messageQueueSize),
				waitGroup:    sync.WaitGroup{},
			},
		},
	}

	newAsyncStructuredLogger.Open(messageQueueSize)

	testutils.AssertNotNil(t, newAsyncStructuredLogger.baseLogger.(*baseAsyncLogger).messageQueue)
}

// TestAsyncStructuredLogger_Close tests that AsyncStructuredLogger.Close closes
// message queue channel.
func TestAsyncStructuredLogger_Close(t *testing.T) {
	mockHandler := &MockHandler{}
	newAsyncStructuredLogger := &AsyncStructuredLogger{
		Logger: &Logger{
			baseLogger: &baseAsyncLogger{
				baseLogger: &baseLogger{
					name:     loggerName,
					handlers: []handler.Interface{mockHandler},
				},
				messageQueue: make(chan logrecord.Interface, messageQueueSize),
				waitGroup:    sync.WaitGroup{},
			},
		},
	}

	newAsyncStructuredLogger.Close()

	testutils.AssertEquals(t, true, isChannelClosed(newAsyncStructuredLogger.baseLogger.(*baseAsyncLogger).messageQueue))
}
