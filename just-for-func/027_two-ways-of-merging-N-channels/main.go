package main

import (
	//"time"
	//"math/rand"
	"fmt"
	"sync"
	"reflect"
)

func main() {
	a := asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := asChan(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c := asChan(20, 21, 22, 23, 24, 25, 26, 27, 28, 29)

	for v := range mergeReflect(a, b, c) {
		fmt.Println(v)
	}
}

// slow
func merge1(chans ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, c := range chans {
			// we can't use select stmt since the number of channels is random
			// also, we need to pass `c` instead of using local `c` which is changes in parent goroutine
			for v := range c {
				out <- v
			}
		}
	}()

	return out
}

// more concurrent
func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		var wg sync.WaitGroup
		wg.Add(len(chans))

		for _, c := range chans {
			// we can't use select stmt since the number of channels is random
			// also, we need to pass `c` instead of using local `c` which is changes in parent goroutine
			go func(c <-chan int) {
				for v := range c {
					out <- v
				}

				wg.Done()
			}(c)
		}

		wg.Wait()
	}()

	return out
}

func mergeReflect(chans ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				// remove `i` from slice
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}

			out <- v.Interface().(int)
		}
	}()

	return out
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)

		for _, v := range vs {
			c <- v
			//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()

	return c
}

func mergeRecur(chans ...<-chan int) <-chan int {
	switch len(chans) {
	case 0:
		c := make(chan int)
		close(c)
		return c
	case 1:
		return chans[0]
	case 2:
		return mergeTwo(chans[0], chans[1])
	default:
		m := len(chans) / 2
		return mergeTwo(mergeRecur(chans[:m]...), mergeRecur(chans[m:]...))
	}
}

func mergeTwo(a, b <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

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
	}()

	return out
}
