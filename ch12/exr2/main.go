package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		defer close(ch1)

		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)

		for i := 10; i < 20; i++ {
			ch2 <- i
		}
	}()

	for count := 0; count < 2; {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				count++
				continue
			}
			fmt.Printf("from goroutine 1: %d\n", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				count++
				continue
			}
			fmt.Printf("from goroutine 2: %d\n", v)
		}
	}
}
