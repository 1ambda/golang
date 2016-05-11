package main

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

/** A method is a function with a special receiver argument */
func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs2(v Vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ExampleMethods() {
	v := Vertex3{3, 4}
	fmt.Println(v.Abs()) /** method */
	fmt.Println(Abs2(v)) /** function */

	// Output:
	// 5
	// 5
}

/** We can declare a method on non-struct types too */
type MyFloat0 float64

func (f MyFloat0) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func ExampleMethodContinued() {
	f := MyFloat0(-math.Sqrt2)

	fmt.Println(f.Abs())

	// Output: 1.4142135623730951
}

func (v Vertex3) Scale1(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex3) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ExamplePointerReceivers() {
	v := Vertex3{3, 4}
	v.Scale1(10)
	fmt.Println(v.Abs()) /** value receiver */
	v.Scale2(10)
	fmt.Println(v.Abs()) /** pointer receiver */

	// Output:
	// 5
	// 50
}

func Scale3(v Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale4(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ExamplePointersAndFunctions() {
	v := Vertex3{3, 4}

	Scale3(v, 10) /** will copy v */
	fmt.Println(v)

	Scale4(&v, 10) /** will use the original v */
	fmt.Println(v)

	// Output:
	// {3 4}
	// {30 40}
}

func ExampleMethodAndPointerIndirection() {
	v1 := Vertex3{3, 4}
	/** Scale4(v1, 10) will cause compile error since Scale4 requires Vertex pointer */

	v1.Scale2(10)
	fmt.Println(v1) /** receivers can take either a value or a pointer */

	p1 := &Vertex3{4, 3}
	p1.Scale2(10)
	fmt.Println(*p1)
	p1.Scale1(10) /** method will copy p1 instead of using p1 */
	fmt.Println(*p1)

	// Output:
	// {30 40}
	// {40 30}
	// {40 30}
}
