package main

import (
	"log"
)

// case1 即使jobCannel沒有東西，它也會卡住！
func case1() {
	jobChannel := make(chan int, 100)

	for j := range jobChannel {
		log.Println("get j: ", j)
	}

	for i := 0; i < 10; i++ {
		jobChannel <- i
	}

	// 	fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan receive]:
}
