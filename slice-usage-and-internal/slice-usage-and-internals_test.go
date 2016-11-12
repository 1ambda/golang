package slice_test

import (
	"fmt"
)

func ExampleSlice() {
	/**
	 * Go's arrays are values.
	 * An array variable denotes the entire array, not a pointer to the first element.
	 * This means that when you assign or pass around an array,
	 * you will make a copy of its content
	 * To avoid the copy, you could pass a pointer to array.
	 */
	initialized1 := [2]string{"Penn", "Teller"}
	initialized2 := [...]string{"Penn", "Teller"}
	use(initialized1, initialized2)

	// Output:

	// Example: initialize slice
	letters := []string{"a", "b", "c", "d"}
	use(letters)

	var bytes []byte = make([]byte, 5)
	test(len(bytes), cap(bytes))
	// 5 5

	// Example: slicing
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	// Comment: `b[1:4]` will share the same storage as b
	test(b[1:4])
	// "ola"
	test(b[:])
	// "golang"
	test(b[:2])
	// "go"
	test(b[2:])
	// "lang"

	/**
	 * A slice is a descriptor of an array segment.
	 * It soncicst of a
	 * - pointer to the array
	 * - the length of the segment
	 * - its capacity
	 *
	 *  Slicing doesn't copy the slice's data.
	 *  It just creates a newslice value that points to the original array.
	 *  Therefore, modifying the elements of a re-slice modifies the elements of the original
	 */

	road := []byte{'r', 'o', 'a', 'd'}
	road2 := road[2:]
	test(road2)
	// "ad"
	road2[1] = 'm'
	test(road2)
	// "am"
	test(road)
	// "roam"

	/**
	 * We can grow a shrinked slice to its capacity by slicing it again.
	 * `s = s[:cap(s)]`
	 *
	 * But a slice cannot be grown beyond its capacity.
	 * Attemping to do so will cause a runtime panic,
	 * just as when indexing outside the bounds of a slice or array
	 */

	/**
	 * To increase the cap of a slice, one must create a new, larger slice
	 * and copy the contents of the original slice into it.
	 */

	gopher := []string{"g", "o", "p", "h", "e", "r"}
	gopher2 := make([]string, len(gopher), (cap(gopher)+1)*2)
	for i := range gopher {
		gopher2[i] = gopher[i]
	}
	test(gopher2)
	// "gopher"
	test(cap(gopher2))
	// 12

	// Comment: we can use `copy` function instead of looping each element
	gopher3 := make([]string, len(gopher), (cap(gopher)+1)*2)
	copy(gopher3, gopher)
	test(cap(gopher2) == cap(gopher3))
	// True

	/**
	 * The `append` func appends elements to the end of a slice,
	 * and grows the slice if a greater cap is needed
	 */

	gopher4 := append(gopher, "gopher")
	test(cap(gopher4) == 12)
	// True

	// Comment: we can splat operator to append slice into slice
	names1 := []string{"John", "Paul"}
	names2 := []string{"George", "Ringo", "Pete"}
	names3 := append(names1, names2...)
	test(len(names3))
	// 5

	// See possible gotcha from here
	// https://blog.golang.org/go-slices-usage-and-internals#TOC_6.
	// Summary: Copy only needed elements before return
	//          otherwise, golang runtime will keep all data
}

func test(vals ...interface{}) {
	fmt.Println("%q", vals)
}

func use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}
