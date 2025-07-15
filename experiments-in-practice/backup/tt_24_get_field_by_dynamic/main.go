package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	// test1()
	// test2()
	test3()

}

func test1() {
	v := Vertex{1, 2}
	fmt.Println(getField(&v, "X"))
	fmt.Println(setField(&v, "X"))
}

func test2() {
	b := BigStruct{
		Name: "繃啾",
		ID:   999,
	}

	result := `{"name":"little johnny","other":{"max_fg":9,"opt":[{"double":1,"times":20},{"double":1,"times":15},{"double":1,"times":12},{"double":1,"times":9},{"double":1,"times":6}],"select_trigger_count":3,"select_trigger_way":"reel","switch":true}}`

	var middleInterface interface{}
	json.Unmarshal([]byte(result), &middleInterface)
	log.Println(middleInterface)

	f := setSTField(&b, "Option")
	f.Set(reflect.ValueOf(middleInterface))

	log.Println(b)
}

type BigStruct struct {
	Name   string
	ID     int
	Option interface{}
}

func test3() {

	result := `{"name":"little johnny","other":{"max_fg":9,"opt":[{"double":1,"times":20},{"double":1,"times":15},{"double":1,"times":12},{"double":1,"times":9},{"double":1,"times":6}],"select_trigger_count":3,"select_trigger_way":"reel","switch":true}}`

	var middleInterface interface{}
	json.Unmarshal([]byte(result), &middleInterface)
	// log.Println(middleInterface)

	var c ComplexStruct

	c.Name = "meow meow"
	c.ID = 12355679

	f := setSTInternalField(&c, "Option")
	f.Set(reflect.ValueOf(middleInterface))
	log.Println(c)
}

type ComplexStruct struct {
	Name string
	ID   int
	Res  struct {
		ErrMsg string
		Data   struct {
			Option interface{}
			Flag   int
		}
	}
}
