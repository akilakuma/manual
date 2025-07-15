package main

import (
	"log"
	"time"
)

// case2 先丟一輪參數進channel，休息一段時間後，再丟下一批
// 同樣在接受到參數後也會運行該method
func case2() {
	jobChannel := make(chan int, 100)

	go case2Worker(jobChannel)

	// 第1次丟參數
	for i := 0; i < 2; i++ {
		jobChannel <- i
	}
	// 休息
	time.Sleep(2 * time.Second)

	// 第2次丟參數
	for i := 4; i < 6; i++ {
		jobChannel <- i
	}
	// 關閉channel
	close(jobChannel)

	// 休息
	time.Sleep(2 * time.Second)
	// 第3次丟參數
	for i := 4; i < 6; i++ {
		jobChannel <- i
	}

	for {
	}

	// 2019/04/22 10:23:48 get j:  0
	// 2019/04/22 10:23:48 get j:  1
	// 2019/04/22 10:23:50 get j:  4
	// 2019/04/22 10:23:50 get j:  5
	// panic: send on closed channel

	// goroutine 1 [running]:
	// main.case2()
	//         /Users/shen_su/go/src/golang-advance-practice/t1_worker_pool/main.go:68 +0x118
	// main.main()
	//         /Users/shen_su/go/src/golang-advance-practice/t1_worker_pool/main.go:20 +0x20
	// exit status 2
}

// case2Worker 接受channel的method，用goruntine執行就不會卡主執行緒
func case2Worker(jobs <-chan int) {
	for j := range jobs {
		log.Println("get j: ", j)
	}
}
