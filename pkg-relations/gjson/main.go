package main

import (
	"log"

	"github.com/tidwall/gjson"
)

func main() {

	// 取得array 裡面的object

	respBody := `{"result":"ok","ret":{"entry":[{"id":509838}]}}`

	temp := gjson.Get(respBody, "ret.entry").Array()

	if len(temp) > 0 {

		log.Println(temp)

		data := temp[0].Map()

		log.Println(data["id"].Int())

	}
}
