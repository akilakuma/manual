package main

import "strings"

var (
	A int
)

func stringIndexWay() {

	if strings.Index("9432", "6") >= 0 {
		A++
	}
}
