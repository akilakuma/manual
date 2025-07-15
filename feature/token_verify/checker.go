package main

import (
	"strconv"
)

func Check(code, pattern string) bool {
	if len(code) != 12 || len(pattern) == 0 {
		return false
	}

	strFront := code[:4]
	strBack := code[4:]

	indexArr := make([]int, 4)
	for i := 0; i < 4; i++ {
		pair := strBack[i*2 : i*2+2]
		val, err := strconv.Atoi(pair)
		if err != nil || val >= len(pattern) {
			return false
		}
		indexArr[i] = val
	}

	for i := 0; i < 4; i++ {
		if pattern[indexArr[i]] != strFront[i] {
			return false
		}
	}
	return true
}
