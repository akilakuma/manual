package main

import "fmt"

type Cellphone struct {
	ID   int
	name string
}

type Iphone struct {
	Cellphone
}

type HTC struct {
	Cellphone
}

func (c *Cellphone) SetID(n int) {
	c.ID = n + 100
}

func (i *Iphone) SetID(n int) {
	i.ID = n + 500
	i.Cellphone.SetID(-1000)
}

const (
	a1 =  iota
	a2
	a3
	a4
	a5
)

func main() {

	newIphone := Iphone{}
	fmt.Println(newIphone)

	newIphone2 := Iphone{
		Cellphone{
			ID:   1,
			name: "iphone",
		},
	}
	fmt.Printf("%#v", newIphone2)
	fmt.Println()

	newIphone2.SetID(99)
	fmt.Printf("%#v", newIphone2)
	fmt.Println()

	fmt.Printf("%#v", newIphone2.ID)

	fmt.Println("a2", a5)
}
