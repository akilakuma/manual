package main

import (
	"crypto/sha1"
	"log"
	"strconv"
	"sync"
	"time"
)



func v3() {
var (
		sum    int
		target = 100000000
	)

	t1 := time.Now()

	// for i := 0; i < target; i++ {

	// 	num := strconv.Itoa(i)

	// 	shalResult := sha1.Sum([]byte(num))

	// 	for j := 0; j < 20; j++ {
	// 		sum += int(shalResult[j])
	// 	}
	// }

	var wg sync.WaitGroup
	core := 4
	total := make(chan int, core)
	for i := 0; i < core; i++ {
		wg.Add(1)
		go func(start int, end int, channel chan int) {
			temp := 0
			for i := start; i < end; i++ {

				num := strconv.Itoa(i)
				shalResult := sha1.Sum([]byte(num))

				for j := 0; j < 20; j++ {
					temp += int(shalResult[j])
				}
			}
			channel <- temp
			wg.Done()
		}(i*target/core, (i+1)*target/core, total)
	}
	wg.Wait()
	var length = len(total)
	for i := 0; i < length; i++ {
		sum += <-total
	}

	t2 := time.Since(t1)
	log.Println("加總數字:", sum, "花費時間", t2)

	// 2020/06/02 16:54:44 加總數字: 255001985043 花費時間 23.777596654s

}
