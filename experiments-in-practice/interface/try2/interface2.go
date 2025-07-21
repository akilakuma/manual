package main

import "fmt"
import "math"

type geometry interface {
	area(float64) float64
	perimeter() float64
}

type shape interface {
	area() float64
	good() int
	hello() string
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

type category1 struct {
	word float64
	data string
	kind int
}

func (s square) area(bonus float64) float64 {
	return s.width * s.height * bonus
}

func (cato category1) area() float64 {
	return cato.word
}

func (cato category1) perimeter() float64 {
	return cato.word
}

func (cato category1) hello() string {
	return cato.data
}

func (s square) perimeter() float64 {
	return 2*s.width + 2*s.height
}

func (c circle) area(bonus float64) float64 {
	return math.Pi * c.radius * c.radius * bonus
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area(10))
	fmt.Println(g.perimeter())
}

func measure2(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	//fmt.Println(s.perimeter()) //偷寫geometry()，事實證明因為shape內沒有，因此報錯
	fmt.Println(s.good()) // 有寫shape裡面的area()和good()，但是上面沒寫 func（s shape) good() int{}的函式
	//所以在main呼叫時會錯
}

func main() {
	// s := square{width: 3, height: 4}
	// c := circle{radius: 6}

	// measure(s)
	// measure(c)

	t := category1{word: 1.2, data: "abc", kind: 5}

	measure2(t)
}
