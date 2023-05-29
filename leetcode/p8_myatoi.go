package main

import (
	"strconv"
	"strings"
)

func myAtoi2(s string) int {
	var (
		numSlice      []string
		isNeg         bool
		whiteSpaceCnt int
		otherCnt      int
	)
	for _, ru := range s {
		if len(numSlice) == 0 && string(ru) == "-" {
			isNeg = true
		}
		switch string(ru) {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			numSlice = append(numSlice, string(ru))
		case " ":
			if whiteSpaceCnt == 0 {
				whiteSpaceCnt = 1
			}
			if len(numSlice) > 0 {
				break
			}
		default:
			otherCnt++
		}
	}
	nStr := strings.Join(numSlice, "")
	nInt, _ := strconv.Atoi(nStr)
	if isNeg {
		return -nInt
	}
	return nInt
}

func myAtoi(s string) int {

	var (
		tmpSlice      []string
		numSlice      []string
		isPos         bool
		isNeg         bool
		isException   bool
		hasOtherWords bool
	)

	r := strings.Split(s, " ")
	for _, v := range r {
		var (
			aWordInside   bool
			isBreakSwitch bool
		)
		if v == "+0" || v == "-0" {
			hasOtherWords = true
			continue
		}
		for index, ru := range v {
			switch string(ru) {
			case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
				//fmt.Println(string(ru))
				if !hasOtherWords {
					tmpSlice = append(tmpSlice, string(ru))
				} else {
					isBreakSwitch = true
				}
			case ".":
				//fmt.Println("point")
				isBreakSwitch = true
			case "+":
				isPos = true
			case "-":
				isNeg = true
				if index == 0 {
					isException = true
				}
			default:
				aWordInside = true
				isBreakSwitch = true
			}
			if isBreakSwitch {
				break
			}
		}
		if len(r) == 1 {
			numSlice = append(numSlice, tmpSlice...)
		} else if aWordInside {
			if isException {
				numSlice = append(numSlice, tmpSlice...)
				tmpSlice = []string{}
			}
			hasOtherWords = true
		} else {
			numSlice = append(numSlice, tmpSlice...)
			tmpSlice = []string{}

		}
	}

	nStr := strings.Join(numSlice, "")
	if nStr == "" {
		return 0
	}
	if isNeg && isPos {
		return 0
	}

	nInt, _ := strconv.Atoi(nStr)
	if isNeg && !isPos {
		nInt = -nInt
	}

	if nInt < -2147483648 {
		nInt = -2147483648
	} else if nInt > 2147483647 {
		nInt = 2147483647
	}
	return nInt
}
