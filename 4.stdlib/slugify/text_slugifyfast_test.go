package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"unsafe"
)

// начало решения

var sepRE = regexp.MustCompile(`[^0-9a-z\-]+`) // 1+ НЕ разрешенных символов (включая пробел)

func slugify1(src string) string {
	res := strings.ToLower(src)
	res = sepRE.ReplaceAllString(res, "-")
	res = strings.Trim(res, "-") // в начале и конце могут остаться хвосты
	return res
}

var wordRE = regexp.MustCompile(`[a-z0-9\-]+`)

func slugify2(src string) string {
	words := wordRE.FindAllString(strings.ToLower(src), -1)
	return strings.Join(words, "-")
}

// средне
func slugify3(src string) string {
	res := strings.ToLower(src)
	res = strings.Map(purifyChar, res)
	words := strings.Fields(res)
	return strings.Join(words, "-")
}

// purifyChar преобразует недопустимые символы в пробелы
func purifyChar(r rune) rune {
	const validChars string = "abcdefghijklmnopqrstuvwxyz01234567890- "
	if strings.IndexRune(validChars, r) == -1 {
		return ' '
	}
	return r
}

// длинно
func slugify4(src string) string {
	var runes []byte
	for i, r := range src {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || (r == ' ') || (r == '-') {
			runes = append(runes, src[i])
		} else {
			runes = append(runes, ' ')
		}
	}
	result := string(runes)

	var arr []string
	for _, s := range strings.Split(result, " ") {
		if s != "" {
			arr = append(arr, s)
		}
	}

	result = strings.Join(arr, "-")
	result = strings.ToLower(result)
	return string(result)

}

// длинно
func slugify5(src string) string {
	runes := make([]byte, len(src))
	prev_bad := true

	ptr := 0
	for _, r := range src {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || (r == '-') {
			if r >= 'A' && r <= 'Z' {
				r += 32 // ToLower
			}
			runes[ptr] = byte(r)
			prev_bad = false
			ptr++
		} else {
			if prev_bad {
				continue
			}
			runes[ptr] = byte('-')
			prev_bad = true
			ptr++
		}
	}

	if ptr > 1 && prev_bad && runes[ptr-1] == '-' {
		return string(runes[:ptr-1])
	} else if ptr > 0 {
		return string(runes[:ptr])
	} else {
		return ""
	}
}

// unsafe
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func slugify6(src string) string {
	src_bytes := *(*[]byte)(unsafe.Pointer(&src))
	runes := make([]byte, len(src))
	prev_bad := true
	ptr := 0
	for _, r := range src_bytes {
		// a-z 97-122 A-Z 65-90 0-9 48-57 - 45
		if r > 64 && r < 91 { // A-Z
			r += 32 // ToLower
		}
		if (r == 45) || (r > 47 && r < 58) || (r > 96 && r < 123) {
			runes[ptr] = r
			prev_bad = false
			ptr++
		} else {
			if prev_bad {
				continue
			}
			runes[ptr] = 45
			prev_bad = true
			ptr++
		}
	}

	if ptr > 1 && prev_bad {
		runes = runes[:ptr-1]
	} else if ptr > 0 {
		runes = runes[:ptr]
	} else {
		runes = nil
	}

	return *(*string)(unsafe.Pointer(&runes))
}

const phrase = "? ?  A 100x Investment (2019) ! Go 1.18   is released! Go - 1.18 is - released! !"

func BenchmarkMatchSlugify1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slugify1(phrase)
	}
}

func BenchmarkMatchSlugify2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slugify2(phrase)
	}
}

func BenchmarkMatchSlugify3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slugify3(phrase)
	}
}

func BenchmarkMatchSlugify4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slugify4(phrase)
	}
}

func BenchmarkMatchSlugify5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slugify5(phrase)
	}
}

func BenchmarkMatchSlugify6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slugify6(phrase)
	}
}

func Test(t *testing.T) {
	// в го есть еще много таких же извращенских способов проверки равных значений,
	// но нормального в стандартной библиотеке нет
	allEqual := len(map[string]bool{
		slugify1(phrase): true,
		slugify2(phrase): true,
		slugify3(phrase): true,
		slugify4(phrase): true,
		slugify5(phrase): true,
		slugify6(phrase): true,
	}) == 1

	fmt.Println(slugify1(phrase))
	fmt.Println(slugify2(phrase))
	fmt.Println(slugify3(phrase))
	fmt.Println(slugify4(phrase))
	fmt.Println(slugify5(phrase))
	fmt.Println(slugify6(phrase))

	fmt.Println(map[string]bool{
		slugify1(phrase): true,
		slugify2(phrase): true,
		slugify3(phrase): true,
		slugify4(phrase): true,
		slugify5(phrase): true,
		slugify6(phrase): true,
	})
	fmt.Printf("All equal: %v\n", allEqual)
	if !allEqual {
		t.Errorf("One or more variants work wrong")
	}
}


/*
ЧЕМПИОНСКИЕ РЕШЕНИЯ
*/

/* 1 место
var charType = [256]int8{}

const (
	upper = 1
	lower = 2
	digit = 4
	dash  = 8
)

func init() {
	for i := 'A'; i <= 'Z'; i++ {
		charType[i] = upper
	}
	for i := 'a'; i <= 'z'; i++ {
		charType[i] = lower
	}
	for i := '0'; i <= '9'; i++ {
		charType[i] = digit
	}
	charType['-'] = dash
}

// Let's go black unsafe magic 😈

func ptrAdd(p *byte, n int) *byte {
	return (*byte)(unsafe.Add(unsafe.Pointer(p), n))
}

func ptrSub(p1, p2 *byte) int {
	return int(uintptr(unsafe.Pointer(p1)) - uintptr(unsafe.Pointer(p2)))
}

func ptrGet(p *byte) (byte, *byte) {
	return *p, (*byte)(unsafe.Add(unsafe.Pointer(p), 1))
}

func ptrSet(p *byte, v byte) *byte {
	*p = v
	return (*byte)(unsafe.Add(unsafe.Pointer(p), 1))
}

func slugify(s string) string {
	if len(s) == 0 {
		return ""
	}

	buf := make([]byte, len(s))

	// See `stringStruct` and `slice` in GOROOT/runtime/string.go and GOROOT/runtime/slice.go.
	src := *(**byte)(unsafe.Pointer(&s))
	dst := *(**byte)(unsafe.Pointer(&buf))

	srcEnd := ptrAdd(src, len(s))
	dstStart := dst
	ch := byte(0)

mainLoop:
	for {
		for {
			if src == srcEnd {
				ch = 0
				break mainLoop
			}

			ch, src = ptrGet(src)

			if t := charType[ch]; t&upper != 0 {
				ch += 32
				break
			} else if t&(lower|digit|dash) != 0 {
				break
			}
		}

		for {
			dst = ptrSet(dst, ch)

			if src == srcEnd {
				break mainLoop
			}

			ch, src = ptrGet(src)

			if t := charType[ch]; t&(lower|digit|dash) != 0 {
				//
			} else if t&upper != 0 {
				ch += 32
			} else {
				break
			}
		}

		dst = ptrSet(dst, '-')
	}

	count := ptrSub(dst, dstStart)

	if count == 0 {
		return ""
	}

	if ch == 0 {
		count--
	}

	// buf = buf[:count]
	(*struct {
		p   *byte
		len int
	})(unsafe.Pointer(&buf)).len = count

	return *(*string)(unsafe.Pointer(&buf))
}
*/

/* самое быстрое БЕЗ УКАЗАТЕЛЕЙ
func slugify(src string) string {
    dst := new(strings.Builder)
    dst.Grow(len(src) + 1)

    for j := 0; j < len(src); {
        for ; j < len(src) && !isValid(src[j]); j++ {
        }
        if j == len(src) {
            break
        }
        dst.WriteByte('-')
        for ; j < len(src) && isValid(src[j]); j++ {
            dst.WriteByte(toLower(src[j]))
        }
    }

    if dst.Len() == 0 {
        return ""
    }
    return dst.String()[1:]
}

func isValid(ch byte) bool {
    return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' || ch == '-'
}

func toLower(ch byte) byte {
    if ch >= 'A' && ch <= 'Z' {
        return ch + 32
    }
    return ch
}
*/