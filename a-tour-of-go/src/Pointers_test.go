package main

import (
	"fmt"
)

func ExamplePointers1() {
	i, j := 42, 2701

	p := &i // address of i to p
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	// Output:
	// 42
	// 21
	// 73
}
