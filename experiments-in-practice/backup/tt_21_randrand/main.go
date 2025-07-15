package main

import (
	"log"
	"math/rand"
	"time"
)

var TmpPool []chan int

func main() {
	go test1()
	time.Sleep(5 * time.Second)

	log.Println("index 1: ")
	RandNum(11)
	log.Println("index 2: ")
	RandNum(12)
	log.Println("index 3: ")
	RandNum(13)

}

func test1() {
	indexSize := 1000
	poolSize := 1000

	for i := 0; i < indexSize; i++ {
		TmpPool = append(TmpPool, make(chan int, poolSize))
	}

	for {
		for i := 10; i < indexSize; i++ {

			select {
			case TmpPool[i] <- randNum(i):
				// log.Println("caseA")
			default:
				log.Println("default")
				time.Sleep(100 * time.Millisecond)
			}

			// log.Println(TmpPool)
		}

	}

}

func randNum(seed int) int {
	var rNum int
	limit := 100
	rand.Seed(int64(time.Now().Nanosecond()) + int64(seed))
	rNum = rand.Int()
	rNum = rNum % limit
	time.Sleep(1 * time.Millisecond)

	return rNum
}

func RandNum(limit int) {
	// 防止遊戲設定錯誤 導致除以0
	if limit <= 0 {
		limit = 1
	}

	if len(TmpPool) > limit {
		select {
		case nn := <-TmpPool[limit]:
			log.Println(nn)
		default:
			log.Println("not number")
		}
	}

	return
}
