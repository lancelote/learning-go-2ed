package main

import "fmt"

func UpdateSlice(slice []string, last string) {
	length := len(slice)
	slice[length-1] = last
	fmt.Println("in UpdateSlice:", slice)
}

func GrowSlice(slice []string, last string) {
	slice = append(slice, last)
	fmt.Println("in GrowSlice:", slice)
}

func main() {
	slice := []string{"a", "b", "c"}

	UpdateSlice(slice, "d")
	fmt.Println("in main after UpdateSlice:", slice)

	GrowSlice(slice, "e")
	fmt.Println("in main after GrowSlice:", slice)
}
