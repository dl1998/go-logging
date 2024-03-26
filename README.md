# Go Logger

[![Go Reference](https://pkg.go.dev/badge/github.com/dl1998/go-logging.svg)](https://pkg.go.dev/github.com/dl1998/go-logging)

Go logger implements logger for Golang, current implementation is majorly inspired by Python logger.

## Installation

```bash
go get github.com/dl1998/go-logging
```

or

```bash
go install github.com/dl1998/go-logging@v1.0.0
```

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

### Custom Logger

Alternatively you could create application logger. To do this you would need to create a new logger.

- Standard logger

```go
applicationLogger := logger.New("application-logger", time.RFC3339)
```

- Structured logger

```go
applicationStructuredLogger := structuredlogger.New("application-logger", time.RFC3339)
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
applicationLogger.AddHandler(newConsoleHandler)
applicationLogger.AddHandler(newConsoleErrorHandler)
applicationLogger.AddHandler(newFileHandler)
```

Now it could be used to log the message, simply by calling respective level of logging and providing message with
arguments.

- Standard logger

```go
applicationLogger.Info("My message: %s.", "logged using application logger")
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

## Class Diagram

![Class Diagram](./docs/architecture/diagrams/png/class_diagram.png)

## Sequence Diagram - Create A New Logger

![Sequence Diagram](./docs/architecture/diagrams/png/create_new_logger.png)