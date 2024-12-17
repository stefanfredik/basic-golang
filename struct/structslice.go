// struct bisa

package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func main() {
	products := []Product{
		{"Laptop", 1000},
		{"Smartphone", 100},
	}

	for _, value := range products {
		fmt.Printf("Products : %s , Price : Rupiah %.2f \n", value.Name, value.Price)
	}
}
