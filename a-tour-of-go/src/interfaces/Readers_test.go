package interfaces

import (
	"fmt"
	"io"
	"strings"
)

// Reader represents the read end of a stream of date
func ExampleReaders() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	// func (T) Read(b []byte) (n int, err error)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)

		// %q: a single-quoted char literal safely escaped with Go syntax
		// see https://golang.org/pkg/fmt/
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}

	// Output:
	// n = 8, err = <nil>, b = [72 101 108 108 111 44 32 82]
	// b[:n] = "Hello, R"
	// n = 6, err = <nil>, b = [101 97 100 101 114 33 32 82]
	// b[:n] = "eader!"
	// n = 0, err = EOF, b = [101 97 100 101 114 33 32 82]
	// b[:n] = ""
}
