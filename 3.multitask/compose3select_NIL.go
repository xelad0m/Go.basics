package main

import (
	"fmt"
	"time"
)

func rangeGen(start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		for i := start; i < stop; i++ {
			time.Sleep(50 * time.Millisecond)
			out <- i
		}
		close(out)
	}()
	return out
}

func merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for in1 != nil || in2 != nil {
			select {
			case val1, ok := <-in1:
				if ok {
					out <- val1
				} else {
					in1 = nil
				}

			case val2, ok := <-in2:
				if ok {
					out <- val2
				} else {
					in2 = nil
				}
			}
		}
	}()
	return out
}

func main() {
	in1 := rangeGen(11, 15)
	in2 := rangeGen(21, 25)

	start := time.Now()
	merged := merge(in1, in2)
	for val := range merged {
		fmt.Print(val, " ")
	}
	fmt.Println()
	// 21 11 22 12 23 13 24 14
	fmt.Println("Took", time.Since(start))
}
