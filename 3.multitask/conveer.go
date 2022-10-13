package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// начало решения

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
		// defer fmt.Println("generate done...")
		defer close(out)
		for {
			word := randomWord(5)
			select {
			case out <- word:
				// fmt.Printf("'%s' generated and sent to chan...\n", word)
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		// defer fmt.Println("unique done...")
		defer close(out)
		for {
			select {
			case word, ok := <-in:
				word_arr := strings.Split(word, "")
				word_dict := make(map[string]bool)
				for _, ch := range word_arr {
					word_dict[ch] = true
				}
				if !ok {
					return
				}

				if len(word_arr) != len(word_dict) {
					// fmt.Printf("'%s' is NOT unique...\n", word)
					continue	// не return!
				}

				select {
				case out <- word:
				case <-cancel:
					return
				}
			}
		}
	}()
	return out
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)

	rev := func(s string) string { // https://github.com/golang/example/blob/master/stringutil/reverse.go
		r := []rune(s)
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	go func() {
		// defer fmt.Println("reverse done...")
		defer close(out)
		for {
			select {
			case word, ok := <-in:
				if !ok {
					return
				}
				rev_word := rev(word)
				// fmt.Printf("'%s' is reversed '%s'...\n", rev_word, word)
				select {
				case out <- fmt.Sprintf("%s -> %s", word, rev_word):
				case <-cancel:
					return
				}
			}
		}
	}()
	return out
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		// defer fmt.Println("merge done...")
		defer close(out)
		for c1 != nil || c2 != nil {
			select {
			case val1, ok := <-c1:
				if ok {
					out <- val1
				} else {
					c1 = nil
				}

			case val2, ok := <-c2:
				if ok {
					out <- val2
				} else {
					c2 = nil
				}

			case <-cancel:
				return
			}
		}
	}()
	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan string, n int) {
	for i := 1; i <= n; i++ {
		// fmt.Printf("Got: %s\n", <-in)
		fmt.Println(<-in)
	}
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
}
