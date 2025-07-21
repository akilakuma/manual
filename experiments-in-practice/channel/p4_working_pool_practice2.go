package main

import (
	"log"
	"sync"
	"time"
)

func practice4() {
	var (
		worker        int = 3
		jobChannel        = make(chan int, 10)
		resultChannel     = make(chan int, 10)
		wg            sync.WaitGroup
	)

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go workerMan(&wg, jobChannel, resultChannel)
	}

	go func() {
		wg.Wait()
		close(resultChannel) // 關閉結果channel，讓printResult能夠結束
	}()

	go printResult(resultChannel)

	for i := 0; i < 100; i++ {
		jobChannel <- i
	}

	close(jobChannel) // 關閉工作channel，讓workerMan能夠結束
	select {}
}

func workerMan(wg *sync.WaitGroup, jobs <-chan int, result chan<- int) {
	defer wg.Done()
	for job := range jobs {
		log.Println("workerMan get job: ", job)
		time.Sleep(1 * time.Second) // 模擬處理時間
		result <- job * 2           // 假設處理後的結果是原來的兩倍
	}
}

func printResult(result <-chan int) {
	for res := range result {
		log.Println("Result: ", res)
	}

}
