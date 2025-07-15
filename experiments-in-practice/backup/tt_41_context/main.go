package main

import (
	"fmt"
	"log"
)

func main() {

	handleChan := make(chan string, 10)

	for i := 0; i < 10; i++ {
		handleChan <- "ball"
	}

	lenOfchan := len(handleChan)
	for i := 0; i < lenOfchan; i++ {
		fmt.Println("c Drop Ball -> ", <-handleChan)
		log.Println("i:", i)

		log.Println("len of channel:", len(handleChan))
	}
}
