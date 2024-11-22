package main

import "fmt"

func main() {
	jenisMinuman := map[string]int{
		"Tuak": 10,
		"Bir":  25,
		"Arak": 5,
	}

	for nama, jumlah := range jenisMinuman {
		fmt.Printf("Minuman %s memiliki stok %d botol. \n", nama, jumlah)
	}
}
