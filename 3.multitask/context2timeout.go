package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// выполняет функцию fn с учетом контекста ctx
func execute(ctx context.Context, fn func() int) (int, error) {
	ch := make(chan int, 1)

	go func() {
		ch <- fn()
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	// работает в течение 100 мс
	work := func() int {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("work done")
		return 42
	}

	// возвращает случайный агрумент из переданных
	randomChoice := func(arg ...int) int {
		i := rand.Intn(len(arg))
		return arg[i]
	}

	// случайный таймаут - 50 мс либо 150 мс
	timeout := time.Duration(randomChoice(50, 150)) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := execute(ctx, work)
	fmt.Println(res, err)
}
