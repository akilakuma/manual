package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func main() {

	loc, _ := time.LoadLocation("EST")

	log.Println(time.Now().In(loc))
	switch os.Args[1] {
	case "cal":
		log.Println("calculate")
	case "deposit":
		log.Println("deposit")
	}

	// test1()
	// test2()
}
func test1() {

	// SubInCasinoResult Casino 電子遊戲
	type SubInCasinoResult struct {
		GroupID        int     `json:"GroupID"`
		WagersTotal    int     `json:"wagersTotal"`
		Commissionable float64 `json:"commissionable"`
		Payoff         float64 `json:"payoff"`
	}
	// CasinoResultMG 電子遊戲MG 特例格式
	type CasinoResultMG struct {
		ErrorCode     int                            `json:"error_code"`
		ErrorMessage  interface{}                    `json:"error_message"`
		ExecutionTime string                         `json:"execution_time"`
		ServerName    string                         `json:"server_name"`
		Data          map[string][]SubInCasinoResult `json:"data"`
	}

	data := `{"error_code":0,"error_message":null,"execution_time":"109 ms","server_name":"backend-56d457dcd-wwg6m","data":{"344259058":[{"GroupID":0,"wagersTotal":0,"commissionable":0,"payoff":0},{"GroupID":1,"wagersTotal":0,"commissionable":0,"payoff":0},{"GroupID":2,"wagersTotal":178,"commissionable":324.8,"payoff":99.87},{"GroupID":3,"wagersTotal":0,"commissionable":0,"payoff":0},{"GroupID":4,"wagersTotal":0,"commissionable":0,"payoff":0}]}}`

	var casinoResultMG CasinoResultMG

	json.Unmarshal([]byte(data), &casinoResultMG)

	log.Println(casinoResultMG)

	for userID, userBet := range casinoResultMG.Data {

		// userIDInt, convErr := strconv.Atoi(userID)

		// 第二層 map 細項電子投注資料
		for _, casinoDetail := range userBet {
			if casinoDetail.Commissionable > 0 {

				log.Println(userID)
				log.Println(casinoDetail.Commissionable)
			}
		}
	}
}

// func test2() {
// 	var realtimeRakeBackEntry []RealtimeRakeBackEntry
// 	conn := getDBConnect("almond_salve")

// 	dateBegin := "2018-11-13" + " 00:00:00"
// 	dateEnd := "2018-11-14" + " 23:59:59"

// 	// conn.LogMode(true)
// 	conn.Table("realtime_rakeback_entry").Select("user_id, value").Where(" user_id = ? AND bet_start >= ? AND bet_end <= ? AND status = ?", dateBegin, dateEnd, "RF").Find(&realtimeRakeBackEntry)

// 	// return realtimeRakeBackEntry

// 	userSelfWithdrawRecord := make(ST.UserSelfWithdrawRecord)

// 	// 從DB 撈取這個人的紀錄, 如果自領三次就會有三筆
// 	records := getRealtimeRakeBackEntry(userID, day)
// 	for _, records := range records {
// 		// 東西都放在value欄位，要json decode 字串
// 		var valueInRealtimeRakeBackEntry ValueInRealtimeRakeBackEntry
// 		json.Unmarshal([]byte(records.Value), &valueInRealtimeRakeBackEntry)

// 		// 處理每ㄧ筆value的內容
// 		for index, item := range valueInRealtimeRakeBackEntry {

// 			if index == 0 {
// 				if item.Commit.Put != 0 && item.Commit.RefID == 0 {
// 					log.Println(records.UserID)
// 				}
// 			}
// 		}

// 	}
// 	return userSelfWithdrawRecord
// }
