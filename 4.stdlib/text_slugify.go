package main

// go test ./text_slugify_test.go -v

import (
	"fmt"
)

func slugify(src string) string {
	runes := make([]byte, len(src))
	prev_bad := true

	ptr := 0
	for _, r := range src {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || (r == '-') {
			if r >= 'A' && r <= 'Z' {
				r += 32 // ToLower
			}
			runes[ptr] = byte(r)
			prev_bad = false
			ptr++
		} else {
			if prev_bad {
				continue
			}
			runes[ptr] = byte('-')
			prev_bad = true
			ptr++
		}
	}

	if ptr > 1 && prev_bad && runes[ptr-1] == '-' {
		return string(runes[:ptr-1])
	} else if ptr > 0 {
		return string(runes[:ptr])
	} else {
		return ""
	}
}

func main() {
	const phrase = "? ?  A 100x Investment (2019) ! Go 1.18   is released! Go - 1.18 is - released! !"
	const want = "a-100x-investment-2019-go-1-18-is-released-go---1-18-is---released"

	// const phrase = ""
	// const want = ""

	got := slugify(phrase)

	fmt.Printf("%s\n%s\n%s\n", phrase, got, want)
}
