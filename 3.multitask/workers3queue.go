package main

import (
	"errors"
	"fmt"
)

var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

/*
// начало решения

// Queue - FIFO-очередь на n элементов
// type queue interface {
// 	MakeQueue() Queue
// 	Get() any
// 	Put() any
// }

type Queue struct {
	size int
	fifo chan int
}

// Get возвращает очередной элемент.
// Если элементов нет и block = false -
// возвращает ошибку.
func (q Queue) Get(block bool) (int, error) {
	if len(q.fifo) != 0 {
		return <-q.fifo, nil
	} else {
		if block {
			return <-q.fifo, nil
		} else {
			return 0, ErrEmpty
		}
	}
}

// Put помещает элемент в очередь.
// Если очередь заполнения и block = false -
// возвращает ошибку.
func (q Queue) Put(val int, block bool) error {
	if len(q.fifo) < q.size {
		q.fifo <- val
		return nil
	} else {
		if block {
			q.fifo <- val
			return nil
		} else {
			return ErrFull
		}
	}
}

// MakeQueue создает новую очередь (конструктор)
func MakeQueue(n int) Queue {
	return Queue{
		size: n,
		fifo: make(chan int, n),
	}
}

// конец решения
*/

/*
// Queue - FIFO-очередь на n элементов
type Queue struct {
    values chan int
}

// Get возвращает очередной элемент.
// Если элементов нет и block = false -
// возвращает ошибку.
func (q Queue) Get(block bool) (int, error) {
    if block {
        return <-q.values, nil
    }
    select {
    case val := <-q.values:
        return val, nil
    default:
        return 0, ErrEmpty
    }
}

// Put помещает элемент в очередь.
// Если очередь заполнения и block = false -
// возвращает ошибку.
func (q Queue) Put(val int, block bool) error {
    if block {
        q.values <- val
        return nil
    }
    select {
    case q.values <- val:
        return nil
    default:
        return ErrFull
    }
}

// MakeQueue создает новую очередь
func MakeQueue(n int) Queue {
    ch := make(chan int, n)
    return Queue{ch}
}
*/


// начало решения

// Queue - FIFO-очередь на n элементов
type Queue chan int

// Get возвращает очередной элемент.
// Если элементов нет и block = false -
// возвращает ошибку.
func (q Queue) Get(block bool) (int, error) {
    if len(q) > 0 || block {
        return <-q, nil
    }
    
    return 0, ErrEmpty
}

// Put помещает элемент в очередь.
// Если очередь заполнения и block = false -
// возвращает ошибку.
func (q Queue) Put(val int, block bool) error {
    if len(q) < cap(q) || block {
        q <- val
        return nil
    }
    
    return ErrFull
}

// MakeQueue создает новую очередь
func MakeQueue(n int) Queue {
    return make(Queue, n)
}

// конец решения

func main() {
	q := MakeQueue(2)

	err := q.Put(1, false)
	fmt.Println("put 1:", err)

	err = q.Put(2, false)
	fmt.Println("put 2:", err)

	err = q.Put(3, false)
	fmt.Println("put 3:", err)

	res, err := q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)
}
