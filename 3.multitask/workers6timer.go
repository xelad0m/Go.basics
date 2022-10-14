package main

import (
	"fmt"
	"time"
)

func main() {
	work := func() {
		fmt.Println("work done")
	}

	var eventTime time.Time

	start := time.Now()
	timer := time.NewTimer(100 * time.Millisecond)
	go func() {
		eventTime = <-timer.C
		work()
	}()

	// достаточно времени, чтобы сработал таймер
	time.Sleep(150 * time.Millisecond)
	fmt.Printf("delayed function started after %v\n", eventTime.Sub(start))

	// таймер еще не успел сработать
	if timer.Stop() {
		// fmt.Printf("delayed function canceled after %v\n", time.Since(start))
		fmt.Println("too late to cancel")
	}
}
