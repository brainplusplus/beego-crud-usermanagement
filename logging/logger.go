package logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"runtime"
	"strings"
	"fmt"
)

type AppLogger struct {
	Hostname string
	*logrus.Logger
}

func MustGetLogger() *AppLogger {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	AppLogger := &AppLogger{hostname, logrus.StandardLogger()}
	return AppLogger
}

func (tx *AppLogger) Info(args ...interface{}) {
	fields(tx).Info(args...)
}

func (tx *AppLogger) Infof(format string, args ...interface{}) {
	fields(tx).Infof(format, args...)
}

func (tx *AppLogger) Debug(args ...interface{}) {
	fields(tx).Debug(args...)
}

func (tx *AppLogger) Debugf(format string, args ...interface{}) {
	fields(tx).Debugf(format, args...)
}

func (tx *AppLogger) Warn(args ...interface{}) {
	fields(tx).Warn(args...)
}

func (tx *AppLogger) Warnf(format string, args ...interface{}) {
	fields(tx).Warnf(format, args...)
}

func (tx *AppLogger) Error(args ...interface{}) {
	fields(tx).Error(args...)
}

func (tx *AppLogger) Errorf(format string, args ...interface{}) {
	fields(tx).Errorf(format, args...)
}

// DebugfWithId write formatted debug level log with added log_id field
func (tx *AppLogger) DebugfWithId(id string, format string, args ...interface{}) {
	fields(tx).WithField("log_id", id).Debugf(format, args...)
}

// InfofWithId write formatted info level log with added log_id field
func (tx *AppLogger) InfofWithId(id string, format string, args ...interface{}) {
	fields(tx).WithField("log_id", id).Infof(format, args...)
}

// InfoWithId write info level log with added log_id field
func (tx *AppLogger) InfoWithId(id string, args ...interface{}) {
	fields(tx).WithField("log_id", id).Info(args...)
}

// ErrorfWithId write formatted error level log with added log_id field
func (tx *AppLogger) ErrorfWithId(id string, format string, args ...interface{}) {
	fields(tx).WithField("log_id", id).Errorf(format, args...)
}

// ErrorWithId write error level log with added log_id field
func (tx *AppLogger) ErrorWithId(id string, args ...interface{}) {
	fields(tx).WithField("log_id", id).Error(args...)
}

func fields(tx *AppLogger) *logrus.Entry {
	file, line := getCaller()
	return tx.Logger.WithField("hostname", tx.Hostname).WithField("time", time.Now().UTC().Format(time.RFC3339)).WithField("source", fmt.Sprintf("%s:%d", file, line))
}

func getCaller() (string, int) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	return file, line
}
