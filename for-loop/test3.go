package main

import "fmt"

func main() {
	sloki := 0

	for {
		if sloki == 15 {
			break
		}

		sloki++
		fmt.Printf("Kamu Udah minum %d sloki. \n", sloki)
	}

	fmt.Printf("Kamu udah habisin %d sloki, Udah cukup minumnya, kamu udah minum terlalu banyak \n", sloki)
}
