package main

import (
	"context"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func main() {

	for {
		for i := 0; i < 10000; i++ {
			go test1()
		}
		time.Sleep(1 * time.Second)
		log.Println("建立1萬個.....")
	}
}
func test1() {
	var (
		timeoutDuration = 100 * time.Millisecond
		ctx             context.Context
		cancel          context.CancelFunc
	)

	ctx, cancel = context.WithTimeout(context.Background(), timeoutDuration)

	go func() {
		defer cancel()

		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(120)
		timeRandom := time.Duration(x) * time.Millisecond

		time.Sleep(timeRandom)

		select {
		case <-ctx.Done():

		default:
		}
	}()
	runtime.Gosched()

	<-ctx.Done()
}
