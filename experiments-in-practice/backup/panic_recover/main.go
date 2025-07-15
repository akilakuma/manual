package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"sync"
)

func main() {

	// recover main
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			// 或者自定義的處理
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
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
			wg.Done()
		}
	}()

	for i := 0; i < 10; i++ {
		// recover  A-2
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Error:", err)
				wg.Done()
			}
		}()

		count := i
		wg.Add(1)
		go testB(wg, count)
	}
}

func testB(wg *sync.WaitGroup, count int) {

	// recover B
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("Error:", err)
	// 		wg.Done()
	// 	}
	// }()
	type a map[string]string
	var aa a

	fmt.Println(count)
	if count == 5 {
		// panic("occur panic!")
		aa["123"] = "bbb"
	}

	fmt.Println("doing something")
	wg.Done()
}
