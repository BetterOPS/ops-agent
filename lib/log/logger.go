package log

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"os"
	"runtime"
	"strings"
)

const (
	SKIP       = 2
	DebugLevel = logrus.DebugLevel
	InfoLevel  = logrus.InfoLevel
	WarnLevel  = logrus.WarnLevel
	ErrorLevel = logrus.ErrorLevel
	FatalLevel = logrus.FatalLevel
	PanicLevel = logrus.PanicLevel
)

var logger *logrus.Logger = logrus.New()

func InitLogger(level string) {
	logger.Out = os.Stdout
	switch level {
	case "info":
		logger.Level = InfoLevel
	case "warn":
		logger.Level = WarnLevel
	case "error":
		logger.Level = ErrorLevel
	case "fatal":
		logger.Level = FatalLevel
	case "panic":
		logger.Level = PanicLevel
	default:
		logger.Level = DebugLevel
	}
	return
}
func Info(args ...interface{}) {
	if logger.Level >= InfoLevel {
		item := logger.WithFields(logrus.Fields{})
		item.Data["__file__"] = lineInfo(SKIP)
		item.Info(args...)
	}
}
func Infof(format string, args ...interface{}) {
	if logger.Level >= InfoLevel {
		item := logger.WithFields(logrus.Fields{})
		item.Data["_file__"] = lineInfo(SKIP)
		item.Infof(format, args...)
	}
}
func Debug(args ...interface{}) {
	if logger.Level >= DebugLevel {
		item := logger.WithFields(logrus.Fields{})
		item.Data["__file__"] = lineInfo(SKIP)
		item.Debug(args...)
	}
}
func Debugf(format string, args ...interface{}) {
	if logger.Level >= DebugLevel {
		item := logger.WithFields(logrus.Fields{})
		item.Data["__file__"] = lineInfo(SKIP)
		item.Debugf(format, args...)
	}
}
func Fatal(args ...interface{}) {
	if logger.Level >= FatalLevel {
		item := logger.WithFields(logrus.Fields{})
		item.Data["__file__"] = lineInfo(SKIP)
		item.Fatal(args...)
	}
}
func Fatalf(format string, args ...interface{}) {
	if logger.Level >= FatalLevel {
		item := logger.WithFields(logrus.Fields{})
		item.Data["__file__"] = lineInfo(SKIP)
		item.Fatalf(format, args...)
	}
}
func Error(args ...interface{}) {
	if logger.Level >= ErrorLevel {
		item := logger.WithFields(logrus.Fields{})
		item.Data["__file__"] = lineInfo(SKIP)
		item.Error(args...)
	}
}
func Errorf(format string, args ...interface{}) {
	if logger.Level >= ErrorLevel {
		item := logger.WithFields(logrus.Fields{})
		item.Data["__file__"] = lineInfo(SKIP)
		item.Errorf(format, args...)
	}
}
func lineInfo(skip int) (res string) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return fmt.Sprintf("%s:%d", "?", 0)
	}
	idx := strings.LastIndex(file, "/")
	if idx >= 0 {
		file = file[idx+1:]
	}
	return fmt.Sprintf("%s:%d", file, line)
}
