package main

import (
	"crypto/sha1"
	"log"
	"strconv"
	"time"
)

func v1() {
	// 2020/06/02 16:54:44 加總數字: 255001985043 花費時間 23.777596654s

	var (
		sum    int
		target = 100000000
	)

	t1 := time.Now()
	for i := 0; i < target; i++ {

		num := strconv.Itoa(i)

		shalResult := sha1.Sum([]byte(num))

		for j := 0; j < 20; j++ {
			sum += int(shalResult[j])
		}
	}
	t2 := time.Since(t1)
	log.Println("加總數字:", sum, "花費時間", t2)
}
