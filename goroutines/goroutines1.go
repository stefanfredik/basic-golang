package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers() {
	for i := 1; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		printNumbers()
	}()

	fmt.Println("Execution!")
	wg.Wait()

	fmt.Println("Main function done")

}
