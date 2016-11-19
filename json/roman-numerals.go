package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	roman "github.com/StefanSchroeder/Golang-Roman"
)

type RomanNumeral int
type Movie struct {
	Title string
	Year  RomanNumeral
}

func main() {

	// Encoding
	movies := []Movie{{"E.T.", 1982}, {"The Matrix", 1999}}
	encoded, err := json.MarshalIndent(movies, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Movies: %s\n", encoded)

	// Decoding
	var m Movie
	input := `{"Title": "Alien", "Year": "MCMLXXIX"}`
	if err := json.NewDecoder(strings.NewReader(input)).Decode(&m); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s was released in %d\n", m.Title, m.Year)
}

func (r *RomanNumeral) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("roman numerals should be string but got %s", data)
	}

	res := roman.Arabic(s)
	if res == -1 {
		return fmt.Errorf("invalid roman numerals: %s", data)
	}

	*r = RomanNumeral(res)
	return nil
}

func (r RomanNumeral) MarshalJSON() ([]byte, error) {
	if r <= 0 {
		return nil, fmt.Errorf("Roman had only natural numbers, but got %d", r)
	}

	s := roman.Roman(int(r))
	return json.Marshal(s)
}
