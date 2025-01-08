package main

import "fmt"

func main() {
	sloki := []int{10, 20, 30, 58, 79}

	fmt.Println("Sebelum : ", sloki)

	// modifikasi slice
	sloki[1] = 5
	fmt.Println("Setelah Modifikasi : ", sloki)

	// memotong slice
	sloki2 := sloki[1:3]

	fmt.Println("Sloki 2 : ", sloki2)

	// menambah 1 element di dalam slice
	sloki2 = append(sloki2, 70)
	fmt.Println("Setelah Menambahkan 1 element : ", sloki2)

	// menambah beberapa element pada slice
	sloki2 = append(sloki2, 80, 90, 100)
	fmt.Println("Setelah Menambahkan beberapa element : ", sloki2)

	// sopy element dari slice ke slice lainya
	sloki3 := make([]int, 2)
	copy(sloki3, sloki2[0:2])
	fmt.Println("Setelah Mencopy element : ", sloki3)

}
