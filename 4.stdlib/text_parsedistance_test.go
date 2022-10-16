package main

import (
	"strconv"
	"strings"
	"testing"
)

// начало решения

// calcDistance возвращает общую длину маршрута в метрах
func calcDistance(directions []string) int {
	distance := 0
	for _, dir := range directions {
		word := getDigits(dir)
		distance += parseDigits(word)
	}
	return distance
}

func getDigits(str string) string {
	for _, s := range strings.Fields(str) {
		if strings.ContainsAny(s, "0123456789") {
			return s
		}
	}
	return ""
}

func parseDigits(str string) int {
	if strings.HasSuffix(str, "km") {
		str = strings.Split(str, "km")[0]
		f, _ := strconv.ParseFloat(str, 64)
		return int(1000 * f)
	} else {
		str = strings.Split(str, "m")[0]
		i, _ := strconv.Atoi(str)
		return i
	}

}

// конец решения


/* Вариант
// calcDistance возвращает общую длину маршрута в метрах
func calcDistance(directions []string) int {
    total := 0
    for _, dir := range directions {
        total += extractDistance(dir)
    }
    return total
}

// extractDistance извлекает из строки расстояние
// в метрах
func extractDistance(s string) int {
    for _, word := range strings.Fields(s) {
        char, _ := utf8.DecodeRuneInString(word)
        if !unicode.IsDigit(char) {
            continue
        }
        if strings.HasSuffix(word, "km") {
            return parseDistance(word[:len(word)-2], 1000)
        }
        return parseDistance(word[:len(word)-1], 1)
    }
    return 0
}

// parseDistance преобразует строковое расстояние
// в целое число с учетом мультипликатора
func parseDistance(distance string, multiplier int) int {
    num, _ := strconv.ParseFloat(distance, 64)
    return int(num * float64(multiplier))
}*/

func Test(t *testing.T) {
	directions := []string{
		"100m to intersection",
		"turn right",
		"straight 300m",
		"enter motorway",
		"straight 5.12km",
		"exit motorway",
		"500m straight",
		"turn sharp left",
		"continue 100m to destination",
	}
	const want = 6120
	got := calcDistance(directions)
	if got != want {
		t.Errorf("%v: got %v, want %v", directions, got, want)
	}
}

func main() {
	Test(&testing.T{})
}
