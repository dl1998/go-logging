package logger

import (
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"github.com/dl1998/go-logging/pkg/logger/logrecord"
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

	record := logrecord.New(loggerName, 0, timeFormat, message, parameters, 3)
	newBaseAsyncLogger.messageQueue <- record

	go newBaseAsyncLogger.startListeningMessages()
	newBaseAsyncLogger.waitGroup.Add(1)
	newBaseAsyncLogger.waitGroup.Wait()

	testutils.AssertEquals(t, record, mockHandler.Parameters[0].(*logrecord.LogRecord))
}

// TestBaseAsyncLogger_WaitToFinishLogging tests
// baseAsyncLogger.WaitToFinishLogging method of the baseAsyncLogger.
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

// TestBaseAsyncLogger_Open tests baseAsyncLogger.Open method of the
// baseAsyncLogger.
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

// TestBaseAsyncLogger_Close tests baseAsyncLogger.Close method of the
// baseAsyncLogger.
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

// TestBaseAsyncLogger_Log tests baseAsyncLogger.Log method of the baseAsyncLogger.
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

	expectedMessage := "test"
	newBaseAsyncLogger.Log(0, expectedMessage)
	actualMessage, ok := <-newBaseAsyncLogger.messageQueue

	testutils.AssertEquals(t, true, ok)
	testutils.AssertEquals(t, expectedMessage, actualMessage.Message())
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
		newBaseAsyncLogger.Log(level.Trace, message, parameters...)
	}
}

// TestNewAsyncLogger tests NewAsyncLogger function of the baseAsyncLogger.
func TestNewAsyncLogger(t *testing.T) {
	newAsyncLogger := NewAsyncLogger(loggerName, timeFormat, messageQueueSize)

	testutils.AssertNotNil(t, newAsyncLogger)
	testutils.AssertEquals(t, loggerName, newAsyncLogger.Name())
	testutils.AssertEquals(t, messageQueueSize, cap(newAsyncLogger.baseLogger.(*baseAsyncLogger).messageQueue))
}

// BenchmarkNewAsyncLogger benchmarks NewAsyncLogger function of the
// baseAsyncLogger.
func BenchmarkNewAsyncLogger(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NewAsyncLogger(loggerName, timeFormat, messageQueueSize)
	}
}

// TestAsyncLogger_WaitToFinishLogging tests AsyncLogger.WaitToFinishLogging
// method of the AsyncLogger.
func TestAsyncLogger_WaitToFinishLogging(t *testing.T) {
	mockHandler := &MockHandler{}
	newAsyncLogger := &AsyncLogger{
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
		newAsyncLogger.WaitToFinishLogging()
	}()

	waited := <-waitChannel

	testutils.AssertEquals(t, true, waited)
}

// TestAsyncLogger_Open tests AsyncLogger.Open method of the AsyncLogger.
func TestAsyncLogger_Open(t *testing.T) {
	mockHandler := &MockHandler{}
	newAsyncLogger := &AsyncLogger{
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

	newAsyncLogger.Open(messageQueueSize)

	testutils.AssertNotNil(t, newAsyncLogger.baseLogger.(*baseAsyncLogger).messageQueue)
}

// TestAsyncLogger_Close tests AsyncLogger.Close method of the AsyncLogger.
func TestAsyncLogger_Close(t *testing.T) {
	mockHandler := &MockHandler{}
	newAsyncLogger := &AsyncLogger{
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

	newAsyncLogger.Close()

	testutils.AssertEquals(t, true, isChannelClosed(newAsyncLogger.baseLogger.(*baseAsyncLogger).messageQueue))
}
