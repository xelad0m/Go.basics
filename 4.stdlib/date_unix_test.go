package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	// "strings"
	"testing"
	"time"
)

// начало решения

// asLegacyDate преобразует время в легаси-дату
func asLegacyDate(t time.Time) string {
	s := t.Unix()
	ns := t.Nanosecond()

	rem := ns % 10
	for {
		if (ns !=0) && (rem == 0) {
			ns = ns / 10
			rem = ns % 10
		} else {
			break
		}
	}

	return fmt.Sprintf("%d.%d", s, ns)
}

// parseLegacyDate преобразует легаси-дату во время.
// Возвращает ошибку, если легаси-дата некорректная.
func parseLegacyDate(d string) (time.Time, error) {
	RE := regexp.MustCompile(`(\d+)\.(\d+)`)
	parts := RE.FindStringSubmatch(d)
	if len(parts) != 3 {
		return time.Time{}, errors.New("Bad time format")
	}

	s, _ := strconv.Atoi(parts[1])
	ns, _ := strconv.Atoi(parts[2])
	dig := len(parts[2])	// разрядность

	// 
	t := time.Unix(int64(s), int64(ns*int(math.Pow10(9 - dig))))
	return t, nil
}

// конец решения

/*
// asLegacyDate преобразует время в легаси-дату
func asLegacyDate(t time.Time) string {
    sec := t.Unix()
    nano := t.UnixNano() - sec*1e9
    if nano == 0 {
        return fmt.Sprintf("%d.0", sec)
    }
    str := fmt.Sprintf("%d.%d", sec, nano)
    return strings.TrimRight(str, "0")
}

// parseLegacyDate преобразует легаси-дату во время.
// Возвращает ошибку, если легаси-дата некорректная.
func parseLegacyDate(d string) (time.Time, error) {
    strSec, strNano, ok := strings.Cut(d, ".")
    if !ok {
        return time.Time{}, fmt.Errorf("invalid date: %v", d)
    }

    sec, err := strconv.ParseInt(strSec, 10, 64)
    if err != nil {
        return time.Time{}, fmt.Errorf("invalid date: %v", d)
    }

    if len(strNano) == 0 {
        return time.Time{}, fmt.Errorf("invalid date: %v", d)
    }
    strNano = padZerosRight(strNano, 9)
    nano, err := strconv.ParseInt(strNano, 10, 64)
    if err != nil {
        return time.Time{}, fmt.Errorf("invalid date: %v", d)
    }

    return time.Unix(sec, nano), nil
}

// padZerosRight отбивает строку нулями справа до указанной длины
func padZerosRight(str string, length int) string {
    if len(str) >= length {
        return str
    }
    return str + strings.Repeat("0", length-len(str))
}
*/

func Test_asLegacyDate(t *testing.T) {
	samples := map[time.Time]string{
		time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC):  "3600.123456789",
		time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC):          "3600.0",
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC):          "0.0",
		time.Date(2022, 5, 24, 14, 45, 22, 951, time.UTC):    "1653403522.951",
		time.Date(2022, 5, 24, 14, 45, 22, 951205, time.UTC): "1653403522.951205",
	}
	for src, want := range samples {
		got := asLegacyDate(src)
		if got != want {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}

func Test_parseLegacyDate(t *testing.T) {
	samples := map[string]time.Time{
		"3600.123456789":       time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC),
		"3600.0":               time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC),
		"0.0":                  time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		"1.123456789":          time.Date(1970, 1, 1, 0, 0, 1, 123456789, time.UTC),
		"1653403522.951205999": time.Date(2022, 5, 24, 14, 45, 22, 951205999, time.UTC),
		// "1653403522":           error,
	}
	for src, want := range samples {
		got, err := parseLegacyDate(src)
		if err != nil {
			t.Fatalf("%v: unexpected error", src)
		}
		if !got.Equal(want) {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}
