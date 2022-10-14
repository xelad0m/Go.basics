package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func work() int {
	if rand.Intn(10) < 8 {
		time.Sleep(10 * time.Millisecond)
	} else {
		time.Sleep(200 * time.Millisecond)
	}
	return 42
}

func withTimeout(fn func() int, timeout time.Duration) (int, error) {
	var result int

	done := make(chan struct{})
	go func() {
		result = fn()
		close(done)
	}()

	select {
	case <-done:
		return result, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

func main() {
	for i := 0; i < 10; i++ {
		start := time.Now()
		timeout := 50 * time.Millisecond
		if answer, err := withTimeout(work, timeout); err != nil {
			fmt.Printf("Took longer than %v. Error: %v\n", time.Since(start), err)
		} else {
			fmt.Printf("Took %v. Result: %v\n", time.Since(start), answer)
		}
	}
}
