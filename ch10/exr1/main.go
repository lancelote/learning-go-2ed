package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed english_rights.txt
var english string

//go:embed dutch_rights.txt
var dutch string

func main() {
	if len(os.Args) == 1 {
		fmt.Println("language is required")
		return
	}

	language := os.Args[1]

	switch language {
	case "english":
		fmt.Println(english)
	case "dutch":
		fmt.Println(dutch)
	default:
		fmt.Println("unknown language")
	}
}
