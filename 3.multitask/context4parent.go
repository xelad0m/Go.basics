package main

import (
	"context"
	"fmt"
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
	// работает в течение 100 мс
	work := func() int {
		time.Sleep(100 * time.Millisecond)
		return 42
	}

	// работает в течение 300 мс
	slow := func() int {
		time.Sleep(300 * time.Millisecond)
		return 13
	}

	// возвращает контекст
	// с умолчательным таймаутом 200 мс
	getDefaultCtx := func() (context.Context, context.CancelFunc) {
		const timeout = 200 * time.Millisecond
		return context.WithTimeout(context.Background(), timeout)
	}

	// work с умолчательным контекстом
	{
		// таймаут 200 мс
		ctx, cancel := getDefaultCtx()
		defer cancel()
		// успеет выполниться
		res, err := execute(ctx, work)
		fmt.Println(res, err)
		// 42 <nil>
	}

	// slow с умолчательным контекстом
	{
		// таймаут 200 мс
		ctx, cancel := getDefaultCtx()
		defer cancel()
		// НЕ успеет выполниться
		res, err := execute(ctx, slow)
		fmt.Println(res, err)
		// 0 context deadline exceeded
	}

	// дочерний контекст с жестким таймаутом
	{
		// родительский контекст с таймаутом 200 мс
		parentCtx, cancel := getDefaultCtx()
		defer cancel()

		// дочерний контекст с таймаутом 50 мс
		childCtx, cancel := context.WithTimeout(parentCtx, 50*time.Millisecond)
		defer cancel()

		// теперь work НЕ успеет выполниться
		res, err := execute(childCtx, work)
		fmt.Println(res, err)
		// 0 context deadline exceeded
	}

	// дочерний контекст с мягким таймаутом
	{
		// родительский контекст с таймаутом 200 мс
		parentCtx, cancel := getDefaultCtx()
		defer cancel()

		// дочерний контекст с таймаутом 500 мс
		childCtx, cancel := context.WithTimeout(parentCtx, 500*time.Millisecond)
		defer cancel()

		// slow все равно НЕ успеет выполниться
		res, err := execute(childCtx, slow)
		fmt.Println(res, err)
		// 0 context deadline exceeded
	}
}
