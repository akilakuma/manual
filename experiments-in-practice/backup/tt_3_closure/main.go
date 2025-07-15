package main

import (
	"fmt"
)

type Collection []interface{}

type FilterFunc func(v interface{}) bool

func (coll Collection) Filter(f FilterFunc) Collection {
	results := make(Collection, 0)
	for _, v := range coll {
		if f(v) {
			results = append(results, v)
		}
	}
	return results
}

type Person struct {
	Name string
	Age  int
}


func main() {
	pList := Collection{
		Person{
			Name: "pikachu",
			Age:  10,
		},
		Person{
			Name: "superman",
			Age:  5,
		},
	}
	results := pList.Filter(func(p interface{}) bool {
		return p.(Person).Age > 9
	})

	fmt.Println(results)

}
