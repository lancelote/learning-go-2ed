package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancelFunc()

	iterations, err := generate1234(ctx)
	fmt.Println(iterations, err)
}

func generate1234(ctx context.Context) (int, error) {
	var n int
	var i int

	for {
		select {
		case <-ctx.Done():
			return i, ctx.Err()
		default:
			n = rand.Intn(100_000_000)
			if n == 1234 {
				return i, nil
			}
			i++
		}
	}
}
