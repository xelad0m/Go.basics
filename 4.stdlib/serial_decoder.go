package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// поточное декодирование json lines

	r := strings.NewReader(`
	{ "name": "Alice", "age": 25 }
	{ "name": "Emma", "age": 23 }
	{ "name": "Grace", "age": 27 }
	`)

	dec := json.NewDecoder(r)
	for {
		var person Person
		err := dec.Decode(&person)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(person)
	}
	// {Alice 25}
	// {Emma 23}
	// {Grace 27}
}
