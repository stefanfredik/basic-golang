package main

import "fmt"

func main() {

	var jenisMinuman []string = []string{"tuak", "bir", "arak", "whisky", "vodka"}

	for _, value := range jenisMinuman {

		if value == "arak" {
			fmt.Printf("%s ga bagus untuk diminum \n", value)
			continue
		}
		fmt.Printf("Jenis minuman : %s. \n", value)
	}
}
