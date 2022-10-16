package main

// go test ./text_slugify_test.go -v

import (
	// "fmt"
	"regexp"
	"strings"
	"testing"
)

/*
func slugify(src string) string {
	res := strings.ToLower(src)

	sep := regexp.MustCompile(`[^0-9a-z\-]+`)		// 1+ НЕ разрешенных символов (включая пробел)
	res = sep.ReplaceAllString(res, "-")
	res = strings.Trim(res, "-")					// в начале и конце могут остаться хвосты
	return res
}
*/

// последовательность допустимых символов
var wordRE = regexp.MustCompile(`[a-z0-9\-]+`)

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
func slugify(src string) string {
    words := wordRE.FindAllString(strings.ToLower(src), -1)
    return strings.Join(words, "-")
}

func Test(t *testing.T) {
	const phrase = "Go: 1.18 - is released!!"
	const want = "go-1-18---is-released"
	got := slugify(phrase)

	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}

func main() {
	Test(&testing.T{})
}
