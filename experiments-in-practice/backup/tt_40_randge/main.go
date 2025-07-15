package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"os/exec"
	"time"
)

// DefaultRandInstance 預設的亂數實體
var DefaultRandInstance *RandManager

func main() {
	DefaultRandInstance = NewRandManager()
	go DefaultRandInstance.KeepGenerateRandNum()

	d := time.Duration(time.Second * 10)
	t := time.NewTicker(d)
	defer t.Stop()

	// loop:
	for {
		select {
		case <-t.C:
			log.Println(len(DefaultRandInstance.randQueue))
			// break loop
		}
	}
}

// RandManager 亂數產生業務
type RandManager struct {
	randQueue chan int
}

// NewRandManager 新的亂數產生器
func NewRandManager() *RandManager {
	return &RandManager{
		randQueue: make(chan int, 100000000),
	}
}

// KeepGenerateRandNum 持續產生亂數
func (r *RandManager) KeepGenerateRandNum() {

	for {
		// r.RandGenerate()
		t1 := time.Now()
		r.RandGenerate()
		// numSlice := r.RandGenerate()
		t2 := time.Now().Sub(t1)
		log.Println("t2:", t2)
		// t3 := time.Now()
		// for _, num := range numSlice {
		// 	// log.Println(num)
		// 	r.randQueue <- num
		// }
		// t4 := time.Now().Sub(t3)
		// log.Println("t4:", t4)
		// log.Println("another send,len of r.randQueue :", len(r.randQueue))
	}
}

// RandGenerate 亂數產生
func (r *RandManager) RandGenerate() []int {
	var out bytes.Buffer
	cmd := exec.Command("head", "-c", "20", "/dev/hwrng")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {

		var isSuccess bool

		for !isSuccess {
			//本機mac開發無/dev/hwrng
			cmd := exec.Command("head", "-c", "300000", "/dev/urandom")
			cmd.Stdout = &out
			err2 := cmd.Run()
			if err2 != nil {
				log.Println("error", "error_test: "+err.Error())
				continue
			}
			isSuccess = true
		}
	}
	n := bytes.Split(out.Bytes(), []byte(" "))

	var numSlice []int
	for _, v := range n {
		if len(v) > 4 {
			randNum := binary.BigEndian.Uint32(v)
			numSlice = append(numSlice, int(randNum))
		}
	}
	// log.Println("generate!")
	return numSlice
}

// GetRand 取得亂數
func (r *RandManager) getRand() int {
	n, ok := <-r.randQueue
	if ok {
		// log.Println("getRand")
		return n
	}
	log.Println("r.randQueue 錯誤")
	return 0
}
