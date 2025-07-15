package main

import (
	"log"

	"github.com/tidwall/gjson"
)

func main() {

	// s := `{"error_code":"0","error_text":"","event":"game_close","data":[{"game_id":300001}]}`
	// s := `{"error_code":"0","error_text":"","event":"game_close_by_parent","data":[{"user_id":12,"user_role":4,"game_id":[300001]}]}`
	// s := `{"error_code":"0","error_text":"","event":"maintain","data":"2020-01-07T16:21:31+08:00"}`
	s := `{"error_code":"0","error_text":"","event":"session","data":[{"user_id":129,"session":"d2f43f4a0118cdb6b761b131833a402c906a46cf89a2382e96f8bac0926a9406","logout":2},{"user_id":130,"session":"d2f43f4a0118cdb6b761b131833a402c906a46cf89a2382e96f8bac0926a9406","logout":2}]}`

	event := gjson.Get(s, "event").String()

	switch event {
	case "game_close":
		info := gjson.Get(s, "data").Array()
		for _, v := range info {
			subDataMap := v.Map()
			if gameID, exists := subDataMap["game_id"]; exists {
				if int(gameID.Int()) == 300001 {
					// 關閉整個捕魚大聖
					log.Println("OK1")
				}
			}
		}
	case "game_close_by_parent":
		//  role 2:品牌,3:營運商,4:合作商

		info := gjson.Get(s, "data").Array()
		for _, v := range info {
			subDataMap := v.Map()
			if gameIDSlice, exists := subDataMap["game_id"]; exists {
				for _, vv := range gameIDSlice.Array() {
					if int(vv.Int()) == 300001 {
						// 關閉整個捕魚大聖
						log.Println("OK2")
						managerID := int(subDataMap["user_id"].Int())
						managerRole := int(subDataMap["user_role"].Int())
						log.Println("managerID", managerID, "managerRole", managerRole)
					}
				}
			}
		}
	case "maintain":
		// 維護
		openTime := gjson.Get(s, "data").String()
		log.Println("OK3", "開放時間", openTime)
	case "session":
		info := gjson.Get(s, "data").Array()
		var userIDSlice []int
		for _, v := range info {
			subDataMap := v.Map()
			if userID, exists := subDataMap["user_id"]; exists {
				uID := int(userID.Int())
				userIDSlice = append(userIDSlice, uID)
			}
		}
		log.Println("踢出名單:", userIDSlice)
	}
}
