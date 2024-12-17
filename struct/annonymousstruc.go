package main

import "fmt"

func main() {
	p := struct {
		Name string
		Age  int
	}{
		Name: "Fredik Stefan",
		Age:  20,
	}

	fmt.Printf("Hallo, nama saya %s, usia saya %d tahun. \n", p.Name, p.Age)

}
