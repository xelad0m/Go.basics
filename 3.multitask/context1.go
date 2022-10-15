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

	// ждет 50 мс, после этого
	// с вероятностью 50% отменяет работу
	maybeCancel := func(cancel func()) {
		time.Sleep(50 * time.Millisecond)
		if rand.Float32() < 0.5 {
			cancel()
		}
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go maybeCancel(cancel)

	res, err := execute(ctx, work)
	fmt.Println(res, err)
}
