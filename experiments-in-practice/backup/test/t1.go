package main

// Package
type Collection []interface{}

// func (coll *Collection) Slice(i int, j int) {
// 	arr := transfer(coll)
// 	*coll = Collection(arr[i:j])
// }

func (coll Collection) Ppp() interface{} {
	len := len(coll)
	if len == 0 {
		return nil
	}

	arr := transfer(&coll)
	coll = arr[:len-1]

	return arr[len-1]
}

func (coll *Collection) Pop() interface{} {
	len := len(*coll)
	if len == 0 {
		return nil
	}

	arr := transfer(coll)
	*coll = arr[:len-1]

	return arr[len-1]
}

func (coll *Collection) Push(i interface{}) {
	len := len(*coll)
	if len == 0 {
		return
	}

	arr := transfer(coll)
	arr = append(arr, i)
	*coll = arr

	return
}

func transfer(coll *Collection) []interface{} {
	return []interface{}(*coll)
}

// Client
type Person struct {
	Name string
	Age  int
}

func AA(c Collection) {
	c = append(c, 123)
}

func BB(c *Collection) {
	*c = append((*c), 123)
}

// func test1() {
// 	pList := Collection{
// 		Person{
// 			Name: "Foo",
// 			Age:  11,
// 		},
// 		Person{
// 			Name: "Bar",
// 			Age:  2,
// 		},
// 		Person{
// 			Name: "Fuzz",
// 			Age:  16,
// 		},
// 	}

// 	p := pList.Ppp()
// 	fmt.Println(pList)
// 	fmt.Println(p)

// 	fmt.Println(reflect.TypeOf(pList))

// 	// s := pList.Ppp()
// 	// fmt.Println(s)

// 	// person := Person{
// 	// 	Name: "rascal",
// 	// 	Age:  55,
// 	// }

// 	// pList.Push(person)
// 	fmt.Println(pList)
// }
