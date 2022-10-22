// БРЕД
package main

import (
	"fmt"
	"strings"
	"io"
)

// TokenReader считывает токены из источника
type TokenReader interface {
	// ReadToken считывает очередной токен
	// Если токенов больше нет, возвращает ошибку io.EOF
	ReadToken() (string, error)
}

type Reader struct {
	src []string
	cur int
}

func (r Reader) ReadToken() (string, error) {
	if r.cur >= len(r.src) {
		return "", io.EOF
	}
	s := r.src[r.cur]
	r.cur = r.cur + 1
	return s, nil
}

func NewWordReader(src string) Reader {
	return Reader{src: strings.Split(src, " "), cur: 0}
}

// TokenWriter записывает токены в приемник
type TokenWriter interface {
	// WriteToken записывает очередной токен
	WriteToken(s string) error
}

type Writer struct {
	dst []string
	total int
}

func (w Writer) WriteToken(s string) error {
	w.dst = append(w.dst, s)
	w.total++
	return nil
}

func (w Writer) Words() []string {
	return w.dst
}

func NewWordWriter() Writer {
	return Writer{total: 0}
}

// начало решения

// FilterTokens читает все токены из src и записывает в dst тех,
// кто проходит проверку predicate
// начало решения

// FilterTokens читает все токены из src и записывает в dst тех,
// кто проходит проверку predicate
func FilterTokens(dst TokenWriter, src TokenReader, predicate func(s string) bool) (int, error) {
	tot := 0

	for {
		s, err := src.ReadToken()
        if err == io.EOF {
            return tot, nil
        }
        
        if err != nil {
            return tot, err
        }			
        
		if predicate(s) {
            werr := dst.WriteToken(s)
            if werr == nil {
                tot++
            } else {
                return tot, werr
            }
		}       
	}
	
}

// конец решения

/*Читаем токены из src, пока не закончатся.
Если токен не проходит predicate, игнорируем его.
Если проходит — пишем в dst.
Попутно считаем количество записанных токенов.
Верное решение #727679651

func FilterTokens(dst TokenWriter, src TokenReader, predicate func(s string) bool) (int, error) {
    total := 0
    for {
        token, err := src.ReadToken()
        if err == io.EOF {
            break
        }
        if err != nil {
            return total, err
        }

        if !predicate(token) {
            continue
        }

        err = dst.WriteToken(token)
        if err != nil {
            return total, err
        }
        total++
    }
    return total, nil
}*/

func main() {
	// Для проверки придется создать конкретные типы,
	// которые реализуют интерфейсы TokenReader и TokenWriter.

	// Ниже для примера используются NewWordReader и NewWordWriter,
	// но вы можете сделать любые на свое усмотрение.

	r := NewWordReader("go is awesome")
	w := NewWordWriter()
	predicate := func(s string) bool {
		return s != "is"
	}
	n, err := FilterTokens(w, r, predicate)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d tokens: %v\n", n, w.Words())
	// 2 tokens: [go awesome]
}
