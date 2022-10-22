package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name      string
	BirthDate time.Time
}

func main() {
	// определяемые типы, срезы, карты

	{
		date, _ := time.Parse("2006-01-02", "2000-05-25")
		alice := Person{
			Name:      "Alice",
			BirthDate: date,
		}

		b, err := json.Marshal(alice)
		fmt.Println(err, string(b))
		// <nil> {"Name":"Alice","BirthDate":"2000-05-25T00:00:00Z"}
		// RFC3339Nano
	}

	{
		nums := []int{1, 3, 5}
		b, err := json.Marshal(nums)
		fmt.Println(err, string(b))
		// <nil> [1,3,5]
	}

	{
		m := map[string]int{
			"one":   1,
			"three": 3,
			"five":  5,
		}
		b, err := json.Marshal(m)
		fmt.Println(err, string(b))
		// <nil> {"five":5,"one":1,"three":3}
	}

	{
		ch := make(chan int)
		_, err := json.Marshal(ch)
		fmt.Println(err)
		// json: unsupported type: chan int
	}

	{
		fn := func() int { return 42 }
		_, err := json.Marshal(fn)
		fmt.Println(err)
		// json: unsupported type: func() int
	}

}
