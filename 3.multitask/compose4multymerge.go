package main

import (
	"fmt"
	// "sync"
	"time"
)

// rangeGen отправляет в канал числа от start до stop-1
func rangeGen(start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			time.Sleep(50 * time.Millisecond)
			out <- i
		}

	}()
	return out
}

// начало решения

// merge выбирает числа из входных каналов и отправляет в выходной
// func merge(channels ...<-chan int) <-chan int {
// 	// объедините все исходные каналы в один выходной
// 	// последовательное объединение НЕ подходит
// 	out := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(len(channels))
// 	for _, c := range channels {
// 		go func(c <-chan int) {
// 			for v := range c {
// 				out <- v
// 			}
// 			wg.Done()
// 		}(c)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
// 	return out
// }

// конец решения

// РЕКУРСИВНЫЙ ВАРИАНТ
// https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6
// начало решения
// func mergeTwo(ch1, ch2 <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for ch1 != nil || ch2 != nil {
// 			select {
// 			case val, ok := <-ch1:
// 				if !ok {
// 					ch1 = nil
// 				} else {
// 					out <- val
// 				}
// 			case val, ok := <-ch2:
// 				if !ok {
// 					ch2 = nil
// 				} else {
// 					out <- val
// 				}
// 			}
// 		}
// 	}()
// 	return out
// }

// func merge(channels ...<-chan int) <-chan int {
// 	switch len(channels) {
// 	case 0:
// 		ch := make(chan int)
// 		close(ch)
// 		return ch
// 	case 1:
// 		return channels[0]
// 	default:
// 		half := len(channels) / 2
// 		return mergeTwo(merge(channels[:half]...), merge(channels[half:]...))
// 	}
// }

// конец решения

// КАНАЛ ЗАВЕРШЕНИЯ
// func merge(channels ...<-chan int) <-chan int {
//     done := make(chan struct{})
//     out := make(chan int)

//     for _, channel := range channels {
//         go func(channel <-chan int) {
//             for value := range channel {
//                 out <- value
//             }
//             done <- struct{}{}
//         }(channel)
//     }

//     go func() {
//         for i := 0; i < len(channels); i++ {
//             <-done
//         }
//         close(out)
//     }()

//     return out
// }	

// начало решения

// SELECT
func merge(channels ...<-chan int) <-chan int {
    out := make(chan int)
	go func() {
		defer close(out)
		for channels != nil {
			for _, channelx := range channels {
				select {
				case val, ok := <-channelx:
					if ok {
						out <- val
					} else {
						channels = nil

					}
				}
			}
		}
	}()
	return out
}

func main() {
	in1 := rangeGen(11, 15)
	in2 := rangeGen(21, 25)
	in3 := rangeGen(31, 35)

	start := time.Now()
	merged := merge(in1, in2, in3)
	for val := range merged {
		fmt.Print(val, " ")
	}
	fmt.Println()
	fmt.Println("Took", time.Since(start))
}
