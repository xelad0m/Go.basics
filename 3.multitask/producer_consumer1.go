package main

import (
	"fmt"
	"strings"
	"unicode"
)

// counter хранит количество цифр в каждом слове.
// ключ карты - слово, а значение - количество цифр в слове.
type counter map[string]int

// countDigitsInWords считает количество цифр в словах фразы
func countDigitsInWords(phrase string) counter {
	words := strings.Fields(phrase)
	counted := make(chan int)
	// начало решения

	go func() {
		// Пройдите по словам,
		for _, word := range words {
			// посчитайте количество цифр в каждом,
			count := countDigits(word)
			// и запишите его в канал counted
			counted <- count
		}
	}()

	// Считайте значения из канала counted
	// и заполните stats.
	stats := counter{}
	for _, word := range words {
		// посчитайте количество цифр в каждом,
		cnt := <-counted
		stats[word] = cnt
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

func main() {
	phrase := "0ne 1wo thr33 4068"
	stats := countDigitsInWords(phrase)
	printStats(stats)
}
