package main

import (
	"fmt"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {

	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	// build chain sequencially from leftmost to right
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	// trigger
	go func(c chan int) { c <- 1 }(right)

	// get result
	fmt.Println(<-leftmost)
}
