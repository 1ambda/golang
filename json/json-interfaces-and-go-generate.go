package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type ShirtSize byte

const (
	NA ShirtSize = iota
	XS
	S
	M
	L
	XL
)

func main() {
	raw := `
{
    "name": "Gopher",
    "birthdate": "2009/11/10",
    "shirt-size": "XS"
}
`

	// 1. initial version
	var p1 Person1
	if err := p1.Parse(raw); err != nil {
		panic("p1.Parse failed")
	}

	// 2. use aux struct
	var p2 Person2
	if err := p2.Parse(raw); err != nil {
		panic("p2.Parse failed")
	}

	// 3. use Marshaler and Unmarshaler
	var p3 Person3
	dec := json.NewDecoder(strings.NewReader(raw))
	if err := dec.Decode(&p3); err != nil {
		log.Fatal("p3.Parse failed", err)
		panic(err)
	}
}

type Date struct{ time.Time }

type Person3 struct {
	Name string    `json:"name"`
	Born Date      `json:"birthdate"`
	Size ShirtSize `json:"shirt-size"`
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("birthdate should be a string, got %s", data)
	}

	t, err := time.Parse("2006/01/02", s)
	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}

	d.Time = t
	return nil
}

func (ss *ShirtSize) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("shirt-size should be a string, got %s", data)
	}

	got, ok := map[string]ShirtSize{"XS": XS, "S": S, "M": M, "L": L, "XL": XL}[s]
	if !ok {
		return fmt.Errorf("invalid ShirtSize: %q", s)
	}

	*ss = got
	return nil
}

type Person2 struct {
	Name string    `json:"name"`
	Born time.Time `json:"birthdate"`
	Size ShirtSize `json:"shirt-size"`
}

func (p *Person2) Parse(s string) error {
	// Let's use an auxiliary struct type to avoid parsing
	// Note: the field tag for Name is not needed;
	// the JSON decoder perform a case insensitive match
	// if the exact form is not found.
	var aux struct {
		Name string
		Born string `json:"birthdate"`
		Size string `json:"shirt-size"`
	}

	dec := json.NewDecoder(strings.NewReader(s))
	if err := dec.Decode(&aux); err != nil {
		return fmt.Errorf("decode person: %v", err)
	}

	p.Name = aux.Name
	born, err := time.Parse("2006/01/02", aux.Born)
	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}
	p.Born = born
	p.Size, err = ParseShirtSize(aux.Size)
	return err
}

type Person1 struct {
	Name string
	Born time.Time
	Size ShirtSize
}

func (p *Person1) Parse(s string) error {
	// use map
	fields := map[string]string{}

	dec := json.NewDecoder(strings.NewReader(s))
	if err := dec.Decode(&fields); err != nil {
		return fmt.Errorf("decode person: %v", err)
	}

	// 1. get name
	p.Name = fields["name"]

	// 2. get birthdate
	born, err := time.Parse("2006/01/02", fields["birthdate"])
	if err != nil {
		return fmt.Errorf("invalid date %v", err)
	}
	p.Born = born

	// 3. get shirt size
	p.Size, err = ParseShirtSize(fields["shirt-size"])

	return nil
}

func ParseShirtSize(s string) (ShirtSize, error) {
	sizes := map[string]ShirtSize{"XS": XS, "S": S, "M": M, "L": L, "XL": XL}
	ss, ok := sizes[s]
	if !ok {
		return NA, fmt.Errorf("invalid ShirtSize %q", s)
	}
	return ss, nil
}
