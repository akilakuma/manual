package log

import (
	"fmt"
	"strings"
)

// Level type
type Level uint32

// Convert the Level to a string. E.g. PanicLevel becomes "panic".
func (level Level) String() string {
	if b, err := level.MarshalText(); err == nil {
		return string(b)
	} else {
		return "unknown"
	}
}

// ParseLevel takes a string level and returns the Logrus log level constant.
func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "DEFAULT":
		return DefaultLevel, nil
	case "DEBUG":
		return DebugLevel, nil
	case "INFO":
		return InfoLevel, nil
	case "NOTICE":
		return NoticeLevel, nil
	case "WARN":
		return WarnLevel, nil
	case "ERROR":
		return ErrorLevel, nil
	case "CRITICAL":
		return CriticalLevel, nil
	case "ALERT":
		return AlertLevel, nil
	case "EMERGENCY":
		return EmergencyLevel, nil
	}

	var l Level

	return l, fmt.Errorf("not a valid logrus Level: %q", lvl)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (level *Level) UnmarshalText(text []byte) error {
	l, err := ParseLevel(string(text))
	if err != nil {
		return err
	}

	*level = l

	return nil
}

func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case DefaultLevel:
		return []byte("DEFAULT"), nil
	case DebugLevel:
		return []byte("DEBUG"), nil
	case InfoLevel:
		return []byte("INFO"), nil
	case NoticeLevel:
		return []byte("NOTICE"), nil
	case WarnLevel:
		return []byte("WARN"), nil
	case ErrorLevel:
		return []byte("ERROR"), nil
	case CriticalLevel:
		return []byte("CRITICAL"), nil
	case AlertLevel:
		return []byte("ALERT"), nil
	case EmergencyLevel:
		return []byte("EMERGENCY"), nil
	}

	return nil, fmt.Errorf("not a valid logrus level %d", level)
}

// A constant exposing all logging levels
var AllLevels = []Level{
	DefaultLevel,
	DebugLevel,
	InfoLevel,
	NoticeLevel,
	WarnLevel,
	ErrorLevel,
	CriticalLevel,
	AlertLevel,
	EmergencyLevel,
}

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// // PanicLevel level, highest level of severity. Logs and then calls panic with the
	// // message passed to Debug, Info, ...
	// PanicLevel Level = iota
	// // FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// // logging level is set to Panic.
	// FatalLevel
	// // ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// // Commonly used for hooks to send errors to an error tracking service.
	// ErrorLevel
	// // WarnLevel level. Non-critical entries that deserve eyes.
	// WarnLevel
	// // InfoLevel level. General operational entries about what's going on inside the
	// // application.
	// InfoLevel
	// // DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	// DebugLevel
	// // TraceLevel level. Designates finer-grained informational events than the Debug.
	// TraceLevel

	DefaultLevel   Level = 0   // 日誌條目沒有指定的嚴重性級別
	DebugLevel     Level = 100 // 調試或跟蹤信息
	InfoLevel      Level = 200 // 常規信息，例如正在進行的狀態或性能
	NoticeLevel    Level = 300 // 正常但重要的事件，例如啟動，關閉或配置更改
	WarnLevel      Level = 400 // 警告事件可能會導致問題
	ErrorLevel     Level = 500 // 錯誤事件可能會導致問題
	CriticalLevel  Level = 600 // 嚴重事件會導致更嚴重的問題或中斷
	AlertLevel     Level = 700 // 一個人必須立即採取行動
	EmergencyLevel Level = 800 // 一個或多個系統無法使用
)
