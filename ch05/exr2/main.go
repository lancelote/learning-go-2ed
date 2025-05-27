package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func getFile(name string) (*os.File, func(), error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}

	return f, func() {
		f.Close()
	}, nil
}

func fileLen(name string) (int, error) {
	f, closer, err := getFile(name)
	if err != nil {
		return 0, err
	}
	defer closer()

	data := make([]byte, 2048)
	counter := 0
	for {
		count, err := f.Read(data)
		counter += count
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			return counter, nil
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	count, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}
