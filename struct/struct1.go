package main

import "fmt"

type Mahasiswa struct {
	namaMahasiswa string
	usia          int
	statusNikah   bool
}

func main() {
	mahasiswa1 := Mahasiswa{namaMahasiswa: "Fredik Stefan", usia: 25, statusNikah: false}

	fmt.Println(mahasiswa1)
	fmt.Println("Nama  : ", mahasiswa1.namaMahasiswa)
	fmt.Println("Usia : ", mahasiswa1.usia)
	fmt.Println("Status Nikah : ", mahasiswa1.statusNikah)

	mahasiswa2 := Mahasiswa{
		namaMahasiswa: "Stefan Fredik",
		usia:          29,
		statusNikah:   false,
	}

	fmt.Println(mahasiswa2)

}
