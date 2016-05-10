package interfaces

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rotate13(b byte) byte {
	if 'a' <= b && b <= 'z' {
		return (b-'a'+13)%26 + 'a'
	}

	if 'A' <= b && b <= 'Z' {
		return (b-'A'+13)%26 + 'A'
	}

	return b
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)

	for i := 0; i < n; i++ {
		b[i] = rotate13(b[i])
	}

	return n, err
}

func ExampleR13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	// Output: You cracked the code!
}
