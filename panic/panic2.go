// panic dan defer. defer akan tetap dicetak walapun panic sudah dieksekusi

package main

import "fmt"

func minumArak(jumlahBotol int) {

	defer fmt.Println("Tidurlah, istirahatlah minum.")

	if jumlahBotol > 10 {
		panic("Stop, kamu sudah minum lebih dari 10. Saatnya istirahat.")
	}
}

func main() {
	minumArak(15)
}
