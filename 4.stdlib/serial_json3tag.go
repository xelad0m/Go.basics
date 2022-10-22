package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string  `json:"name"`
	Age       int     `json:"age"`
	Weight    float64 `json:"-"`
	IsAwesome bool    `json:"is_awesome"`
}

func main() {
	// теги

	alice := Person{
		Name:      "Alice",
		Age:       25,
		Weight:    55.5,
		IsAwesome: true,
	}

	b, err := json.MarshalIndent(alice, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	/*
		{
		    "name": "Alice",
		    "age": 25,
		    "is_awesome": true
		}
	*/
}
	