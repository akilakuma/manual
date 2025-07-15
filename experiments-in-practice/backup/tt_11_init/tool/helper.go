package tool

import (
	//  myLog "golang-learning/tt_11_init/log"
	"log"
	"time"
)

func init() {
	log.Println("you are in tool's init")
}

func init() {
	log.Println("you are in tool's init2")
}

// GetDayToday 取得今日時間
func GetDayToday() string {
	// myLog.SetLog()
	return time.Now().UTC().Add(8 * time.Hour).Format("2006-01-02")
}

// getDayYesterday 取得昨日時間
func getDayYesterday() string {
	return time.Now().UTC().Add(-16 * time.Hour).Format("2006-01-02")
}
