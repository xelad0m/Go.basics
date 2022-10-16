package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
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

func Test(t *testing.T) {
	// в го есть еще много таких же извращенских способов проверки равных значений,
	// но нормального в стандартной библиотеке нет
	allEqual := len(map[string]bool{
		slugify1(phrase): true,
		slugify2(phrase): true,
		slugify3(phrase): true,
		slugify4(phrase): true,
		slugify5(phrase): true,
	}) == 1

	fmt.Printf("All equal: %v\n", allEqual)
	if !allEqual {
		t.Errorf("One or more variants work wrong")
	}
}
