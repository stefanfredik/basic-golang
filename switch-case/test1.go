package main

import "fmt"

func main() {
	jumlahSloki := 25
	switch jumlahSloki {

	case 1:
		fmt.Println("Peminum pemula.")

	case 5:
		fmt.Println("Peminum menengah.")

	case 10:
		fmt.Println("Pemimum senior.")
	default:
		fmt.Println("Peminum master.")
	}
}
