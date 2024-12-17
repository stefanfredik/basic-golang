package main

import (
	"fmt"
)

type Orang struct {
	name   string
	age    int32
	gender string
}

func main() {
	o := &Orang{name: "Fredik Stefan", age: 20, gender: "Man"}

	orang2 := Orang{name: "Kopi FLores", age: 29, gender: "Man"}

	fmt.Println("Nama : ", o.name)
	fmt.Println("Orang 2 : ", orang2.name)
}
