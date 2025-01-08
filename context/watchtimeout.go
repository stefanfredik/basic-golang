package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task selesai")
	case <-ctx.Done():
		fmt.Println("Task dibatalkan : ", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go doSomething(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
}
