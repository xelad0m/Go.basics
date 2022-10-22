package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
)

// начало решения
type rndReader struct {
	src []byte
	max int
	cur int
}

// Read implements the io.Reader interface.
func (r *rndReader) Read(b []byte) (n int, err error) {
	if r.cur >= r.max {
		return 0, io.EOF
	}
	n = copy(b, r.src[r.cur:])
	r.cur += n
	return
}

// RandomReader создает читателя, который возвращает случайные байты,
// но не более max штук
func RandomReader(max int) io.Reader {
	src := make([]byte, max)
	rand.Read(src)
	return &rndReader{src: src, max: max, cur: 0}
}

// конец ВООБШЕ НЕПРАВИЛЬНОГО (?), но решения

/*
type randomReader struct{}

func (r *randomReader) Read(p []byte) (n int, err error) {
    return rand.Read(p)
}

func RandomReader(max int) io.Reader {
    rd := &randomReader{}
    return io.LimitReader(rd, int64(max))
}
*/


/*
type randomReader struct {
    max int
    n   int
}

func (r *randomReader) Read(p []byte) (n int, err error) {
    if r.n >= r.max {
        return 0, io.EOF
    }
    if len(p) > r.max {
        p = p[:r.max]
    }
    n, err = rand.Read(p)
    r.n += n
    return
}

func RandomReader(max int) io.Reader {
    return &randomReader{max: max, n: 0}
}
*/

/*
func RandomReader(max int) io.Reader {
	randBytes := make([]byte, max)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(randBytes)
}
*/

func main() {
	rand.Seed(0)

	rnd := RandomReader(5)
	rd := bufio.NewReader(rnd)
	for {
		b, err := rd.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", b)
	}
	fmt.Println()
	// 1 148 253 194 250
}
