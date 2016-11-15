package main

type Vertex struct {
	a float64
	b float64
}

func main() {

	var m1 map[string]int = make(map[string]int)
	m1["key"] = 1

	if v1, ok := m1["key"]; ok {
		pass(true)
		_ = v1
	}

	// map literal
	var m2 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	_ = m2

	// using anonymous struct is cheaper and safer,
	// than using `map[string]interface{}`
	point := struct {
		X, Y int
	}{1, 2}
	_ = point
}

func pass(expr bool) {
	if !expr {
		panic(0)
	}
}
