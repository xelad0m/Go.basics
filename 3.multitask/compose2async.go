package main

import (
	"fmt"
	"sync"
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
	var wg sync.WaitGroup
	wg.Add(2)

	out := make(chan int)

	// первая горутина читает из in1 в out
	go func() {
		defer wg.Done()
		for val := range in1 {
			out <- val
		}
	}()

	// вторая горутина читает из in2 в out
	go func() {
		defer wg.Done()
		for val := range in2 {
			out <- val
		}
	}()

	// ждем, пока исчерпаются оба входных канала,
	// после чего закрываем выходной
	go func() {
		wg.Wait()
		close(out)
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
	// 21 11 22 23 12 24 13 14
	fmt.Println("Took", time.Since(start))
}
