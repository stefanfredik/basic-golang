package main

import "fmt"

// function sederhana

func sayHello() {
	fmt.Println("Hallo brother, mari minum tuak.")
}

// function satu parameter
func mariMinum(nama string) {
	fmt.Printf("Hallo brother %s \n", nama)
}

// parameter lebih dari satu
func hobiMinum(nama string, suka string) {
	fmt.Printf("Hallo brother %s, kesukaanmu minum %s. \n", nama, suka)
}

// satu nilai return
func hitungJumlahSloki(ukuranBotol float32) float32 {
	sloki := 44

	return ukuranBotol / float32(sloki)
}

// lebih dari 1 return
func hitungSlokiPerBotol(jenisMinuman string) (int, float32) {
	var ukuran int
	var jumlahSloki float32

	listMinuman := map[string]int{
		"tuak":   100,
		"bir":    450,
		"arak":   300,
		"whisky": 700,
		"vodka":  450,
	}

	for index, value := range listMinuman {
		sloki := 44
		if jenisMinuman == index {
			ukuran = value
			jumlahSloki = float32(value / sloki)
			break
		}
	}

	return ukuran, jumlahSloki
}

// fungsi dengan parameter nilai default

func mariMinumTuak(nama string) {
	if nama == "" {
		nama = "Brother"
	}

	fmt.Printf("Hallo %s, mari minum tuak! \n", nama)
}

//variadic function : menerima parameter tidak terbatas

func minumanKesukaan(minuman ...string) {

	fmt.Print("Aku sukan minum ")
	for _, value := range minuman {
		fmt.Printf("%s, ", value)
	}
	fmt.Println("")
}

// fungsi sebagai parameter

func hitungJumlahbotol()

func main() {
	sayHello()
	mariMinum("Petrus")
	hobiMinum("Petrus", "tuak")

	bir1Botol := 400
	fmt.Printf("Jumlah sloki 1 botol : %f sloki. \n", hitungJumlahSloki(float32(bir1Botol)))

	ukuran, jumlahSloki := hitungSlokiPerBotol("tuak")

	fmt.Printf("Ukuran botol %d ml dengan jumlah sloki sebanyak %.0f sloki. \n", ukuran, jumlahSloki)
	mariMinumTuak("")
	minumanKesukaan("tuak", "bir", "arak", "whisky", "vodka")

	// function anonym
	// function yang bisa langsung dijalankan dan hanya bisa dipanggil sekali

	// langsung dijalankan
	func() {
		fmt.Println("Hallo brother, kapan kita bisa mabuk lagi ")
	}()

	// disimpan dalam variabel
	minum := func(nama string) {
		fmt.Printf("Hallo %s, Satu botol bir manalah cukup", nama)
	}

	minum("Petrus")
}
