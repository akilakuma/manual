package main

import (
	"fmt"
		Random "math/rand"
		"time"
)

func main() {

		// random 0~99
	r := GetRandomVary()
	target := r.Intn(10)

	fmt.Println(target)
}



// GetRandomVary 變數產生實例
func GetRandomVary() *Random.Rand {

	s1 := Random.NewSource(time.Now().UnixNano())
	return Random.New(s1)
}
