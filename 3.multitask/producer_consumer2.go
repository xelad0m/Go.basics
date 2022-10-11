package main

import (
	"fmt"
	"strings"
	"unicode"
)

// nextFunc возвращает следующее слово из генератора
type nextFunc func() string

// counter хранит количество цифр в каждом слове.
// ключ карты - слово, а значение - количество цифр в слове.
type counter map[string]int

// pair хранит слово и количество цифр в нем
type pair struct {
	word  string
	count int
}

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWords(next nextFunc) counter {
	pending := make(chan string)
	counted := make(chan pair)

	// начало решения

	// отправляет слова на подсчет
	go func() {
		// Пройдите по словам и отправьте их
		for {
			word := next()
			// в канал pending
			pending <- word
			if word == "" {
				break
			}
		}
	}()

	// считает цифры в словах
	go func() {
		// Считайте слова из канала pending,
		for {
			word := <-pending
			// посчитайте количество цифр в каждом,
			p := pair{word, countDigits(word)}
			// и запишите его в канал counted
			counted <- p
			if word == "" {
				break
			}
		}
	}()

	// Считайте значения из канала counted
	// и заполните stats.
	stats := counter{}
	for {
		// var p pair
		p := <-counted
		if p.word == "" {
			break
		}
		stats[p.word] = p.count
	}

	// В результате stats должна содержать слова
	// и количество цифр в каждом.

	// конец решения

	return stats
}

// countDigits возвращает количество цифр в строке
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats печатает слова и количество цифр в каждом
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// wordGenerator возвращает генератор, который выдает слова из фразы
func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}
