package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Stringer interface {
	String() string
}

func main() {
	p := Point{2, 3}
	var x Stringer = p
	// Method calls via interface types are dynamically dispatched (virtual function call)
	fmt.Println(x)
	// Method calls via non-interface types are statically dispatched
	fmt.Println(Point{3, 5}) // fmt.Println knows about Stringer
}
