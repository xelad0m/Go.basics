package main

import (
	"testing"
	"time"
)

// начало решения

func isLeapYear(year int) bool {
	d := time.Date(year, 2, 29, 0, 0, 0, 0, time.UTC)
	// в невисокосном году дата распарсится как 01 марта
	return d.Month() == time.February

	// или
	// return time.Date(year, 2, 29, 12, 0, 0, 0, time.Local).Day() == 29
	
	// или классический вариант
	// return year%4 == 0 && (year%100 != 0 || year%400 == 0)

}

// конец решения

func Test(t *testing.T) {
	if !isLeapYear(2020) {
		t.Errorf("2020 is a leap year")
	}
	if isLeapYear(2022) {
		t.Errorf("2022 is NOT a leap year")
	}
}

func main() {
	Test(&testing.T{})
}
