package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	{
		fmt.Println("---")

		re := regexp.MustCompile(`\d+`)
		s := "2050-11-05 is November 5th, 2050"

		ok := re.MatchString(s)
		fmt.Println("MatchString")
		fmt.Println(ok)
		fmt.Println()
		// true

		first := re.FindString(s)
		fmt.Println("FindString")
		fmt.Println(first)
		fmt.Println()
		// 2050

		idx := re.FindStringIndex(s)
		fmt.Println("FindStringIndex")
		fmt.Println(idx)
		fmt.Println()
		// [0 4]

		three := re.FindAllString(s, 3)
		fmt.Println("FindAllString")
		fmt.Println(three)
		fmt.Println()
		// [2050 11 05]

		indices := re.FindAllStringIndex(s, 3)
		fmt.Println("FindAllStringIndex")
		fmt.Println(indices)
		fmt.Println()
		// [[0 4] [5 7] [8 10]]
	}

	{
		fmt.Println("---")

		re := regexp.MustCompile(`(\d\d\d\d)-(\d\d)-(\d\d)`)
		s := "2050-11-05 is November 5th, 2050"

		match := re.FindString(s)
		fmt.Println("FindString")
		fmt.Println(match)
		fmt.Println()
		// 2050-11-05

		groups := re.FindStringSubmatch(s)
		fmt.Println("FindStringSubmatch")
		fmt.Println(groups)
		fmt.Println()
		// [2050-11-05 2050 11 05]

		indices := re.FindStringSubmatchIndex(s)
		fmt.Println("FindStringSubmatchIndex")
		fmt.Println(indices)
		fmt.Println()
		// [0 10 0 4 5 7 8 10]
	}

	{
		fmt.Println("---")

		re := regexp.MustCompile(`\s*\d+\s*`)
		s := "one 01 two 02 three 03"

		parts := re.Split(s, -1)
		fmt.Println("Split")
		fmt.Printf("%#v\n", parts)
		fmt.Println()
		// []string{"one", "two", "three", ""}
	}

	{
		fmt.Println("---")

		re := regexp.MustCompile(`(\d\d\d\d)-(\d\d)-(\d\d)`)
		src := "2050-11-05 is November 5th, 2050"

		res := re.ReplaceAllString(src, "$3.$2.$1")
		fmt.Println("ReplaceAllString")
		fmt.Println(res)
		fmt.Println()
		// 05.11.2050 is November 5th, 2050

		fn := func(src string) string {
			parts := strings.Split(src, "-")
			reversed := []string{parts[2], parts[1], parts[0]}
			return strings.Join(reversed, ".")
		}
		res = re.ReplaceAllStringFunc(src, fn)
		fmt.Println("ReplaceAllStringFunc")
		fmt.Println(res)
		fmt.Println()
		// 05.11.2050 is November 5th, 2050
	}
}
