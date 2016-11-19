package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// make zero value useful

	// 1. pointers: we can call even `nil` receiver!
	var t *tree
	equal(t.sum() == 0)

	// 2. slices
	var s []int
	for i := range s { //interates zero times
		_ = i
		panic("won't be executed")
	}

	for i := 0; i < 10; i++ {
		// append on nil slices
		s = append(s, i) // fast enought since slice increase its cap twice if needed
	}

	// IMPORTANT: s[0] will cause panic

	// 3. maps
	var headers map[string]string
	equal(len(headers) == 0)

	for k, v := range headers { // iterates zero times
		_, _ = k, v
		panic("won't be executed")
	}
	headerValue, headerExist := headers["Accept"]
	equal(headerValue == "")
	equal(headerExist == false)

	// IMPORTANT: headers["abc"] = "def" will cause panic

	// good example making use of nil map
	newGet := func(url string, headers map[string]string) (*http.Request, error) {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		for k, v := range headers {
			req.Header.Set(k, v)
		}

		return req, nil
	}
	_ = newGet
	// we can pass `nil` instead of `map[string]string{}`
	// IMPORTANT: Use `nil` maps as read-only empty maps

	// 4. channels
	var c1 chan int
	equal(c1 == nil)
	// IMPORTANT: `<- c` blocks forever (reading nil channel)
	// IMPORTANT: `c <- x` blocks forever (writing to nil channel)
	// IMPORTANT: `close(c)` causes panic (closing nil channel)

	// we can check channnel is closed or not
	c2 := make(chan int)
	close(c2)
	cv, ok := <-c2
	equal(ok == false)
	_ = cv

	// IMPORTANT: use nil chans to disable a select case
	// see merge1, merge2, merge3

	// 5. functions: nil funcs for default value
	newServerFunc := func(logger func(string, ...interface{})) {
		if logger == nil {
			logger = log.Printf // nil can also imply default behavior
		}
	}
	_ = newServerFunc

	// 6. interfaces: The nil interface is used as a signal for `default`` for the interface

	doSum := func(s Summer) int {
		if s == nil {
			return 0
		}
		return s.sum()
	}

	// (*tree, nil) and default behavior of `*tree`
	var tr *tree
	equal(doSum(tr) == 0)
	// (ints, nil) and default behavior of `ints`
	var is ints
	equal(doSum(is) == 0)
	// (nil, nil) which implies we need really `default behavior` for `Summer`
	equal(doSum(nil) == 0)

}

func equal(expr bool) {
	if !expr {
		panic("error")
	}
}

type Summer interface {
	sum() int
}

type ints []int

func (i ints) sum() int {
	s := 0
	for _, v := range i {
		s += v
	}

	return s
}

func merge3(out chan<- int, a, b <-chan int) {
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil
				continue
			}
			out <- v
		case v, ok := <-b:
			if !ok {
				b = nil
				continue
			}
			out <- v
		}
	}

	close(out)
}

func merge2(out chan<- int, a, b <-chan int) {
	var aClosed, bClosed bool

	for !aClosed || !bClosed {
		select {
		case v, ok := <-a: // PROBLEM: will cause busy loop
			if !ok {
				aClosed = true
				continue
			}
			out <- v
		case v, ok := <-b: // PROBLEM: will cause busy loop
			if !ok {
				bClosed = true
				continue
			}
			out <- v
		}
	}

	// without closing `out`, this blocks forever (deadlock)
	close(out)
}

// PROBLEM: run forever if `out` is closed
func merge1(out chan<- int, a, b <-chan int) {
	for {
		select {
		case v := <-a:
			out <- v
		case v := <-b:
			out <- v
		}
	}
}

type tree struct {
	v int
	l *tree
	r *tree
}

// nil receivers are useful
func (t *tree) sum() int {
	if t == nil {
		return 0
	}

	return t.v + t.l.sum() + t.r.sum()
}

func (t *tree) String() string {
	if t == nil {
		return ""
	}

	return fmt.Sprint(t.l, t.v, t.r)
}

func (t *tree) Find(v int) bool {
	if t == nil {
		return false
	}

	return t.v == v || t.l.Find(v) || t.r.Find(v)
}

// problem of sum1
// - code repretition: `if v != nil...`
// - panic when `t` is nil
func (t *tree) sum1() int {
	sum := t.v

	if t.l != nil {
		sum += t.l.sum1()
	}

	if t.r != nil {
		sum += t.r.sum1()
	}

	return sum
}
