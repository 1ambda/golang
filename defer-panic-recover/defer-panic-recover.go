package main

import (
	"fmt"
	"io"
	"os"
)

func buggyCopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	// If os.create fails, `src` will not be closed
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	dst.Close()
	src.Close()

	return io.Copy(dst, src)
}

func copyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {

	/**
	 * A `defer` statement pushes a function call onto a list.
	 * The list of saved calls is executed after the surrounding function returns.
	 *
	 * The behavior of defer statements is straightforward and predictable.
	 * There are 3 simple rules.
	 */

	// 1. A deffered function's arguments are evaluated when the defer stmt is evaluated
	a := func() {
		i := 0
		defer fmt.Println(i) // i == 0
		i++
		return
	}
	_ = a

	// 2. Deferred function calls are executed
	// in LIFO order after the surrounding function returns

	b := func() {
		for i := 0; i < 4; i++ {
			defer fmt.Print(i) // 3210
		}
	}
	_ = b

	// 3. Deferred functinos may read and assign to
	// the returning function's named return values

	c := func(i int) int {
		defer func() { i++ }()
		return i
	}(1)
	_ = c

	/**
		 * `Panic` is a built in function that stops the ordinary flow of control and
		 * begins panicking. When the function `F` calls panic, executino of F stops,
		 * any deffered funcs in F are executed normally, and then F returns to its caller.
		 *
		 * To the caller, F then behaves like a call to panic.
		 * The process continues up the stack until all functions in
	   * the current goroutine have returned, at which point the program crashes.
		 *
		 * Panics can be initiated by invoking panic directly.
		 * The can also be caused by runtime errors, such as out-of-bouns array accesses.
	*/

	/**
	 * `Recover` is a built-in function tat regains control of a panicking goroutines.
	 * Recover is only useful inside deferred functions. During normal execution,
	 * a call to recover will return nil and have no other effect.
	 *
	 * If the current gorutine is panicking, a call to recover will capture the value given
	 * to panic and resume normal execution.
	 */

	f()
	fmt.Println("Returned normally from f.")

	// Output:
	// (f) Caliing g.
	// (g) P 0
	// (g) P 1
	// (g) P 2
	// (g) P 3
	// (g) Panicking!
	// (g) Defer in g 3
	// (g) Defer in g 2
	// (g) Defer in g 1
	// (g) Defer in g 0
	// (f) Recovered in f 4
	// (main) returned normally from f
}

func f() {
	// If we remove the deferred function here,
	// `f` will be not recovered and reaches the top of the goroutine's stack,
	// terminating the program.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
