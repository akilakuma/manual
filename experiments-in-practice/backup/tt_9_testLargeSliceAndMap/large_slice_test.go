package test

import (
	"fmt"
	"log"
	"runtime"
	"testing"
	"time"
)

func TestLargeSlice(t *testing.T) {

	log.Println("100萬個slice的記憶體使用狀況")
	PrintMemUsage()

	timeS := time.Now()
	LargeSlice()
	timeE := time.Since(timeS)
	fmt.Print("花費時間  ")
	fmt.Print(timeE.Seconds())
	fmt.Println("秒")

	PrintMemUsage()

}

func TestLargeSliceWithStruct(t *testing.T) {

	fmt.Println()
	log.Println("100萬個map[int]struct的記憶體使用狀況")
	PrintMemUsage()

	timeS := time.Now()
	LargeSliceWithStruct()
	timeE := time.Since(timeS)
	fmt.Print("花費時間  ")
	fmt.Print(timeE.Seconds())
	fmt.Println("秒")

	PrintMemUsage()

}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
