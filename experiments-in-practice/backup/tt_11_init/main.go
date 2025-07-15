package main

import (
	myLog "golang-learning/tt_11_init/log"
	TOOL "golang-learning/tt_11_init/tool"
	"log"
)

func main() {
	// log.Println("不看日期了")

	d := TOOL.GetDayToday()
	log.Println("今天是" + d)
	myLog.SetLog()
}
