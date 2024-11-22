package main

import "fmt"

func main() {
	umur := 17

	if umur >= 18 {
		fmt.Println("Mantap, mari kita mabuk")
	} else if umur == 17 {
		fmt.Println("Sabar bro, tahun depan kita mabuk")
	} else {
		fmt.Println("Nunggu dulu, kamu belum bisa mabuk")
	}
}
