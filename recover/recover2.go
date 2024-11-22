// kombinasi recover panic dan defer

package main

import "fmt"

func proccess() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print("Recover dari error : ", r)
		}
	}()

	fmt.Println("Memulai proccess ...")

	panic("Terjadi kesalahan \n")

	fmt.Println("Program selesai")
}

func main() {
	proccess()
	fmt.Println("Program tetap berjalan.")
}
