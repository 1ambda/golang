package main

func main() {
	test(closure1(1)() == 2)

	inner, two := closure2()
	test(inner() == 101)
	test(two == 2)

	test(closure3(1) == 3)
}

func closure1(n int) func() int {
	outer_var := 2

	// lexically scoped:
	// functions can access values that were in the same scoped when defining the function
	c := func() int {
		return outer_var + n
	}

	// but captured variable can be changed when original variable is changed
	outer_var = 1

	return c
}

func closure2() (func() int, int) {
	outer_var := 2

	// but original value doens't change even though the captured variable is changed
	inner := func() int {
		outer_var += 99 // this code do not change original `outer_var`, just redefining it
		return outer_var
	}

	return inner, outer_var // func, 2
}

func closure3(n int) int {
	outer_var := 2

	// change local (out of func) variable
	func() {
		outer_var = outer_var + n
	}()

	return outer_var
}

func test(expr bool) {
	if !expr {
		panic(0)
	}
}
