package main

import (
	"fmt"
	"time"
)

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы

// через отдельный канал для индексов
// func gather(funcs []func() any) []any {
// 	// начало решения
// 	done := make(chan any, len(funcs))
// 	idxs := make(chan int, len(funcs))	// заведем канал для сохранения порядка

// 	// выполните все переданные функции,
// 	for i := 0; i < len(funcs); i++ {
// 		go func(i int) {
// 			done <- funcs[i]()
// 			idxs <- i
// 		}(i)
// 	}
// 	// соберите результаты в срез
// 	res := make([]any, len(funcs))
// 	for range funcs {
// 		res[<-idxs] = <-done
// 	}
// 	// и верните его
// 	return res
// 	// конец решения
// }

// через "канал завершения"
// func gather(funcs []func() any) []any {
//     done := make(chan struct{})

//     // заранее создаем срез размером по количеству переданных функций
//     results := make([]any, len(funcs))
//     // запускаем функции в отдельных горутинах
//     for idx, fn := range funcs {
//         idx, fn := idx, fn
//         go func() {
//             // записываем результат i-й функции в i-ю позицию среза
//             results[idx] = fn()
//             done <- struct{}{}
//         }()
//     }
//     // дожидаемся, пока отработают все горутины
//     for i := 0; i < len(funcs); i++ {
//         <-done
//     }
//     return results
// }

// индексы во вспомогательной структуре
func gather(funcs []func() any) []any {
	// результат вызова i-й функции из переданных
	type result struct {
		idx int
		val any
	}

	n := len(funcs)

	// запускаем по горутине на каждую функцию
	// и складываем результат в канал
	ready := make(chan result, n)
	for idx, fn := range funcs {
		idx := idx
		fn := fn
		go func() {
			ready <- result{idx, fn()}
		}()
	}

	// начитываем результаты из канала
	// и готовим итоговый срез с результатами
	results := make([]any, n)
	for i := 0; i < n; i++ {
		res := <-ready
		results[res.idx] = res.val
	}
	return results
}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}
	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
