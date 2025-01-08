package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	result := doWork(ctx)
	fmt.Println(result)
}

func doWork(ctx context.Context) string {
	select {
	case <-time.After(3 * time.Second):
		return "Pekerjaan Selesai"
	case <-ctx.Done():
		return "Pekerjaan dibatalkan." + ctx.Err().Error()
	}
}
