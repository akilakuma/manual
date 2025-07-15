package main

import (
	"encoding/json"
	"log"
)

type UncomputeFisherWagers struct {
	Status    string `json:"status"`
	ErrorCode string `json:"errorCode"`
	Data      struct {
		Amount string `json:"amount"`
	} `json:"data"`
	Version string `json:"version"`
}

func main() {
	testAPIfmt()

}

func testAPIfmt() {


	var apiData = []byte(`{"status":"000","errorCode":"00","data":{"amount":"126427"},"version":"1.69.3"}`)
	//  var apiData2 = []byte(`{"status":"000","errorCode":"00","data":[],"version":"1.69.3"}`)
	// var apiData3 = []byte(`{"status":"000","errorCode":"00","data":{"amount":"0"},"version":"1.69.3"}`)

	var response UncomputeFisherWagers

	// json.Unmarshal(apiData, &response)
	json.Unmarshal(apiData, &response)

	// fmt.Println("response")
	// fmt.Printf("%s\n", string(apiData))

	log.Println(response.Data)

	if response.Data.Amount == "" || response.Data.Amount == "0" {
		log.Println("is nil or zero")
	} else {
		log.Println("is has")
	}
}
