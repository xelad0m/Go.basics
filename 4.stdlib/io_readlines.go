package main

import (
	"fmt"
	"os"

	// "strings"
	"bufio"
)

// начало решения

/*
// readLines возвращает все строки из указанного файла
func readLines(name string) ([]string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	if len(data) < 1 {
		return []string{}, nil
	} else {
		return strings.Split(string(data[:len(data)-1]), "\n"), nil
	}
}
*/

// конец решения

/*
func readLines(name string) ([]string, error) {
    data, err := os.ReadFile(name)
    if err != nil {
        return nil, err
    }
    lines := strings.Split(string(data), "\n")
    if len(lines) > 0 && lines[len(lines)-1] == "" {
        lines = lines[:len(lines)-1]
    }
    return lines, nil
}
*/

func readLines(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := []string{}
	for scanner.Scan() {						// в конце файла ИЛИ при обшибке возвращает false...
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {		// ...поэтому ошибка обрабатывается так
		return []string{}, err
	}

	return lines, nil
}

func main() {
	lines, err := readLines("/etc/passwd")
	if err != nil {
		panic(err)
	}
	for idx, line := range lines {
		fmt.Printf("%d: %s\n", idx+1, line)
	}
}
