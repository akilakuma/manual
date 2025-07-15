package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

// UserMap 測試用Map
type UserMap map[int]int

var muLock sync.RWMutex

// Copy 複製一份新的UserMap
func (macMap *UserMap) Copy() UserMap {

	macMapDeReff := *macMap
	muLock.Lock()
	defer muLock.Unlock()

	newMap := make(UserMap, 0)
	for k, v := range macMapDeReff {
		newMap[k] = v
	}
	// return macMapDeReff
	return newMap
}

func main() {
	test()
	race()
}

func test() {
	m := make(UserMap)

	for i := 0; i < 100; i++ {
		key := rand.Intn(128)
		m[i] = key
	}
	n := m
	// n := m.Copy()
	go func() {
		for {
			key := rand.Intn(128)
			// muLock.Lock()
			m[key] = key
			// muLock.Unlock()
			// log.Println("set once")
		}
	}()
	// muLock.Lock()
	for key, value := range n {
		log.Println(key, value)
	}
	// muLock.Unlock()
}

func race() {
	a := 0
	times := 3
	c := make(chan bool)

	for i := 0; i < times; i++ {
		go func() {
			a++
			c <- true
		}()
	}

	for i := 0; i < times; i++ {
		<-c
	}
	fmt.Printf("a = %d\n", a)
}
