package main

import (
	"errors"
	"fmt"
	"time"
)

func withWorkers(n int, fn func()) (handle func() error, wait func()) {
	// канал с токенами
	free := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		free <- struct{}{}
	}

	// выполняет fn, но не более n одновременно
	handle = func() error {
		select {
		case <-free:
			go func() {
				fn()
				free <- struct{}{}
			}()
			return nil
		default:
			return errors.New("busy")
		}
	}

	// ожидает, пока все запущенные fn отработают
	wait = func() {
		for i := 0; i < n; i++ {
			<-free
		}
	}

	return handle, wait
}

func main() {
	work := func() {
		time.Sleep(100 * time.Millisecond)
	}

	handle, wait := withWorkers(2, work)

	start := time.Now()

	err := handle()
	fmt.Println("1st call, error:", err)

	err = handle()
	fmt.Println("2nd call, error:", err)

	err = handle()
	fmt.Println("3rd call, error:", err)

	err = handle()
	fmt.Println("4th call, error:", err)

	wait()

	fmt.Println("4 calls took", time.Since(start))
}
