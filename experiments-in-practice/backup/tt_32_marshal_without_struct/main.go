package main

import (
	"encoding/json"
	"log"
)

type freeTypeMap map[string]interface{}

type SubA struct {
	A1 int `json:"a1"`
	A2 int `json:"a2"`
}

type SubB struct {
	B1 int `json:"b1"`
	B2 int `json:"b2"`
}

type A1ST struct {
	A SubA `json:"sub_a"`
	B SubB `json:"sub_b,omitempty"`
	C int  `json:"c"`
}

type A2ST struct {
	A SubA `json:"sub_a"`
	B SubB `json:"-"`
	C int  `json:"c"`
}

func main() {

	sa := SubA{
		A1: 1,
		A2: 2,
	}
	sb := SubB{
		B1: 1,
		B2: 2,
	}

	fObject := make(freeTypeMap)
	fObject["sub_a"] = sa
	fObject["sub_b"] = sb

	a, err := json.Marshal(&fObject)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(a))

	gObject := make(freeTypeMap)
	gObject["sub_a"] = sa

	b, err2 := json.Marshal(&gObject)
	if err2 != nil {
		log.Println(err2)
	}
	log.Println(string(b))

}
