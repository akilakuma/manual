package main

import (
	"encoding/json"
	"log"
)

type FormatInfo struct {
	Member  int     `json:"member"`
	Bet     float64 `json:"bet"`
	Amount  float64 `json:"amount"`
	Sport   float64 `json:"sport"`
	Live    float64 `json:"live"`
	Casino  float64 `json:"CASIno"`
	Lottery float64 `json:"lottery"`
	Fish    float64 `json:"fish"`
	Card    float64 `json:"card"`
}

func main() {
	s := `{"member":1,"bet":40983.6,"amount":5594.26,"sport":0,"live":0,"casino":5594.26,"lottery":0,"fish":0,"card":0}`

	var formatInfo FormatInfo
	json.Unmarshal([]byte(s), &formatInfo)

	log.Println(formatInfo)

}
