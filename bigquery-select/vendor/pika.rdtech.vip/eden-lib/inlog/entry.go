package inlog

import (
	"github.com/sirupsen/logrus"
	"pika.rdtech.vip/eden-lib/inerror"
)

type Entry struct {
	Entry *logrus.Entry
	level Level
}

func (l *Entry) WithFields(fields map[string]interface{}) (e *Entry) {

	//	clean 保留字
	delete(fields, "code")
	delete(fields, "err_msg")
	delete(fields, "extrainfo")
	delete(fields, "time")
	delete(fields, "service")
	delete(fields, "origing_err")
	delete(fields, "error")
	return &Entry{
		Entry: l.Entry.WithFields(logrus.Fields(ConvertInt64ToString(fields))),
	}

}

func (l *Entry) WithError(err error) (e *Entry) {
	switch v := err.(type) {
	case inerror.Error:
		e = &Entry{
			Entry: l.Entry.WithFields(logrus.Fields{
				"code":       v.Code,
				"err_msg":    v.Msg,
				"extrainfo":  ConvertInt64ToString(v.ExtraInfo),
				"time":       v.Time,
				"service":    v.Service,
				"origin_err": v.OriginErr,
			}),
			level: l.level,
		}
	default:
		e = &Entry{
			Entry: l.Entry.WithError(err),
			level: l.level,
		}
	}
	return
}

func (l *Entry) SetLogLevel(lv Level) (e *Entry) {
	l.level = lv
	return l
}

func (l *Entry) Println(args ...interface{}) {
	if l.level <= LevelDefault {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelDefault,
		}).Println(args...)
	}
}

func (l *Entry) Printf(format string, args ...interface{}) {
	if l.level <= LevelDefault {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelDefault,
		}).Printf(format, args...)
	}
}

func (l *Entry) Info(args ...interface{}) {
	if l.level <= LevelInfo {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelInfo,
		}).Println(args...)
	}
}
func (l *Entry) Infof(format string, args ...interface{}) {
	if l.level <= LevelInfo {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelInfo,
		}).Printf(format, args...)
	}
}
func (l *Entry) Debug(args ...interface{}) {
	if l.level <= LevelDebug {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelDebug,
		}).Println(args...)
	}
}
func (l *Entry) Debugf(format string, args ...interface{}) {
	if l.level <= LevelDebug {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelDebug,
		}).Printf(format, args...)
	}
}
func (l *Entry) Notice(args ...interface{}) {
	if l.level <= LevelNotice {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelNotice,
		}).Println(args...)
	}
}
func (l *Entry) Noticef(format string, args ...interface{}) {
	if l.level <= LevelNotice {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelNotice,
		}).Printf(format, args...)
	}
}
func (l *Entry) Warn(args ...interface{}) {
	if l.level <= LevelWarn {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelWarn,
		}).Println(args...)
	}
}
func (l *Entry) Warnf(format string, args ...interface{}) {
	if l.level <= LevelWarn {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelWarn,
		}).Printf(format, args...)
	}
}
func (l *Entry) Error(args ...interface{}) {
	if l.level <= LevelError {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelError,
		}).Println(args...)
	}
}
func (l *Entry) Errorf(format string, args ...interface{}) {
	if l.level <= LevelError {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelError,
		}).Printf(format, args...)
	}
}
func (l *Entry) Critical(args ...interface{}) {
	if l.level <= LevelCritical {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelCritical,
		}).Println(args...)
	}
}
func (l *Entry) Criticalf(format string, args ...interface{}) {
	if l.level <= LevelCritical {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelCritical,
		}).Printf(format, args...)
	}
}
func (l *Entry) Alert(args ...interface{}) {
	if l.level <= LevelAlert {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelAlert,
		}).Println(args...)
	}
}
func (l *Entry) Alertf(format string, args ...interface{}) {
	if l.level <= LevelAlert {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelAlert,
		}).Printf(format, args...)
	}
}
func (l *Entry) Emergency(args ...interface{}) {
	if l.level <= LevelEmergency {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelEmergency,
		}).Println(args...)
	}
}
func (l *Entry) Emergencyf(format string, args ...interface{}) {
	if l.level <= LevelEmergency {
		l.Entry.WithFields(logrus.Fields{
			"severity": LevelEmergency,
		}).Printf(format, args...)
	}
}
