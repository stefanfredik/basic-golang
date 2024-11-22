// panic merupakan fungsi untuk mengehntikan blok program secara paksa jika terjadi kesalahan

package main

import "fmt"

func divide(a, b int) int {
	if b == 0 {
		panic("Ga boleh ada bagi nol.")
	}

	return a / b
}

func main() {
	fmt.Println("Hasil : ", divide(10, 5))
	fmt.Println("Hasil : ", divide(5, 0))
}
