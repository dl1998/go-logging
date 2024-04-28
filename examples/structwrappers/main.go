// Example that shows how to use struct, http.Request, http.Response wrappers
// with the standard/structured logger.
package main

import (
	"github.com/dl1998/go-logging/pkg/common/level"
	standard "github.com/dl1998/go-logging/pkg/logger"
	standardFormatter "github.com/dl1998/go-logging/pkg/logger/formatter"
	standardHandler "github.com/dl1998/go-logging/pkg/logger/handler"
	structured "github.com/dl1998/go-logging/pkg/structuredlogger"
	structuredFormatter "github.com/dl1998/go-logging/pkg/structuredlogger/formatter"
	structuredHandler "github.com/dl1998/go-logging/pkg/structuredlogger/handler"
	"net/http"
	"net/url"
	"time"
)

var (
	exampleStructTemplate = "Name: {Name}, Age: {Age}"
	exampleStructMapping  = map[string]string{
		"name": "Name",
		"age":  "Age",
	}

	exampleStruct = ExampleStruct{
		Name: "John",
		Age:  25,
	}

	exampleRequest = &http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: "http",
			Host:   "localhost",
			Path:   "/example",
		},
	}

	exampleResponse = &http.Response{
		StatusCode: http.StatusOK,
		Status:     "OK",
		Request:    exampleRequest,
	}
)

type ExampleStruct struct {
	Name string
	Age  int
}

func main() {
	exampleStandardLogger()
	exampleStructuredLogger()
}

// exampleStandardLogger is a sample function to show how to use struct,
// http.Request, http.Response wrappers with the standard logger.
func exampleStandardLogger() {
	applicationLogger := standard.New("main", time.DateTime)

	applicationFormatter := standardFormatter.New("%(datetime)\t[%(level)]\t%(message)")

	applicationHandler := standardHandler.NewConsoleErrorHandler(level.All, level.Null, applicationFormatter)

	applicationLogger.AddHandler(applicationHandler)

	applicationLogger.SetRequestTemplate("[{Method}] {URL}")
	applicationLogger.SetResponseTemplate("[{StatusCode}] {Status}")

	applicationLogger.WrapStruct(level.Info, exampleStructTemplate, exampleStruct)
	applicationLogger.WrapRequest(level.Info, exampleRequest)
	applicationLogger.WrapResponse(level.Info, exampleResponse)
}

// exampleStructuredLogger is a sample function to show how to use struct,
// http.Request, http.Response wrappers with the structured logger.
func exampleStructuredLogger() {
	applicationLogger := structured.New("main", time.DateTime)

	applicationFormatter := structuredFormatter.NewJSON(
		map[string]string{
			"timestamp": "%(timestamp)",
			"level":     "%(level)",
			"name":      "%(name)",
		},
		false,
	)

	applicationHandler := structuredHandler.NewConsoleErrorHandler(level.All, level.Null, applicationFormatter)

	applicationLogger.AddHandler(applicationHandler)

	applicationLogger.SetRequestMapping(map[string]string{
		"ExampleMethod": "Method",
		"ExampleURL":    "URL",
	})
	applicationLogger.SetResponseMapping(map[string]string{
		"ExampleStatusCode": "StatusCode",
		"ExampleStatus":     "Status",
	})

	applicationLogger.WrapStruct(level.Info, exampleStructMapping, exampleStruct, "hostname", "localhost")
	applicationLogger.WrapRequest(level.Info, exampleRequest, "hostname", "localhost")
	applicationLogger.WrapResponse(level.Info, exampleResponse, "hostname", "localhost")
}
