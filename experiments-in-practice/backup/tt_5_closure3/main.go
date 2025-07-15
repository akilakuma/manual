package main

import (
	"fmt"
)

// Collection 介面規範
type Collection interface {
	CollectFilter(v interface{}) bool
}

// CollectionFunc 轉介處理
type CollectionFunc func(v interface{}) bool

// CollectFilter 實際操作動作
func (c CollectionFunc) CollectFilter(v interface{}) bool {
	return c(v)
}

func FilterFunc(v []interface{}, f func(v interface{}) bool) []interface{} {
	return Filter(v, CollectionFunc(f))
}

func Filter(v []interface{}, c Collection) []interface{} {
	results := make([]interface{}, 0)
	if v != nil && len(v) > 0 {
		for _, data := range v {
			a := c.CollectFilter(data)
			if a {
				results = append(results, data)
			}
		}
	}

	return results
}

func MoreThan10(v interface{}) bool {
	return v.(Person).Age > 9
}

func lessThan6(v interface{}) bool {
	return v.(Person).Age < 6
}

type Person struct {
	Name string
	Age  int
}

type anything []interface{}

func main() {
	pList := anything{
		Person{
			Name: "pikachu",
			Age:  10,
		},
		Person{
			Name: "superman",
			Age:  5,
		},
	}
	r := FilterFunc(pList, MoreThan10)
	s := FilterFunc(pList, lessThan6)

	fmt.Println("年紀大於10")
	fmt.Println(r)
	fmt.Println("年紀小於6")
	fmt.Println(s)
}
