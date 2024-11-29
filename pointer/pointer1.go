package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Println("Nilai p : ", p)
	fmt.Println("Nilai *p :", *p)

	*p = 20
	fmt.Println("Nilai x: ", x)
}
