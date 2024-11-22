package main

import "fmt"

func main() {
	var panjang int = 100
	var lebar int = 50
	var luas = hitungLuas(panjang, lebar)
	fmt.Println(luas)
}

func hitungLuas(p int, l int) int {
	luas := (p * l)
	return luas
}
