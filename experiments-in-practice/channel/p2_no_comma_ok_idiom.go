package main

import (
	"log"
	"time"
)

/*
-說明：
	data, ok := <-b.contactChan
	這個用法叫做Comma-ok Idiom

	如果沒有用這樣的寫法，遇到channel已經關閉，又使用data的method，會造成panic

-output：
	2025/07/15 14:16:53 receive message
	2025/07/15 14:16:55 panic occur, recover handle it!
	2025/07/15 14:16:55 practice end......

*/

func practice2() {

	ss := S{
		output: make(chan int, 10),
	}

	ss.write(10)
	time.Sleep(2 * time.Second)

	// 關閉channel
	close(ss.output)
	ss.write(20)

	log.Println("practice end......")

	time.Sleep(100 * time.Second)
}

type S struct {
	output chan int
}

func (s *S) write(message int) {

	defer func() {
		if recover() != nil {
			log.Println("panic occur, recover handle it!")
		}
	}()

	select {
	case s.output <- message:
		log.Println("receive message")
	default:
		log.Println("select default")
	}
}
