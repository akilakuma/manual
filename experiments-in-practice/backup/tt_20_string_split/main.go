package main

import (
	"fmt"
	"strings"
)

func main() {
	para := []string{"1:1", "100"}

	for _, v := range para {
		Cut(v)
	}
}

func Cut(para string) {
	a := strings.Split(para, ":")
	fmt.Println(a[0])
	fmt.Println(a[1])
}
