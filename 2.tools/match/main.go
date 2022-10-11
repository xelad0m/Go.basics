// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	pattern, src, err := readInput()
	if err != nil {
		fail(err)
	}
	isMatch := match(pattern, src)
	if !isMatch {
		os.Exit(0)
	}
	fmt.Println(src)
}

// match returns true if src matches pattern,
// false otherwise.
func match(pattern string, src string) bool {
	return strings.Contains(src, pattern)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (pattern, src string, err error) {
	flag.StringVar(&pattern, "p", "", "pattern to match against")
	flag.Parse()
	if pattern == "" {
		return pattern, src, errors.New("missing pattern")
	}
	src = strings.Join(flag.Args(), "")
	if src == "" {
		return pattern, src, errors.New("missing string to match")
	}
	return pattern, src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}
