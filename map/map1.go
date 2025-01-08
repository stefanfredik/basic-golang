package main

import "fmt"

func main() {

	// inisiasi dan deklarsi map
	var fred map[string]interface{} = make(map[string]interface{})

	fmt.Println(fred)

	stefan := make(map[string]string)
	fmt.Println(stefan)

	nilai := make(map[string]int)
	fmt.Println(nilai)

	jumlahUang := map[string]int{
		"wisuda":      5000000,
		"yudisium":    1000000,
		"sertifikasi": 500000,
	}

	fmt.Println(jumlahUang)

	// Menambahkan elemen map
	fred["nama"] = "Fredik Stefan"
	fred["umur"] = 20
	fred["jenisKelamin"] = "Pria"
	fred["domisili"] = "Bali"

	stefan["nama"] = "Stefan Fredik"
	stefan["jenisKelamin"] = "Pria"
	stefan["domisili"] = "Bali"

	nilai["programming"] = 100
	nilai["filsafat"] = 100

	fmt.Println()
	fmt.Println(fred)
	fmt.Println(stefan)
	fmt.Println(nilai)

	// mengakses nilai map
	fmt.Println()
	fmt.Println("Nama : ", fred["nama"])
	fmt.Println("Nama : ", stefan["nama"])
	fmt.Println("Nilai Programming : ", nilai["programming"])
	fmt.Println("Uang Wisuda : ", jumlahUang["wisuda"])
	// menghapus value
	delete(nilai, "filsafat")
	fmt.Println(nilai)

	// memeriksa keberadaan key
	value, exists := nilai["programming"]

	if exists {
		fmt.Println("Nilai Programming : ", value)
	}

	// memeriksa keberadaan key
	v, e := nilai["filsafat"]

	if e {
		fmt.Println("Nilai Filsafat : ", v)
	} else {
		fmt.Println("Sorry key tidak ditemukan ")
	}

}
