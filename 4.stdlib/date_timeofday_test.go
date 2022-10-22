package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// начало решения

// TimeOfDay описывает время в пределах одного дня
type TimeOfDay struct {
	wrapped time.Time
}

// Hour возвращает часы в пределах дня
func (t TimeOfDay) Hour() int {
	return t.wrapped.Hour()
}

// Minute возвращает минуты в пределах часа
func (t TimeOfDay) Minute() int {
	return t.wrapped.Minute()
}

// Second возвращает секунды в пределах минуты
func (t TimeOfDay) Second() int {
	return t.wrapped.Second()
}

// String возвращает строковое представление времени
// в формате чч:мм:сс TZ (например, 12:34:56 UTC)
func (t TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d:%02d %s", t.Hour(), t.Minute(), t.Second(), t.wrapped.Location())
}

// Equal сравнивает одно время с другим.
// Если у t и other разные локации - возвращает false.
func (t TimeOfDay) Equal(other TimeOfDay) bool {
	return (t.Hour() == other.Hour()) && (t.Minute() == other.Minute()) && (t.Second() == other.Second()) && (t.wrapped.Location().String() == other.wrapped.Location().String())

}

var ErrDifferentLocations = errors.New("different locations")

// Before возвращает true, если время t предшествует other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) Before(other TimeOfDay) (bool, error) {
	if t.wrapped.Location().String() != other.wrapped.Location().String() {
		return false, ErrDifferentLocations
	}

	before := false
	if t.Hour() < other.Hour() {
		before = true
	} else if (t.Hour() == other.Hour()) && (t.Minute() < other.Minute()) {
		before = true
	} else if (t.Minute() == other.Minute()) && (t.Second() < other.Second()) {
		before = true
	}

	return before, nil
}

// After возвращает true, если время t идет после other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) After(other TimeOfDay) (bool, error) {
	if t.wrapped.Location().String() != other.wrapped.Location().String() {
		return false, ErrDifferentLocations
	}

	after := false
	if t.Hour() > other.Hour() {
		after = true
	} else if (t.Hour() == other.Hour()) && (t.Minute() > other.Minute()) {
		after = true
	} else if (t.Minute() == other.Minute()) && (t.Second() > other.Second()) {
		after = true
	}
	// ... или просто использовать wrapped.After(), если полагаемся на равенство года/месяца/числа
	
	return after, nil
}

// MakeTimeOfDay создает время в пределах дня
func MakeTimeOfDay(hour, min, sec int, loc *time.Location) TimeOfDay {
	return TimeOfDay{time.Date(0, 0, 0, hour, min, sec, 0, loc)}
}

// конец решения

func Test(t *testing.T) {
	t1 := MakeTimeOfDay(17, 45, 22, time.UTC)
	t2 := MakeTimeOfDay(20, 3, 4, time.UTC)

	fmt.Println(t1.String())
	if t1.Equal(t2) {
		t.Errorf("%v should not be equal to %v", t1, t2)
	}

	before, _ := t1.Before(t2)
	if !before {
		t.Errorf("%v should be before %v", t1, t2)
	}

	after, _ := t1.After(t2)
	if after {
		t.Errorf("%v should NOT be after %v", t1, t2)
	}
}
