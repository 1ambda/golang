package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
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

func formatting1() {
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
}

func formatting2() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func castType() {
	i := 42
	f := float64(i)
	u := uint(f)

	fmt.Printf("%v %v %v \n", i, f, u)
}

func main() {
	formatting2()
}
