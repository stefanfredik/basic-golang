package main

import "fmt"

// Define a struct Person with fields Name and Age
type Person struct {
	Name string
	Age  int
}

// Method to set the Name field of Person
func (p *Person) SetName(name string) {
	p.Name = name
}

// Method to get the Name field of Person
func (p *Person) GetName() string {
	return p.Name
}

func main() {
	// Create an instance of Person
	person := &Person{}
	// Set the name of the person
	person.SetName("Fredik Stefan")
	// Print the name of the person
	fmt.Println("Hallo, my name is  : ", person.GetName())
}
