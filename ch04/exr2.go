package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var slice = make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		slice = append(slice, rand.Intn(100))
	}

	for _, x := range slice {
		switch {
		case x%2 == 0 && x%3 == 0:
			fmt.Println("Six!")
		case x%2 == 0:
			fmt.Println("Two!")
		case x%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
