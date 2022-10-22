package main

import (
	"encoding/json"
	"fmt"
)

// начало решения

// Genre описывает жанр фильма
type Genre string

// Movie описывает фильм
type Movie struct {
	Title  string  `json:"name"`
	Year   int     `json:"released_at"`
	Genres []Genre `json:"tags"`
}

func (g *Genre) UnmarshalJSON(data []byte) error {
	// Go рекомендует игнорировать значения null
	if string(data) == "null" {
		return nil
	}
	
	// декодируем исходную строку через карту с приведением типа
	var gnr any
	err := json.Unmarshal(data, &gnr)
	if err != nil {
		fmt.Println("err")
		return err
	}

	m := gnr.(map[string]any)

	name := m["name"].(string)
	*g = Genre(name)
	return nil
}

// конец решения

/* // более красиво

func (g *Genre) UnmarshalJSON(data []byte) error {
    if string(data) == "null" {
        return nil
    }
    var obj map[string]string
    if err := json.Unmarshal(data, &obj); err != nil {
        return err
    }
    if val, ok := obj["name"]; ok {
        *g = Genre(val)
    }
    return nil
}
*/

func main() {
	const src = `{
		"name": "Interstellar",
		"released_at": 2014,
		"director": "Christopher Nolan",
		"tags": [
			{ "name": "Adventure" },
			{ "name": "Drama" },
			{ "name": "Science Fiction" }
		],
		"duration": "2h49m",
		"rating": "★★★★★"
	}`

	var m Movie
	err := json.Unmarshal([]byte(src), &m)
	fmt.Println(err)
	// nil
	fmt.Println(m)
	// {Interstellar 2014 [Adventure Drama Science Fiction]}
}
