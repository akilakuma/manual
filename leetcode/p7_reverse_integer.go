package main

import (
	"strconv"
	"strings"
)

// constraints
// -231 <= x <= 231 - 1
// 如果超過範圍，回傳0
func reverse(x int) int {

	var isNegative bool
	if x < 0 {
		isNegative = true
		x = 0 - x
	}

	xStr := strconv.Itoa(x)
	var tmpSlice []string
	for i := len(xStr) - 1; i >= 0; i-- {
		tmpSlice = append(tmpSlice, string(xStr[i]))
	}

	revStr := strings.Join(tmpSlice, "")

	r, convErr := strconv.ParseInt(revStr, 10, 32)
	if convErr != nil {
		return 0

	}
	if isNegative {
		return -int(r)
	}

	return int(r)
}
