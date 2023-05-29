package log

import (
	"context"
	"io"
	"time"
)

type Fields map[string]interface{}

type Logger interface {
	AddHook(hook Hook)
	Exit(code int)
	GetLevel() Level
	IsLevelEnabled(level Level) bool
	// ReplaceHooks(hooks LevelHooks) LevelHooks
	SetFormatter(formatter Formatter)
	SetLevel(level Level)
	SetNoLock()
	SetOutput(output io.Writer)
	SetReportCaller(reportCaller bool)
	WithContext(ctx context.Context) *Entry
	WithError(err error) *Entry
	WithField(key string, value interface{}) *Entry
	WithFields(fields Fields) *Entry
	WithTime(t time.Time) *Entry
	Writer() *io.PipeWriter
	WriterLevel(level Level) *io.PipeWriter

	Defaultf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Noticef(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Criticalf(format string, args ...interface{})
	Alertf(format string, args ...interface{})
	Emergencyf(format string, args ...interface{})

	Default(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Notice(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Critical(args ...interface{})
	Alert(args ...interface{})
	Emergency(args ...interface{})

	Defaultln(args ...interface{})
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Noticeln(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})
	Criticalln(args ...interface{})
	Alertln(args ...interface{})
	Emergencyln(args ...interface{})
}
