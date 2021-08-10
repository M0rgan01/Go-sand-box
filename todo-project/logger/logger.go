package logger

import (
	stdlog "log"
	"os"
)

// StdLogger is a struct that wraps the general logger provided by the Go
// standard library and causes it to meet the log.Logger interface
type StdLogger struct {
	Level int
}

// Current contains the logger used for the package level logging functions
var current StdLogger

var LevelMap = map[int]string{
	TraceLevel: "TRACE",
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERR",
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
}

const (
	TraceLevel = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

const DefaultLogLevel = InfoLevel
const SccLogLevel = "SCC_LOG_LEVEL"

func init() {
	levelAsString := GetLogLevel()
	numLevel, ok := getKeyFromValue(levelAsString)

	if !ok {
		numLevel = DefaultLogLevel
		stdlog.Printf("error when setting log level: %s is not a valid log level", levelAsString)
	}

	current = StdLogger{
		Level: numLevel,
	}
}

func GetLogLevel() string {
	logLevel, exists := os.LookupEnv(SccLogLevel)
	if !exists {
		logLevel = LevelMap[DefaultLogLevel]
	}
	return logLevel
}

func (l StdLogger) logf(level int, template string, args ...interface{}) {
	logLevel := "[" + LevelMap[level] + "]   "
	switch {
	case (level == TraceLevel) && (l.Level <= level):
		stdlog.Printf(logLevel+template, args...)
		break
	case (level == DebugLevel) && (l.Level <= level):
		stdlog.Printf(logLevel+template, args...)
		break
	case (level == InfoLevel) && (l.Level <= level):
		stdlog.Printf(logLevel+template, args...)
		break
	case (level == WarnLevel) && (l.Level <= level):
		stdlog.Printf(logLevel+template, args...)
		break
	case (level == ErrorLevel) && (l.Level <= level):
		stdlog.Printf(logLevel+template, args...)
		break
	case (level == PanicLevel) && (l.Level <= level):
		stdlog.Panicf(logLevel+template, args...)
	case (level == FatalLevel) && (l.Level <= level):
		stdlog.Fatalf(logLevel+template, args...)
	}
}

func (l StdLogger) log(level int, args ...interface{}) {
	logLevel := "[" + LevelMap[level] + "]   "
	switch {
	case (level == TraceLevel) && (l.Level <= level):
		stdlog.Print(append([]interface{}{logLevel}, args...)...)
		break
	case (level == DebugLevel) && (l.Level <= level):
		stdlog.Print(append([]interface{}{logLevel}, args...)...)
		break
	case (level == InfoLevel) && (l.Level <= level):
		stdlog.Print(append([]interface{}{logLevel}, args...)...)
		break
	case (level == WarnLevel) && (l.Level <= level):
		stdlog.Print(append([]interface{}{logLevel}, args...)...)
		break
	case (level == ErrorLevel) && (l.Level <= level):
		stdlog.Print(append([]interface{}{logLevel}, args...)...)
		break
	case (level == PanicLevel) && (l.Level <= level):
		stdlog.Panic(append([]interface{}{logLevel}, args...)...)
	case (level == FatalLevel) && (l.Level <= level):
		stdlog.Fatal(append([]interface{}{logLevel}, args...)...)
	}
}

// Trace logs a message at the Trace level
func Trace(msg ...interface{}) {
	current.log(TraceLevel, msg...)
}

// Tracef formats a message according to a format specifier and logs the
// message at the Trace level
func Tracef(template string, args ...interface{}) {
	current.logf(TraceLevel, template, args...)
}

// Debug logs a message at the Debug level
func Debug(msg ...interface{}) {
	current.log(DebugLevel, msg...)
}

// Debugf formats a message according to a format specifier and logs the
// message at the Debug level
func Debugf(template string, args ...interface{}) {
	current.logf(DebugLevel, template, args...)
}

// Info logs a message at the Info level
func Info(msg ...interface{}) {
	current.log(InfoLevel, msg...)
}

// Infof formats a message according to a format specifier and logs the
// message at the Info level
func Infof(template string, args ...interface{}) {
	current.logf(InfoLevel, template, args...)
}

// Warn logs a message at the Warn level
func Warn(msg ...interface{}) {
	current.log(WarnLevel, msg...)
}

// Warnf formats a message according to a format specifier and logs the
// message at the Warning level
func Warnf(template string, args ...interface{}) {
	current.logf(WarnLevel, template, args...)
}

// Error logs a message at the Error level
func Error(msg ...interface{}) {
	current.log(ErrorLevel, msg...)
}

// Errorf formats a message according to a format specifier and logs the
// message at the Error level
func Errorf(template string, args ...interface{}) {
	current.logf(ErrorLevel, template, args...)
}

// Panic logs a message at the Panic level and panics
func Panic(msg ...interface{}) {
	current.log(PanicLevel, msg...)
}

// Panicf formats a message according to a format specifier and logs the
// message at the Panic level and then panics
func Panicf(template string, args ...interface{}) {
	current.logf(PanicLevel, template, args...)
}

// Fatal logs a message at the Fatal level and exists the application
func Fatal(msg ...interface{}) {
	current.log(FatalLevel, msg...)
}

// Fatalf formats a message according to a format specifier and logs the
// message at the Fatal level and exits the application
func Fatalf(template string, args ...interface{}) {
	current.logf(FatalLevel, template, args...)
}

func getKeyFromValue(value string) (int, bool) {
	for k, v := range LevelMap {
		if v == value {
			return k, true
		}
	}
	return -1, false
}
