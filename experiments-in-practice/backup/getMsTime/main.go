package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	a := time.Now()
	log.Println(a)
}

func test1() {
	start := time.Now()

	a := 0
	for i := 0; i < 10000000; i++ {
		a++
	}

	elapsed := time.Since(start)
	ms := float64(elapsed) / float64(time.Millisecond)
	fmt.Println(a)
	fmt.Println(ms)
}
