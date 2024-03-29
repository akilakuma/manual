package log

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	bufferPool *sync.Pool

	// qualified package name, cached at first use
	logrusPackage string

	// Positions in the call stack when tracing to report the calling method
	minimumCallerDepth int

	// Used for caller information initialisation
	callerInitOnce sync.Once
)

const (
	maximumCallerDepth int = 25
	knownLogrusFrames  int = 4
)

func init() {
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	// start at the bottom of the stack before the package-name cache is primed
	minimumCallerDepth = 1
}

// Defines the key when adding errors using WithError.
var ErrorKey = "error"

// An entry is the final or intermediate Logrus logging entry. It contains all
// the fields passed with WithField{,s}. It's finally logged when Trace, Debug,
// Info, Warn, Error, Fatal or Panic is called on it. These objects can be
// reused and passed around as much as you wish to avoid field duplication.
type Entry struct {
	L *Log

	// Contains all the fields set by the user.
	Data Fields

	// Time at which the log entry was created
	Time time.Time

	// Level the log entry was logged at: Trace, Debug, Info, Warn, Error, Fatal or Panic
	// This field will be set on entry firing and the value will be equal to the one in Log struct field.
	Level Level

	// Calling method, with package name
	Caller *runtime.Frame

	// Message passed to Trace, Debug, Info, Warn, Error, Fatal or Panic
	Message string

	// When formatter is called in entry.log(), a Buffer may be set to entry
	Buffer *bytes.Buffer

	// Contains the context set by the user. Useful for hook processing etc.
	Context context.Context

	// err may contain a field formatting error
	err string
}

func NewEntry(l *Log) *Entry {
	return &Entry{
		L: l,
		// Default is three fields, plus one optional.  Give a little extra room.
		Data: make(Fields, 6),
	}
}

// Returns the bytes representation of this entry from the formatter.
func (entry *Entry) Bytes() ([]byte, error) {
	return entry.L.Formatter.Format(entry)
}

// Returns the string representation from the reader and ultimately the
// formatter.
func (entry *Entry) String() (string, error) {
	serialized, err := entry.Bytes()

	if err != nil {
		return "", err
	}

	str := string(serialized)

	return str, nil
}

// Add an error as single field (using the key defined in ErrorKey) to the Entry.
func (entry *Entry) WithError(err error) *Entry {
	return entry.WithField(ErrorKey, err)
}

// Add a context to the Entry.
func (entry *Entry) WithContext(ctx context.Context) *Entry {
	dataCopy := make(Fields, len(entry.Data))

	for k, v := range entry.Data {
		dataCopy[k] = v
	}

	return &Entry{L: entry.L, Data: dataCopy, Time: entry.Time, err: entry.err, Context: ctx}
}

// Add a single field to the Entry.
func (entry *Entry) WithField(key string, value interface{}) *Entry {
	return entry.WithFields(Fields{key: value})
}

// Add a map of fields to the Entry.
func (entry *Entry) WithFields(fields Fields) *Entry {
	data := make(Fields, len(entry.Data)+len(fields))

	for k, v := range entry.Data {
		data[k] = v
	}

	fieldErr := entry.err

	for k, v := range fields {
		isErrField := false

		if t := reflect.TypeOf(v); t != nil {
			switch t.Kind() { // nolint
			case reflect.Func:
				isErrField = true
			case reflect.Ptr:
				isErrField = t.Elem().Kind() == reflect.Func
			}
		}

		if isErrField {
			tmp := fmt.Sprintf("can not add field %q", k)

			if fieldErr != "" {
				fieldErr = entry.err + ", " + tmp
			} else {
				fieldErr = tmp
			}
		} else {
			data[k] = v
		}
	}

	return &Entry{L: entry.L, Data: data, Time: entry.Time, err: fieldErr, Context: entry.Context}
}

// Overrides the time of the Entry.
func (entry *Entry) WithTime(t time.Time) *Entry {
	dataCopy := make(Fields, len(entry.Data))
	for k, v := range entry.Data {
		dataCopy[k] = v
	}

	return &Entry{L: entry.L, Data: dataCopy, Time: t, err: entry.err, Context: entry.Context}
}

// getPackageName reduces a fully qualified function name to the package name
// There really ought to be to be a better way...
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")

		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

// getCaller retrieves the name of the first non-logrus calling function
func getCaller() *runtime.Frame {
	// cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, maximumCallerDepth)
		_ = runtime.Callers(0, pcs)

		// dynamic get the package name and the minimum caller depth
		for i := 0; i < maximumCallerDepth; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "getCaller") {
				logrusPackage = getPackageName(funcName)
				break
			}
		}

		minimumCallerDepth = knownLogrusFrames
	})

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if pkg != logrusPackage {
			return &f //nolint:scopelint
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}

func (entry Entry) HasCaller() (has bool) {
	return entry.L != nil &&
		entry.L.ReportCaller &&
		entry.Caller != nil
}

// This function is not declared with a pointer value because otherwise
// race conditions will occur when using multiple goroutines
func (entry Entry) log(level Level, msg string) {
	var buffer *bytes.Buffer

	// Default to now, but allow users to override if they want.
	//
	// We don't have to worry about polluting future calls to Entry#log()
	// with this assignment because this function is declared with a
	// non-pointer receiver.
	if entry.Time.IsZero() {
		entry.Time = time.Now()
	}

	entry.Level = level
	entry.Message = msg
	entry.L.mu.Lock()
	if entry.L.ReportCaller {
		entry.Caller = getCaller()
	}
	entry.L.mu.Unlock()

	entry.fireHooks()

	buffer = bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()

	defer bufferPool.Put(buffer)

	entry.Buffer = buffer

	entry.write()

	entry.Buffer = nil

	// To avoid Entry#log() returning a value that only would make sense for
	// panic() to use in Entry#Panic(), we avoid the allocation by checking
	// directly here.
	// if level <= PanicLevel {
	// 	panic(&entry)
	// }
}

func (entry *Entry) fireHooks() {
	entry.L.mu.Lock()
	defer entry.L.mu.Unlock()
	err := entry.L.Hooks.Fire(entry.Level, entry)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to fire hook: %v\n", err)
	}
}

func (entry *Entry) write() {
	entry.L.mu.Lock()
	defer entry.L.mu.Unlock()
	serialized, err := entry.L.Formatter.Format(entry)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to obtain reader, %v\n", err)
		return
	}

	if _, err = entry.L.Out.Write(serialized); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to log, %v\n", err)
	}
}

func (entry *Entry) Log(level Level, args ...interface{}) {
	if entry.L.IsLevelEnabled(level) {
		entry.log(level, fmt.Sprint(args...))
	}
}

func (entry *Entry) Default(args ...interface{}) {
	entry.Log(DefaultLevel, args...)
}
func (entry *Entry) Debug(args ...interface{}) {
	entry.Log(DebugLevel, args...)
}
func (entry *Entry) Info(args ...interface{}) {
	entry.Log(InfoLevel, args...)
}
func (entry *Entry) Notice(args ...interface{}) {
	entry.Log(NoticeLevel, args...)
}
func (entry *Entry) Warn(args ...interface{}) {
	entry.Log(WarnLevel, args...)
}
func (entry *Entry) Error(args ...interface{}) {
	entry.Log(ErrorLevel, args...)
}
func (entry *Entry) Critical(args ...interface{}) {
	entry.Log(CriticalLevel, args...)
}
func (entry *Entry) Alert(args ...interface{}) {
	entry.Log(AlertLevel, args...)
}
func (entry *Entry) Emergency(args ...interface{}) {
	entry.Log(EmergencyLevel, args...)
}

// Entry Printf family functions

func (entry *Entry) Logf(level Level, format string, args ...interface{}) {
	if entry.L.IsLevelEnabled(level) {
		entry.Log(level, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Defaultf(format string, args ...interface{}) {
	entry.Logf(DefaultLevel, format, args...)
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	entry.Logf(DebugLevel, format, args...)
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	entry.Logf(InfoLevel, format, args...)
}

func (entry *Entry) Noticef(format string, args ...interface{}) {
	entry.Logf(NoticeLevel, format, args...)
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	entry.Logf(WarnLevel, format, args...)
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	entry.Logf(ErrorLevel, format, args...)
}

func (entry *Entry) Criticalf(format string, args ...interface{}) {
	entry.Logf(CriticalLevel, format, args...)
}

func (entry *Entry) Alertf(format string, args ...interface{}) {
	entry.Logf(AlertLevel, format, args...)
}

func (entry *Entry) Emergencyf(format string, args ...interface{}) {
	entry.Logf(EmergencyLevel, format, args...)
}

// Entry Println family functions

func (entry *Entry) Logln(level Level, args ...interface{}) {
	if entry.L.IsLevelEnabled(level) {
		entry.Log(level, entry.sprintlnn(args...))
	}
}

func (entry *Entry) Defaultln(args ...interface{}) {
	entry.Logln(DefaultLevel, args...)
}
func (entry *Entry) Debugln(args ...interface{}) {
	entry.Logln(DebugLevel, args...)
}
func (entry *Entry) Infoln(args ...interface{}) {
	entry.Logln(InfoLevel, args...)
}
func (entry *Entry) Noticeln(args ...interface{}) {
	entry.Logln(NoticeLevel, args...)
}
func (entry *Entry) Warnln(args ...interface{}) {
	entry.Logln(WarnLevel, args...)
}
func (entry *Entry) Errorln(args ...interface{}) {
	entry.Logln(ErrorLevel, args...)
}
func (entry *Entry) Criticalln(args ...interface{}) {
	entry.Logln(CriticalLevel, args...)
}
func (entry *Entry) Alertln(args ...interface{}) {
	entry.Logln(AlertLevel, args...)
}
func (entry *Entry) Emergencyln(args ...interface{}) {
	entry.Logln(EmergencyLevel, args...)
}

// Sprintlnn => Sprint no newline. This is to get the behavior of how
// fmt.Sprintln where spaces are always added between operands, regardless of
// their type. Instead of vendoring the Sprintln implementation to spare a
// string allocation, we do the simplest thing.
func (entry *Entry) sprintlnn(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	return msg[:len(msg)-1]
}
