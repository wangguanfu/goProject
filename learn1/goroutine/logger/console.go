package logger

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
	level int
}


func NewConsoleLogger(config map[string]string) (log LogInterface, err error) {
	logLevel, ok := config["level"]
	if !ok {
		err = fmt.Errorf("not found log_level ")
		return
	}

	level := getLogLevel(logLevel)
	log = &ConsoleLogger{
		level: level,
	}
	return
}

func (c *ConsoleLogger) Init() {

}

func (c *ConsoleLogger) SetLevel(level int) {
	if level < LevelDebug || level > LevelFatal {
		level = LevelDebug
	}
	c.level = level
}


func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	if c.level > LevelDebug {
		return
	}

	logData := writeLog(LevelDebug, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s:%d) %s\n", logData.TimeStr,
		logData.LevelStr, logData.Filename, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	if c.level > LevelTrace {
		return
	}

	logData := writeLog(LevelTrace, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s:%d) %s\n", logData.TimeStr,
		logData.LevelStr, logData.Filename, logData.FuncName, logData.LineNo, logData.Message)
}
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > LevelInfo {
		return
	}

	logData := writeLog(LevelInfo, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s:%d) %s\n", logData.TimeStr,
		logData.LevelStr, logData.Filename, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > LevelWarn {
		return
	}

	logData := writeLog(LevelWarn, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s:%d) %s\n", logData.TimeStr,
		logData.LevelStr, logData.Filename, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > LevelError {
		return
	}

	logData := writeLog(LevelError, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s:%d) %s\n", logData.TimeStr,
		logData.LevelStr, logData.Filename, logData.FuncName, logData.LineNo, logData.Message)
}
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > LevelFatal {
		return
	}

	logData := writeLog(LevelFatal, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s:%d) %s\n", logData.TimeStr,
		logData.LevelStr, logData.Filename, logData.FuncName, logData.LineNo, logData.Message)
}

func (c *ConsoleLogger) Close() {

}