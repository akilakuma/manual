package log

import (
	TOOL "golang-learning/tt_11_init/tool"
	"log"
)

// SetLog 寫log
func SetLog() {
	d := TOOL.GetDayToday()
	log.Println("寫入log")
	log.Println(d)

	// 寫log
}
