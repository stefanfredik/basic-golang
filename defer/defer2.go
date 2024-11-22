package main

import (
	"fmt"
	"os"
)

func writeFile(filename string) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Gagal membuat file.")
		return
	}

	defer file.Close()

	file.WriteString("Hallo bro, mari minum \n \n Mabukkkkkk")

	fmt.Println("File berhasil ditulis.")
}

func main() {
	writeFile("mabuk.txt")
}
