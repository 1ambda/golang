package main

import "fmt"

func main() {
	// pointers, slices, maps, channels, functions and interfaces
	// these types have `nil` as the zero value

	// nil slice has `nil` array
	var slice1 []byte
	equal(len(slice1) == 0)
	equal(cap(slice1) == 0)
	equal(slice1 == nil)

	// channels, maps, and functions are actually pointing their implementation
	var map1 map[string]int
	equal(map1 == nil)
	var func1 func()
	equal(func1 == nil)
	var chan1 chan int
	equal(chan1 == nil)

	// every interface has its underlying (type, value)
	var interface1 fmt.Stringer // Stringer (nil, nil)
	equal(interface1 == nil)    // (nil, nil) equals nil

	// IMPORTANT: but (*Person, nil) is not nil
	var p1 *Person
	var interface2 fmt.Stringer = p1
	equal(interface2 != nil)

	// IMPORTANT: do not declare concrete error
	err := wrapGetMyError()
	equal(err != nil)
}

func equal(expr bool) {
	if !expr {
		panic("error")
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %d", p.Name, p.Age)
}

type MyError struct{}

func (e MyError) Error() string {
	return "MyError :("
}

func getMyError() *MyError {
	return nil
}

func wrapGetMyError() error {
	return getMyError() // return `error (*MyError, nil)`. this is not nil!
}
