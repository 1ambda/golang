package interfaces

import (
	"fmt"
	"math"
)

type MyFloat float64
type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ExampleInterfaces() {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // MyFloat implements Abser
	fmt.Println(a)
	a = &v // *Vertex implements Abser (not Vertex4)
	fmt.Println(a)

	// Output:
	// -1.4142135623730951
	// &{3 4}
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func ExampleInterfacesValue() {
	var i I
	var t *T

	i = t       // method `M` will be called with a nil receiver
	describe(i) // An interface value that holds a nil concrete value is itself non-nil
	fmt.Println(i == nil)
	i.M()

	i = &T{"hello"}
	describe(i)
	fmt.Println(i == nil)
	i.M()

	// Output:
	// (<nil>, *interfaces.T)
	// false
	// <nil>
	// (&{hello}, *interfaces.T)
	// false
	// hello
}

func ExampleNilInterfaceValues() {
	var i I
	describe(i) // A nil interface value holds niether value nor concrete type
	fmt.Println(i == nil)

	// i.M() will cause runtime error

	// Output:
	// (<nil>, <nil>)
	// true
}

func describeAny(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func ExampleEmptyInterface() {
	// An empty interface may hold values of any type
	// since every type implements ay least zero methods

	var i interface{}
	describeAny(i)

	i = 42
	describeAny(i)

	i = "hello"
	describeAny(i)

	// Output:
	// (<nil>, <nil>)
	// (42, int)
	// (hello, string)
}

func ExampleTypeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f := i.float(64) will cause panic

	// Output:
	// hello
	// hello true
	// 0 false
}

func SwitchType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		// v is of the same interface type and value as i
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func ExampleTypeSwitches() {
	SwitchType(21)
	SwitchType("hello")
	SwitchType(true)

	// Output:
	// Twice 21 is 42
	// "hello" is 5 bytes long
	// I don't know about type bool!
}
