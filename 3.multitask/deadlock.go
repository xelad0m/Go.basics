package main

import (
	"fmt"
)

func work(done chan struct{}, out chan int) {
    out <- 42
    done <- struct{}{}
}

func main() {
    out := make(chan int)
    done := make(chan struct{})

    go work(done, out)    // (1)

    <-done                // (2)

    fmt.Println(<-out)    // (3)
}