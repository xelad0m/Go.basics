package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

// начало решения

// Task описывает задачу, выполненную в определенный день
type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

// ParsePage разбирает страницу журнала
// и возвращает задачи, выполненные за день
func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	date, err := parseDate(lines[0])

	if err != nil {
		return nil, err
	}

	tasks, err := parseTasks(date, lines[1:])
	sortTasks(tasks)
	return tasks, err
}

// parseDate разбирает дату в формате дд.мм.гггг
func parseDate(src string) (time.Time, error) {
	layout := "02.01.2006"
	return time.Parse(layout, src)
}

// parseTasks разбирает задачи из записей журнала
func parseTasks(date time.Time, lines []string) ([]Task, error) {
	time_title_RE := regexp.MustCompile(`(\d+:\d+) - (\d+:\d+) (.+)`)
	tmp := make(map[string]time.Duration)

	for _, line := range lines {
		match := time_title_RE.FindStringSubmatch(line)
		if len(match) < 4 {
			return nil, errors.New("bad string")
		}

		start, err1 := time.Parse("15:04", match[1])
		end, err2 := time.Parse("15:04", match[2])

		if (err1 != nil) || (err2 != nil) {
			return nil, errors.New("wrong time format")
		}
		if !end.After(start) {	// автор хочет, чтоб 0 продолжительность тоже отсекалась
			return nil, errors.New("wrong time format")
		}

		tmp[match[3]] += end.Sub(start)
	}

	v := make([]Task, 0, len(tmp))

	for title, dur := range tmp {
		v = append(v, Task{date, dur, title})
	}

	return v, nil
}

// sortTasks упорядочивает задачи по убыванию длительности
func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i].Dur > tasks[j].Dur })
}

// конец решения

func main() {
	page := `15.04.2022
8:00 - 8:30 Завтрак
8:30 - 9:30 Оглаживание кота
9:30 - 10:00 Интернеты
10:00 - 14:00 Напряженная работа
14:00 - 14:45 Обед
14:45 - 15:00 Оглаживание кота
15:00 - 19:00 Напряженная работа
19:00 - 19:30 Интернеты
19:30 - 22:30 Безудержное веселье
22:30 - 23:00 Оглаживание кота`

	parseTasks(time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), []string{"11:00 - 12:00 task"})

	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}
	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}

	// ожидаемый результат
	/*
		Мои достижения за 2022-04-15
		- Напряженная работа: 8h0m0s
		- Безудержное веселье: 3h0m0s
		- Оглаживание кота: 1h45m0s
		- Интернеты: 1h0m0s
		- Обед: 45m0s
		- Завтрак: 30m0s
	*/
}
