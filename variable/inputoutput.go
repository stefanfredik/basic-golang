package main

import "fmt"

func main() {

	var nama string
	fmt.Print("Masukkan nama anda : ")
	//fmt.Scanln(&nama) //ini tidak bisa menerima spasi
	fmt.Scan(&nama)

	fmt.Println("Hallo," + nama)

}
