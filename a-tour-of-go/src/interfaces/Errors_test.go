package interfaces

import (
	"fmt"
	"time"
)

/*

The error type is a built-in interface similar to fmt.Stringer

type error interface {
  Error() string
}

*/

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func raiseMyError() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func ExampleMyError() {
	if err := raiseMyError(); err != nil {
		fmt.Println(err)
	}

	// Output looks like `at 2016-05-10 00:27:03.91289924 +0900 KST, it didn't work`
}
