package main

import (
	"log"
	"time"
)

func main() {

	ss := S{
		output: make(chan int, 10),
	}

	ss.write(10)
	time.Sleep(2 * time.Second)

	close(ss.output)
	ss.write(20)

	log.Println("very good!")

	time.Sleep(100 * time.Second)
}

type S struct {
	output chan int
}

func (s *S) write(message int) {

	defer func() {
		if recover() != nil {
			log.Println("hello man!")
		}
	}()

	select {
	case s.output <- message:
		log.Println("receive good")

	default:
		log.Println("跑到default")
	}
}
