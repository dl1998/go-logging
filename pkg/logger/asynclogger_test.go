package logger

import (
	"github.com/dl1998/go-logging/internal/testutils"
	"github.com/dl1998/go-logging/pkg/logger/handler"
	"github.com/dl1998/go-logging/pkg/logger/logrecord"
	"sync"
	"testing"
	"time"
)

var (
	name             = "test"
	messageQueueSize = 1
	testTimeFormat   = time.RFC3339
)

// TestBaseAsyncLogger_WaitToFinishLogging tests
// baseAsyncLogger.WaitToFinishLogging method of the baseAsyncLogger.
func TestBaseAsyncLogger_WaitToFinishLogging(t *testing.T) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := &baseAsyncLogger{
		baseLogger: &baseLogger{
			name:     name,
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
			name:     name,
			handlers: []handler.Interface{mockHandler},
		},
		messageQueue: make(chan logrecord.Interface, messageQueueSize),
		waitGroup:    sync.WaitGroup{},
	}

	newBaseAsyncLogger.Open(messageQueueSize)

	testutils.AssertNotNil(t, newBaseAsyncLogger.messageQueue)
}

// BenchmarkBaseAsyncLogger_Open benchmarks baseAsyncLogger.Open method of the
// baseAsyncLogger.
func BenchmarkBaseAsyncLogger_Open(b *testing.B) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := &baseAsyncLogger{
		baseLogger: &baseLogger{
			name:     name,
			handlers: []handler.Interface{mockHandler},
		},
		messageQueue: make(chan logrecord.Interface, messageQueueSize),
		waitGroup:    sync.WaitGroup{},
	}

	for index := 0; index < b.N; index++ {
		newBaseAsyncLogger.Open(messageQueueSize)
	}
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
			name:     name,
			handlers: []handler.Interface{mockHandler},
		},
		messageQueue: make(chan logrecord.Interface, messageQueueSize),
		waitGroup:    sync.WaitGroup{},
	}

	newBaseAsyncLogger.Close()

	testutils.AssertEquals(t, true, isChannelClosed(newBaseAsyncLogger.messageQueue))
}

// BenchmarkBaseAsyncLogger_Close benchmarks baseAsyncLogger.Close method of the
// baseAsyncLogger.
func BenchmarkBaseAsyncLogger_Close(b *testing.B) {
	loggers := make([]*baseAsyncLogger, b.N)
	mockHandler := &MockHandler{}
	newBaseLogger := &baseLogger{
		name:     name,
		handlers: []handler.Interface{mockHandler},
	}

	for index := range loggers {
		loggers[index] = &baseAsyncLogger{
			baseLogger:   newBaseLogger,
			messageQueue: make(chan logrecord.Interface, messageQueueSize),
			waitGroup:    sync.WaitGroup{},
		}
	}

	b.ResetTimer()

	for _, logger := range loggers {
		logger.Close()
	}
}

// TestBaseAsyncLogger_Log tests baseAsyncLogger.Log method of the baseAsyncLogger.
func TestBaseAsyncLogger_Log(t *testing.T) {
	mockHandler := &MockHandler{}
	newBaseAsyncLogger := &baseAsyncLogger{
		baseLogger: &baseLogger{
			name:     name,
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
			name:     name,
			handlers: []handler.Interface{mockHandler},
		},
		messageQueue: make(chan logrecord.Interface, b.N),
		waitGroup:    sync.WaitGroup{},
	}

	b.ResetTimer()

	for index := 0; index < b.N; index++ {
		newBaseAsyncLogger.Log(0, "test")
	}
}

// TestNewAsyncLogger tests NewAsyncLogger function of the baseAsyncLogger.
func TestNewAsyncLogger(t *testing.T) {
	newAsyncLogger := NewAsyncLogger(name, testTimeFormat, messageQueueSize)

	testutils.AssertNotNil(t, newAsyncLogger)
	testutils.AssertEquals(t, name, newAsyncLogger.Name())
	testutils.AssertEquals(t, messageQueueSize, cap(newAsyncLogger.baseLogger.(*baseAsyncLogger).messageQueue))
}

// BenchmarkNewAsyncLogger benchmarks NewAsyncLogger function of the
// baseAsyncLogger.
func BenchmarkNewAsyncLogger(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NewAsyncLogger(name, testTimeFormat, messageQueueSize)
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
					name:     name,
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

// BenchmarkAsyncLogger_WaitToFinishLogging benchmarks
// AsyncLogger.WaitToFinishLogging method of the AsyncLogger.
func BenchmarkAsyncLogger_WaitToFinishLogging(b *testing.B) {
	mockHandler := &MockHandler{}
	newAsyncLogger := &AsyncLogger{
		Logger: &Logger{
			baseLogger: &baseAsyncLogger{
				baseLogger: &baseLogger{
					name:     name,
					handlers: []handler.Interface{mockHandler},
				},
				messageQueue: make(chan logrecord.Interface, messageQueueSize),
				waitGroup:    sync.WaitGroup{},
			},
		},
	}

	for index := 0; index < b.N; index++ {
		newAsyncLogger.WaitToFinishLogging()
	}
}

// TestAsyncLogger_Open tests AsyncLogger.Open method of the AsyncLogger.
func TestAsyncLogger_Open(t *testing.T) {
	mockHandler := &MockHandler{}
	newAsyncLogger := &AsyncLogger{
		Logger: &Logger{
			baseLogger: &baseAsyncLogger{
				baseLogger: &baseLogger{
					name:     name,
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

// BenchmarkAsyncLogger_Open benchmarks AsyncLogger.Open method of the
// AsyncLogger.
func BenchmarkAsyncLogger_Open(b *testing.B) {
	mockHandler := &MockHandler{}
	newAsyncLogger := &AsyncLogger{
		Logger: &Logger{
			baseLogger: &baseAsyncLogger{
				baseLogger: &baseLogger{
					name:     name,
					handlers: []handler.Interface{mockHandler},
				},
				messageQueue: make(chan logrecord.Interface, messageQueueSize),
				waitGroup:    sync.WaitGroup{},
			},
		},
	}

	for index := 0; index < b.N; index++ {
		newAsyncLogger.Open(messageQueueSize)
	}
}

// TestAsyncLogger_Close tests AsyncLogger.Close method of the AsyncLogger.
func TestAsyncLogger_Close(t *testing.T) {
	mockHandler := &MockHandler{}
	newAsyncLogger := &AsyncLogger{
		Logger: &Logger{
			baseLogger: &baseAsyncLogger{
				baseLogger: &baseLogger{
					name:     name,
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

// BenchmarkAsyncLogger_Close benchmarks AsyncLogger.Close method of the
// AsyncLogger.
func BenchmarkAsyncLogger_Close(b *testing.B) {
	loggers := make([]*AsyncLogger, b.N)
	mockHandler := &MockHandler{}
	newBaseLogger := &baseLogger{
		name:     name,
		handlers: []handler.Interface{mockHandler},
	}

	for index := range loggers {
		loggers[index] = &AsyncLogger{
			Logger: &Logger{
				baseLogger: &baseAsyncLogger{
					baseLogger:   newBaseLogger,
					messageQueue: make(chan logrecord.Interface, messageQueueSize),
					waitGroup:    sync.WaitGroup{},
				},
			},
		}
	}

	b.ResetTimer()

	for _, logger := range loggers {
		logger.Close()
	}
}
