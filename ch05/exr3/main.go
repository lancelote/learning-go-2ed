package main

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(line string) string {
		return prefix + " " + line
	}
}

func main() {
	helloPrefix := prefixer("Hello")

	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Alice"))
}
