package main

import (
	"fmt"
	"time"
)

func greet(name string, ch chan string) {
	time.Sleep(10 * time.Second)
	ch <- "Hello," + name
}

func main() {

	ch := make(chan string)
	go greet("Alice", ch)
	message := <-ch

	fmt.Println(message)

	fmt.Println("Executed!")
}
