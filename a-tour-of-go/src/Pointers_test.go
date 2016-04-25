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

type Vertex struct {
	X int
	Y int
}

func ExampleStruct1() {
	fmt.Println(Vertex{1, 2})

	// Output: {1 2}
}

func ExampleStruct2() {
	v := Vertex{1, 2}
	v.X = 4

	fmt.Println(v.X)

	// Output: 4
}

func ExampleStructLiterals1() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1} // Y: 0 is implicit
		v3 = Vertex{}     // X: 0 and Y: 0
		p  = &Vertex{1, 3}
	)

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(p)

	// Output:
	// {1 2}
	// {1 0}
	// {0 0}
	// &{1 3}
}

func ExampleArrays1() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Output:
	// Hello World
	// [Hello World]
	// [2 3 5 7 11 13]
}

func ExampleSlices1() { /** slice is a dynamically-sized */
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	// Output: [3 5 7]
}

func ExampleSlices2() {
	/** slice does not store any data,
	 * it just describes a section of an underlying array
	 */

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" // update copied slice will affect the original array
	fmt.Println(names)
	fmt.Println(a, b)

	// Output:
	// [John Paul George Ringo]
	// [John Paul] [Paul George]
	// [John XXX George Ringo]
	// [John XXX] [XXX George]
}

func ExampleSliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// Output:
	// [2 3 5 7 11 13]
	// [true false true true false true]
	// [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
}

func ExampleSliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	// Output:
	// [3 5 7]
	// [3 5]
	// [5]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func ExampleSliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[:0]
	printSlice(s)

	s = s[:4] // extends its length
	printSlice(s)

	s = s[2:] // drop its first two values
	printSlice(s)

	// Output:
	// len=0 cap=6 []
	// len=4 cap=6 [2 3 5 7]
	// len=2 cap=4 [5 7]
}

func ExampleNilSlices() {
	var s []int // nil slice has a lanegth and capacity of 0 and has no underlying array
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}

	// Output:
	// [] 0 0
	// nil!
}

func printSliceWithString(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func ExampleCreatingASliceWithMake() {
}
