package main

import "fmt"

func main() {
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	var firstTwo = greetings[:2]
	var middleThree = greetings[1:4]
	var lastTwo = greetings[3:]

	fmt.Println(greetings)
	fmt.Println(firstTwo)
	fmt.Println(middleThree)
	fmt.Println(lastTwo)
}
