package main

import "fmt"

func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Reover dari panic :  ", r)

		}
	}()

	if b == 0 {
		panic("Tidak bisa membagi dengan nol!")
	}

	fmt.Println("Hasil : ", a/b)
}

func main() {
	safeDivide(10, 2)
	safeDivide(10, 0)

	fmt.Println("Program tetap berjalan")
}
