package log

import (
	"context"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type Log struct {
	// The logs are `io.Copy`'d to this in a mutex. It's common to set this to a
	// file, or leave it default which is `os.Stderr`. You can also set this to
	// something more adventurous, such as logging to Kafka.
	Out io.Writer
	// Hooks for the log instance. These allow firing events based on logging
	// levels and log entries. For example, to send errors to an error tracking
	// service, log to StatsD or dump the core on fatal errors.
	Hooks LevelHooks
	// All log entries pass through the formatter before logged to Out. The
	// included formatters are `TextFormatter` and `JSONFormatter` for which
	// TextFormatter is the default. In development (when a TTY is attached) it
	// logs with colors, but to a file it wouldn't. You can easily implement your
	// own that implements the `Formatter` interface, see the `README` or included
	// formatters for examples.
	Formatter Formatter

	// Flag for whether to log caller info (off by default)
	ReportCaller bool

	// The logging level the log should log at. This is typically (and defaults
	// to) `logrus.Info`, which allows Info(), Warn(), Error() and Fatal() to be
	// logged.
	Level Level
	// Used to sync writing to the log. Locking is enabled by Default
	mu MutexWrap
	// Reusable empty entry
	entryPool sync.Pool
	// Function to exit the application, defaults to `os.Exit()`
	ExitFunc exitFunc
}

type exitFunc func(int)

type MutexWrap struct {
	lock     sync.Mutex
	disabled bool
}

func (mw *MutexWrap) Lock() {
	if !mw.disabled {
		mw.lock.Lock()
	}
}

func (mw *MutexWrap) Unlock() {
	if !mw.disabled {
		mw.lock.Unlock()
	}
}

func (mw *MutexWrap) Disable() {
	mw.disabled = true
}

// New
func New() *Log {
	return &Log{
		Out:          os.Stderr,
		Formatter:    new(TextFormatter),
		Hooks:        make(LevelHooks),
		Level:        InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
}
func (l *Log) newEntry() *Entry {
	entry, ok := l.entryPool.Get().(*Entry)
	if ok {
		return entry
	}

	return NewEntry(l)
}

func (l *Log) releaseEntry(entry *Entry) {
	entry.Data = map[string]interface{}{}
	l.entryPool.Put(entry)
}

// WithField allocates a new entry and adds a field to it.
// Debug, Print, Info, Warn, Error, Fatal or Panic must be then applied to
// this new returned entry.
// If you want multiple fields, use `WithFields`.
func (l *Log) WithField(key string, value interface{}) *Entry {
	entry := l.newEntry()
	defer l.releaseEntry(entry)

	return entry.WithField(key, value)
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (l *Log) WithFields(fields Fields) *Entry {
	entry := l.newEntry()
	defer l.releaseEntry(entry)

	return entry.WithFields(fields)
}

// Add an error as single field to the log entry.  All it does is call
// `WithError` for the given `error`.
func (l *Log) WithError(err error) *Entry {
	entry := l.newEntry()
	defer l.releaseEntry(entry)

	return entry.WithError(err)
}

// Add a context to the log entry.
func (l *Log) WithContext(ctx context.Context) *Entry {
	entry := l.newEntry()
	defer l.releaseEntry(entry)

	return entry.WithContext(ctx)
}

// Overrides the time of the log entry.
func (l *Log) WithTime(t time.Time) *Entry {
	entry := l.newEntry()
	defer l.releaseEntry(entry)

	return entry.WithTime(t)
}

func (l *Log) Logf(level Level, format string, args ...interface{}) {
	if l.IsLevelEnabled(level) {
		entry := l.newEntry()
		entry.Logf(level, format, args...)
		l.releaseEntry(entry)
	}
}

// Defaultf Defaultf
func (l *Log) Defaultf(format string, args ...interface{}) {
	l.Logf(DefaultLevel, format, args...)
}

// Debugf Debugf
func (l *Log) Debugf(format string, args ...interface{}) {
	l.Logf(DebugLevel, format, args...)
}

// Infof Infof
func (l *Log) Infof(format string, args ...interface{}) {
	l.Logf(InfoLevel, format, args...)
}

// Noticef Noticef
func (l *Log) Noticef(format string, args ...interface{}) {
	l.Logf(NoticeLevel, format, args...)
}

// Warnf Warnf
func (l *Log) Warnf(format string, args ...interface{}) {
	l.Logf(WarnLevel, format, args...)
}

// Errorf Errorf
func (l *Log) Errorf(format string, args ...interface{}) {
	l.Logf(ErrorLevel, format, args...)
}

// Criticalf Criticalf
func (l *Log) Criticalf(format string, args ...interface{}) {
	l.Logf(CriticalLevel, format, args...)
}

// Alertf Alertf
func (l *Log) Alertf(format string, args ...interface{}) {
	l.Logf(AlertLevel, format, args...)
}

// Emergencyf Emergencyf
func (l *Log) Emergencyf(format string, args ...interface{}) {
	l.Logf(EmergencyLevel, format, args...)
}

// Log Log
func (l *Log) Log(level Level, args ...interface{}) {
	if l.IsLevelEnabled(level) {
		entry := l.newEntry()
		entry.Log(level, args...)
		l.releaseEntry(entry)
	}
}

// Default Default
func (l *Log) Default(args ...interface{}) {
	l.Log(DefaultLevel, args...)
}

// Debug Debug
func (l *Log) Debug(args ...interface{}) {
	l.Log(DebugLevel, args...)
}

// Info Info
func (l *Log) Info(args ...interface{}) {
	l.Log(InfoLevel, args...)
}

// Notice Notice
func (l *Log) Notice(args ...interface{}) {
	l.Log(NoticeLevel, args...)
}

// Warn Warn
func (l *Log) Warn(args ...interface{}) {
	l.Log(WarnLevel, args...)
}

// Error Error
func (l *Log) Error(args ...interface{}) {
	l.Log(ErrorLevel, args...)
}

// Critical Critical
func (l *Log) Critical(args ...interface{}) {
	l.Log(CriticalLevel, args...)
}

// Alert Alert
func (l *Log) Alert(args ...interface{}) {
	l.Log(AlertLevel, args...)
}

// Emergency Emergency
func (l *Log) Emergency(args ...interface{}) {
	l.Log(EmergencyLevel, args...)
}

// Logln Logln
func (l *Log) Logln(level Level, args ...interface{}) {
	if l.IsLevelEnabled(level) {
		entry := l.newEntry()
		entry.Logln(level, args...)
		l.releaseEntry(entry)
	}
}

// Defaultln Defaultln
func (l *Log) Defaultln(args ...interface{}) {
	l.Logln(DefaultLevel, args...)
}

// Debugln Debugln
func (l *Log) Debugln(args ...interface{}) {
	l.Logln(DebugLevel, args...)
}

// Infoln Infoln
func (l *Log) Infoln(args ...interface{}) {
	l.Logln(InfoLevel, args...)
}

// Noticeln Noticeln
func (l *Log) Noticeln(args ...interface{}) {
	l.Logln(NoticeLevel, args...)
}

// Warnln Warnln
func (l *Log) Warnln(args ...interface{}) {
	l.Logln(WarnLevel, args...)
}

// Errorln Errorln
func (l *Log) Errorln(args ...interface{}) {
	l.Logln(ErrorLevel, args...)
}

// Criticalln Criticalln
func (l *Log) Criticalln(args ...interface{}) {
	l.Logln(CriticalLevel, args...)
}

// Alertln Alertln
func (l *Log) Alertln(args ...interface{}) {
	l.Logln(AlertLevel, args...)
}

// Emergencyln Emergencyln
func (l *Log) Emergencyln(args ...interface{}) {
	l.Logln(EmergencyLevel, args...)
}

// Exit Exit
func (l *Log) Exit(code int) {
	runHandlers()

	if l.ExitFunc == nil {
		l.ExitFunc = os.Exit
	}

	l.ExitFunc(code)
}

//When file is opened with appending mode, it's safe to
//write concurrently to a file (within 4k message on Linux).
//In these cases user can choose to disable the lock.
func (l *Log) SetNoLock() {
	l.mu.Disable()
}

func (l *Log) level() Level {
	return Level(atomic.LoadUint32((*uint32)(&l.Level)))
}

// SetLevel sets the log level.
func (l *Log) SetLevel(level Level) {
	atomic.StoreUint32((*uint32)(&l.Level), uint32(level))
}

// GetLevel returns the log level.
func (l *Log) GetLevel() Level {
	return l.level()
}

// AddHook adds a hook to the log hooks.
func (l *Log) AddHook(hook Hook) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Hooks.Add(hook)
}

// IsLevelEnabled checks if the log level of the log is greater than the level param
func (l *Log) IsLevelEnabled(level Level) bool {
	return l.level() <= level
}

// SetFormatter sets the log formatter.
func (l *Log) SetFormatter(formatter Formatter) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Formatter = formatter
}

// SetOutput sets the log output.
func (l *Log) SetOutput(output io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Out = output
}

func (l *Log) SetReportCaller(reportCaller bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.ReportCaller = reportCaller
}

// ReplaceHooks replaces the log hooks and returns the old ones
func (l *Log) ReplaceHooks(hooks LevelHooks) LevelHooks {
	l.mu.Lock()
	oldHooks := l.Hooks
	l.Hooks = hooks
	l.mu.Unlock()

	return oldHooks
}
