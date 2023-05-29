package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

var romanIndexMap = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

// constraints
// 1 <= num <= 3999
func intToRoman(num int) string {

	numStr := strconv.Itoa(num)
	var (
		i        int
		romanStr string
	)
	for _, ru := range numStr {
		nInt, _ := strconv.Atoi(string(ru))
		romanStr += numFunction(nInt, len(numStr)-i)
		i++
	}
	return romanStr
}

func numFunction(num, digit int) string {
	fmt.Println(num, digit)
	var pattern string
	switch num {
	case 1:
		pattern = "A"
	case 2:
		pattern = "AA"
	case 3:
		pattern = "AAA"
	case 4:
		pattern = "AB"
	case 5:
		pattern = "B"
	case 6:
		pattern = "BA"
	case 7:
		pattern = "BAA"
	case 8:
		pattern = "BAAA"
	case 9:
		pattern = "AZ"
	}

	var np string
	switch digit {
	case 1:
		// 個位
		np = strings.Replace(pattern, "A", "I", -1)
		np = strings.Replace(np, "B", "V", -1)
		np = strings.Replace(np, "Z", "X", -1)
	case 2:
		// 十位
		np = strings.Replace(pattern, "A", "X", -1)
		np = strings.Replace(np, "B", "L", -1)
		np = strings.Replace(np, "Z", "C", -1)
	case 3:
		// 百位
		np = strings.Replace(pattern, "A", "C", -1)
		np = strings.Replace(np, "B", "D", -1)
		np = strings.Replace(np, "Z", "M", -1)
	case 4:
		// 千位
		np = strings.Replace(pattern, "A", "M", -1)
		np = strings.Replace(np, "B", "", -1)
		np = strings.Replace(np, "Z", "", -1)
	}
	return np
}
