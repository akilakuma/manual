package main

import (
	"fmt"
	"time"
)

type Foo struct {
	Bar string
}

func (foo *Foo) Fuzz() {
	fmt.Println(foo.Bar)
}

func main() {
	test1()
	go test2(2)
	go test2(0)

	
	for true {
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func test1() {
	foo := &Foo{
		Bar: "YOOOO",
	}
	foo.Fuzz()
	fmt.Println(&foo)
}

func test2(sleepTime int) {

	foo := &Foo{
		Bar: "NANANA",
	}

	if sleepTime == 0 {
		foo.Bar = "LU!"
	}

	time.Sleep(time.Duration(sleepTime) * time.Second)
	foo.Fuzz()
	fmt.Println(&foo)
}
