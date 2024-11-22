package main

import "fmt"

func main() {
	targetBotol := 25

	botolSelesai := 0

	for botolSelesai = 0; botolSelesai < targetBotol; botolSelesai++ {
		fmt.Println("Jumlah botol yang udah diminum : ", botolSelesai)

		if botolSelesai == 15 {
			fmt.Printf("Kamu udah minum %d botol. Istirahatlah dulu. nanti lanjut lagi", botolSelesai)
			break
		}
	}

}
