package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)

	fmt.Println(slice)
	fmt.Println("Panjang : ", len(slice), "Kapasitas : ", cap(slice))

}
