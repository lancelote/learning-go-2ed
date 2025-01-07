package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	person1 := Employee{"John", "Doe", 42}
	person2 := Employee{
		firstName: "Foo",
		lastName:  "Bar",
		id:        34,
	}
	var person3 Employee
	person3.firstName = "Baz"
	person3.lastName = "Bug"
	person3.id = 43

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
}
