package main

import (
	"golang-learning/tt_17_gjson/unmarshaljt"
	"log"

	"github.com/tidwall/gjson"
)

const setting = `{"game_body":{"Ratio":{"CaseMap":{},"FreeGameRatio":3,"Max":-1},"Round":{"FreeGameTimes":2,"Max":-1},"Transsymbol":{"OriginResult":null,"TransResult":{}}},"game_set":{"Category":"3x5","SpinRule":"spin_3x5","CalculateScoreRule":"score_3x5_line_left_and_right","CalculateSymbolRule":"symbol_3x5_divergence","LineType":50,"ID":2,"MainGameCode":["1,1,1,2,10,2,3,4,2,1,3,9,9,10","1,1,1,2,10,2,3,4,2,1,3,9,9,10","1,1,1,2,10,2,3,4,2,1,3,9,9,10","1,1,1,2,10,2,3,4,2,1,3,9,9,10","1,1,1,2,10,2,3,4,2,1,3,9,9,10"],"FreeGameCode":["1,1,4,5,6","2,2,4,3,2","1,2,1,3,4","2,2,4,2,5","5,2,4,3,5"],"WildID":9,"FGScatterID":10,"FGNumsThreshold":3,"XAxis":5,"YAxis":3,"SymbolScore":{"6":{"3":30,"4":40,"5":50},"7":{"5":50,"3":30,"4":40},"8":{"4":40,"5":50,"3":30},"1":{"3":30,"4":40,"5":50},"2":{"3":30,"4":40,"5":50},"3":{"3":30,"4":40,"5":50},"4":{"5":50,"3":30,"4":40},"5":{"3":30,"4":40,"5":50}}},"game_mem":{"Credit":10100}}`

func main() {

	gjson.Get(setting, "game_set").Get("CalculateSymbolRule").String()
	// log.Println(CalculateSymbolRule)

	FreeGameCode := gjson.Get(setting, "game_set").Get("FreeGameCode").String()

	unmarshaljt.SetStringSlice(FreeGameCode)

	// GJson.Get(setting, "CalculateSymbolRule").String(),
	// FreeGameCode:        GJson.Get(setting, "FreeGameCode").Value().([]string),
	// SymbolScore:         GJson.Get(setting, "SymbolScore").Value().(map[string]map[string]int),

	// println(value.String())

	a := gjson.Get(setting, "game_set").Get("FreeGameCode").Array()
	// log.Println("num", a.Num)
	// log.Println("Raw", a.Raw)
	// log.Println("Str", a.Str)
	// log.Println("type", a.Type)

	for _, v := range a {
		u := v.String()
		log.Println(u)
	}

	log.Println()

	b := gjson.Get(setting, "game_set").Get("SymbolScore").Map()
	for k, v := range b {
		u := v.Map()
		log.Println("k",k , "u",u)
	}
}
