package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 20)

	var wgWriter sync.WaitGroup
	wgWriter.Add(2)

	var wgReader sync.WaitGroup
	wgReader.Add(1)

	go func() {
		defer wgWriter.Done()

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wgWriter.Done()

		for i := 10; i < 20; i++ {
			ch <- i
		}
	}()

	go func() {
		wgWriter.Wait()
		close(ch)
	}()

	go func() {
		defer wgReader.Done()

		for v := range ch {
			fmt.Println(v)
		}
	}()

	wgReader.Wait()
}
