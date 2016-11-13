package main

import (
	"fmt"
)

// `error` type is an interface
//
// type error interface {
//   Error() string
// }

// We can implement our own error type
type myErrorString struct {
	s string
}

func (e *myErrorString) Error() string {
	return e.s
}

func new(text string) error {
	return &myErrorString{text}
}

// Since `error` is an interface, we can use arbitrary data structures
type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("math: square root of negative number %g", float64(f))
}

/**
 * We can use type assertion like
 *
 * if nerr, ok := err.(NetworkError); ok && nerr.Temporate() { ... }
 */
type NetworkError interface {
	error
	Timeout() bool
	Temporary() bool
}

// See more: https://golang.org/doc/faq#nil_error

func main() {

}
