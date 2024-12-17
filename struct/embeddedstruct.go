package main

import "fmt"

type Address struct {
	City    string
	Zipcode int
}

type Man struct {
	Name string
	Age  int
	Address
}

func main() {
	fred := Man{
		Name: "Fredik Stefan",
		Age:  20,
		Address: Address{
			City:    "Denpasar",
			Zipcode: 10001,
		},
	}

	fmt.Println("Nama : ", fred.Name)
	fmt.Println("Kota Asal : ", fred.City)
	fmt.Println("Kode Pos : ", fred.Zipcode)

}
