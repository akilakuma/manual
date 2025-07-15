package main

import "log"

// case3 如果goruntine開出去的method(如下面的case3Worker)，會將收到的內容一個一個依序處理。
// 該method做的就是個queue的概念
// 那麼開多個goruntine，就是產生多個queue在處理事情，開20個goruntine，就是產生20個queue處理。
// j 不確定會進到哪個queue裡面，但會依序進入到queue，queue拉出來也會照著接收到的順序
func case3() {
	workChannel := make(chan int, 1000)

	for i := 0; i < 4; i++ {
		go case3Worker(i, workChannel)
	}

	for j := 0; j < 100; j++ {
		workChannel <- j
	}

	for {
	}

	// 2019/04/22 10:44:40 worker : 1 handle:  0
	// 2019/04/22 10:44:40 worker : 1 handle:  4
	// 2019/04/22 10:44:40 worker : 2 handle:  1
	// 2019/04/22 10:44:40 worker : 2 handle:  6
	// 2019/04/22 10:44:40 worker : 2 handle:  7
	// 2019/04/22 10:44:40 worker : 0 handle:  2
	// 2019/04/22 10:44:40 worker : 0 handle:  9
	// 2019/04/22 10:44:40 worker : 0 handle:  10
	// 2019/04/22 10:44:40 worker : 0 handle:  11
	// 2019/04/22 10:44:40 worker : 0 handle:  12
	// 2019/04/22 10:44:40 worker : 0 handle:  13
	// 2019/04/22 10:44:40 worker : 0 handle:  14
	// 2019/04/22 10:44:40 worker : 0 handle:  15
	// 2019/04/22 10:44:40 worker : 0 handle:  16
	// 2019/04/22 10:44:40 worker : 3 handle:  3
	// 2019/04/22 10:44:40 worker : 3 handle:  18
	// 2019/04/22 10:44:40 worker : 3 handle:  19
	// 2019/04/22 10:44:40 worker : 3 handle:  20
	// 2019/04/22 10:44:40 worker : 3 handle:  21
	// 2019/04/22 10:44:40 worker : 2 handle:  8
	// 2019/04/22 10:44:40 worker : 3 handle:  22

}

func case3Worker(wID int, jobs <-chan int) {
	for j := range jobs {
		log.Println("worker :", wID, "handle: ", j)
	}
}
