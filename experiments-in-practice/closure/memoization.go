package main

import (
	"fmt"
)

func example3() {
	fn, getCounter := memoFib()

	fmt.Println(fn(10), getCounter())
	fmt.Println(fn(50), getCounter())

}

func memoFib() (func(int) int, func() int) {
	memo := make(map[int]int)
	counter := 0

	var fib func(int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		if val, ok := memo[n]; ok {
			return val
		}
		counter++ // 計算次數 +1
		memo[n] = fib(n-1) + fib(n-2)
		return memo[n]
	}

	return fib, func() int {
		return counter
	}
}
