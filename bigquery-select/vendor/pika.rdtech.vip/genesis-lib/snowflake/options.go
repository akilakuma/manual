package snowflake

import "time"

//Option Client setting
type Option interface {
	apply(*Setting)
}

type funcOption struct {
	f func(*Setting)
}

func (fo *funcOption) apply(c *Setting) {
	fo.f(c)
}

func newFuncOption(f func(*Setting)) *funcOption {
	return &funcOption{
		f: f,
	}
}

//SetStartTime 設定起始時間
func SetStartTime(t time.Time) Option {
	return newFuncOption(func(s *Setting) {
		s.sonySet.StartTime = t
	})
}

//SetMachineID 設定機器辨識碼
func SetMachineID(callback func() (uint16, error)) Option {
	return newFuncOption(func(s *Setting) {
		s.sonySet.MachineID = callback
	})
}
