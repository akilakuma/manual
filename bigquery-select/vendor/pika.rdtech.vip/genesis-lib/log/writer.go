package log

import (
	"bufio"
	"io"
	"runtime"
)

// Writer at INFO level. See WriterLevel for details.
func (l *Log) Writer() *io.PipeWriter {
	return l.WriterLevel(InfoLevel)
}

// WriterLevel returns an io.Writer that can be used to write arbitrary text to
// the l at the given log level. Each line written to the writer will be
// printed in the usual way using formatters and hooks. The writer is part of an
// io.Pipe and it is the callers responsibility to close the writer when done.
// This can be used to override the standard library log easily.
func (l *Log) WriterLevel(level Level) *io.PipeWriter {
	return NewEntry(l).WriterLevel(level)
}

func (entry *Entry) Writer() *io.PipeWriter {
	return entry.WriterLevel(InfoLevel)
}

func (entry *Entry) WriterLevel(level Level) *io.PipeWriter {
	reader, writer := io.Pipe()

	var printFunc func(args ...interface{})

	switch level {
	case DefaultLevel:
		printFunc = entry.Default
	case DebugLevel:
		printFunc = entry.Debug
	case InfoLevel:
		printFunc = entry.Info
	case NoticeLevel:
		printFunc = entry.Notice
	case WarnLevel:
		printFunc = entry.Warn
	case ErrorLevel:
		printFunc = entry.Error
	case CriticalLevel:
		printFunc = entry.Critical
	case AlertLevel:
		printFunc = entry.Alert
	case EmergencyLevel:
		printFunc = entry.Emergency
	default:
		printFunc = entry.Default
	}

	go entry.writerScanner(reader, printFunc)
	runtime.SetFinalizer(writer, writerFinalizer)

	return writer
}

func (entry *Entry) writerScanner(reader *io.PipeReader, printFunc func(args ...interface{})) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		printFunc(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		entry.Errorf("Error while reading from Writer: %s", err)
	}

	reader.Close() // nolint
}

func writerFinalizer(writer *io.PipeWriter) {
	writer.Close() // nolint
}
