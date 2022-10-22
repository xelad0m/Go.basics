package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// начало решения

// Duration описывает продолжительность фильма
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	d1 := time.Duration(d)
	h := int64(d1.Hours())
	m := int64((d1 - time.Duration(h*int64(time.Hour))).Minutes())

	if h == 0 {
		return []byte(fmt.Sprintf(`"%dm"`, m)), nil
	} else if m == 0 {
		return []byte(fmt.Sprintf(`"%dh"`, h)), nil
	} else {
		return []byte(fmt.Sprintf(`"%dh%dm"`, h, m)), nil
	}
}

// Rating описывает рейтинг фильма
type Rating int

func (r Rating) MarshalJSON() ([]byte, error) {
	// ★☆
	rt := []rune{}
	for i := 0; i < 5; i++ {
		if i < int(r) {
			rt = append(rt, '★')
		} else {
			rt = append(rt, '☆')
		}
	}
	return []byte(fmt.Sprintf(`"%s"`, string(rt))), nil
}

// Movie описывает фильм
type Movie struct {
	Title    string   // "Interstellar",
	Year     int      // 2014,
	Director string   // "Christopher Nolan",
	Genres   []string // []string{"Adventure", "Drama", "Science Fiction"},
	Duration Duration // Duration(2*time.Hour + 49*time.Minute),
	Rating   Rating   // 5,
}

// MarshalMovies кодирует фильмы в JSON.
// - если indent = 0 - использует json.Marshal
// - если indent > 0 - использует json.MarshalIndent
//   с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {
	m := []Movie{}
	for _, mv := range movies {
		m = append(m, mv)
	}

	if indent == 0 {
		b, err := json.Marshal(m)
		return string(b), err
	} else {
		ind := []byte{}
		for i := 0; i < indent; i++ {
			ind = append(ind, ' ')
		}
		b, err := json.MarshalIndent(m, "", string(ind))
		return string(b), err
	}
}

// конец решения

/*
// Duration описывает продолжительность фильма
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
    str := time.Duration(d).String()
    if strings.HasSuffix(str, "m0s") {
        str = str[:len(str)-2]
    }
    if strings.HasSuffix(str, "h0m") {
        str = str[:len(str)-2]
    }
    b := make([]byte, 0, len(str)+2)
    b = append(b, '"')
    b = append(b, []byte(str)...)
    b = append(b, '"')
    return b, nil
}

// Rating описывает рейтинг фильма
type Rating int

func (r Rating) MarshalJSON() ([]byte, error) {
    str := strings.Repeat("★", int(r)) + strings.Repeat("☆", 5-int(r))
    b := make([]byte, 0, 7)
    b = append(b, '"')
    b = append(b, []byte(str)...)
    b = append(b, '"')
    return b, nil
}

// Movie описывает фильм
type Movie struct {
    Title    string
    Year     int
    Director string
    Genres   []string
    Duration Duration
    Rating   Rating
}

// MarshalMovies кодирует фильмы в JSON.
// Если indent > 0 - форматирует с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {
    var b []byte
    var err error
    if indent <= 0 {
        b, err = json.Marshal(movies)
    } else {
        padding := strings.Repeat(" ", indent)
        b, err = json.MarshalIndent(movies, "", padding)
    }
    if err != nil {
        return "", nil
    }
    return string(b), nil
}
*/

func main() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(0*time.Hour + 49*time.Minute),
		Rating:   5,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour + 0*time.Minute),
		Rating:   4,
	}

	b, err := MarshalMovies(4, m1, m2)
	fmt.Println(err)
	// nil
	fmt.Println(string(b))
	/*
		[
		    {
		        "Title": "Interstellar",
		        "Year": 2014,
		        "Director": "Christopher Nolan",
		        "Genres": [
		            "Adventure",
		            "Drama",
		            "Science Fiction"
		        ],
		        "Duration": "2h49m",
		        "Rating": "★★★★★"
		    },
		    {
		        "Title": "Sully",
		        "Year": 2016,
		        "Director": "Clint Eastwood",
		        "Genres": [
		            "Drama",
		            "History"
		        ],
		        "Duration": "1h36m",
		        "Rating": "★★★★☆"
		    }
		]
	*/
}
