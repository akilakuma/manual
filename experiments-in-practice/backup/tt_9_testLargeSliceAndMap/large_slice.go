package test

import (
	"math/rand"
	"time"
)

type userData struct {
	curreny string
	parents []int
	randId  int
}

func LargeSlice() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// log.Println(r1.Intn(10000))

	var largeS []int

	for i := 0; i < 1000000; i++ {
		largeS = append(largeS, r1.Intn(10000))
	}

}

func LargeSliceWithStruct() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// log.Println(r1.Intn(10000))

	largeS := make(map[int]userData)

	for i := 0; i < 1000000; i++ {
		largeS[i] = userData{
			curreny: "CNY",
			parents: []int{1, 2, 3, 4},
			randId:  r1.Intn(10000),
		}
	}

}
