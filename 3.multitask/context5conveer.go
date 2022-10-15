package main

import (
	"context"
	"fmt"
	"strings"
	"unicode"
)

// информация о количестве цифр в каждом слове
type counter map[string]int

// слово и количество цифр в нем
type pair struct {
	word  string
	count int
}

// начало решения

// считает количество цифр в словах
func countDigitsInWords(ctx context.Context, words []string) counter {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	pending := submitWords(ctx, words)
	counted := countWords(ctx, pending)
	return fillStats(ctx, counted)
}

// отправляет слова на подсчет
func submitWords(ctx context.Context, words []string) <-chan string {
	out := make(chan string, 10)

	go func() {
		defer close(out)
		for _, word := range words {
			select {
			case out <- word:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

// считает цифры в словах
func countWords(ctx context.Context, in <-chan string) <-chan pair {
	out := make(chan pair, 10)

	go func() {
		defer close(out)
		for word := range in {
			count := countDigits(word)
			select {
			case out <- pair{word, count}:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

// готовит итоговую статистику
func fillStats(ctx context.Context, in <-chan pair) counter {
	/*Но в данном случае контекст для отмены не нужен, т.к. функция не пишет в каналы и без 
	горутинг, когда закончится канал, закончится и цикл*/
	stats := counter{}

	for p := range in {
		stats[p.word] = p.count
		select {
        case <-ctx.Done():
            return stats
        default:
            stats[p.word] = p.count
        }
		// ну или так тоже сойдет...
		// if len(ctx.Done()) > 0 {
		// 	return stats
		// }
	}

	return stats
}

// конец решения

// считает количество цифр в слове
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	words := strings.Fields(phrase)

	ctx := context.Background()
	stats := countDigitsInWords(ctx, words)
	fmt.Println(stats)
}
