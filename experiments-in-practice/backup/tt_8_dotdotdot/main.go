package main

import (
	"fmt"
	"time"
)

func main() {

	// ... 任意傳入參數數量
	// testDotDot("a")
	// testDotDot("b", 1, 2)

	// closure
	// c := testClosure(1, func(a int) int {
	// 	return a * 10
	// })
	// fmt.Println(c)

	testGorutineContext()

}

// testDotDot
func testDotDot(s string, v ...interface{}) {
	fmt.Println(s)
	fmt.Println(v)

	if len(v) > 0 {
		fmt.Println(v[0])
	}

	if len(v) > 1 {
		fmt.Println(v[1])
	}

}

// testClosure
func testClosure(i int, f func(int) int) int {
	return i + f(i)
}

// testContext 在Gorutine執行之前，會先確認執行的狀態，保存這些goruntine運行狀態的就是context
func testContext() {}

type Context interface {
	Done() <-chan struct{}
	Err() error
	Deadline() (deadline time.Time, ok bool)
	Value(key interface{}) interface{}
}
