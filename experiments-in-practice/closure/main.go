package main

import "fmt"

func main() {
	// closure
	//example()

	// rate limiter
	//example2()

	// memoization
	example3()
}

func example() {
	nextSet := intSeq()

	fmt.Println(nextSet())

	for i := 0; i < 3; i++ {
		fmt.Println(nextSet())
	}

	// re initialize nextSet
	nextSet = intSeq()
	fmt.Println("re initialize nextSet")
	fmt.Println(nextSet())

}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
