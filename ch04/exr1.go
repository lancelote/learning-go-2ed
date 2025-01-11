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

	fmt.Println(slice)
}
