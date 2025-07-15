package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"sync"
)

/*
-說明：
	示範panic recover
	以及需要的話，印出stack trace
	如果有開goroutine，記住recover需要擺在開出去的goroutine function內

-output：
	execute testA
	execute testB
	panic recover Error: occur panic!

*/

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			// stack trace print
			log.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	var wg sync.WaitGroup
	testA(&wg)
	wg.Wait()
}

func testA(wg *sync.WaitGroup) {

	// recover  A-1
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	fmt.Println("execute testA")

	wg.Add(1)
	go testB(wg)

}

func testB(wg *sync.WaitGroup) {

	// recover B
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover Error:", err)
			wg.Done()
		}
	}()

	fmt.Println("execute testB")
	panic("occur panic!")

	fmt.Println("doing something")
	wg.Done()
}
