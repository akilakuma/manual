package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

/*
	方便了解nsq customer 收到的訊息是什麼

	// 監聽topic
	// 掛channel
*/

func main() {
	sendNSQ()
	receiveNSQ()

	for {
		time.Sleep(10 * time.Second)
	}
}

func receiveNSQ() {

	// 建立空白設定檔。
	config := nsq.NewConfig()

	q1, _ := nsq.NewConsumer("EDEN_CATE_USER_TRANS_PRE", "receive", config)
	q1.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {

		fmt.Println("💎:", string(message.Body))

		// 訊息處理。
		//gameName := (gjson.Get(string(message.Body), "current.game_name")).String()
		//closeTime := (gjson.Get(string(message.Body), "current.closed_at")).Int()

		//fmt.Println(gameName, closeTime)

		return nil
	}))

	//err := q.ConnectToNSQLookupd("127.0.0.1:4160")
	// 單一一台nsqd
	err := q1.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		fmt.Println("ConnectToNSQD:", err)
	}

}

// CustomerUserCateTrans 收NSQ進來的資料型態
type CustomerUserCateTrans struct {
	PreGame string `json:"pre_game"`
	UserId  int64  `json:"user_id"`
}

func sendNSQ() {

	var cuc = CustomerUserCateTrans{
		PreGame: "LDDR-LDDR-1563429804",
		UserId:  1178704773388431360,
	}
	marshData, _ := json.Marshal(cuc)

	// 建立空白設定檔。
	config := nsq.NewConfig()

	q1, _ := nsq.NewProducer("127.0.0.1:4150", config)
	q1.Publish("EDEN_CATE_USER_TRANS_PRE", marshData)

}
