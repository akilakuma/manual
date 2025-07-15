package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	temp := randPickByNum(9, 10)
	log.Println(temp)
}

func randPickByNum(pickNum, length int) []int {
	var (
		takeNum   int
		targetInt []int
		storeMap  = make(map[int]bool, 0)
	)
	// 如果需要的數目大於目標數量
	if pickNum >= length {
		for i := 0; i < length; i++ {
			targetInt = append(targetInt, i)
		}
	} else {
		for takeNum < pickNum {
			rand.Seed(time.Now().UnixNano())
			x := rand.Intn(length)

			if _, exists := storeMap[x]; !exists {
				takeNum++
				targetInt = append(targetInt, x)
				storeMap[x] = true
			}
		}
	}

	return targetInt
}
