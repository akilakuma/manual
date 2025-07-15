package main

import (
	"log"
)

func case4() {
	workChannel := make(chan int, 1000)
	resultChannel := make(chan int, 20)

	for i := 0; i < 2; i++ {
		go case4Worker(i, workChannel, resultChannel)
	}

	for j := 0; j < 5; j++ {
		workChannel <- j
	}

	close(workChannel)

	for j := 0; j < 3; j++ {
		// 因為有讀取，所以會卡住gorutine，所以後面不用再寫一個無限迴圈卡著
		log.Println("getResult: ", <-resultChannel)
	}

	// 	2019/04/22 11:13:47 worker : 0 handle:  1
	// 2019/04/22 11:13:47 worker : 0 handle:  2
	// 2019/04/22 11:13:47 worker : 0 handle:  3
	// 2019/04/22 11:13:47 worker : 0 handle:  4
	// 2019/04/22 11:13:47 getResult:  1
	// 2019/04/22 11:13:47 getResult:  2
	// 2019/04/22 11:13:47 getResult:  3

}

// case4Worker 注意這裡的chan !!!
// jobs <-chan int ，是指傳入的channel是int類型，並且是這個method的範圍內是receive-only type，限制只能接收
// result chan<- int ，是指傳入的channel是int類型，這個method的範圍內不限操作
func case4Worker(wID int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		log.Println("worker :", wID, "handle: ", j)

		// 重點地方，如果不寫入，外面有讀取的話會卡死，寫入的數量一定要比讀取多
		result <- j
	}
}
