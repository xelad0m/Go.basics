package main

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

// начало решения

// prettify возвращает отформатированное
// строковое представление карты
func prettify(m map[string]int) string {
	if len(m) == 0 {
		return "{}"
	}

	prefix := "{\n"
	suffix := "}"
	if len(m) == 1 {
		prefix = "{ "
		suffix = " }"
	}

	var b strings.Builder

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		if len(m) > 1 { 
			b.WriteString("    ")
		}
		b.WriteString(key)
		b.WriteString(": ")
		b.WriteString(strconv.Itoa(m[key]))
		if len(m) > 1 {
			b.WriteString(",\n")
		}
	}
	return prefix + b.String() + suffix
}

// конец решения

/*
// prettify возвращает отформатированное
// строковое представление карты
func prettify(m map[string]int) string {
    if len(m) == 0 {
        return "{}"
    }
    if len(m) == 1 {
        for key, val := range m {
            return fmt.Sprintf("{ %v: %v }", key, val)
        }
    }
    keys := extractKeys(m)
    return asOrdered(keys, m)
}

// extractKeys возвращает упорядоченные
// по алфавиту ключи карты
func extractKeys(m map[string]int) []string {
    keys := make([]string, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    return keys
}

// asOrdered форматирует карту,
// выводя ключи в указанном порядке
func asOrdered(keys []string, m map[string]int) string {
    var b strings.Builder
    b.WriteString("{\n")
    for _, key := range keys {
        b.WriteString("    ")
        b.WriteString(key)
        b.WriteString(": ")
        b.WriteString(strconv.Itoa(m[key]))
        b.WriteString(",\n")
    }
    b.WriteString("}")
    return b.String()
}*/

func Test(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	const want = "{\n    one: 1,\n    three: 3,\n    two: 2,\n}"
	got := prettify(m)
	if got != want {
		t.Errorf("%v\ngot:\n%v\n\nwant:\n%v", m, got, want)
	}
}

func main() {
	Test(&testing.T{})
}
