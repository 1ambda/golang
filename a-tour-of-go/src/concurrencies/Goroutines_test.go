package concurrencies

import (
	"fmt"
	"time"
)

/*

 A `goroutine` is a lightweight thread managed by the Go runtime

 The evaluation of `f`, `x`, `y` and `z` happens in the current goroutine and
 the execution of `f` happens in the new goroutine

 Goroutines run in the same address space, so access to shared memory must be synchronized
 (The `sync` package provides useful primitives)

*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func ExampleGoroutinesSay() {
	go say("world")
	say("hello")

	// Output might be
	// hello
	// world
	// world
	// hello
	// hello
	// world
	// world
	// hello
	// hello
	// world
}

/*

 Channels are a typed conduit through which you can send and receive values
 By default, sends are receives block until the other side is ready.

*/

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum // send sum to c
}

func ExampleChannels() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	left := s[:len(s)/2]
	right := s[len(s)/2:]

	go sum(left, c)
	go sum(right, c)

	x, y := <-c, <-c

	fmt.Println(x + y)

	// Output: 12
}

func ExampleBufferedChannels() {
	// Send to a buffered channel block only when the buffer is full.
	// Receives block when the buffer is empty
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Output:
	// 1
	// 2
}

func fibo1(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func ExampleRangeAndClose() {
	c := make(chan int, 10)

	fmt.Println(cap(c))

	go fibo1(cap(c), c)

	// we can test whether a channel is closed or not by `v, ok := <- c`

	// range loop receives values from the channel repeatedly until it is closed
	for i := range c {
		fmt.Println(i)
	}

	/*
			Only the sender should close a channel, never the receiver.
			Sending on a closed channel will cause a panic

		  Channels aren't like files. You don't usually need to close them.
		  Closing is only necessary when the receiver must be told there are no more values coming,
		  such as to terminate a `range` loop
	*/

	// Output:
	// 10
	// 0
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
}

/*
  The `select` statement lets a goroutine wait on multiple communication operations
  A `select` blocks until one of its cases can run, then it executes that case.
  It chooses one at random if multiple are ready.
*/

func fibo2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func ExampleSelect() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibo2(c, quit)

	// Output:
	// 0
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
	// quit
}

func ExampleDefaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom!")
			return
		default:
			fmt.Println(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
