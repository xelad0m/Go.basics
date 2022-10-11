package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"sync"
)

/*
// Работает, но смешивается логика (say) и многопоточность (sync). 
// В таком виде say не сможет использоваться в однопоточном режиме.

func main() {
    var wg sync.WaitGroup

    wg.Add(1)
    go say(&wg, 1, "go is awesome")

    wg.Add(1)
    go say(&wg, 2, "cats are cute")

    wg.Wait()
}

func say(wg *sync.WaitGroup, id int, phrase string) {
    for _, word := range strings.Fields(phrase) {
        fmt.Printf("Worker #%d says: %s...\n", id, word)
        dur := time.Duration(rand.Intn(100)) * time.Millisecond
        time.Sleep(dur)
    }
    wg.Done()
}

*/

func main() {
    var wg sync.WaitGroup
    wg.Add(2)	// устанавливает счетчик горутин в группе

    go func() {
        defer wg.Done()	// уменьшает счетчик горутин (даже если исключение)
        say(1, "go is awesome")
    }()

    go func() {
        defer wg.Done()
        say(2, "cats are cute")
    }()

    wg.Wait()
}

func say(id int, phrase string) {
    for _, word := range strings.Fields(phrase) {
        fmt.Printf("Worker #%d says: %s...\n", id, word)
        dur := time.Duration(rand.Intn(100)) * time.Millisecond
        time.Sleep(dur)
    }
}