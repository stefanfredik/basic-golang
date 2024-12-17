package main

import "fmt"

type PersegiPanjang struct {
	Panjang float32
	Lebar   float32
}

func (p *PersegiPanjang) Luas() float32 {
	return p.Panjang * p.Lebar
}

func main() {
	persegiPanjang1 := PersegiPanjang{
		Panjang: 10,
		Lebar:   5,
	}

	persegiPanjang2 := PersegiPanjang{
		Panjang: 15,
		Lebar:   20,
	}

	fmt.Println("Luas persegi panjang 1 : ", persegiPanjang1.Luas())
	fmt.Println("Luas persegi panjang 2 : ", persegiPanjang2.Luas())
}
