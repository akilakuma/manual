package main

import (
	"context"
	"fmt"
	"time"
)

func testGorutineContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go watch(ctx, "[監控1]")
	go watch(ctx, "[監控2]")
	go watch(ctx, "[監控3]")

	time.Sleep(10 * time.Second)
	fmt.Println("通知監控停止")
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("已停止")
			return
		default:
			fmt.Println("監控中")
			time.Sleep(2 * time.Second)
		}
	}
}

func Dead() <-chan struct{} {
	return nil
}

type HandleFunc func(*Context)
