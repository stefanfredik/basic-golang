package main

import (
	"encoding/json"
	"fmt"
)

var buah [5]string

func main() {
	buah = [5]string{"mangga", "nangka", "rambutan", "jeruk", "durian"}

	for i, data := range buah {
		fmt.Printf("Fruit %d: %s\n", i+1, data)
	}

	jsonData, err := json.Marshal(buah)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

}
