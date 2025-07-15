package unmarshaljt

import (
	"encoding/json"
	// "github.com/tidwall/gjson"
)

type SliceType struct {
	stringSlice []string `json:"ss"`
}

type SliceTypeJson struct {
	ss string `json:"ss"`
}

func SetStringSlice(para string) {

	var ss SliceTypeJson
	ss.ss = para

	// log.Println("ss", ss)

	a, _ := json.Marshal(ss)

	var uss SliceType
	json.Unmarshal(a, &uss)

	// log.Println(uss)
}
