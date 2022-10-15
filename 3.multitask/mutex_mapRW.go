package main

import (
	"fmt"
	"sync"
)

// начало решения
type Counter struct {
	lock sync.RWMutex		// имеет методы блокировки для чтения или чтения/записи
	dict map [string] int
}

func (c *Counter) Increment(str string) {
	c.lock.Lock()			// тут блокируется и чтение и запись
	defer c.lock.Unlock()
	c.dict[str]++
}

func (c *Counter) Value(str string) int {
	c.lock.RLock()			// тут блокировка разрешает параллельное чтение
	defer c.lock.RUnlock()
	return c.dict[str]
	// val := c.dict[str]
	// return val
}

func (c *Counter) Range(fn func(key string, val int)) {
	c.lock.RLock()			// тут блокировка разрешает параллельное чтение	
	defer c.lock.RUnlock()	// т.к. чтение в range... то нужно защитить весь цикл
	for key, val := range c.dict {
		fn(key, val)
	}
}

func NewCounter() *Counter {
	return &Counter{
		lock: sync.RWMutex{}, 
		dict: make(map[string]int),
	}
}

// конец решения

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(3)

	increment := func(key string, val int) {
		defer wg.Done()
		for ; val > 0; val-- {
			counter.Increment(key)
		}
	}

	go increment("one", 100)
	go increment("two", 200)
	go increment("three", 300)

	wg.Wait()

	fmt.Println("two:", counter.Value("two"))

	fmt.Print("{ ")
	counter.Range(func(key string, val int) {
		fmt.Printf("%s:%d ", key, val)
	})
	fmt.Println("}")
}

