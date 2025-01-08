package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) Birthday() {
	p.Age++
}

func main() {
	person := NewPerson("Alice", 30)
	fmt.Println("Before birthday:", person.Name, person.Age)
	person.Birthday()
	fmt.Println("After birthday:", person.Name, person.Age)
}
