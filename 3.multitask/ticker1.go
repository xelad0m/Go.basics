package main

import (
	"fmt"
	"time"
)

func main() {
	work := func(at time.Time) {
		fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
	}

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		for {
			at := <-ticker.C
			work(at)
		}
	}()

	// хватит на 5 тиков
	time.Sleep(260 * time.Millisecond)
}