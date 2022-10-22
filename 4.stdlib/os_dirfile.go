package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"bufio"
)

// алфавит планеты Нибиру
const alphabet = "aeiourtnsl"

// Census реализует перепись населения.
// Записи о рептилоидах хранятся в каталоге census, в отдельных файлах,
// по одному файлу на каждую букву алфавита.
// В каждом файле перечислены рептилоиды, чьи имена начинаются
// на соответствующую букву, по одному рептилоиду на строку.
type Census struct {
	files map[rune]string
}

// Count возвращает общее количество переписанных рептилоидов.
func (c *Census) Count() int {
	return 0
}

// Add записывает сведения о рептилоиде.
func (c *Census) Add(name string) {
	first_letter := []rune(name)[0]
	p := c.files[first_letter]

	file, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(name)
	writer.WriteByte('\n')

	err = writer.Flush()
	if err != nil {
		file.Close()
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}
	
}

// Close закрывает файлы, использованные переписью.
func (c *Census) Close() {
	// os.RemoveAll("census")
}

// NewCensus создает новую перепись и пустые файлы
// для будущих записей о населении.
func NewCensus() *Census {
	os.Mkdir("census", 0755)

	touch := func(path string) {
		p := filepath.FromSlash(path)
		data := []byte{}
		os.WriteFile(p, data, 0644)
	}

	files := make(map[rune]string)
	for _, rune := range alphabet {
		p := filepath.Join("census", string(rune)+".txt")
		touch(p)
		files[rune] = p
	}

	return &Census{files: files}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// randomName возвращает имя очередного рептилоида.
func randomName(n int) string {
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(chars)
}

func main() {
	rand.Seed(0)
	census := NewCensus()
	defer census.Close()
	for i := 0; i < 1024; i++ {
		reptoid := randomName(5)
		census.Add(reptoid)
	}
	fmt.Println(census.Count())
}
