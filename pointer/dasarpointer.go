package main

import "fmt"

func main() {
	var a int = 42

	var p *int = &a

	fmt.Println("Nilai a : ", a)
	fmt.Println("Alamat a : ", &a)
	fmt.Println("Pointer p : ", p)
	fmt.Println("Nilai melalui p : ", *p)

	*p = 100
	fmt.Println("Nilai setelah diubah melalui pointer : ", a)

	denganPointer()
	fmt.Println(" ------------------------------------ ")
	tanpaPointer()
}

func denganPointer() {
	nilai := 100

	var nilaiMhs1 *int = &nilai

	fmt.Println("Nilai  : ", nilai)
	fmt.Println("Nilai Mahasiswa 1 : ", *nilaiMhs1)

	*nilaiMhs1 = 90
	fmt.Println("Nilai Mhs1 : ", *nilaiMhs1)
}

func tanpaPointer() {
	nilai := 100

	var nilaiMhs1 int = nilai
	fmt.Println("Nilai : ", nilai)
	fmt.Println("Nilai Mahasiswa 1 : ", nilaiMhs1)

	nilaiMhs1 = 90
	fmt.Println("Nilai Mhs1 : ", nilaiMhs1)
}
