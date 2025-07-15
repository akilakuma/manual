package main

import (
	"fmt"
)

func example() {
	code := "ddea02631215"
	pattern := "67d8309e4026e15a7a9fa79a0a33ed01e490ff15bea67e8dd474e21e93c5d7cd"

	if Check(code, pattern) {
		fmt.Println("✅ Code is valid!")
	} else {
		fmt.Println("❌ Code is invalid.")
	}
}
