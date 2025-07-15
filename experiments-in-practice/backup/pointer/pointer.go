package main

import (
	"fmt"
	"reflect"
)

/*
	搭配以下兩篇文章說明，有全新領悟
	http://wp.mlab.tw/?p=176
	https://studygolang.com/articles/13120?fr=sidebar
*/

func main() {
	test2()
}

func test1() {
	var i int = 10
	ip := &i
	fmt.Println(reflect.TypeOf(ip))    // ip 此時就是一個指標的型態
	fmt.Printf("原始指针的内存地址是：%p\n", &i)  // 變數i本身會有個記憶體位置
	fmt.Printf("原始指针的内存地址是：%p\n", &ip) // 這裡的ip有個新位置
	modify(ip)
	fmt.Println(i)
}

func modify(ip2 *int) {
	fmt.Println(ip2)
	fmt.Printf("ip2的内存地址是：%p\n", &ip2) // 這裡的ip2也是新產生的位置

	// ip2前面的*，是表示提取ip2的value，跟指定型態沒關係
	// 這樣的用法就是吃準ip2已經明明白白，確確實實是一個指標，所以這樣使用。
	*ip2 = 1
}

func test2() {
	originMap := make(map[int]int, 0)
	fmt.Printf("原始指针的内存地址是：%p\n", &originMap)

	test3(&originMap)
	test4(originMap)
}

func test3(myMap *map[int]int) {
	fmt.Printf("原始指针的内存地址是：%p\n", myMap)
}

func test4(myMap map[int]int) {
	fmt.Printf("原始指针的内存地址是：%p\n", &myMap)
}
