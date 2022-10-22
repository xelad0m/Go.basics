package main

// go test ./text_slugify_test.go -v

import (
	"fmt"
	"unsafe"
)

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

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func slugify6(src string) string {
	src_bytes := *(*[]byte)(unsafe.Pointer(&src))
	runes := make([]byte, len(src))										// пишем результат сюда
	prev_bad := true

	ptr := 0															// индекс текущего в результате
	for _, r := range src_bytes {
		// 'a-z' 97-122 'A-Z' 65-90 '0-9' 48-57 '-' 45 (выяснилось, что не ускорит, считает при компиляции)
		if r > 64 && r < 91 { 											// A-Z
			r += 32 													// ToLower
		}
		if (r == 45) || (r > 47 && r < 58) || (r > 96 && r < 123) {		// разрешенный символ
			runes[ptr] = r
			prev_bad = false											// флаг начала разрешенного фрагмента
			ptr++
		} else {
			if prev_bad {												// уже был плохой - пропускаем
				continue
			}
			runes[ptr] = 45
			prev_bad = true												// плохой добавлен
			ptr++
		}
	}

	if ptr > 1 && prev_bad {	// если в конце были недопустимые символы
		runes = runes[:ptr-1]	// то надо убрать '-' из хвоста
	} else if ptr > 0 {
		runes = runes[:ptr]		// отчикать лишнее
	} else {
		runes = nil				// приведение nil даст пустую строку (\0 \х00 не канает)
	}

	return *(*string)(unsafe.Pointer(&runes))
}

func main() {
	const phrase = "? ?  A 100x Investment (2019) ! Go 1.18   is released! Go - 1.18 is - released! !"
	const want = "a-100x-investment-2019-go-1-18-is-released-go---1-18-is---released"

	// const phrase = ""
	// const want = ""

	got := slugify6(phrase)

	fmt.Printf("%s\n%s\n%s\n%v\n", phrase, got, want, got == want)
}
