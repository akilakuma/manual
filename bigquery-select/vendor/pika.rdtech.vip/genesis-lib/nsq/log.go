package nsq

import (
	"regexp"
	"strings"

	"pika.rdtech.vip/eden-lib/inlog"
)

var parser *regexp.Regexp = regexp.MustCompile(`^(\S{3})(.*)`)

//InLogLevel InLogLevel
type InLogLevel int

// Log levels
const (
	LogLevelDebug InLogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

//Logger Logger
type Logger interface {
	Output(int, string) error
	OverWriteLogLevel()
}

//NewLogger NewLogger
func NewLogger(l *inlog.Logger) Logger {
	return &nsqLog{
		Log: l,
	}
}

type nsqLog struct {
	Log *inlog.Logger
	IsOverwriteLogLevel bool
}

// OverWriteLogLevel 強制把log輸入的level都設為Info
// 請注意，使用此method會導致config的SetLogLevel 失效
func (n *nsqLog) OverWriteLogLevel() {
	n.IsOverwriteLogLevel = true
}

func (n *nsqLog) Output(i int, s string) (err error) {

	matchs := parser.FindStringSubmatch(s)
	var log, ss string
	if len(matchs) != 3 {
		ss = "ERR"
		log = s
	} else {
		ss = matchs[1]
		log = strings.TrimSpace(matchs[2])
	}

	if n.IsOverwriteLogLevel {
		ss = "INF"
	}

	switch ss {
	case "INF":
		n.Log.Info(log)
	case "WRN":
		n.Log.Warn(log)
	case "ERR":
		n.Log.Error(log)
	default:
		n.Log.Debug(log)
	}
	return
}
