package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	rand.Seed(0)

	in := generate(100, 3)
	counter := map[string]int{}

	var wg sync.WaitGroup
	wg.Add(2)

	var lock sync.Mutex
	go count(&wg, &lock, in, counter)
	go count(&wg, &lock, in, counter)

	wg.Wait()

	fmt.Println(counter)
}

// count calculates word frequencies
func count(wg *sync.WaitGroup, lock *sync.Mutex, in <-chan string, counter map[string]int) {
	defer wg.Done()
	for word := range in {
		lock.Lock()
		counter[word]++
		lock.Unlock()
	}
}

// generate sends nWords random words of length wordLen
// to the out channel
func generate(nWords, wordLen int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for ; nWords > 0; nWords-- {
			out <- randomWord(wordLen)
		}
	}()
	return out
}

// randomWord returns a random word of length n
func randomWord(n int) string {
	const vowels = "eaiou"
	const consonants = "rtnslcdpm"
	chars := make([]byte, n)
	for i := 0; i < n; i += 2 {
		chars[i] = consonants[rand.Intn(len(consonants))]
	}
	for i := 1; i < n; i += 2 {
		chars[i] = vowels[rand.Intn(len(vowels))]
	}
	return string(chars)
}
