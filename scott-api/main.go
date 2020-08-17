package main

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/syhlion/greq"
	"github.com/syhlion/requestwork.v2"
)

var connection *greq.Client

func init() {
	threads := 50
	worker := requestwork.New(threads)
	connection = greq.New(worker, time.Duration(2)*time.Second, false)
}

func main() {
	// 打API
	GetUserData(123421, "normal")
}

// GetUserData 取得使用者資料
func GetUserData(userID int64, action string) ([]byte, bool, error) {
	// header 帶入遊戲商ID1
	connection.SetHeader("play", "1")

	url := "http://localhost/api/v1/user"

	var funPara = struct {
		UserID int64  `json:"user_id"`
		Type   string `json:"type"`
	}{
		UserID: userID,
		Type:   action,
	}

	// application/json
	marshalPara, marshalErr := json.Marshal(funPara)
	if marshalErr != nil {
		log.Println("marshalErr err:", marshalErr)
	}

	data, code, err := connection.PutRaw(url, bytes.NewBuffer(marshalPara))
	if err != nil {
		log.Println("api err:", err)
	}

	log.Println("code:", code, " data:", string(data))
}
