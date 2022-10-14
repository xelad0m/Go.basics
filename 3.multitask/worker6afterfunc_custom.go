package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения аналог time.AfterFunc
/*в лоб*/
func delay(dur time.Duration, fn func()) func() {
	chancel_ch := make(chan struct{}, 1)

	cancel_fn := func() {
		if len(chancel_ch) == 0 {
			chancel_ch <- struct{}{}
		}
	}

	go func() {
		time.Sleep(dur)
		select {
		case <-chancel_ch:
			return
		default:
			fn()
		}
	}()

	return cancel_fn
}

// конец решения

/*Можно реализовать delay() с помощью таймера. Но я решил обойтись только каналами.

Идея следующая:

Создаем два канала – done для успешного срабатывания и canceled для отмены.
Запускаем горутину wait(), которая ждет dur времени, после чего сигнализирует done.
Запускаем горутину tick(), которая через select выбирает первый сработавший канал (done или canceled). Если done – выполняет отложенную функцию fn.
Возвращаем функцию cancel, которая при вызове закрывает одноименный канал.
Предотвращаем повторное закрытие канала отмены через select.*/
// func delay(dur time.Duration, fn func()) func() {
//     done := make(chan struct{})
//     canceled := make(chan struct{})

//     // ждет dur времени, после чего сигнализирует done
//     wait := func() {
//         time.Sleep(dur)
//         close(done)
//     }

//     // выполняет fn по готовности, либо выходит по отмене
//     tick := func() {
//         select {
//         case <-done:
//             fn()
//         case <-canceled:
//             return
//         }
//     }

//     // отменяет запуск
//     cancel := func() {
//         select {
//         case <-canceled:
//         default:
//             close(canceled)
//         }
//     }

//     go tick()
//     go wait()

//     return cancel
// }

/*Реализация delay() на базе таймера:

Создаем таймер и канал отмены.
Запускаем горутину, которая через select выбирает из канала таймера или отмены, смотря что сработает раньше. Если канал таймера — выполняет отложенную функцию fn.
Возвращаем функцию cancel, которая при вызове останавливает таймер и закрывает одноименный канал.*/

// func delay(duration time.Duration, fn func()) func() {
//     canceled := make(chan struct{})

//     timer := time.NewTimer(duration)
//     go func() {
//         select {
//         case <-timer.C:
//             fn()
//         case <-canceled:
//         }
//     }()

//     cancel := func() {
//         if !timer.Stop() {
//             return
//         }
//         close(canceled)
//     }
//     return cancel
// }

func main() {
	rand.Seed(time.Now().Unix())

	work := func() {
		fmt.Println("work done")
	}

	for i := 0; i < 25; i++ {
		cancel := delay(100*time.Millisecond, work)

		time.Sleep(10 * time.Millisecond)
		if rand.Float32() < 0.5 {
			cancel()
			cancel()	
			cancel()	// множественный вызов не должен ломать логику отмены
			fmt.Println("delayed function canceled")
		}
	}
	time.Sleep(100 * time.Millisecond)
}
