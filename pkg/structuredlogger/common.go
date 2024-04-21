package structuredlogger

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	"github.com/dl1998/go-logging/pkg/structuredlogger/handler"
	"github.com/dl1998/go-logging/pkg/structuredlogger/logrecord"
)

// baseLoggerInterface defines low level logging interface.
type baseLoggerInterface interface {
	Log(level level.Level, parameters ...any)
	Name() string
	SetName(name string)
	Handlers() []handler.Interface
	AddHandler(handlerInterface handler.Interface)
	RemoveHandler(handlerInterface handler.Interface)
}

// baseLogger struct contains basic fields for the logger.
type baseLogger struct {
	name       string
	timeFormat string
	handlers   []handler.Interface
}

// convertParametersToMap converts parameters to map[string]interface{}.
func convertParametersToMap(parameters ...any) map[string]interface{} {
	var parametersMap = make(map[string]interface{})
	parametersCount := len(parameters)

	if parametersCount == 1 {
		parametersMap = parameters[0].(map[string]interface{})
	} else if parametersCount > 1 {
		if parametersCount%2 != 0 {
			parametersCount--
		}
		for index := 0; index < parametersCount; index += 2 {
			parametersMap[parameters[index].(string)] = parameters[index+1]
		}
	}

	return parametersMap
}

// Log logs interpolated message with the provided level.Level.
func (logger *baseLogger) Log(logLevel level.Level, parameters ...any) {
	var parametersMap = convertParametersToMap(parameters...)

	logRecord := logrecord.New(logger.name, logLevel, logger.timeFormat, parametersMap, 4)

	for _, registeredHandler := range logger.handlers {
		registeredHandler.Write(logRecord)
	}
}

// Name return baseLogger name.
func (logger *baseLogger) Name() string {
	return logger.name
}

// SetName sets a new name for the baseLogger.
func (logger *baseLogger) SetName(name string) {
	logger.name = name
}

// Handlers returns a list of the registered handler.Interface objects for the
// baseLogger.
func (logger *baseLogger) Handlers() []handler.Interface {
	return logger.handlers
}

// AddHandler register a new handler.Interface for the baseLogger.
func (logger *baseLogger) AddHandler(handlerInterface handler.Interface) {
	logger.handlers = append(logger.handlers, handlerInterface)
}

// RemoveHandler removes a handler.Interface from the baseLogger handlers.
func (logger *baseLogger) RemoveHandler(handlerInterface handler.Interface) {
	newSlice := make([]handler.Interface, 0)
	for _, element := range logger.handlers {
		if element != handlerInterface {
			newSlice = append(newSlice, element)
		}
	}
	logger.handlers = newSlice
}
