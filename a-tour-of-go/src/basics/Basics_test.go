package basics

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
)

func f1() {
	const Pi = 3.14
}

func f2() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

func f3() {
	fmt.Printf("Now you have %g problems", math.Sqrt(7))
}

func add(x, y int) int {
	return x + y
}

func swap(s1, s2 string) (string, string) {
	return s2, s1
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func ExampleFormatting1() {
	var (
		ToBe   bool       = false
		MaxInt uint64     = (1 << 64) - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	const f = "type: %T, value: %v\n"

	println(ToBe, MaxInt, z)

	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	// Output:
	// type: bool, value: false
	// type: uint64, value: 18446744073709551615
	// type: complex128, value: (2+3i)
}

func ExampleFormatting2() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// Output: 0 0 false ""
}

func ExampleCastType() {
	i := 42
	f := float64(i)
	u := uint(f)

	fmt.Printf("%v %v %v \n", i, f, u)

	// Output: 42 42 42
}

func sum(xs []int) int {
	acc := 0

	for _, v := range xs {
		acc += v
	}

	return acc
}

func ExampleWhileUsingFor() {
	acc := 1

	for acc < 1000 {
		acc += 1
	}

	fmt.Println(acc)

	// Output: 1000
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim // can't use v here
}

func Sqrt1(x float64) float64 {
	z := float64(1)

	for i := 1; i < 10; i++ {
		z = z - (math.Pow(z, 2)-x)/(2*z)
	}

	return z
}

func printOS() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX!")
	case "linux":
		fmt.Println("LINUX!")
	default:
		fmt.Printf("Maybe Window, %s", os)
	}
}

func deferTest() {
	// deferred call's ars are evaluated immediately
	// ut function call is not execued until the surrouding function returns
	defer fmt.Println("world")

	fmt.Println("hello")
}

func ExampleDeferStack() {
	// deferred calls are pushed onto a stack
	// when a function returnes, its deferred calls are executed in LIFO order

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("return")

	// Output:
	// return
	// 9
	// 8
	// 7
	// 6
	// 5
	// 4
	// 3
	// 2
	// 1
	// 0
}
