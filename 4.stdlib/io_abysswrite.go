package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// начало решения

// AbyssWriter пишет данные в никуда,
// но при этом считает количество записанных байт
type AbyssWriter struct {
	w   io.Writer
	tot int
}

// Writer is the interface that wraps the basic Write method.
func (e *AbyssWriter) Write(p []byte) (int, error) {
	n, err := e.w.Write(p)
	e.tot += n
	if err != nil {
		return n, err
	}
	if n != len(p) {
		return n, io.ErrShortWrite
	}
	return n, nil
}

// Total возвращает общее количество записанных байт
func (e *AbyssWriter) Total() int {
	return e.tot
}

// NewAbyssWriter создает новый AbyssWriter
func NewAbyssWriter() *AbyssWriter {
	return &AbyssWriter{w: ioutil.Discard, tot: 0}
}

// конец НЕКОРРЕКТНОГО решения

/*
// Писать вообще никуда не надо, просто крутить счетчик

// AbyssWriter пишет данные в никуда,
// но при этом считает количество записанных байт
type AbyssWriter struct {
    total int
}

// Write пишет данные в никуда
func (w *AbyssWriter) Write(p []byte) (n int, err error) {
    w.total += len(p)
    return len(p), nil
}

// Total возвращает общее количество записанных байт
func (w *AbyssWriter) Total() int {
    return w.total
}

// NewAbyssWriter создает новый AbyssWriter
func NewAbyssWriter() *AbyssWriter {
    return &AbyssWriter{}
}
*/
func main() {
	r := strings.NewReader("go is awesome")
	w := NewAbyssWriter()
	written, err := io.Copy(w, r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("written %d bytes\n", written)
	fmt.Println(written == int64(w.Total()))
}
