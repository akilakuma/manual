package main

import "fmt"

const (
	mutexLocked      = 1 << iota // 2進位：1 , 10進位：1, mutex is locked
	mutexWoken                   // 2進位：10 , 10進位：2
	mutexStarving                // 2進位：100 , 10進位：4
	mutexWaiterShift = iota      // 10進位：3

	starvationThresholdNs = 1e6
)

func main() {
	// fmt.Printf("%b", mutexLocked)
	// fmt.Println()
	// fmt.Printf("%b", mutexWoken)
	// fmt.Println()
	// fmt.Printf("%b", mutexStarving)
	// fmt.Println()
	// fmt.Printf("%b", mutexWaiterShift)
	// fmt.Println()
	// fmt.Println(mutexWaiterShift)

	// old :=4
	// fmt.Printf("%b", old&(mutexLocked|mutexStarving))
	// fmt.Println()

	// new := 1
	// fmt.Printf("%b", 1<<mutexWaiterShift)
	// fmt.Println()
	// new += 1 << mutexWaiterShift
	// fmt.Printf("%b", new)
	// fmt.Println()
	// new += 1 << mutexWaiterShift
	// fmt.Printf("%b", new)
	// fmt.Println()

	new := 4
	new |= mutexLocked
	fmt.Printf("%b", new)
	fmt.Println()
}

func table() {
	fmt.Printf("%b", mutexLocked)
	fmt.Println()
	fmt.Printf("%b", mutexWoken)
	fmt.Println()
	fmt.Printf("%b", mutexStarving)
	fmt.Println()
	fmt.Printf("%b", mutexWaiterShift)
	fmt.Println()
	fmt.Println(mutexWaiterShift)
}
