package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	errorColor    = "\033[38;5;9m"
	panicColor    = "\033[38;5;124m"
	fatalColor    = "\033[38;5;1m"
	blue          = "\033[34m"
	yellow        = "\033[1;33m"
	filePathColor = "\033[38;5;250m"
	lightOrange   = "\033[38;5;11m"
	defaultColor  = "\033[0m"
	message       = "\033[38;5;255m"
)

type LogLevel string

var (
	Info  LogLevel = "info"
	Debug LogLevel = "debug"
	Error LogLevel = "error"
)

func NewLogger() *Logger {
	return &Logger{
		level:    Info,
		printLog: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		errLog:   log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

type Logger struct {
	level    LogLevel
	errLog   *log.Logger
	printLog *log.Logger
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) GetLevel() LogLevel {
	return l.level
}

func (l *Logger) Info(v ...any) {
	if l.level == Error {
		return
	}

	msg := []any{message}
	msg = append(msg, v...)
	msg = append(msg, defaultColor)
	l.printLog.SetPrefix(blue + "[Info] " + filePathColor)
	l.printLog.Output(2, fmt.Sprint(msg...))
}

func (l *Logger) Json(data any) {
	jStr, err := json.Marshal(data)
	if err != nil {
		l.Error(err)
		return
	}
	l.Info(string(jStr))
}

func (l *Logger) Infof(format string, v ...any) {
	if l.level == Error {
		return
	}

	format = message + format + defaultColor + "\n"
	l.printLog.SetPrefix(blue + "[Info] " + filePathColor)
	l.printLog.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(v ...any) {
	if l.level != Debug {
		return
	}

	msg := []any{message}
	msg = append(msg, v...)
	msg = append(msg, defaultColor)
	l.printLog.SetPrefix(lightOrange + "[Debug] " + filePathColor)
	l.printLog.Output(2, fmt.Sprint(msg...))
}

func (l *Logger) Debugf(format string, v ...any) {
	if l.level != Debug {
		return
	}

	format = message + format + defaultColor + "\n"
	l.printLog.SetPrefix(lightOrange + "[Debug] " + filePathColor)
	l.printLog.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...any) {
	msg := []any{message}
	msg = append(msg, v...)
	msg = append(msg, defaultColor)
	l.errLog.SetPrefix(errorColor + "[Error] " + filePathColor)
	l.errLog.Output(2, fmt.Sprint(msg...))
}

func (l *Logger) Errorf(format string, v ...any) {
	format = message + format + defaultColor + "\n"
	l.errLog.SetPrefix(errorColor + "[Error] " + filePathColor)
	l.errLog.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Warning(v ...any) {
	if l.level == Error {
		return
	}

	msg := []any{message}
	msg = append(msg, v...)
	msg = append(msg, defaultColor)
	l.printLog.SetPrefix(yellow + "[Warning] " + filePathColor)
	l.printLog.Output(2, fmt.Sprint(msg...))
}

func (l *Logger) Warningf(format string, v ...any) {
	if l.level == Error {
		return
	}

	format = message + format + defaultColor + "\n"
	l.printLog.SetPrefix(yellow + "[Warning] " + filePathColor)
	l.printLog.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...any) {
	msg := []any{message}
	msg = append(msg, v...)
	msg = append(msg, defaultColor, "\n")
	l.errLog.SetPrefix(panicColor + "[Panic] " + filePathColor)

	s := fmt.Sprint(msg...)
	l.errLog.Output(3, s)
	panic(s)
}

func (l *Logger) Panicf(format string, v ...any) {

	l.errLog.SetPrefix(panicColor + "[Panic] " + filePathColor)

	s := fmt.Sprintf(format, v...)
	l.errLog.Output(3, s)
	panic(s)
}

func (l *Logger) Fatal(v ...any) {
	msg := []any{message}
	msg = append(msg, v...)
	msg = append(msg, defaultColor, "\n")
	l.errLog.SetPrefix(fatalColor + "[Fatal] " + filePathColor)
	l.errLog.Output(2, fmt.Sprint(msg...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...any) {
	format = message + format + defaultColor + "\n"
	l.errLog.SetPrefix(fatalColor + "[Fatal] " + filePathColor)
	l.errLog.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
