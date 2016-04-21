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

func sum(xs []int) int {
	acc := 0

	for _, v := range xs {
		acc += v
	}

	return acc
}

func whileUsingFor() {
	acc := 1

	for acc < 1000 {
		acc += 1
	}

	fmt.Println(acc)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) floa64 {
	if v:= 
}

func main() {
	fmt.Println(sqrt(-3))
}
