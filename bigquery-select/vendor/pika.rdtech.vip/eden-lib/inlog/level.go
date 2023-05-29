package inlog

type Level uint32

const (
	// 日誌條目沒有指定的嚴重性級別。
	LevelDefault Level = 0
	// 調試或跟踪信息。
	LevelDebug Level = 100
	// 常規信息，例如正在進行的狀態或性能。
	LevelInfo Level = 200
	// 正常但重要的事件，例如啟動，關閉或配置更改。
	LevelNotice Level = 300
	// 警告事件可能會導致問題。
	LevelWarn Level = 400
	// 錯誤事件可能會導致問題。
	LevelError Level = 500
	// 嚴重事件會導致更嚴重的問題或中斷。
	LevelCritical Level = 600
	// 一個人必須立即採取行動。
	LevelAlert Level = 700
	// 一個或多個系統無法使用。
	LevelEmergency Level = 800
)

func LevelResolve(l Level) (s string) {
	switch l {
	case LevelDefault:
		s = "DEFAULT"
	case LevelDebug:
		s = "DEBUG"
	case LevelInfo:
		s = "INFO"
	case LevelNotice:
		s = "NOTICE"
	case LevelWarn:
		s = "WARN"
	case LevelError:
		s = "ERROR"
	case LevelCritical:
		s = "CRITICAL"
	case LevelAlert:
		s = "ALERT"
	case LevelEmergency:
		s = "EMERENCY"
	}
	return
}
