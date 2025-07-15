package main

import (
	"encoding/json"
	"log"
)

func main() {
	// test1()
	test2()
}

func test1() {
	// 來源資料
	var result = `{"max_fg":9,"opt":[{"double":1,"times":20},{"double":1,"times":15},{"double":1,"times":12},{"double":1,"times":9},{"double":1,"times":6}],"select_trigger_count":3,"select_trigger_way":"reel","switch":true}`

	// 先unmarshal
	var middleInterface interface{}
	json.Unmarshal([]byte(result), &middleInterface)
	log.Println(middleInterface)

	// 要塞進去的目標struct
	b := myTotalST{
		Name:  "little johnny",
		Other: middleInterface,
	}

	// marshal 在一起
	c, _ := json.Marshal(b)
	log.Println(string(c))

}

type myTotalST struct {
	Name  string      `json:"name"`
	Other interface{} `json:"other"`
}

func test2() {
	game := gameType{
		MaxFg: 9,
		Opt: []struct {
			Double int
			Times  int
		}{
			{
				Double: 1,
				Times:  1,
			},
			{
				Double: 2,
				Times:  2,
			},
		},
		SelectTriggerCount: 3,
		SelectTriggerWay:   "reel",
		Switch:             true,
	}

	// 要塞進去的目標struct
	b := myTotalST{
		Name:  "little johnny",
		Other: game,
	}

	// marshal 在一起
	c, _ := json.Marshal(b)
	log.Println(string(c))
}

type gameType struct {
	MaxFg int
	Opt   []struct {
		Double int
		Times  int
	}
	SelectTriggerCount int
	SelectTriggerWay   string
	Switch             bool
}
