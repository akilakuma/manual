package main

import (
	"crypto/sha1"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	V2sum    int
	V2rwlock *sync.RWMutex
	V2wg     sync.WaitGroup
)

func v2() {
	t1 := time.Now()
	V2rwlock = new(sync.RWMutex)

	SplitAndAdd(100000000, 1000)
	V2wg.Wait()
	log.Println("花費時間", time.Since(t1), "answer:", V2sum)

}

func SplitAndAdd(total, splitNum int) {

	c := total / splitNum

	log.Println("c=", c)

	for i := 0; i < c; i++ {
		start := i * splitNum
		end := (i + 1) * splitNum
		V2wg.Add(1)
		go V2SubADD(start, end)
	}

	// 餘數處理
	d := total % splitNum
	if d > 0 {
		start := splitNum * c
		end := total
		V2wg.Add(1)
		go V2SubADD(start, end)
	}
}

func V2SubADD(start, end int) {

	var sum int
	for i := start; i < end; i++ {

		num := strconv.Itoa(i)

		shalResult := sha1.Sum([]byte(num))

		for j := 0; j < 20; j++ {
			sum += int(shalResult[j])
		}
	}

	// V2rwlock.Lock()
	V2sum += sum
	// V2rwlock.Unlock()
	V2wg.Done()
}
