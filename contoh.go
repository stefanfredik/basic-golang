package main

import (
	"fmt"
	"time"
)

func main() {

}

func hitung(tahunLahir int) int {
	tahunSekarang := time.Now().Year()
	usia := tahunSekarang - tahunLahir
	fmt.Println("Hallo Gess				")
	return usia

}
