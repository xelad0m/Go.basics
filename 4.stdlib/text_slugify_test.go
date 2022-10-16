package main

// go test ./text_slugify_test.go -v

import (
	// "fmt"
	"strings"
	"testing"
)

// начало решения

// НИКОГДА не использовать сплит, только филд! сплит "чето   там" -> ["чето", "", "", "", "там"]

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
/*func slugify(src string) string {
	var runes []byte
	for i, r := range src {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || (r == ' ') || (r == '-') {
			runes = append(runes, src[i])
		} else {
			runes = append(runes, ' ')
		}
	}
	result := string(runes)
	fmt.Printf("'%s'\n", result)

	var arr []string
	for _, s := range strings.Split(result, " ") {
		if s != "" {
			arr = append(arr, s)
		}
	}

	fmt.Printf("'%s'\n", arr)
	result = strings.Join(arr, "-")

	result = strings.ToLower(result)
	return string(result)

}*/

// конец решения

// кратко
/*func slugify(src string) string {
	src = strings.ToLower(src)
	words := strings.FieldsFunc(src, func(r rune) bool {
		return (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') && r != '-'
	})
	return strings.Join(words, "-")
}*/

// средне
func slugify(src string) string {
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

func Test(t *testing.T) {
	const phrase = "Go 1.18 is released!"
	const want = "go-1-18-is-released"
	got := slugify(phrase)

	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}

func main() {
	Test(&testing.T{})
}
