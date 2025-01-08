package main

import "fmt"

type Vehicle struct {
	Brand string
	Model string
}

type Car struct {
	Vehicle
	NumberOfDoors int
}

func main() {
	car := Car{
		Vehicle: Vehicle{
			Brand: "Toyota",
			Model: "Corrola",
		},
		NumberOfDoors: 4,
	}

	fmt.Println("Merk Mobil : ", car.Brand)
	fmt.Println("Model Mobil : ", car.Model)
	fmt.Println("Jumlah Pintu : ", car.NumberOfDoors)
}
