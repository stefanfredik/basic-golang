package main

import "fmt"

func main() {
	minum := []int{10, 20, 40, 50, 60, 80}

	fmt.Printf("Jumlah botol Minum dalam %d  sesi \n", len(minum))

	jumlahBotol := 0
	for index, value := range minum {
		fmt.Printf("Sesi %d habisin %d botol \n", index+1, value)
		jumlahBotol += value
	}

	fmt.Println("Total botol yang dipake untuk minum : ", jumlahBotol)

	jenisMinuman := []string{"arak", "bir", "tuak", "whisky", "vodca"}

	fmt.Printf("Total minum : %d botol \n", len(jenisMinuman))

	for _, value := range jenisMinuman {
		fmt.Println("List minuman buat mabuk : ", value)
	}

}
