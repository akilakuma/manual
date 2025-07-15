package main

import (
	"log"
	"strconv"
)

func main() {
	paraList := []string{"abc", "1", "99"}

	for _, v := range paraList {
		err := dosomething(v)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("everything is fine")
		}
	}
}

func dosomething(para string) (methodErr interface{}) {

	defer func() {
		if err := recover(); err != nil {
			methodErr = err
		}
	}()

	ConvertStrToInt(para, true, "dosomething")
	return nil
}

// ConvertStrToInt string轉成Int
// needPanic:true 緊要資料有問題觸發panic，needPanic:false 不要緊資料，有錯誤回傳預設值無妨
// funcName 印出log之類可以考慮使用
func ConvertStrToInt(s string, needPanic bool, funcName string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		if needPanic {
			panic(funcName + "string轉成Int 重要參數發生錯誤,傳入: " + s)
		}
	}
	return r
}
