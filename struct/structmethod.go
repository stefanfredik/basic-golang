package main

import "fmt"

type Rectangle struct {
	Width  float32
	Height float32
}

func (r Rectangle) Area() float64 {
	return float64(r.Width) * float64(r.Height)
}

func (r *Rectangle) Resize(width, height float64) {
	r.Width = float32(width)
	r.Height = float32(height)
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}

	fmt.Println("Luas Area : ", rect.Area())

	rect.Resize(20, 40)
	fmt.Println("Update Area : ", rect.Area())

}
