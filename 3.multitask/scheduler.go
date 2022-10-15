package main

import (
	"fmt"
	"time"
)

// начало решения

func schedule(dur time.Duration, fn func()) func() {
	ticker := time.NewTicker(dur)
	cancel_ch := make(chan struct{}, 1)

	cancel_fn := func() {
		ticker.Stop()
		if len(cancel_ch) == 0 {
			cancel_ch <- struct{}{}
		}
	}

	go func() {
		defer close(cancel_ch)
		for {
			select {
			case <-cancel_ch:
				return
			case <-ticker.C:
				fn()
			}
		}
	}()

	return cancel_fn
}

// конец решения

/*Реализуем schedule() поверх тикера:

Создаем тикер с периодом dur, и канал canceled.
Запускаем горутину tick, которая через for+select либо начитывает тики и выполняет fn, либо выходит по canceled.
Возвращаем функцию cancel, которая при вызове закрывает одноименный канал.*/
// func schedule(dur time.Duration, fn func()) func() {
// 	ticker := time.NewTicker(dur)
// 	canceled := make(chan struct{})

// 	tick := func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				fn()
// 			case <-canceled:	// В канал отмены незачем что-то записывать, достаточно его закрыть. После этого всегда будет срабатывать эта ветка
// 				return
// 			}
// 		}
// 	}

// 	cancel := func() {
// 		select {
// 		case <-canceled:
// 			return
// 		default:
// 			ticker.Stop()
// 			close(canceled)
// 		}
// 	}

// 	go tick()
// 	return cancel
// }

func main() {
	work := func() {
		at := time.Now()
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
	}

	cancel := schedule(10*time.Millisecond, work)
	defer func() { cancel(); cancel(); cancel() }()

	time.Sleep(120 * time.Millisecond)
}
