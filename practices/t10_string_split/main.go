package main

import (
	"fmt"
	"strconv"
)

const (
	pattern = "67d8309e4026e15a7a9fa79a0a33ed01e490ff15bea67e8dd474e21e93c5d7cd"
	code    = "ddea02631215"
	code2   = "ddea02631315"
)

func main() {
	// pattern := "67d8309e4026e15a7a9fa79a0a33ed01e490ff15bea67e8dd474e21e93c5d7cd"
	// code := "ddea02631215"

	isVgaPlayer := CheckVgaPlayer1(code, pattern)
	fmt.Println("isVgaPlayer = ", isVgaPlayer)
}

func CheckVgaPlayer1(code string, pattern string) bool {
	var (
		// code字串前段4碼
		strFront string = code[:4]
		// code字串後段8碼
		strBack string = code[4:]
		// 切割後段8碼字串暫時儲存用變數
		str string = ""
		// 新的int陣列
		indexArr [4]int = [4]int{}
		// 新的int陣列中的index
		index int = 0
		// 是否為一般玩家(預設為true)
		isVgaPlayer bool = true
	)

	for i, v := range strBack {
		val := string(v)
		// 切割出的內容，非偶數位則跳過
		if val == "0" && (i+1)%2 != 0 {
			continue
		}
		// 拼接字串
		str += val
		if (i+1)%2 == 0 {
			index = (i+1)/2 - 1
			indexArr[index], _ = strconv.Atoi(str)
			str = ""
		}
	}

	for i, v := range indexArr {
		// 檢查是否與sid符合
		if pattern[v] != strFront[i] {
			isVgaPlayer = false
			break
		}
	}

	return isVgaPlayer
}

func CheckVgaPlayer2(code, pattern string) bool {

	// 如果不是12個英數字
	if len(code) > 12 {
		return false
	}

	var codeStrSlice [12]string
	for i, letter := range code {
		codeStrSlice[i] = string(letter)
	}

	// 0 --> 4,5 --> 4+(0*2) ,  4+(0*2)+1
	// 1 --> 6,7 --> 4+(1*2) ,  4+(1*2)+1
	// 2 --> 8,9 --> 4+(2*2) ,  4+(2*2)+1
	// 3 --> 10,11 --> 4+(3*2) ,  4+(3*2)+1

	for i := 0; i < 4; i++ {
		key1 := 4 + (i * 2)
		key2 := 4 + (i * 2) + 1

		targetStr := codeStrSlice[key1] + codeStrSlice[key2]
		targetInt, conErr := strconv.Atoi(targetStr)
		if conErr != nil {
			return false
		}

		if codeStrSlice[i] != string(pattern[targetInt]) {
			return false
		}

	}
	return true
}
