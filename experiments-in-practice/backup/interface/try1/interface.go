package main

import "fmt"
import "math"

/*
	golang的interface，看來是一個很容易實踐『物件導向』所謂『多型』的部分
*/

// 寫法跟struct 有點像
// 定義了area()和perimeter()兩個function，其中area要傳入一個參數，並且為float64的型態
type geometry interface {
	area(float64) float64
	perimeter() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 這個area() 跟下面的area()是不一樣的，golang實現了多型的奧義！！！
func area() int {
	return 123456
}

// 前面這個(s square) 指定傳入一個square類型的struct，但這樣的function只供有適用介面的struct使用
func (s square) area(bonus float64) float64 {
	return s.width * s.height * bonus
}

func (s square) perimeter() float64 {
	return 2 * s.width + 2 * s.height
}

// 前面這個(c circle) 為指定傳入一個circle類型的struct
func (c circle) area(bonus float64) float64 {
	return math.Pi * c.radius * c.radius *bonus
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// square 不認得geometry
// geometry 也不認得square或circle
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area(10))
	fmt.Println(g.perimeter())
}

func main() {
	fmt.Println(area())
	s := square{width:3,height:4}
	c := circle{radius:6}

	measure(s)
	measure(c)
}
