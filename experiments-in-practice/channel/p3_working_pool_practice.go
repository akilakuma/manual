package main

import (
	"log"
	"time"
)

/*
	『working pool』 :

	概念：
		搭配『range channel』的概念。
		一般的型態，range 撈完裡面的東西就結束了，接著脫離range的範圍，不會再進入這段code。
		但range channel不會，一旦channel有東西進來就會立刻處理，這段會一直stand by。
		直到channel被close。
*/

func practice3() {

	case1()
	case2()
	case3()
	case4()
}

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

	// 休息
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

	// 2019/04/22 11:13:47 worker : 0 handle:  1
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
