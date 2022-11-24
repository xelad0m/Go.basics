package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1]
	//fmt.Printf("%q\n", strings.Split(str, " "))

	count := 0
	for _, s := range strings.Split(str, " ") {
		if len(s) != 0 {
			count++
		}
	}

	fmt.Println(count)
}
