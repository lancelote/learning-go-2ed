package main

import "fmt"

func main() {
	var message = "Hi 👩 and 👨"
	var runes = []rune(message)
	fmt.Println(string(runes[3]))
}
