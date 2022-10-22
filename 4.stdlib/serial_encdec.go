package main

import (
	"bufio"
	"encoding/json"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// поточное кодирование json lines

	people := []Person{
		{"Alice", 25},
		{"Emma", 23},
		{"Grace", 27},
	}

	w := bufio.NewWriter(os.Stdout)
	enc := json.NewEncoder(w)
	for _, person := range people {
		err := enc.Encode(person)
		if err != nil {
			panic(err)
		}
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}

	// {"name":"Alice","age":25}
	// {"name":"Emma","age":23}
	// {"name":"Grace","age":27}
}
