package main

import "fmt"

func main() {
	var umur int
	fmt.Print("Masukan umur anda : ")
	fmt.Scanln(&umur)

	if umur > 17 {
		fmt.Printf("Mantap, umurmu udah %d Mari kita mabuk, mau arak atau tuak : ", umur)
		var minum string

		fmt.Scanln(&minum)
		if minum == "tuak" {
			fmt.Println("Siapakan gelas untuk minum tuak")
		}

		if minum == "arak" {
			fmt.Println("Siapkan sloki buat minum arak.")
		}

	} else {
		fmt.Println("Sorry, umurmu belum cukup buat minum arak")
	}
}
