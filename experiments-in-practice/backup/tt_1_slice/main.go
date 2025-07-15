package main

import (
	"fmt"
	"log"
)

type data struct {
	name string
}

func main() {

	// test1()
	// test2()
	test3()
}

func test1() {
	kumaSlice1 := []int{100, 200, 300}
	log.Println(len(kumaSlice1), cap(kumaSlice1), kumaSlice1)
	// 2018/08/13 21:04:38 3 3 [100 200 300]

	kumaSlice2 := kumaSlice1[1:]
	log.Println(len(kumaSlice2), cap(kumaSlice2), kumaSlice2)
	// 2018/08/13 21:04:38 2 2 [200 300]

	kumaSlice3 := kumaSlice2[1:]
	log.Println(len(kumaSlice3), cap(kumaSlice3), kumaSlice3)
	// 2018/08/13 21:04:38 1 1 [300]

	for k := range kumaSlice2 {
		kumaSlice2[k] += 2
	}

	log.Println(kumaSlice1)
	// 2018/08/13 21:13:57 [100 202 302]
	log.Println(kumaSlice2)
	// 2018/08/13 21:13:57 [202 302]
	log.Println(kumaSlice3)
	// 2018/08/13 21:13:57 [302]
}

func test2() {
	kumaSlice1 := []int{100, 200, 300}
	kumaSlice2 := kumaSlice1[1:]
	// kumaSlice3 := kumaSlice2[1:]

	kumaSlice3 := append(kumaSlice2, 400)

	for k := range kumaSlice2 {
		kumaSlice2[k] += 2
	}

	log.Println(kumaSlice1)
	// 2018/08/13 21:13:57 [100 202 302]
	log.Println(kumaSlice2)
	// 2018/08/13 21:13:57 [202 302]
	log.Println(kumaSlice3)
	// 2018/08/13 21:18:07 [200 300 400]
}

func test3() {

	s := []int{5}

	s = append(s, 7)
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0])

	s = append(s, 9)
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0])

	x := append(s, 11) // 只有擴到x，但s還是舊的
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0], "ptr(x) =", &x[0])
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[1], "ptr(x) =", &x[1])
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[2], "ptr(x) =", &x[2])

	b := &s[2]
	*b = 100
	fmt.Println(s)
	fmt.Println(x)
	fmt.Println()

	y := append(s, 12)
	fmt.Println(y)
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0], "ptr(y) =", &y[0])

	z := append(x, 12)
	fmt.Println(z)
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0], "ptr(z) =", &z[0])

	a := append(s, 0)
	fmt.Println(a)
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0], "ptr(z) =", &a[0])

	/*
		cap(s) = 2 ptr(s) = 0xc4200160a0
		cap(s) = 4 ptr(s) = 0xc420018160
		cap(s) = 4 ptr(s) = 0xc420018160 ptr(x) = 0xc420018160
		cap(s) = 4 ptr(s) = 0xc420018168 ptr(x) = 0xc420018168
		cap(s) = 4 ptr(s) = 0xc420018170 ptr(x) = 0xc420018170
		[5 7 100]
		[5 7 100 11]

		[5 7 100 12]
		cap(s) = 4 ptr(s) = 0xc420018160 ptr(y) = 0xc420018160
		[5 7 100 12 12]
		cap(s) = 4 ptr(s) = 0xc420018160 ptr(z) = 0xc420012180
		[5 7 100 0]
		cap(s) = 4 ptr(s) = 0xc420018160 ptr(z) = 0xc420018160
	*/
}
