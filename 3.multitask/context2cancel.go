package main

import (
	"context"
	"fmt"
)

// начало решения

// генерит целые числа от start и до бесконечности
func generate(ctx context.Context, start int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := start; ; i++ {
			select {
			case out <- i:
			case <-ctx.Done():		// defer cancel() напишет в ctx.Done()
				return
			}
		}
	}()

	return out
}

// конец решения

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	generated := generate(ctx, 11)
	for num := range generated {
		fmt.Print(num, " ")
		if num > 14 {
			break
		}
	}
	fmt.Println()
}
