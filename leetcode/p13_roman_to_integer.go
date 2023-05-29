package main

import (
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

var romanPatternMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}
var romanPatternExceptMap = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

// constraints
// 1 <= num <= 3999
func romanToInt(num string) int {

	var totalValue int
	for k, v := range romanPatternExceptMap {
		if strings.Contains(num, k) {
			totalValue += v
			num = strings.Replace(num, k, "", -1)
		}
	}
	for _, ru := range num {
		totalValue += romanPatternMap[string(ru)]
	}

	return totalValue
}
