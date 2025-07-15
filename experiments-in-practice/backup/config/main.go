package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
)

func main() {
	loadConfig()

}

// EnvConfig 環境變數
var EnvConfig ProjectConfig

// ProjectConfig 專案Config
type ProjectConfig struct {
	Control struct {
		GinPort         int `json:"gin_port,omitempty"`
		GinReadTimeout  int `json:"gin_read_timeout,omitempty"`
		GinWriteTimeout int `json:"gin_write_timeout,omitempty"`
	} `json:"control,omitempty"`
	Center struct {
		GinPort         int      `json:"gin_port,omitempty"`
		GinReadTimeout  int      `json:"gin_read_timeout,omitempty"`
		GinWriteTimeout int      `json:"gin_write_timeout,omitempty"`
		RatioList       []string `json:"ratio_list,omitempty"`
		RoomMaxNumber   int      `json:"room_max_number,omitempty"`
		DefaultWeapon   int      `json:"default_weapon,omitempty"`
		MaxMember       int      `json:"max_member,omitempty"`
	} `json:"center,omitempty"`
	Game struct {
		MaxSfNumber  int `json:"max_sf_number,omitempty"`
		MaxBfNumber  int `json:"max_bf_number,omitempty"`
		SfPeriodTime int `json:"sf_period_time,omitempty"`
		BfPeriodTime int `json:"bf_period_time,omitempty"`
	} `json:"game,omitempty"`
}

func loadConfig() {
	getDefaultConfig()
	fmt.Printf("%+v", EnvConfig)
	fmt.Println()
	getEnvConfig()
	fmt.Printf("%+v", EnvConfig)

}

// getDefaultConfig 取得預設的設定檔
func getDefaultConfig() {
	path := "/Users/shen_su/go/src/golang-learning/config/default.json"
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err.Error() + ",path:" + path)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var projectConfig ProjectConfig
	gjson.Unmarshal(byteValue, &projectConfig)

	EnvConfig = projectConfig
}

// getEnvConfig 取得環境的設定檔
func getEnvConfig() {

	path := "/Users/shen_su/go/src/golang-learning/config/env.json"
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err.Error() + ",path:" + path)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// ====================================== control ===========================================
	// control: gin server設定
	if exists := gjson.Get(string(byteValue), "control.gin_port").Exists(); exists {
		EnvConfig.Control.GinPort = int(gjson.Get(string(byteValue), "control.gin_port").Int())
	}
	// control: gin server設定
	if exists := gjson.Get(string(byteValue), "control.gin_read_timeout").Exists(); exists {
		EnvConfig.Control.GinReadTimeout = int(gjson.Get(string(byteValue), "control.gin_read_timeout").Int())
	}
	// control: gin server設定
	if exists := gjson.Get(string(byteValue), "control.gin_write_timeout").Exists(); exists {
		EnvConfig.Control.GinWriteTimeout = int(gjson.Get(string(byteValue), "control.gin_write_timeout").Int())
	}
	// ====================================== center ===========================================
	// center: gin server設定
	if exists := gjson.Get(string(byteValue), "center.gin_port").Exists(); exists {
		EnvConfig.Center.GinPort = int(gjson.Get(string(byteValue), "center.gin_port").Int())
	}
	// center: gin server設定
	if exists := gjson.Get(string(byteValue), "center.gin_read_timeout").Exists(); exists {
		EnvConfig.Center.GinReadTimeout = int(gjson.Get(string(byteValue), "center.gin_read_timeout").Int())
	}
	// center: gin server設定
	if exists := gjson.Get(string(byteValue), "center.gin_write_timeout").Exists(); exists {
		EnvConfig.Center.GinWriteTimeout = int(gjson.Get(string(byteValue), "center.gin_write_timeout").Int())
	}
	// center: 開分級距表
	if exists := gjson.Get(string(byteValue), "center.ratio_list").Exists(); exists {
		resultArr := gjson.Get(string(byteValue), "center.ratio_list").Array()
		var ratioSlice []string
		for _, v := range resultArr {
			ratioSlice = append(ratioSlice, v.String())
		}
		EnvConfig.Center.RatioList = ratioSlice
	}
	// center: server上最大房間總數
	if exists := gjson.Get(string(byteValue), "center.room_max_number").Exists(); exists {
		EnvConfig.Center.RoomMaxNumber = int(gjson.Get(string(byteValue), "center.room_max_number").Int())
	}
	// center: 預設的武器等級
	if exists := gjson.Get(string(byteValue), "center.default_weapon").Exists(); exists {
		EnvConfig.Center.DefaultWeapon = int(gjson.Get(string(byteValue), "center.default_weapon").Int())
	}
	// center: 每個房間的容納最多人數
	if exists := gjson.Get(string(byteValue), "center.max_member").Exists(); exists {
		EnvConfig.Center.MaxMember = int(gjson.Get(string(byteValue), "center.max_member").Int())
	}
	// ====================================== game ===========================================
	// 	game: 小魚最大數量
	if exists := gjson.Get(string(byteValue), "game.max_sf_number").Exists(); exists {
		EnvConfig.Game.MaxSfNumber = int(gjson.Get(string(byteValue), "game.max_sf_number").Int())
	}
	// 	game: 大魚最大數量
	if exists := gjson.Get(string(byteValue), "game.max_bf_number").Exists(); exists {
		EnvConfig.Game.MaxBfNumber = int(gjson.Get(string(byteValue), "game.max_bf_number").Int())
	}
	// 	game: 小魚定時出魚週期時間
	if exists := gjson.Get(string(byteValue), "game.sf_period_time").Exists(); exists {
		EnvConfig.Game.SfPeriodTime = int(gjson.Get(string(byteValue), "game.sf_period_time").Int())
	}
	// game: 大魚定時出魚週期時間
	if exists := gjson.Get(string(byteValue), "game.bf_period_time").Exists(); exists {
		EnvConfig.Game.BfPeriodTime = int(gjson.Get(string(byteValue), "game.bf_period_time").Int())
	}

}
