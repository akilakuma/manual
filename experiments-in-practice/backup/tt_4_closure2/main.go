package main

import (
	"fmt"
	"strconv"
)

type Collection []interface{}

type MapFunc func(v interface{}) interface{}

func (coll Collection) Map(f MapFunc) Collection {
	results := make(Collection, 0)
	for _, v := range coll {
		items := f(v)
		results = append(results, items)

	}
	return results
}

type Person struct {
	Name string
	Age  int
}

type ClonePerson struct {
	ID string
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
	results := pList.Map(func(p interface{}) interface{} {
		person := p.(Person)
		clone := ClonePerson{
			ID: person.Name + strconv.Itoa(person.Age),
		}
		return clone
	})

	fmt.Println(results)

}
