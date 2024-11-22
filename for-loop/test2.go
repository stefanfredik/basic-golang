package main

import "fmt"

func main() {
	sloki := 0

	for sloki < 5 {
		fmt.Printf("Mantap, kamu Sudah minum %d sloki, tambah 1 lagi \n", sloki)
		sloki++
	}

	fmt.Printf("Sudah cukup, kamu udah minum %d sloki. istirahatlah dulu", sloki)
}
