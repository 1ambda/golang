package main

func main() {

	// we can put one stmt before the condition
	b, c := 1, 3

	if a := b + c; a > 45 {
		panic(0)
	}

	// type assertion inside if
	var val interface{}
	val = "foo"

	if str, ok := val.(string); !ok {
		panic(str)
	}

}

func test(expr bool) {
	if !expr {
		panic(0)
	}
}
