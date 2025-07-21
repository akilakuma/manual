package main

import "fmt"

func main() {
	// defer 的執行順序是後進先出LIFO
	fmt.Println("example1")
	example1()
	// C
	// B
	// A

	// 命名返回變數，defer會影響到
	fmt.Println("example2")
	fmt.Println(example2())

	// 非命名返回變數，return 已經將value copy，defer不影響
	fmt.Println("example3")
	fmt.Println(example3())

}

func example1() {
	defer func() {
		fmt.Println("A")
	}()

	defer func() {
		fmt.Println("B")
	}()

	defer func() {
		fmt.Println("C")
	}()

}

func example2() (x int) {
	x = 1
	defer func() {
		x = x + 1
	}()
	return
}

func example3() int {
	x := 1
	defer func() {
		x = x + 1
	}()
	return x
}
