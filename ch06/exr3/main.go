package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	people := make([]Person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		people = append(people, MakePerson("John", "Doe", 42))
	}

	for _, p := range people {
		fmt.Println(p.FirstName)
	}
}
