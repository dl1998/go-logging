# Go Logger

[![Go Reference](https://pkg.go.dev/badge/github.com/dl1998/go-logging.svg)](https://pkg.go.dev/github.com/dl1998/go-logging)
[![Go Report Card](https://goreportcard.com/badge/github.com/dl1998/go-logging)](https://goreportcard.com/report/github.com/dl1998/go-logging)
[![Coverage Status](https://coveralls.io/repos/github/dl1998/go-logging/badge.svg)](https://coveralls.io/github/dl1998/go-logging)

Go logger implements logger for Golang, current implementation is majorly inspired by Python logger.

## Installation

```bash
go get github.com/dl1998/go-logging
```

or

```bash
go install github.com/dl1998/go-logging@[version]
```

***Note: replace `[version]` with the version you want to install.***

## Usage

Check examples provided in the [examples](./examples).

Logger supports 11 logging levels + 2 (when not set):

- All (special level, cannot be used for logging)
- Trace
- Debug
- Verbose
- Info
- Notice
- Warning
- Severe
- Error
- Alert
- Critical
- Emergency
- Null (special level, cannot be used for logging)

### Default Logger

Default logger could be used like in the following example:

- Standard logger

  ```go
  logger.Warning("Message for logging: %s.", "my message")
  ```

- Structured logger

  ```go
  structuredlogger.Warning("message", "My message.")
  ```
  
  or
  
  ```go
  structuredlogger.Warning(map[string]string{
      "message": "My message.",
  })
  ```

By default, root logger prints on console only, and starting from Warning level. It could be changed by setting logging
level:

- Standard logger

  ```go
  logger.Configure(logger.NewConfiguration(logger.WithFromLevel(level.All)))
  ```

- Structured logger

  ```go
  structuredlogger.Configure(logger.NewConfiguration(logger.WithFromLevel(level.All)))
  ```

After changing log level to "All" it will print messages for any level.

You could also change the format of the default structured logger by setting the format (default: json).

```go
structuredlogger.Configure(logger.NewConfiguration(logger.WithFormat("key-value")))
```

All options available for the configuration are:

 For Standard Logger

  | Method               |               Default               | Description                                                                        |
  |----------------------|:-----------------------------------:|------------------------------------------------------------------------------------|
  | WithErrorLevel       |             level.Error             | Set logging level used to log raised or captured error.                            |
  | WithPanicLevel       |           level.Critical            | Set logging level used to log panic.                                               |
  | WithRequestTemplate  |     "Request: [{Method}] {URL}"     | Set template for the http.Request wrapper.                                         |
  | WithResponseTemplate | "Response: [{StatusCode}] {Status}" | Set template for the http.Response wrapper.                                        |
  | WithFromLevel        |            level.Warning            | Set logging level from which logger should log messages.                           |
  | WithToLevel          |             level.Null              | Set logging level till which logger should log messages.                           |
  | WithTemplate         |    "%(level):%(name):%(message)"    | Set template for logging message.                                                  |
  | WithFile             |                 ""                  | Set file where to log messages, if not set, then logging to file will be disabled. |
  | WithName             |               "root"                | Set logger name.                                                                   |
  | WithTimeFormat       |            time.RFC3339             | Set time format for logging message.                                               |

 For Structured Logger

  | Method                |                                                       Default                                                       | Description                                                                                                                           |
  |-----------------------|:-------------------------------------------------------------------------------------------------------------------:|---------------------------------------------------------------------------------------------------------------------------------------|
  | WithErrorLevel        |                                                     level.Error                                                     | Set logging level used to log raised or captured error.                                                                               |
  | WithPanicLevel        |                                                   level.Critical                                                    | Set logging level used to log panic.                                                                                                  |
  | WithRequestMapping    |                         map[string]string {<br/>"url": "URL",<br/>"method": "Method",<br/>}                         | Set mapping for the http.Request wrapper.                                                                                             |
  | WithResponseMapping   |                 map[string]string {<br/>"status": "Status",<br/>"status-code": "StatusCode",<br/>}                  | Set mapping for the http.Response wrapper.                                                                                            |
  | WithFromLevel         |                                                    level.Warning                                                    | Set logging level from which logger should log messages.                                                                              |
  | WithToLevel           |                                                     level.Null                                                      | Set logging level till which logger should log messages.                                                                              |
  | WithTemplate          | map[string]string {<br/>"timestamp": "%(timestamp)",<br/>"level":     "%(level)",<br/>"name":      "%(name)",<br/>} | Set template for logging structure.                                                                                                   |
  | WithFile              |                                                         ""                                                          | Set file where to log messages, if not set, then logging to file will be disabled.                                                    |
  | WithFormat            |                                                       "json"                                                        | Set format for structured logging.<br/><br/>Could be one of the following<br/><ul><li>json</li><li>key-value</li></ul>                |
  | WithPretty            |                                                        false                                                        | Set if json message should be pretty printed.<br/>*Option works only with "json" format.*                                             |
  | WithKeyValueDelimiter |                                                         "="                                                         | Set key-value delimiter (eg. "key=value", where '=' is the delimiter).<br/>*Option works only with "key-value" format.*               |
  | WithPairSeparator     |                                                         " "                                                         | Set key-value separator (eg. "key1=value1,key2=value2", where ',' is the separator).<br/>*Option works only with "key-value" format.* |
  | WithName              |                                                       "root"                                                        | Set logger name.                                                                                                                      |
  | WithTimeFormat        |                                                    time.RFC3339                                                     | Set time format for logging message.                                                                                                  |

### Custom Logger

Alternatively you could create application logger. To do this you would need to create a new logger.

- Standard logger

  ```go
  applicationLogger := logger.New("application-logger", time.RFC3339)
  ```

- Standard async logger

  ```go
  applicationLogger := logger.NewAsyncLogger("application-logger", time.RFC3339, 100)
  ```

- Structured logger

  ```go
  applicationLogger := structuredlogger.New("application-logger", time.RFC3339)
  ```

- Structured async logger

  ```go
  applicationLogger := structuredlogger.NewAsyncLogger("application-logger", time.RFC3339, 100)
  ```

After this you need to set up it, for this create a new formatter that says how to log the message by providing a
template.

#### Formatter

Available template options:

|    Option    |      Scope      | Description                                                                  |
|:------------:|:---------------:|------------------------------------------------------------------------------|
|   %(name)    |      Both       | Logger name.                                                                 |
|   %(level)   |      Both       | Log level name.                                                              |
|  %(levelnr)  |      Both       | Log level number.                                                            |
| %(datetime)  |      Both       | Current date and/or time formatted using time format. Default: time.RFC3339. |
| %(timestamp) |      Both       | Current timestamp.                                                           |
|   %(fname)   |      Both       | Name of the file from which logger has been called.                          |
|   %(fline)   |      Both       | Line in the file in which logger has been called.                            |
|  %(message)  | standard logger | Log message.                                                                 |

- Standard logger

  ```go
  applicationFormatter := formatter.New("%(datetime) [%(level)] %(message)")
  ```

- Structured logger
    - JSON format

      ```go
      applicationFormatter := formatter.NewJSON(map[string]string{
          "time":    "%(timestamp)",
          "level":   "%(level)",
      }, false)
      ```

    - Key-Value format

      ```go
      applicationFormatter := formatter.NewKeyValue(map[string]string{
          "time":    "%(timestamp)",
          "level":   "%(level)",
      }, "=", " ")
      ```

After creation of the formatter, you need to create a new handler that tells where to write log messages.

#### Handler

There are three predefined types of handler (for standard and structured logger each):

- Console Handler - it takes log level starting from which it would log messages, log level till which it would log
  messages, and formatter that tells how to log message. It logs messages to standard output.

  ```go
  newConsoleHandler := handler.NewConsoleHandler(level.Debug, level.Null, applicationFormatter)
  ```

- Console Error Handler - it takes log level starting from which it would log messages, log level till which it would
  log messages, and formatter that tells how to log message. It logs messages to error output.

  ```go
  newConsoleErrorHandler := handler.NewConsoleErrorHandler(level.Debug, level.Null, applicationFormatter)
  ```

- File Handler - it takes log level starting from which it would log messages, log level till which it would
  log messages, formatter that tells how to log message, and path to the file where to log those data.

  ```go
  newFileHandler := handler.NewFileHandler(level.Debug, level.Null, applicationFormatter, "system.log")
  ```

You could create your custom handler:

```go
customHandler := handler.New(level.Debug, level.Null, applicationFormatter, os.Stdout)
```

It takes two additional arguments writer for standard messages and for error messages. Standard message logs till
"Error" level, after this error writer is used.

After handler has been created it shall be registered.

```go
// Register console stdout handler.
applicationLogger.AddHandler(newConsoleHandler)
// Register console stderr handler.
applicationLogger.AddHandler(newConsoleErrorHandler)
// Register file handler.
applicationLogger.AddHandler(newFileHandler)
```

Now it could be used to log the message, simply by calling respective level of logging and providing message with
arguments.

- Standard logger

  ```go
  applicationLogger.Info("My message: %s.", "logged using application logger")
  ```

- Standard async logger

  ```go
  applicationLogger.Info("My message: %s.", "logged using application async logger")
  
  // Wait for all messages to be logged before exiting the program.
  applicationLogger.WaitToFinishLogging()
  ```

- Structured logger
  - Varargs

    ```go
    applicationLogger.Info("message", "Logged using structured logger with varargs.")
    ```

  - Map

    ```go
    applicationLogger.Info(map[string]string{
        "message": "Logged using structured logger with map.",
    })
    ```

- Structured async logger
  - Varargs

    ```go
    applicationLogger.Info("message", "Logged using structured logger with varargs.")
    
    // Wait for all messages to be logged before exiting the program.
	applicationLogger.WaitToFinishLogging()
    ```

  - Map

    ```go
    applicationLogger.Info(map[string]string{
        "message": "Logged using structured logger with map.",
    })
    
    // Wait for all messages to be logged before exiting the program.
	applicationLogger.WaitToFinishLogging()
    ```

#### Async Loggers - Additional Information

Async loggers are used to log messages asynchronously. It is useful when you want to log messages without blocking the
main thread. However, you need to wait for all messages to be logged before exiting the program. You can do this by
calling the `WaitToFinishLogging` method, it will block the main thread until all messages are logged. Alternatively,
you can close the logger by calling the `Close` method, it will close the message queue without waiting for all messages
to be logged. This is useful when you want to exit the program without waiting for all messages to be logged. After
calling the `Close` method, you can open the logger again by calling the `Open` method, it accepts the new message queue
size as an argument. `Open` method will open the logger with the new message queue size and start listening for the
messages.

Example that waits for all messages to be logged, then close the logger and open it again with a new message queue size:

```go
for index := 0; index < 1000; index++ {
    applicationLogger.Info("Counter: %d.", index)
}

// Wait for all messages to be logged before exiting the program.
applicationLogger.WaitToFinishLogging()

// Close the logger.
applicationLogger.Close()

// Open the logger with a new message queue size.
if err := applicationLogger.Open(100); err != nil {
    panic(err)
}
```

*Note: if you assign a new message queue size that is smaller than the number of messages sent to the queue, the logger
will add messages to the queue until it is not full, then it will wait (blocking the process) until the message from the
queue will be processed and free up the space in the message queue.*

### Wrappers

#### Error / Panic

You could wrap error or raise a new error and log error message using the logger. By default, it will log error message
using the `level.Error` level. However, it could be changed by setting the error level in the logger configuration.

- Standard logger

  ```go
  var err error
  
  // Raise Error with default error level (level.Error)
  err = applicationLogger.RaiseError("exit code: %d", 1)
  
  // Change error level
  applicationLogger.SetErrorLevel(level.Alert)
  
  // Capture Error with new error level (level.Alert)
  applicationLogger.CaptureError(err)
  ```

- Structured logger

  ```go
  var err error
  
  // Raise Error with default error level (level.Error) and additional fields
  err = applicationLogger.RaiseError("exit code: 1", "hostname", "localhost")
  
  // Change error level
  applicationLogger.SetErrorLevel(level.Alert)
  
  // Capture Error with new error level (level.Alert) and additional fields
  applicationLogger.CaptureError(err, "hostname", "localhost")
  ```
  
Similarly, you could panic and log panic message using the logger. By default, it will log panic message using the
`level.Critical` level. However, it could be changed by setting the panic level in the logger configuration.

- Standard logger

  ```go
  // Change panic level
  applicationLogger.SetPanicLevel(level.Emergency)

  // Raise Panic with new panic level (level.Emergency)
  applicationLogger.Panic("exit code: %d", 1)
  ```

- Structured logger

  ```go
  // Change panic level
  applicationLogger.SetPanicLevel(level.Emergency)
  
  // Raise Panic with new panic level (level.Emergency) and additional fields
  applicationLogger.Panic("exit code: 1", "hostname", "localhost")
  ```

#### Struct

You could wrap a struct and log its public fields using the logger. To do this, you need to provide template (standard
logger), mapping (structured logger) of the struct fields to the logger fields. Optionally, for structured logger you
could also provide additional fields that will be logged with the struct fields.

- Standard logger

  ```go
  type MyStruct struct {
      String string
      Int int
  }
  
  myStruct := MyStruct{
      String: "example",
      Int: 10,
  }
  
  applicationLogger.WrapStruct(level.Info, "{String}: {Int}", myStruct)
  ```

- Structured logger

  ```go
  type MyStruct struct {
      String string
      Int int
  }
  
  myStruct := MyStruct{
      String: "example",
      Int: 10,
  }
  
  applicationLogger.WrapStruct(level.Info, map[string]string{
      "log-string": "String",
      "log-int": "Int",
  }, myStruct, "hostname", "localhost")
  ```

#### http.Request and http.Response

You could wrap http.Request and http.Response and log their fields using the logger. By default, logger has predefined
template (standard logger), mapping (structured logger) for http.Request and http.Response fields. However, you could
change it by providing your own template / mapping. Optionally, for structured logger you could also provide additional
fields that will be logged with the http.Request and http.Response fields.

- Standard logger

  ```go
  request, _ := http.NewRequest("GET", "http://example.com", nil)
  response, _ := http.Get("http://example.com")
  
  // Set custom template for http.Request and http.Response
  applicationLogger.SetRequestTemplate("[{Method}] {URL}")
  applicationLogger.SetResponseTemplate("[{StatusCode}] {Status}")
  
  applicationLogger.WrapRequest(level.Info, request)
  applicationLogger.WrapResponse(level.Info, response)
  ```

- Structured logger

  ```go
  request, _ := http.NewRequest("GET", "http://example.com", nil)
  response, _ := http.Get("http://example.com")
  
  // Set custom mapping for http.Request and http.Response
  applicationLogger.SetRequestMapping(map[string]string{
      "Method": "Method",
      "Url": "URL",
  })
  applicationLogger.SetResponseMapping(map[string]string{
      "StatusCode": "StatusCode",
      "Status": "Status",
  })
  
  applicationLogger.WrapRequest(level.Info, request, "hostname", "localhost")
  applicationLogger.WrapResponse(level.Info, response, "hostname", "localhost")
  ```

### Reading Configuration from File

You could also read configuration from a file. Configuration file should be in one of the following formats: `*.json`,
`*.yaml`, `*.xml`. Configuration file should contain the following fields:

```text
- Loggers (array of loggers)
  - Name (string)
  - Time Format (string)
  - Error Level (string)
  - Panic Level (string)
  - Request Template (string)
  - Response Template (string)
  - Request Mapping (map of string to string)
  - Response Mapping (map of string to string)
  - Message Queue Size (int)
  - Handlers (array of handlers)
    - Type (string)
    - From Level (string)
    - To Level (string)
    - File (string)
    - Formatter (string)
      - Type (string)
      - Pretty Print (bool)
      - Pair Separator (string)
      - Key Value Delimiter (string)
      - Template (template)
        - String Value (string)
        - Map Value (map of string to string)
```

Example of the configuration files:

- JSON

  ```json
  {
    "loggers": [
      {
        "name": "example-logger",
        "time-format": "2006-01-02 15:04:05",
        "error-level": "error",
        "panic-level": "critical",
        "request-template": "Request: [{Method}] {URL}",
        "response-template": "Response: [{StatusCode}] {Status}",
        "request-mapping": {
          "method": "Method",
          "url": "URL"
        },
        "response-mapping": {
          "status-code": "StatusCode",
          "status": "Status"
        },
        "message-queue-size": 100,
        "handlers": [
          {
            "type": "stdout",
            "from-level": "all",
            "to-level": "severe",
            "formatter": {
              "type": "json",
              "pretty-print": false,
              "template": {
                "string": "%(datetime) - %(level) - %(message)",
                "map": {
                  "timestamp": "%(datetime)",
                  "level": "%(level)",
                  "name": "%(name)"
                }
              }
            }
          },
          {
            "type": "stderr",
            "from-level": "error",
            "to-level": "null",
            "formatter": {
              "type": "key-value",
              "pair-separator": " ",
              "key-value-delimiter": ":",
              "template": {
                "string": "%(datetime) - %(level) - %(message)",
                "map": {
                  "timestamp": "%(datetime)",
                  "level": "%(level)",
                  "name": "%(name)"
                }
              }
            }
          },
          {
            "type": "file",
            "from-level": "all",
            "to-level": "null",
            "file": "example.log",
            "formatter": {
              "type": "json",
              "pretty-print": true,
              "template": {
                "string": "%(datetime) - %(level) - %(message)",
                "map": {
                  "timestamp": "%(datetime)",
                  "level": "%(level)",
                  "name": "%(name)"
                }
              }
            }
          }
        ]
      }
    ]
  }
  ```

- YAML

  ```yaml
  loggers:
    - name: example-logger
      time-format: "2006-01-02 15:04:05"
      error-level: error
      panic-level: critical
      request-template: "Request: [{Method}] {URL}"
      response-template: "Response: [{StatusCode}] {Status}"
      request-mapping:
        method: Method
        url: URL
      response-mapping:
        status-code: StatusCode
        status: Status
      message-queue-size: 100
      handlers:
        - type: stdout
          from-level: all
          to-level: severe
          formatter:
            type: json
            pretty-print: false
            template:
            string: "%(datetime) - %(level) - %(message)"
            map:
              timestamp: "%(datetime)"
              level: "%(level)"
              name: "%(name)"
        - type: stderr
          from-level: error
          to-level: "null"
          formatter:
            type: key-value
            pair-separator: " "
            key-value-delimiter: ":"
            template:
            string: "%(datetime) - %(level) - %(message)"
            map:
              timestamp: "%(datetime)"
              level: "%(level)"
              name: "%(name)"
        - type: file
          from-level: all
          to-level: "null"
          file: example.log
          formatter:
            type: json
            pretty-print: true
            template:
            string: "%(datetime) - %(level) - %(message)"
            map:
              timestamp: "%(datetime)"
              level: "%(level)"
              name: "%(name)"
  ```
  
- XML

  ```xml
  <root>
    <loggers>
      <logger>
        <name>example-logger</name>
        <time-format>2006-01-02 15:04:05</time-format>
        <error-level>error</error-level>
        <panic-level>critical</panic-level>
        <request-template>Request: [{Method}] {URL}</request-template>
        <response-template>Response: [{StatusCode}] {Status}</response-template>
        <request-mapping>
          <method>Method</method>
          <url>URL</url>
        </request-mapping>
        <response-mapping>
          <status-code>StatusCode</status-code>
          <status>Status</status>
        </response-mapping>
        <message-queue-size>100</message-queue-size>
        <handlers>
          <handler>
            <type>stdout</type>
            <from-level>all</from-level>
            <to-level>severe</to-level>
            <formatter>
              <type>json</type>
              <pretty-print>false</pretty-print>
              <template>
                <string>%(datetime) - %(level) - %(message)</string>
                <map>
                  <timestamp>%(datetime)</timestamp>
                  <level>%(level)</level>
                  <name>%(name)</name>
                </map>
              </template>
            </formatter>
          </handler>
          <handler>
            <type>stderr</type>
            <from-level>error</from-level>
            <to-level>null</to-level>
            <formatter>
              <type>key-value</type>
              <pair-separator> </pair-separator>
              <key-value-delimiter>:</key-value-delimiter>
              <template>
                <string>%(datetime) - %(level) - %(message)</string>
                <map>
                  <timestamp>%(datetime)</timestamp>
                  <level>%(level)</level>
                  <name>%(name)</name>
                </map>
              </template>
            </formatter>
          </handler>
          <handler>
            <type>file</type>
            <from-level>all</from-level>
            <to-level>null</to-level>
            <file>example.log</file>
            <formatter>
              <type>json</type>
              <pretty-print>true</pretty-print>
              <template>
                <string>%(datetime) - %(level) - %(message)</string>
                <map>
                  <timestamp>%(datetime)</timestamp>
                  <level>%(level)</level>
                  <name>%(name)</name>
                </map>
              </template>
            </formatter>
          </handler>
        </handlers>
      </logger>
    </loggers>
  </root>
  ```

To create a logger from the configuration file, you need to:

1. Create a new Parser with the Configuration object. You shall use parser from the `logger` or `structuredlogger`.
   1. Create a new Configuration object manually and initialize parser with it.
      1. Parse configuration file to receive the Configuration. You could do this by calling the `ReadFromJSON`,
      `ReadFromYAML`, `ReadFromXML` methods respectively, it will return the Configuration object.

          ```go
          // Parse configuration from JSON file.
          newConfiguration, err := parser.ReadFromJSON("path/to/configuration/file.json")
          if err != nil {
              panic(err)
          }
          
          // Parse configuration from YAML file.
          newConfiguration, err := parser.ReadFromYAML("path/to/configuration/file.yaml")
          if err != nil {
              panic(err)
          }
          
          // Parse configuration from XML file.
          newConfiguration, err := parser.ReadFromXML("path/to/configuration/file.xml")
          if err != nil {
              panic(err)
          }
          ```
      2. Create a new Parser with the Configuration object. You shall use parser from the `logger` or `structuredlogger`
      packages respectively, depending on which one you need.

          ```go
          newParser := parser.NewParser(newConfiguration)
          ```
   2. Create Parser from the configuration file directly.

      ```go
      // Create a new Parser from JSON configuration file.
      newParser, err := parser.ParseJSON("path/to/configuration/file.json")
      if err != nil {
          panic(err)
      }
      
      // Create a new Parser from YAML configuration file.
	  newParser, err := parser.ParseYAML("path/to/configuration/file.yaml")
      if err != nil {
          panic(err)
      }
		
      // Create a new Parser from XML configuration file.
	  newParser, err := parser.ParseXML("path/to/configuration/file.xml")
      if err != nil {
          panic(err)
      }
      ```

2. Get a logger from the Parser.

    ```go
    // Standard Logger
    newLogger := newParser.GetLogger("example-logger")
    
    // Async Logger
    newLogger := newParser.GetAsyncLogger("example-logger")
    ```
