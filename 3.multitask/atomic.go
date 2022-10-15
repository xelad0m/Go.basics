package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var total int32

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				atomic.AddInt32(&total, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("total", atomic.LoadInt32(&total))
}
