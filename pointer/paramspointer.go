package main

import "fmt"

func changeValue(val *int) {
	*val = 42
}

func main() {
	x := 10
	fmt.Println()
}
