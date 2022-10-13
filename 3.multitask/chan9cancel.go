package main

import (
	"fmt"
)

func rangeGen(cancel <-chan struct{}, start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			select {			// как switch, но для каналов, выбирает любой (случайный) канал и тех, что открыты
			case out <- i:		// пока cancel открыт, но пустой, будет идти по этой ветке
			case <-cancel:		// когда из cancel что-то прочитается, или он будет закрыт
				return		
			}
		}
	}()
	return out
}

func main() {

	cancel := make(chan struct{})				// канал отмены
	defer close(cancel)							// отработает при выходе из области видимости (в т.ч. по ошибке/исключению)

	generated := rangeGen(cancel, 41, 46)
	for val := range generated {
		fmt.Println(val)
		if val == 42 {							// в данном случае, явно в канал отмены мы не пишем,
			break								// а горутина подвиснет (в out и cancel ничего нет), до отработки defer
		}
	}											// при выходе из main() отработает close(cancel) и горутина пройдет по его ветке select
}

// на практике канал отмены часто обозначают done (как и канал завершения, который висит на блокировке без select)
// канал отмены кажется более гибким подходом