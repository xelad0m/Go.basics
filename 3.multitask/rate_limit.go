package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCanceled error = errors.New("canceled")

// начало решения (куча лишнего, ниже правильно)
/*
func withRateLimit(limit int, fn func()) (handle func() error, cancel func()) {
	dur := time.Duration(1000/limit) * time.Millisecond

	ticker := time.NewTicker(dur)
	cancel_ch := make(chan struct{})
	done := make(chan struct{})

	cancel_fn := func() {
		select {
		case <-cancel_ch: // когда канал закрыт, эта ветка всегда исполняется (уот так уот...)
			return
		default:
			ticker.Stop()
			close(cancel_ch)
		}
	}

	tick := func() {
		select {
		case <-cancel_ch:
			return
		case <-ticker.C:
			go fn()
			done <- struct{}{}
		}
	}

	handle_fn := func() error {
		go tick()
		select {
		case <-cancel_ch:
			return ErrCanceled
		case <-done:
			return nil
		}
	}

	return handle_fn, cancel_fn
}
*/
// конец решения

func withRateLimit(limit int, fn func()) (handle func() error, cancel func()) {
	ticker := time.NewTicker(time.Second / time.Duration(limit))
	canceled := make(chan struct{})

	handle = func() error {
		select {
		case <-ticker.C:
			go fn()
			return nil
		case <-canceled:
			return ErrCanceled
		}
	}

	cancel = func() {
		select {
		case <-canceled:
		default:
			ticker.Stop()
			close(canceled)
		}
	}

	return handle, cancel
}

func main() {
	work := func() {
		fmt.Print(".")
		time.Sleep(1000 * time.Millisecond) // не важно, сколько выполняется, запуск по расписанию в параллель
	}

	handle, cancel := withRateLimit(5, work)
	defer cancel()

	start := time.Now()
	const n = 10
	for i := 0; i < n; i++ {
		handle()
	}
	fmt.Println()
	fmt.Printf("%d queries took %v\n", n, time.Since(start))
}
