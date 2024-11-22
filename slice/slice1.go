package main

import (
	"fmt"
)

func main() {
	jenisBotol := [5]int{10, 20, 30, 12, 5}
	jumlahSloki := []int{10, 23, 61, 89, 90}
	slice := jenisBotol[0:2]

	// check data type of jenisBotol using type
	fmt.Println("jenisBotol : ", jenisBotol)
	fmt.Println("jumlahSlokir : ", jumlahSloki)
	fmt.Println("slice : ", slice)
}
