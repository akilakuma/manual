package main

import "fmt"

func main() {

	p1 := testInt()
	p2 := testString()
	fmt.Println(detectType(p1)) // Output: int
	fmt.Println(detectType(p2)) // Output: string
}

func testInt() interface{} {
	return 123
}

func testString() interface{} {
	return "abc"
}

func detectType(p interface{}) string {
	switch p.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
