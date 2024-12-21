package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	go printNumbers()

	time.Sleep(6 * time.Second)

	fmt.Println("Main function done")

}
