package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	SetAPIRequest()
}

type APIDATA map[string]interface{}

func SetAPIRequest() {
	log.Println("hello main~")

	u, err := url.Parse("http://localhost:8501/rpc")
	if err != nil {
		log.Fatal(err)
	}
	// u.Scheme = "http"
	// u.Host = "google.com"
	q := u.Query()
	q.Set("service", "balance")
	q.Set("method", "Balance.GetBalance")

	api := &APIDATA{
		"user_id": 22002,
	}

	q.Set("params", api)
	u.RawQuery = q.Encode()
	fmt.Println(u)

	spaceClient1 := http.Client{
		Timeout: time.Second * 5,
	}

	apiUtil("POST", "http://localhost:8501/rpc", "", nil, &spaceClient1)
}

func apiUtil(method string, apiURL string, host string,
	parameter io.Reader, spaceClient *http.Client) {

	// Step 1. 組裝要發的內容
	req, err := http.NewRequest(method, apiURL, parameter)

	if err != nil {
		log.Println(err)
	}

	// 視情況選擇設Host
	req.Host = host

	// Step 2. 實際執行
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	// Step 3. 接收response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	// 若沒有錯誤則關閉連線，避免memory leak
	defer res.Body.Close()

	var response string

	json.Unmarshal(body, &response)

	fmt.Printf("%s\n", string(body))
}
