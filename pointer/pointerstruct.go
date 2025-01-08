package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Fredik", Age: 20}

	personPointer := &person

	fmt.Println("Sebelum : ", personPointer.Name, personPointer.Age)

	personPointer.Age = 30

	fmt.Println("Setelah : ", person.Age, "Person : ", person.Age)

}
