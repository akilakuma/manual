package main

import (
	"log"
	"time"
)

func example2() {

	limit := rateLimiter(5 * time.Second)

	for i := 0; i < 10; i++ {
		if limit() {
			// 執行某些操作
			log.Println("ok")
		} else {
			log.Println("not yet")
		}
		time.Sleep(1 * time.Second)
	}
}

func rateLimiter(limit time.Duration) func() bool {
	lastTime := time.Now().Add(-limit)
	return func() bool {
		if time.Now().Sub(lastTime) >= limit {
			lastTime = time.Now()
			return true
		}
		return false
	}
}
