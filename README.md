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

Logger supports 11 logging levels + 1 (when non set):

- None (special level, cannot be used for logging)
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

Default logger could be used like in the following example:

```go
logging.Warning("Message for logging: %s.", "my message")
```

By default, root logger prints on console only, and starting from Warning level. It could be changed by setting logging 
level:

```go
logging.SetLevel(loglevel.None)
```

After changing log level to "None" it will print messages for any level.

Alternatively you could create application logger. To do this you would need to create a new logger.

```go
applicationLogger := logger.New("application-logger")
```

After this you need to set up it, for this create a new formatter that says how to log the message by providing a
template.

```go
applicationFormatter := formatter.New("%(isotime) [%(level)] %(message)")
```

After creation of the formatter, you need to create a new handler that tells where to write log messages.

There are two predefined types of handler:

- Console Handler - it takes log level starting from which it would log messages, and formatter that tells how to log
message.

```go
newConsoleHandler := handler.NewConsoleHandler(loglevel.Debug, applicationFormatter)
```

- File Handler - it takes log level starting from which it would log messages, formatter that tells how to log message,
and path to the file where to log those data.

```go
newFileHandler := handler.NewFileHandler(loglevel.Debug, applicationFormatter, "system.log")
```

After handler has been created it shall be registered.

```go
applicationLogger.AddHandler(newConsoleHandler)
applicationLogger.AddHandler(newFileHandler)
```

Now it could be used to log the message, simply by calling respective level of logging and providing message with
arguments.

```go
applicationLogger.Info("My message: %s.", "logged using application logger")
```

## Class Diagram

![Class Diagram](./docs/architecture/diagrams/png/class_diagram.png)