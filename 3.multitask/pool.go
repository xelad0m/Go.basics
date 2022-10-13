package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// say prints the phrase on behalf of the worker
// from the pool
func say(pool chan<- int, id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
	// возвращаем идентификатор в пул
	pool <- id
}

func main() {
	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}

	// пул идентификаторов для 2 горутин (или по количеству ядер CPU)
	pool := make(chan int, 2)
	pool <- 1
	pool <- 2

	for _, phrase := range phrases {
		// получаем идентификатор из пула,
		// если есть свободные
		id := <-pool
		go say(pool, id, phrase)
	}

	// дожидаемся, пока все горутины закончат работу
	// (то есть все идентификаторы вернутся в пул)
	<-pool
	<-pool
}
