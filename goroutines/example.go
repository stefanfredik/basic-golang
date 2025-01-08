package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readFile(filepath string, resultChan chan<- string, errorChan chan<- error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		errorChan <- err
		return
	}
	resultChan <- string(data)
}

func main() {
	filePath := "goroutines1.go"
	resultChan := make(chan string)
	errorChan := make(chan error)

	fmt.Println("Proses membaca file : ")

	go readFile(filePath, resultChan, errorChan)

	select {
	case result := <-resultChan:
		fmt.Println("File content:")
		fmt.Println(result)
	case err := <-errorChan:
		log.Fatalf("Failed to read file: %v", err)
	}

}
