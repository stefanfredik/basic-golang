package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 1, Y: 2}
	p3 := Point{X: 3, Y: 4}

	fmt.Println(p1 == p2)
	fmt.Println(p1 == p3)
}
