package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	test1()
	test2()

}

func test1() {
	for i := 0; i < 10000; i++ {
		// Step 1. 組裝要發的內容
		req, err := http.NewRequest("GET", "http://172.17.117.23/api/user/562876273/gaming/balance?platform_id=66", nil)
		if err != nil {
			log.Println(err)

		}

		// 視情況選擇設Host
		req.Host = "prod.rd6"

		// 宣告http Client
		// Maximum of 5 secs API 連線逾時
		spaceClient := http.Client{
			Timeout: time.Second * 60,
		}

		// Step 2. 實際執行
		res, getErr := spaceClient.Do(req)
		if getErr != nil {
			log.Println(getErr)

			continue
		}

		// Step 3. 接收response
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Println(readErr)
		}
		// 若沒有錯誤則關閉連線，避免memory leak
		defer res.Body.Close()

		log.Println(string(body))
	}

}

func test2() {
	for i := 0; i < 10000; i++ {
		res, err := http.Get("http://172.17.117.23/api/user/562876273/gaming/balance?platform_id=66")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		sitemap, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", string(sitemap))
		log.Println(" ")
		time.Sleep(1)
	}
}
