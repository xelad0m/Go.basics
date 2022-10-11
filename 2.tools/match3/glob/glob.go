/*
glob package matches strings against patterns.

Supports wildcards:
	-  matches any single character
	*  matches everything
*/
package glob

import (
	"regexp"
)

// specials is a set of characters
// considered special by regexp.
var specials = makeSpecials([]rune{
	'\\', '.', '?', '+', '*', '^', '$',
	'|', '{', '}', '[', ']', '(', ')',
})

// match returns true if src matches pattern,
// false otherwise.
func Match(pattern string, src string) (bool, error) {
	pat := translate(pattern)
	re, err := regexp.Compile(pat)
	if err != nil {
		return false, err
	}
	isMatch := re.MatchString(src)
	return isMatch, nil
}

// translate converts match pattern
// to regexp pattern
func translate(pattern string) string {
	rePat := make([]rune, 0, len(pattern))
	for _, char := range pattern {
		switch char {
		case '*':
			rePat = append(rePat, '.', '*')
		case '?':
			rePat = append(rePat, '.')
		default:
			rePat = append(rePat, escape(char)...)
		}
	}
	return string(rePat)
}

// escape escapes characters
// considered special by regexp.
func escape(char rune) []rune {
	if _, isSpecial := specials[char]; isSpecial {
		return []rune{'\\', char}
	}
	return []rune{char}

}

// makeSpecials creates a set of characters
// considered special by regexp.
func makeSpecials(chars []rune) map[rune]bool {
	specials := make(map[rune]bool, len(chars))
	for _, char := range chars {
		specials[char] = true
	}
	return specials
}
