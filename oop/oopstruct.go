package main

import "fmt"

type Person struct {
	Nama string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hallo : ", p.Nama, ",Selamat sore!")
}

func main() {
	person1 := Person{Nama: "Fredik", Age: 20}
	person1.Greet()
}
