package main

import "fmt"

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

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	person1 := MakePerson("John", "Doe", 42)
	fmt.Println(person1)

	person2 := MakePersonPointer("Alice", "Doe", 42)
	fmt.Println(person2)
}
