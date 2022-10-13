package main

import (
	"fmt"
)

func rangeGen(start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		for i := start; i < stop; i++ {
			out <- i							// горутина виснет тут если из канала перестают читать
		}										// нужен сигнал отмены
		close(out)
	}()
	return out
}

func main() {
	generated := rangeGen(41, 46)
	for val := range generated {
		fmt.Println(val)
		if val == 42 {
			break
		}
	}
}
