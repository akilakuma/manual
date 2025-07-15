package main

import (
	"fmt"
	goSet "github.com/deckarep/golang-set"
)

/*
	golang-set 轉成interface就可以compare set，再轉回來
*/

func main() {

	var (
		source = []int{314256310, 137127114, 373760725, 490558493, 312216009, 452105985, 490541907, 436123127, 303155584, 374836401}
		target = []int{314256310, 137127114, 373760725, 490558493, 312216009, 452105985}
	)

	a := DifferSet(source, target)
	fmt.Println(a)

}

func DifferSet(source, target []int) []int {
	sourceSet := goSet.NewSetFromSlice(IntToInterface(source))
	targetSet := goSet.NewSetFromSlice(IntToInterface(target))
	operatedSet := sourceSet.Difference(targetSet)
	result := operatedSet.ToSlice()
	return InterfaceToInt(result)
}

// IntToInterface []int 轉換成 []interface{}
func IntToInterface(soruce []int) []interface{} {

	r := make([]interface{}, len(soruce))
	for i := range r {
		r[i] = interface{}(soruce[i])
	}
	return r
}

// InterfaceToInt []interface{} 轉換成 []int
func InterfaceToInt(soruce []interface{}) []int {

	r := make([]int, len(soruce))
	for i := range soruce {
		r[i] = soruce[i].(int)
	}
	return r
}
