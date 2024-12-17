// struct dengan memory

package main

import "fmt"

type Person struct {
	nama         string
	jenisKelamin string
	usia         int32
	tinggi       float32
}

func main() {
	fred := Person{
		nama:         "Fredik Stefan",
		jenisKelamin: "Pria",
		usia:         20,
		tinggi:       115.5,
	}

	fmt.Printf("Hallo, nama saya %s dengan jenis kelamin %s berusia %d tahun dengan tinggi badan %f cm \n", fred.nama, fred.jenisKelamin, fred.usia, fred.tinggi)

	fred.nama = "Stefan Fredik"

	fmt.Println("Nama Baru : ", fred.nama)
}
