package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func titleCase(s string) string {
	return strings.Title(strings.ToLower(s))
}

func main () {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)			// по пробелу

	for scanner.Scan() {
		word := scanner.Text()
		fmt.Printf("%s ", titleCase(word))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

/* // как вариант
func main() {
    r := bufio.NewReader(os.Stdin)
    s, _ := r.ReadString('\n')
    fmt.Print(strings.Title(strings.ToLower(s)))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		rns := []rune(scanner.Text())
		rns[0] = rune(unicode.ToUpper(rns[0]))
        
        fmt.Printf("%v%v ", string(rns[0]),strings.ToLower(string(rns[1:])))
	}
}
*/