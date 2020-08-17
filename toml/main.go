package main
// load 讀取環境變數
// func load() (envConfig *ProjectConfig) {
// 	// 讀取預設的設定檔
// 	defaultPath := "config.toml"
// 	defaultToml, readErr := ioutil.ReadFile(defaultPath)
// 	if readErr != nil {
// 		log.Println("🐼🐼read file error:", readErr)
// 		return
// 	}

// 	unmarshalErr := toml.Unmarshal(defaultToml, &envConfig)
// 	if unmarshalErr != nil {
// 		log.Println("🐼🐼unmarshal error:", unmarshalErr)
// 		return
// 	}
// }